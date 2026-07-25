package driver

import (
	"common/pkg/constant"
	"context"
	"log/slog"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Config interface {
	GetDebug() bool
	GetSlowSqlThreshold() *durationpb.Duration
	GetSampleRate() float64
}

type observedDriver struct {
	wrapped dialect.Driver
	logger  *slog.Logger
	policy  policy
}

func Wrap(drv dialect.Driver, logger *slog.Logger, config Config) dialect.Driver {
	return &observedDriver{
		wrapped: drv,
		logger:  logger,
		policy:  newPolicy(config),
	}
}

func (d *observedDriver) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = d.wrapped.Exec(ctx, query, args, v)
	d.recordSQL(ctx, query, args, time.Since(start), err)
	return err
}

func (d *observedDriver) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = d.wrapped.Query(ctx, query, args, v)
	d.recordSQL(ctx, query, args, time.Since(start), err)
	return err
}

func (d *observedDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.wrapped.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &observedTx{
		Tx:     tx,
		logger: d.logger,
		policy: d.policy,
		txID:   uuid.New().String(),
	}, nil
}

func (d *observedDriver) Close() error {
	return d.wrapped.Close()
}

func (d *observedDriver) Dialect() string {
	return d.wrapped.Dialect()
}

func (d *observedDriver) recordSQL(ctx context.Context, query string, args any, latency time.Duration, err error) {
	level, shouldRecord, withCaller := d.policy.sqlLevel(latency, err)
	if !shouldRecord {
		return
	}
	attrs := sqlAttrs(query, args, latency, err, "")
	if withCaller {
		if caller, ok := findCaller(); ok {
			attrs = append(attrs,
				slog.String(constant.LogFieldFile, caller.file),
				slog.Int(constant.LogFieldLine, caller.line),
			)
		}
	}
	d.logger.LogAttrs(ctx, level, "sql executed", attrs...)
}

type observedTx struct {
	dialect.Tx
	logger *slog.Logger
	policy policy
	txID   string
}

func (t *observedTx) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = t.Tx.Exec(ctx, query, args, v)
	t.recordSQL(ctx, query, args, time.Since(start), err)
	return err
}

func (t *observedTx) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = t.Tx.Query(ctx, query, args, v)
	t.recordSQL(ctx, query, args, time.Since(start), err)
	return err
}

func (t *observedTx) Commit() error {
	start := time.Now()
	err := t.Tx.Commit()
	latency := time.Since(start)
	level, shouldRecord := t.policy.txLevel(latency, err, false)
	if !shouldRecord {
		return err
	}
	attrs := []slog.Attr{
		slog.String(constant.LogFieldKind, constant.LogKindSQL),
		slog.String(constant.LogFieldTxID, t.txID),
		slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
	}
	if err != nil {
		attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
	}
	t.logger.LogAttrs(context.Background(), level, "tx commit", attrs...)
	return err
}

func (t *observedTx) Rollback() error {
	start := time.Now()
	err := t.Tx.Rollback()
	latency := time.Since(start)
	level, shouldRecord := t.policy.txLevel(latency, err, true)
	if !shouldRecord {
		return err
	}
	t.logger.LogAttrs(context.Background(), level, "tx rollback",
		slog.String(constant.LogFieldKind, constant.LogKindSQL),
		slog.String(constant.LogFieldTxID, t.txID),
		slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
		slog.Any(constant.LogFieldErr, err),
	)
	return err
}

func (t *observedTx) recordSQL(ctx context.Context, query string, args any, latency time.Duration, err error) {
	level, shouldRecord, withCaller := t.policy.sqlLevel(latency, err)
	if !shouldRecord {
		return
	}
	attrs := sqlAttrs(query, args, latency, err, t.txID)
	if withCaller {
		if caller, ok := findCaller(); ok {
			attrs = append(attrs,
				slog.String(constant.LogFieldFile, caller.file),
				slog.Int(constant.LogFieldLine, caller.line),
			)
		}
	}
	t.logger.LogAttrs(ctx, level, "sql executed", attrs...)
}

const defaultSlowThreshold = 500 * time.Millisecond

