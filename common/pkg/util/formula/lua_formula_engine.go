package formula

import (
	"context"
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

const (
	luaFormulaCacheGlobal         = "__bass_lua_formula_cache"
	luaFormulaDefaultVariableSize = 16
)

// LuaFormulaEngine 是可复用的 Lua 公式计算引擎。
type LuaFormulaEngine struct {
	pool sync.Pool
}

// NewLuaFormulaEngine 创建 Lua 公式计算引擎。
func NewLuaFormulaEngine() *LuaFormulaEngine {
	engine := &LuaFormulaEngine{}
	engine.pool.New = func() any {
		state := lua.NewState(lua.Options{SkipOpenLibs: true})
		lua.OpenMath(state)
		return &luaFormulaRuntime{
			state:         state,
			variables:     state.NewTable(),
			variableNames: make([]string, 0, luaFormulaDefaultVariableSize),
		}
	}
	return engine
}

// Evaluate 执行 Lua 表达式，变量通过 ctx 表传入。
func (e *LuaFormulaEngine) Evaluate(ctx context.Context, req LuaEvalRequest) (float64, error) {
	if req.Formula.ID == "" {
		return 0, fmt.Errorf("lua formula id empty")
	}
	if req.Formula.Expression == "" {
		return 0, fmt.Errorf("lua formula expression empty")
	}
	return e.EvaluatePrepared(ctx, LuaPreparedEvalRequest{
		Formula: &LuaPreparedFormula{
			id:         req.Formula.ID,
			expression: req.Formula.Expression,
			cacheKey:   req.Formula.CacheKey(),
			chunk:      "return function(ctx) return " + req.Formula.Expression + " end",
		},
		Variables: req.Variables,
	})
}

// Prepare 预处理 Lua 表达式，提前生成缓存键并校验语法。
func (e *LuaFormulaEngine) Prepare(ctx context.Context, formula LuaFormula) (*LuaPreparedFormula, error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	if formula.ID == "" {
		return nil, fmt.Errorf("lua formula id empty")
	}
	if formula.Expression == "" {
		return nil, fmt.Errorf("lua formula expression empty")
	}

	prepared := &LuaPreparedFormula{
		id:         formula.ID,
		expression: formula.Expression,
		cacheKey:   formula.CacheKey(),
		chunk:      "return function(ctx) return " + formula.Expression + " end",
	}
	runtime := e.pool.Get().(*luaFormulaRuntime)
	baseTop := runtime.state.GetTop()
	defer func() {
		runtime.ClearVariables()
		runtime.state.SetTop(baseTop)
		e.pool.Put(runtime)
	}()

	value := runtime.state.GetGlobal(luaFormulaCacheGlobal)
	cacheTable, ok := value.(*lua.LTable)
	if !ok {
		cacheTable = runtime.state.NewTable()
		runtime.state.SetGlobal(luaFormulaCacheGlobal, cacheTable)
	}
	if function, ok := cacheTable.RawGetString(prepared.cacheKey).(*lua.LFunction); ok && function != nil {
		return prepared, nil
	}
	compiled, err := runtime.state.LoadString(prepared.chunk)
	if err != nil {
		return nil, fmt.Errorf("lua formula compile failed: %w", err)
	}
	runtime.state.Push(compiled)
	if err := runtime.state.PCall(0, 1, nil); err != nil {
		return nil, fmt.Errorf("lua formula load failed: %w", err)
	}
	loaded := runtime.state.Get(-1)
	runtime.state.Pop(1)
	function, ok := loaded.(*lua.LFunction)
	if !ok {
		return nil, fmt.Errorf("lua formula compiled result invalid")
	}
	cacheTable.RawSetString(prepared.cacheKey, function)
	return prepared, nil
}

// EvaluatePrepared 执行预处理后的 Lua 表达式，适合游戏结算热路径使用。
func (e *LuaFormulaEngine) EvaluatePrepared(ctx context.Context, req LuaPreparedEvalRequest) (float64, error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}
	}
	if req.Formula == nil {
		return 0, fmt.Errorf("lua prepared formula empty")
	}

	runtime := e.pool.Get().(*luaFormulaRuntime)
	baseTop := runtime.state.GetTop()
	defer func() {
		runtime.ClearVariables()
		runtime.state.SetTop(baseTop)
		e.pool.Put(runtime)
	}()

	value := runtime.state.GetGlobal(luaFormulaCacheGlobal)
	cacheTable, ok := value.(*lua.LTable)
	if !ok {
		cacheTable = runtime.state.NewTable()
		runtime.state.SetGlobal(luaFormulaCacheGlobal, cacheTable)
	}
	function, ok := cacheTable.RawGetString(req.Formula.cacheKey).(*lua.LFunction)
	if !ok || function == nil {
		compiled, err := runtime.state.LoadString(req.Formula.chunk)
		if err != nil {
			return 0, fmt.Errorf("lua formula compile failed: %w", err)
		}
		runtime.state.Push(compiled)
		if err := runtime.state.PCall(0, 1, nil); err != nil {
			return 0, fmt.Errorf("lua formula load failed: %w", err)
		}
		loaded := runtime.state.Get(-1)
		runtime.state.Pop(1)
		function, ok = loaded.(*lua.LFunction)
		if !ok {
			return 0, fmt.Errorf("lua formula compiled result invalid")
		}
		cacheTable.RawSetString(req.Formula.cacheKey, function)
	}

	for name, value := range req.Variables {
		runtime.variableNames = append(runtime.variableNames, name)
		runtime.variables.RawSetString(name, lua.LNumber(value))
	}
	runtime.state.Push(function)
	runtime.state.Push(runtime.variables)
	if err := runtime.state.PCall(1, 1, nil); err != nil {
		return 0, fmt.Errorf("lua formula eval failed: %w", err)
	}

	result := runtime.state.Get(-1)
	number, ok := result.(lua.LNumber)
	if !ok {
		return 0, fmt.Errorf("lua formula result must be number")
	}
	return float64(number), nil
}

