package driver

import (
	"common/pkg/constant"
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
)

var m = constant.Dev

type Config interface {
	GetDebug() bool
	GetSlowSqlThreshold() *durationpb.Duration
	GetSampleRate() float64
}

type driver struct {
	l      *log.Helper
	driver dialect.Driver
	config Config
}

func NewDriver(l log.Logger, mode string, drv dialect.Driver, config Config) dialect.Driver {
	m = mode
	return &driver{
		l:      log.NewHelper(l),
		driver: drv,
		config: config,
	}
}

// ===================== Exec =====================

func (d *driver) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()

	err = d.driver.Exec(ctx, query, args, v)
	cost := time.Since(start)

	if shouldLog(d.config, cost) {
		logSQL(ctx, d.l, cost, query, args)
	}
	return
}

// ===================== Query =====================

func (d *driver) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()

	err = d.driver.Query(ctx, query, args, v)
	cost := time.Since(start)

	if shouldLog(d.config, cost) {
		logSQL(ctx, d.l, cost, query, args)
	}
	return
}

// ===================== Tx =====================

func (d *driver) Tx(ctx context.Context) (dialect.Tx, error) {
	t, err := d.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &tx{
		l:      d.l,
		Tx:     t,
		config: d.config,
	}, nil
}

func (d *driver) Close() error {
	return d.driver.Close()
}

func (d *driver) Dialect() string {
	return d.driver.Dialect()
}

// ===================== Tx Wrapper =====================

type tx struct {
	l *log.Helper
	dialect.Tx
	config Config
}

func (t *tx) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()

	err = t.Tx.Exec(ctx, query, args, v)
	cost := time.Since(start)

	if shouldLog(t.config, cost) {
		logSQL(ctx, t.l, cost, query, args)
	}
	return
}

func (t *tx) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()

	err = t.Tx.Query(ctx, query, args, v)
	cost := time.Since(start)

	if shouldLog(t.config, cost) {
		logSQL(ctx, t.l, cost, query, args)
	}
	return
}

// ===================== Log 控制 =====================

func shouldLog(cfg Config, cost time.Duration) bool {
	threshold := 500 * time.Millisecond
	if cfg.GetSlowSqlThreshold() != nil {
		threshold = cfg.GetSlowSqlThreshold().AsDuration()
	}

	if cfg.GetDebug() || cost > threshold {
		return true
	}
	return rand.Float64() < cfg.GetSampleRate()
}

// ===================== 日志输出（核心优化） =====================

func logSQL(ctx context.Context, l *log.Helper, cost time.Duration, query string, args any) {
	argv, ok := args.([]any)

	var sql string
	if ok {
		sql = formatSQLArgs(query, argv)
	} else {
		sql = query
	}
	sql = strings.ReplaceAll(sql, "\n", " ")
	sql = strings.ReplaceAll(sql, "\r", " ")
	sql = strings.ReplaceAll(sql, "\t", " ")
	sql = strings.TrimSpace(sql)

	file, line := getCaller()

	// ===== dev：可读 =====
	if m == constant.Dev {
		l.WithContext(ctx).Debugf(
			"[SQL] %dms | %s:%d | %s",
			cost.Milliseconds(),
			file,
			line,
			sql,
		)
		return
	}

	// ===== prod：结构化 =====
	l.WithContext(ctx).Debugw(
		"sql executed",
		"cost_ms", cost.Milliseconds(),
		"sql", sql,
		"file", file,
		"line", line,
	)
}

// ===================== SQL 参数拼接 =====================

func formatSQLArgs(query string, args []any) string {
	if len(args) == 0 {
		return query
	}

	var buf strings.Builder
	buf.Grow(len(query) + len(args)*8)

	argIdx := 0

	for i := 0; i < len(query); {
		c := query[i]

		if c == '?' && argIdx < len(args) {
			buf.WriteString(argToString(args[argIdx]))
			argIdx++
			i++
			continue
		}

		if c == '$' {
			j := i + 1
			num := 0

			for j < len(query) && query[j] >= '0' && query[j] <= '9' {
				num = num*10 + int(query[j]-'0')
				j++
			}
			if num > 0 && num <= len(args) {
				buf.WriteString(argToString(args[num-1]))
				i = j
				continue
			}
		}

		buf.WriteByte(c)
		i++
	}

	return buf.String()
}

func argToString(arg any) string {
	v := reflect.ValueOf(arg)

	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "NULL"
		}
		return argToString(v.Elem().Interface())
	}

	switch a := arg.(type) {
	case string:
		return "'" + strings.ReplaceAll(a, "'", "\\'") + "'"
	case int, int64, int32, float64, float32:
		return fmt.Sprintf("%v", a)
	case time.Time:
		return "'" + a.Format("2006-01-02 15:04:05") + "'"
	case bool:
		if a {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("'%v'", a)
	}
}

// ===================== 调用栈 =====================

var callerCache sync.Map

func getCaller() (string, int) {
	var pcs [64]uintptr
	n := runtime.Callers(0, pcs[:])

	for i := 0; i < n; i++ {
		pc := pcs[i]

		if v, ok := callerCache.Load(pc); ok {
			c := v.(caller)
			return c.file, c.line
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)

		// 过滤 runtime / 第三方
		if strings.Contains(file, "runtime") || strings.Contains(file, "go/pkg/mod") {
			continue
		}

		if strings.Contains(file, "internal/data/repo") {
			callerCache.Store(pc, caller{file, line})
			return file, line
		}
	}

	return "unknown", 0
}

type caller struct {
	file string
	line int
}
