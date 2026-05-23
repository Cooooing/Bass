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
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
)

// ============================================================
// 配置
// ============================================================

type Config interface {
	GetDebug() bool
	GetSlowSqlThreshold() *durationpb.Duration
	GetSampleRate() float64
}

// ============================================================
// 驱动
// ============================================================

type driver struct {
	l       *log.Helper
	wrapped dialect.Driver
	config  Config
	mode    string
}

func NewDriver(l log.Logger, mode string, drv dialect.Driver, config Config) dialect.Driver {
	return &driver{
		l:       log.NewHelper(l),
		wrapped: drv,
		config:  config,
		mode:    mode,
	}
}

func (d *driver) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = d.wrapped.Exec(ctx, query, args, v)
	cost := time.Since(start)
	if shouldLog(d.config, cost) {
		logSQL(ctx, d.l, d.mode, cost, query, args, err, "")
	}
	return
}

func (d *driver) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = d.wrapped.Query(ctx, query, args, v)
	cost := time.Since(start)
	if shouldLog(d.config, cost) {
		logSQL(ctx, d.l, d.mode, cost, query, args, err, "")
	}
	return
}

func (d *driver) Tx(ctx context.Context) (dialect.Tx, error) {
	t, err := d.wrapped.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &tx{
		l:      d.l,
		Tx:     t,
		config: d.config,
		mode:   d.mode,
		txID:   uuid.New().String(),
	}, nil
}

func (d *driver) Close() error {
	return d.wrapped.Close()
}

func (d *driver) Dialect() string {
	return d.wrapped.Dialect()
}

// ============================================================
// 事务
// ============================================================

type tx struct {
	l *log.Helper
	dialect.Tx
	config Config
	mode   string
	txID   string
}

func (t *tx) Exec(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = t.Tx.Exec(ctx, query, args, v)
	cost := time.Since(start)
	if shouldLog(t.config, cost) {
		logSQL(ctx, t.l, t.mode, cost, query, args, err, t.txID)
	}
	return
}

func (t *tx) Query(ctx context.Context, query string, args, v any) (err error) {
	start := time.Now()
	err = t.Tx.Query(ctx, query, args, v)
	cost := time.Since(start)
	if shouldLog(t.config, cost) {
		logSQL(ctx, t.l, t.mode, cost, query, args, err, t.txID)
	}
	return
}

func (t *tx) Commit() error {
	start := time.Now()
	err := t.Tx.Commit()
	cost := time.Since(start)
	if err != nil || cost > slowThreshold(t.config) {
		t.l.Warnw("tx commit",
			"tx_id", t.txID,
			"cost_ms", cost.Milliseconds(),
			"error", err,
		)
	}
	return err
}

func (t *tx) Rollback() error {
	start := time.Now()
	err := t.Tx.Rollback()
	cost := time.Since(start)
	if err != nil {
		t.l.Warnw("tx rollback",
			"tx_id", t.txID,
			"cost_ms", cost.Milliseconds(),
			"error", err,
		)
	}
	return err
}

// ============================================================
// 日志判定
// ============================================================

const defaultSlowThreshold = 500 * time.Millisecond

func slowThreshold(cfg Config) time.Duration {
	if t := cfg.GetSlowSqlThreshold(); t != nil {
		return t.AsDuration()
	}
	return defaultSlowThreshold
}

func shouldLog(cfg Config, cost time.Duration) bool {
	if cfg.GetDebug() {
		return true
	}
	if cost > slowThreshold(cfg) {
		return true
	}
	rate := cfg.GetSampleRate()
	if rate <= 0 {
		return false
	}
	if rate >= 1 {
		return true
	}
	return rand.Float64() < rate
}

// ============================================================
// 日志输出
// ============================================================

var whitespaceReplacer = strings.NewReplacer("\n", " ", "\r", " ", "\t", " ")

func logSQL(ctx context.Context, l *log.Helper, mode string, cost time.Duration, query string, args any, err error, txID string) {
	argv, _ := args.([]any)
	sql := formatSQLArgs(query, argv)
	sql = whitespaceReplacer.Replace(sql)

	file, line := getCaller()

	txPrefix := ""
	if txID != "" {
		txPrefix = "Tx(" + txID + ")."
	}

	if mode == constant.Dev {
		if err != nil {
			l.WithContext(ctx).Warnf("[SQL] %dms | %s:%d | %sExec: err=%v | %s",
				cost.Milliseconds(), file, line, txPrefix, err, sql)
		} else {
			l.WithContext(ctx).Debugf("[SQL] %dms | %s:%d | %sExec: %s",
				cost.Milliseconds(), file, line, txPrefix, sql)
		}
		return
	}

	// 生产环境：结构化 JSON。
	if err != nil {
		l.WithContext(ctx).Warnw("sql executed",
			"cost_ms", cost.Milliseconds(),
			"sql", sql,
			"file", file,
			"line", line,
			"tx_id", txID,
			"error", err,
		)
	} else {
		l.WithContext(ctx).Debugw("sql executed",
			"cost_ms", cost.Milliseconds(),
			"sql", sql,
			"file", file,
			"line", line,
			"tx_id", txID,
		)
	}
}