// EvaluatePreparedBatch 批量执行预处理后的 Lua 表达式，适合同一次游戏结算内重复计算。
func (e *LuaFormulaEngine) EvaluatePreparedBatch(ctx context.Context, req LuaPreparedBatchEvalRequest) ([]float64, error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	if req.Formula == nil {
		return nil, fmt.Errorf("lua prepared formula empty")
	}

	runtime := e.pool.Get().(*luaFormulaRuntime)
	baseTop := runtime.state.GetTop()
	defer func() {
		runtime.ClearVariables()
		runtime.state.SetTop(baseTop)
		e.pool.Put(runtime)
	}()

	value := runtime.state.GetGlobal(luaFormulaCacheGlobal)
	cacheTable, ok := value.(*lua.LTable)
	if !ok {
		cacheTable = runtime.state.NewTable()
		runtime.state.SetGlobal(luaFormulaCacheGlobal, cacheTable)
	}
	function, ok := cacheTable.RawGetString(req.Formula.cacheKey).(*lua.LFunction)
	if !ok || function == nil {
		compiled, err := runtime.state.LoadString(req.Formula.chunk)
		if err != nil {
			return nil, fmt.Errorf("lua formula compile failed: %w", err)
		}
		runtime.state.Push(compiled)
		if err := runtime.state.PCall(0, 1, nil); err != nil {
			return nil, fmt.Errorf("lua formula load failed: %w", err)
		}
		loaded := runtime.state.Get(-1)
		runtime.state.Pop(1)
		function, ok = loaded.(*lua.LFunction)
		if !ok {
			return nil, fmt.Errorf("lua formula compiled result invalid")
		}
		cacheTable.RawSetString(req.Formula.cacheKey, function)
	}

	results := make([]float64, 0, len(req.Variables))
	for _, variables := range req.Variables {
		runtime.ClearVariables()
		for name, value := range variables {
			runtime.variableNames = append(runtime.variableNames, name)
			runtime.variables.RawSetString(name, lua.LNumber(value))
		}
		runtime.state.Push(function)
		runtime.state.Push(runtime.variables)
		if err := runtime.state.PCall(1, 1, nil); err != nil {
			return nil, fmt.Errorf("lua formula eval failed: %w", err)
		}

		result := runtime.state.Get(-1)
		runtime.state.Pop(1)
		number, ok := result.(lua.LNumber)
		if !ok {
			return nil, fmt.Errorf("lua formula result must be number")
		}
		results = append(results, float64(number))
	}
	return results, nil
}
