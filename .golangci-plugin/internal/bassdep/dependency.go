// Package bassdep 校验 Bass 服务模块的依赖、配置、并发和基础设施边界。
package bassdep

import (
	"go/ast"
	"strings"

	"github.com/cooooing/bass-lint-plugin/internal/project"
	"github.com/cooooing/bass-lint-plugin/internal/rules"
	"golang.org/x/tools/go/analysis"
)

const databaseSQLPackage = "database/sql"

// Analyzer 执行模块分层和高风险运行时边界检查。
var Analyzer = &analysis.Analyzer{
	Name: "bassdep",
	Doc:  "Bass 服务依赖、配置、并发与基础设施边界约束",
	Run: func(pass *analysis.Pass) (any, error) {
		checker := checker{pass: pass}
		checker.check()
		return nil, nil
	},
}

type checker struct {
	pass *analysis.Pass
}

func (c checker) check() {
	for _, astFile := range c.pass.Files {
		file := project.NewFile(c.pass, astFile)
		if !file.IsAppSource() {
			continue
		}
		c.checkImports(file)
		c.checkCalls(file)
	}
}

func (c checker) checkImports(file project.File) {
	for _, spec := range file.AST.Imports {
		path := file.ImportPath(spec)
		c.checkModuleImport(file, spec, path)
		c.checkInfrastructureImport(file, spec, path)
	}
	c.checkBoundaryTypes(file)
}

// checkBoundaryTypes 防止接口或服务签名把数据实现类型泄漏到上层。
func (c checker) checkBoundaryTypes(file project.File) {
	if !file.IsLayer(project.LayerBizRepo) && !file.IsLayer(project.LayerService) {
		return
	}
	for _, declaration := range file.AST.Decls {
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			c.checkFunctionTypes(file, value)
		case *ast.GenDecl:
			c.checkDeclarationTypes(file, value)
		}
	}
}

func (c checker) checkFunctionTypes(file project.File, function *ast.FuncDecl) {
	if function.Type.Params != nil {
		for _, field := range function.Type.Params.List {
			c.checkBoundaryType(file, field.Type)
		}
	}
	if function.Type.Results != nil {
		for _, field := range function.Type.Results.List {
			c.checkBoundaryType(file, field.Type)
		}
	}
}

func (c checker) checkDeclarationTypes(file project.File, declaration *ast.GenDecl) {
	for _, spec := range declaration.Specs {
		switch value := spec.(type) {
		case *ast.TypeSpec:
			c.checkBoundaryType(file, value.Type)
		case *ast.ValueSpec:
			c.checkBoundaryType(file, value.Type)
		}
	}
}

func (c checker) checkBoundaryType(file project.File, expression ast.Expr) {
	if expression == nil {
		return
	}
	ast.Inspect(expression, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		alias, ok := selector.X.(*ast.Ident)
		if !ok {
			return true
		}
		for _, spec := range file.AST.Imports {
			path := file.ImportPath(spec)
			if file.ImportAlias(spec) != alias.Name || !c.isRepositoryImplementation(path, file.Module) {
				continue
			}
			if file.IsLayer(project.LayerBizRepo) {
				rules.RepoLeak.Reportf(c.pass, selector.Pos(), alias.Name+"."+selector.Sel.Name)
			} else {
				rules.APIType.Reportf(c.pass, selector.Pos(), alias.Name+"."+selector.Sel.Name)
			}
		}
		return true
	})
}

func (c checker) checkModuleImport(file project.File, spec *ast.ImportSpec, path string) {
	module, isModuleInternal := c.moduleInternal(path)
	if !isModuleInternal {
		return
	}
	if module != file.Module {
		rules.DepCrossModule.Reportf(c.pass, spec.Pos(), path)
		return
	}
	target := c.internalImportLayer(path, module)
	if file.Layer != "" && !c.layerAllowed(file.Layer, target) {
		rules.DepLayer.Reportf(c.pass, spec.Pos(), file.Layer, path, target)
	}
}

func (c checker) checkInfrastructureImport(file project.File, spec *ast.ImportSpec, path string) {
	if file.HasPath("/internal/biz/") && c.isInfrastructure(path) {
		rules.DepBusiness.Reportf(c.pass, spec.Pos(), path)
	}
	if file.IsLayer(project.LayerBizRepo) && c.isRepositoryImplementation(path, file.Module) {
		rules.RepoLeak.Reportf(c.pass, spec.Pos(), path)
	}
	if file.IsLayer(project.LayerService) && c.isRepositoryImplementation(path, file.Module) {
		rules.APIType.Reportf(c.pass, spec.Pos(), path)
	}
	if path == databaseSQLPackage && !file.HasPath("/internal/data/", "/migration/", "/cmd/", "/module/") {
		rules.Database.Reportf(c.pass, spec.Pos(), path)
	}
}

func (c checker) checkCalls(file project.File) {
	for _, declaration := range file.AST.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Body == nil {
			continue
		}
		ast.Inspect(function.Body, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.CallExpr:
				c.checkCall(file, function, value)
			case *ast.GoStmt:
				c.checkGo(file, value)
			}
			return true
		})
	}
}

