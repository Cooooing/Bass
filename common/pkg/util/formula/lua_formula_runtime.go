package formula

import lua "github.com/yuin/gopher-lua"

type luaFormulaRuntime struct {
	state         *lua.LState
	variables     *lua.LTable
	variableNames []string
}

func (r *luaFormulaRuntime) ClearVariables() {
	for _, name := range r.variableNames {
		r.variables.RawSetString(name, lua.LNil)
	}
	r.variableNames = r.variableNames[:0]
}
