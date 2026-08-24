package formula

import (
	"fmt"
	"hash/fnv"
)

// LuaFormula 表示一个 Lua 表达式公式。
type LuaFormula struct {
	ID         string
	Expression string
}

// LuaEvalRequest 表示一次公式计算请求。
type LuaEvalRequest struct {
	Formula   LuaFormula
	Variables map[string]float64
}

// LuaPreparedFormula 表示已经预处理的 Lua 公式。
type LuaPreparedFormula struct {
	id         string
	expression string
	cacheKey   string
	chunk      string
}

// LuaPreparedEvalRequest 表示一次预处理公式计算请求。
type LuaPreparedEvalRequest struct {
	Formula   *LuaPreparedFormula
	Variables map[string]float64
}

// LuaPreparedBatchEvalRequest 表示一次预处理公式批量计算请求。
type LuaPreparedBatchEvalRequest struct {
	Formula   *LuaPreparedFormula
	Variables []map[string]float64
}

// CacheKey 生成公式缓存键，同一 ID 的表达式变更会重新编译。
func (f LuaFormula) CacheKey() string {
	hash := fnv.New64a()
	_, _ = hash.Write([]byte(f.Expression))
	return fmt.Sprintf("%s:%x", f.ID, hash.Sum64())
}

// Formula 返回原始 Lua 公式。
func (f *LuaPreparedFormula) Formula() LuaFormula {
	if f == nil {
		return LuaFormula{}
	}
	return LuaFormula{
		ID:         f.id,
		Expression: f.expression,
	}
}