type policy struct {
	debug         bool
	slowSQL       time.Duration
	sampleRate    float64
	hasConfigured bool
}

func newPolicy(config Config) policy {
	if config == nil {
		return policy{
			slowSQL: defaultSlowThreshold,
		}
	}
	threshold := defaultSlowThreshold
	if config.GetSlowSqlThreshold() != nil && config.GetSlowSqlThreshold().AsDuration() > 0 {
		threshold = config.GetSlowSqlThreshold().AsDuration()
	}
	return policy{
		debug:         config.GetDebug(),
		slowSQL:       threshold,
		sampleRate:    config.GetSampleRate(),
		hasConfigured: true,
	}
}

func (p policy) sqlLevel(latency time.Duration, err error) (slog.Level, bool, bool) {
	if err != nil {
		return slog.LevelError, true, true
	}
	if latency > p.slowSQL {
		return slog.LevelWarn, true, true
	}
	if p.debug || p.sampled() {
		return slog.LevelDebug, true, false
	}
	return slog.LevelDebug, false, false
}

func (p policy) txLevel(latency time.Duration, err error, rollback bool) (slog.Level, bool) {
	if err != nil {
		return slog.LevelError, true
	}
	if rollback {
		return slog.LevelDebug, false
	}
	if latency > p.slowSQL {
		return slog.LevelWarn, true
	}
	return slog.LevelDebug, false
}

func (p policy) sampled() bool {
	if !p.hasConfigured {
		return false
	}
	if p.sampleRate <= 0 {
		return false
	}
	if p.sampleRate >= 1 {
		return true
	}
	return rand.Float64() < p.sampleRate
}

func sqlAttrs(query string, args any, latency time.Duration, err error, txID string) []slog.Attr {
	attrs := []slog.Attr{
		slog.String(constant.LogFieldKind, constant.LogKindSQL),
		slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
		slog.String(constant.LogFieldSQL, normalizeSQL(query)),
		slog.Int(constant.LogFieldArgsCount, argsCount(args)),
	}
	if txID != "" {
		attrs = append(attrs, slog.String(constant.LogFieldTxID, txID))
	}
	if err != nil {
		attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
	}
	return attrs
}

func normalizeSQL(query string) string {
	return strings.Join(strings.Fields(query), " ")
}

func argsCount(args any) int {
	switch v := args.(type) {
	case nil:
		return 0
	case []any:
		return len(v)
	default:
		return 1
	}
}

var (
	callerCache sync.Map
	selfPkg     string
)

type callerInfo struct {
	file string
	line int
}

func init() {
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return
	}
	name := fn.Name()
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		selfPkg = name[:idx]
	}
}

func findCaller() (callerInfo, bool) {
	var pcs [32]uintptr
	n := runtime.Callers(2, pcs[:])
	for i := 0; i < n; i++ {
		pc := pcs[i]
		if v, ok := callerCache.Load(pc); ok {
			return v.(callerInfo), true
		}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)
		if skipCaller(fn.Name(), file) {
			continue
		}
		caller := callerInfo{
			file: file,
			line: line,
		}
		callerCache.Store(pc, caller)
		return caller, true
	}
	return callerInfo{}, false
}

func skipCaller(funcName string, file string) bool {
	normalizedFile := strings.ReplaceAll(file, "\\", "/")
	if strings.HasPrefix(funcName, "runtime.") || strings.HasPrefix(funcName, "runtime/") {
		return true
	}
	if strings.HasPrefix(funcName, "time.") || strings.HasPrefix(funcName, "time/") {
		return true
	}
	if strings.HasPrefix(funcName, "database/sql.") || strings.HasPrefix(funcName, "database/sql/") {
		return true
	}
	if selfPkg != "" && strings.HasPrefix(funcName, selfPkg+".") {
		return true
	}
	if strings.Contains(normalizedFile, "/go/pkg/mod/") || strings.Contains(normalizedFile, "/pkg/mod/") {
		return true
	}
	if strings.Contains(normalizedFile, "/entgo.io/") {
		return true
	}
	if strings.Contains(normalizedFile, "/internal/data/gen/") {
		return true
	}
	return false
}
