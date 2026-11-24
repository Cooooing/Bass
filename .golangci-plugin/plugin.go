// Package plugin 提供 Bass 项目级 GolangCI-Lint 插件。
package plugin

import (
	"github.com/cooooing/bass-lint-plugin/internal/bassdep"
	"github.com/cooooing/bass-lint-plugin/internal/bassenum"
	"github.com/cooooing/bass-lint-plugin/internal/bassschema"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("bassdep", newPlugin(bassdep.Analyzer))
	register.Plugin("bassenum", newPlugin(bassenum.Analyzer))
	register.Plugin("bassschema", newPlugin(bassschema.Analyzer))
}

// lintPlugin 让每个规则族独立加载和开关，避免将不相关的架构规则混为同一 linter。
type lintPlugin struct {
	analyzer *analysis.Analyzer
}

func newPlugin(analyzer *analysis.Analyzer) func(any) (register.LinterPlugin, error) {
	return func(any) (register.LinterPlugin, error) {
		return lintPlugin{analyzer: analyzer}, nil
	}
}

func (p lintPlugin) GetLoadMode() string { return register.LoadModeSyntax }

func (p lintPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{p.analyzer}, nil
}