// ============================================================
// SQL 参数格式化
// ============================================================

func formatSQLArgs(query string, args []any) string {
	if len(args) == 0 {
		return query
	}

	var buf strings.Builder
	buf.Grow(len(query) + len(args)*16)

	argIdx := 0
	for i := 0; i < len(query); {
		c := query[i]

		// PostgreSQL 占位符：$1、$2 等。
		if c == '$' {
			j := i + 1
			num := 0
			for j < len(query) && query[j] >= '0' && query[j] <= '9' {
				num = num*10 + int(query[j]-'0')
				j++
			}
			if num > 0 && num <= len(args) {
				buf.WriteString(formatArg(args[num-1]))
				i = j
				continue
			}
		}

		// MySQL 占位符：?。
		if c == '?' && argIdx < len(args) {
			buf.WriteString(formatArg(args[argIdx]))
			argIdx++
			i++
			continue
		}

		buf.WriteByte(c)
		i++
	}

	return buf.String()
}

const (
	maxStringLen = 512
	maxByteLen   = 64
)

func formatArg(arg any) string {
	if arg == nil {
		return "NULL"
	}

	v := reflect.ValueOf(arg)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "NULL"
		}
		return formatArg(v.Elem().Interface())
	}

	switch a := arg.(type) {
	case string:
		if len(a) > maxStringLen {
			a = a[:maxStringLen] + "..."
		}
		return "'" + strings.ReplaceAll(a, "'", "''") + "'"
	case []byte:
		if len(a) > maxByteLen {
			return fmt.Sprintf("0x%x...(len=%d)", a[:maxByteLen], len(a))
		}
		return fmt.Sprintf("0x%x", a)
	case time.Time:
		return "'" + a.Format("2006-01-02 15:04:05") + "'"
	case bool:
		if a {
			return "true"
		}
		return "false"
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", a)
	default:
		return fmt.Sprintf("'%v'", a)
	}
}

// ============================================================
// 调用方解析
// ============================================================

var (
	callerCache sync.Map // 程序计数器到调用方信息的缓存
	selfPkg     string   // 当前包导入路径，初始化时检测
)

type callerInfo struct {
	file string
	line int
}

// 启动时检测当前包导入路径。
func init() {
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		name := fn.Name()
		if idx := strings.LastIndex(name, "."); idx >= 0 {
			selfPkg = name[:idx]
		}
	}
}

// getCaller 遍历调用栈，查找触发 SQL 的 data 层栈帧。
// 会跳过 runtime、当前包、第三方模块、ent 生成代码和标准库数据库包。
//
// 典型调用链：
//
//	service.Register
//	  → repo.Create          ← 目标调用方
//	    → ent.UserCreate.Save
//	      → ent.UserCreate.sqlExec
//	        → driver.Exec
//	          → driver.logSQL
//	            → driver.getCaller
//	              → runtime.Callers
func getCaller() (string, int) {
	var pcs [32]uintptr
	n := runtime.Callers(2, pcs[:]) // 跳过 runtime.Callers 和 getCaller。

	for i := 0; i < n; i++ {
		pc := pcs[i]

		if v, ok := callerCache.Load(pc); ok {
			ci := v.(callerInfo)
			return ci.file, ci.line
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		file, line := fn.FileLine(pc)
		name := fn.Name()

		if shouldSkip(name, file) {
			continue
		}

		ci := callerInfo{file: file, line: line}
		callerCache.Store(pc, ci)
		return file, line
	}

	return "unknown", 0
}

// shouldSkip 判断某个栈帧是否应被跳过。
//
// 跳过顺序（从栈顶到栈底）：
//  1. runtime 内部栈帧
//  2. 当前 driver 包
//  3. 第三方模块（go/pkg/mod）
//  4. ent 生成代码
//  5. 标准库数据库包
//
// 第一个通过检查的栈帧就是用户 data 层代码。
func shouldSkip(funcName, file string) bool {
	// runtime 内部栈帧。
	if strings.HasPrefix(funcName, "runtime.") || strings.HasPrefix(funcName, "runtime/") {
		return true
	}
	// 当前 driver 包。
	if selfPkg != "" && strings.HasPrefix(funcName, selfPkg+".") {
		return true
	}
	// 模块缓存中的第三方模块。
	if strings.Contains(file, "/go/pkg/mod/") {
		return true
	}
	// ent 生成代码：函数名通常类似：
	//   github.com/project/internal/data/gen/ent.(*UserCreate).Save
	// "/ent." 模式可以稳定匹配包边界。
	if strings.Contains(funcName, "/ent.") {
		return true
	}
	// 标准库数据库包。
	if strings.HasPrefix(funcName, "database/") {
		return true
	}
	return false
}