func (c checker) checkCall(file project.File, function *ast.FuncDecl, call *ast.CallExpr) {
	alias, name, ok := c.selectorCall(call)
	if !ok {
		return
	}
	qualified := alias + "." + name
	if file.UsesImportAlias(alias, "os") && (name == "Getenv" || name == "LookupEnv") && !file.HasPath("/cmd/", "/module/", "/internal/config/") {
		rules.ConfigEnv.Reportf(c.pass, call.Pos(), qualified)
	}
	if file.IsInternal() && file.UsesImportAlias(alias, "context") && (name == "Background" || name == "TODO") && !c.isRootContextAllowed(file, function) {
		rules.ContextRoot.Reportf(c.pass, call.Pos(), qualified)
	}
	if file.UsesImportAlias(alias, "time") && name == "Sleep" && !file.HasPath("/internal/server/", "/internal/data/") {
		rules.TimeSleep.Reportf(c.pass, call.Pos())
	}
	if c.isInfrastructureConstructor(file, alias, name) && !file.HasPath("/cmd/", "/module/", "/internal/data/") {
		rules.DepConstruct.Reportf(c.pass, call.Pos(), qualified)
	}
}

func (c checker) isRootContextAllowed(file project.File, function *ast.FuncDecl) bool {
	return file.HasPath("/internal/server/", "/internal/data/client/") ||
		(file.HasPath("/internal/data/repo/") && strings.HasPrefix(function.Name.Name, "New"))
}

func (c checker) checkGo(file project.File, statement *ast.GoStmt) {
	if !file.HasPath("/internal/server/", "/internal/data/") {
		rules.GoRoutine.Reportf(c.pass, statement.Pos())
	}
	function, ok := statement.Call.Fun.(*ast.FuncLit)
	if !ok || c.hasStopReference(function) {
		return
	}
	rules.GoCancel.Reportf(c.pass, statement.Pos())
}

func (c checker) selectorCall(call *ast.CallExpr) (string, string, bool) {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", "", false
	}
	ident, ok := selector.X.(*ast.Ident)
	if !ok {
		return "", "", false
	}
	return ident.Name, selector.Sel.Name, true
}

func (c checker) hasStopReference(function *ast.FuncLit) bool {
	found := false
	ast.Inspect(function.Body, func(node ast.Node) bool {
		if found {
			return false
		}
		if selector, ok := node.(*ast.SelectorExpr); ok && selector.Sel.Name == "Wait" {
			found = true
			return false
		}
		if ident, ok := node.(*ast.Ident); ok && c.isStopName(ident.Name) {
			found = true
			return false
		}
		return true
	})
	return found
}

func (c checker) isStopName(name string) bool {
	lower := strings.ToLower(name)
	return lower == "ctx" || strings.HasSuffix(lower, "ctx") || strings.Contains(lower, "cancel") ||
		strings.Contains(lower, "done") || strings.Contains(lower, "stop") || strings.Contains(lower, "waitgroup") || lower == "wg"
}

func (c checker) isInfrastructureConstructor(file project.File, alias, name string) bool {
	for _, spec := range file.AST.Imports {
		if file.ImportAlias(spec) != alias {
			continue
		}
		switch file.ImportPath(spec) {
		case databaseSQLPackage:
			return name == "Open" || name == "OpenDB"
		case "github.com/nats-io/nats.go":
			return name == "Connect"
		case "google.golang.org/grpc":
			return name == "Dial" || name == "DialContext" || name == "NewClient"
		case "github.com/hashicorp/consul/api":
			return name == "NewClient"
		}
	}
	return false
}

func (c checker) moduleInternal(path string) (string, bool) {
	module, rest, found := strings.Cut(path, "/")
	if !found {
		return "", false
	}
	if _, exists := project.Modules()[module]; !exists || (rest != "internal" && !strings.HasPrefix(rest, "internal/")) {
		return "", false
	}
	return module, true
}

func (c checker) internalImportLayer(path, module string) string {
	path = strings.TrimPrefix(path, module+"/")
	return project.InternalLayer("/app/"+module+"/"+path, module)
}

func (c checker) layerAllowed(source, target string) bool {
	allowed, exists := map[string]map[string]bool{
		project.LayerEnum:       {project.LayerEnum: true},
		project.LayerConfig:     {project.LayerConfig: true, project.LayerEnum: true},
		project.LayerBizModel:   {project.LayerBizModel: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerBizRepo:    {project.LayerBizRepo: true, project.LayerBizModel: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerBizUsecase: {project.LayerBizUsecase: true, project.LayerBizRepo: true, project.LayerBizModel: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerBiz:        {project.LayerBiz: true, project.LayerBizUsecase: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerData:       {project.LayerData: true, project.LayerBizRepo: true, project.LayerBizModel: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerService:    {project.LayerService: true, project.LayerBizUsecase: true, project.LayerBizModel: true, project.LayerConfig: true, project.LayerEnum: true},
		project.LayerServer:     {project.LayerServer: true, project.LayerBiz: true, project.LayerBizUsecase: true, project.LayerConfig: true},
	}[source]
	return exists && allowed[target]
}

func (c checker) isInfrastructure(path string) bool {
	return path == databaseSQLPackage || path == "entgo.io/ent" || strings.HasPrefix(path, "entgo.io/ent/") ||
		strings.HasPrefix(path, "github.com/nats-io/") || strings.HasPrefix(path, "github.com/hashicorp/consul") ||
		strings.HasPrefix(path, "google.golang.org/grpc") || strings.HasPrefix(path, "github.com/go-kratos/kratos/v3/transport")
}

func (c checker) isRepositoryImplementation(path, module string) bool {
	return path == databaseSQLPackage || path == "entgo.io/ent" || strings.HasPrefix(path, "entgo.io/ent/") ||
		strings.HasPrefix(path, "github.com/jackc/pgx") || strings.HasPrefix(path, "github.com/lib/pq") ||
		strings.HasPrefix(path, "ariga.io/atlas") || strings.HasPrefix(path, module+"/internal/data/")
}
