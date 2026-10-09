// Package plugin 提供 bassarch 定制 linter。
package plugin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("bassarch", func(conf any) (register.LinterPlugin, error) {
		return bassPlugin{}, nil
	})
}

type bassPlugin struct{}

func (bassPlugin) GetLoadMode() string { return register.LoadModeSyntax }

func (bassPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{bassAnalyzer}, nil
}

var bassAnalyzer = &analysis.Analyzer{
	Name: "bassarch",
	Doc:  "Bass internal 分层依赖和包级函数约束",
	Run:  runBassArch,
}

func runBassArch(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		filename := filepath.ToSlash(pass.Fset.Position(f.Pos()).Filename)
		if strings.HasSuffix(filename, "_test.go") {
			continue
		}
		mod, layer := classifyFile(filename)

		if mod != "" && layer != "" {
			for _, imp := range f.Imports {
				pkg, err := strconv.Unquote(imp.Path.Value)
				if err != nil {
					continue
				}
				for _, msg := range depViolations(pkg, mod, layer) {
					pass.Reportf(imp.Pos(), "%s", msg)
				}
			}
		}
		if mod != "" && internalLayer(filename, mod) == "enum" && !isGeneratedFile(filename) {
			validateEnumFile(pass, f)
		}

		if mod != "" && strings.Contains(filename, "/internal/") && !isGeneratedFile(filename) {
			for _, decl := range f.Decls {
				fd, ok := decl.(*ast.FuncDecl)
				if !ok || fd.Recv != nil || isAllowedPackageFunc(fd.Name.Name) {
					continue
				}
				pass.Reportf(fd.Pos(),
					"结构规范: internal 禁止无 receiver 的包级函数: %s；请挂到对应 struct 方法上",
					fd.Name.Name)
			}
		}
	}
	return nil, nil
}

func classifyFile(filename string) (mod, layer string) {
	if i := strings.LastIndex(filename, "/app/"); i >= 0 {
		rest := filename[i+len("/app/"):]
		if j := strings.Index(rest, "/"); j > 0 {
			mod = rest[:j]
		}
	}
	return mod, internalLayer(filename, mod)
}

func depViolations(pkg, mod, layer string) []string {
	first := pkg
	if before, _, ok := strings.Cut(pkg, "/"); ok {
		first = before
	}
	_, isAppMod := modules()[first]
	if !isAppMod || !under(pkg, first, "internal") {
		return nil
	}
	if first != mod {
		return []string{"依赖规范: 禁止跨服务导入其他 app 模块的 internal 包: " + pkg + "；请走 common 通道或 RPC 客户端"}
	}

	targetLayer := internalLayer(pkg, mod)
	if layerAllows(layer, targetLayer) {
		return nil
	}
	return []string{"依赖规范: " + layer + " 仅可依赖允许的 internal 层，禁止依赖 " + pkg}
}

func under(pkg, first, sub string) bool {
	prefix := first + "/" + sub
	return pkg == prefix || strings.HasPrefix(pkg, prefix+"/")
}

func internalLayer(path, mod string) string {
	if i := strings.Index(path, "/app/"+mod+"/"); i >= 0 {
		path = path[i+len("/app/"+mod)+1:]
	}
	path = strings.TrimPrefix(path, mod+"/")
	if !strings.HasPrefix(path, "internal/") {
		return ""
	}
	path = strings.TrimPrefix(path, "internal/")
	switch {
	case path == "enum" || strings.HasPrefix(path, "enum/"):
		return "enum"
	case path == "config" || strings.HasPrefix(path, "config/"):
		return "config"
	case path == "biz/model" || strings.HasPrefix(path, "biz/model/"):
		return "biz/model"
	case path == "biz/base" || strings.HasPrefix(path, "biz/base/"):
		return "biz/base"
	case path == "biz/repo" || strings.HasPrefix(path, "biz/repo/"):
		return "biz/repo"
	case path == "biz/usecase" || strings.HasPrefix(path, "biz/usecase/"):
		return "biz/usecase"
	case path == "biz" || strings.HasPrefix(path, "biz/"):
		return "biz"
	case path == "data" || strings.HasPrefix(path, "data/"):
		return "data"
	case path == "service" || strings.HasPrefix(path, "service/"):
		return "service"
	case path == "server" || strings.HasPrefix(path, "server/"):
		return "server"
	default:
		return "other"
	}
}

func layerAllows(source, target string) bool {
	allowed, ok := allowedLayers[source]
	return ok && allowed[target]
}

var allowedLayers = map[string]map[string]bool{
	"enum":        {"enum": true},
	"config":      {"config": true, "enum": true},
	"biz/model":   {"biz/model": true, "config": true, "enum": true},
	"biz/base":    {"biz/base": true, "config": true, "enum": true},
	"biz/repo":    {"biz/repo": true, "biz/model": true, "config": true, "enum": true},
	"biz/usecase": {"biz/usecase": true, "biz/base": true, "biz/repo": true, "biz/model": true, "config": true, "enum": true},
	"biz":         {"biz": true, "biz/usecase": true, "config": true, "enum": true},
	"data":        {"data": true, "biz/base": true, "biz/repo": true, "biz/model": true, "config": true, "enum": true},
	"service":     {"service": true, "biz/usecase": true, "biz/model": true, "config": true, "enum": true},
	"server":      {"server": true, "biz": true, "biz/usecase": true, "config": true},
}

func validateEnumFile(pass *analysis.Pass, file *ast.File) {
	var enumTypes []*ast.TypeSpec
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok && isStringEnum(typeSpec) {
				enumTypes = append(enumTypes, typeSpec)
			}
		}
	}
	if len(enumTypes) != 1 {
		pass.Reportf(file.Package, "枚举规范: 每个 internal/enum 文件有且仅有一个 `type Xxxx string` 枚举，当前为 %d 个", len(enumTypes))
		return
	}

	typeName := enumTypes[0].Name.Name
	if got, want := strings.TrimSuffix(filepath.Base(pass.Fset.Position(enumTypes[0].Pos()).Filename), ".go"), snakeCase(typeName); got != want {
		pass.Reportf(enumTypes[0].Pos(), "枚举规范: %s 必须位于 %s.go，当前文件为 %s.go", typeName, want, got)
	}
	values := enumConstValues(pass, file, typeName)
	validateEnumMapping(pass, file, typeName, values)
}

func isStringEnum(spec *ast.TypeSpec) bool {
	ident, ok := spec.Type.(*ast.Ident)
	return ok && ident.Name == "string"
}

func snakeCase(value string) string {
	runes := []rune(value)
	var out strings.Builder
	for i, r := range runes {
		if unicode.IsUpper(r) {
			previousLower := i > 0 && unicode.IsLower(runes[i-1])
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if i > 0 && (previousLower || nextLower) {
				out.WriteByte('_')
			}
			out.WriteRune(unicode.ToLower(r))
			continue
		}
		out.WriteRune(r)
	}
	return out.String()
}

func enumConstValues(pass *analysis.Pass, file *ast.File, typeName string) []string {
	var values []string
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		for _, spec := range gen.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || !isTypedEnumValue(value, typeName) {
				continue
			}
			if len(value.Names) != 1 || len(value.Values) != 1 || !isStringLiteral(value.Values[0]) {
				pass.Reportf(value.Pos(), "枚举规范: %s 的每个值必须写为 `Name %s = \"value\" // 注释`", typeName, typeName)
				continue
			}
			if !hasTrailingComment(pass, file, value) {
				pass.Reportf(value.Pos(), "枚举规范: %s.%s 必须在定义同行追加注释", typeName, value.Names[0].Name)
			}
			values = append(values, value.Names[0].Name)
		}
	}
	return values
}

func isTypedEnumValue(spec *ast.ValueSpec, typeName string) bool {
	ident, ok := spec.Type.(*ast.Ident)
	return ok && ident.Name == typeName
}

func isStringLiteral(expr ast.Expr) bool {
	literal, ok := expr.(*ast.BasicLit)
	return ok && literal.Kind == token.STRING
}

func hasTrailingComment(pass *analysis.Pass, file *ast.File, spec *ast.ValueSpec) bool {
	line := pass.Fset.Position(spec.End()).Line
	for _, group := range file.Comments {
		for _, comment := range group.List {
			if pass.Fset.Position(comment.Slash).Line == line && comment.Slash > spec.End() {
				return true
			}
		}
	}
	return false
}

func validateEnumMapping(pass *analysis.Pass, file *ast.File, typeName string, values []string) {
	mapName := typeName + "Map"
	mapping := findEnumMapping(file, mapName)
	if mapping == nil {
		pass.Reportf(file.Package, "枚举规范: %s 必须定义 `var %s = enum.NewMapping[...]`", typeName, mapName)
		return
	}
	entries := mappingEntries(mapping)
	if len(entries) != len(values) {
		pass.Reportf(mapping.Pos(), "枚举规范: %s 的映射数量 %d 必须与枚举值数量 %d 一致", mapName, len(entries), len(values))
	}
	for i, entry := range entries {
		if i >= len(values) {
			break
		}
		key, ok := entry.Key.(*ast.Ident)
		if !ok || key.Name != values[i] {
			pass.Reportf(entry.Pos(), "枚举规范: %s 的映射顺序必须与 %s const 定义一致", mapName, typeName)
		}
	}
	validateProtoMapping(pass, file, mapping, mapName, entries)
}

func findEnumMapping(file *ast.File, mapName string) ast.Expr {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.VAR {
			continue
		}
		for _, spec := range gen.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || len(value.Names) != 1 || value.Names[0].Name != mapName || len(value.Values) != 1 {
				continue
			}
			if isNewMappingCall(value.Values[0]) {
				return value.Values[0]
			}
		}
	}
	return nil
}

func isNewMappingCall(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	return calledName(call.Fun) == "NewMapping" && hasProtoMappingTypeArgs(call.Fun)
}

func hasProtoMappingTypeArgs(expr ast.Expr) bool {
	indices, ok := expr.(*ast.IndexListExpr)
	if !ok || len(indices.Indices) != 2 {
		return false
	}
	_, ok = indices.Indices[1].(*ast.SelectorExpr)
	return ok
}

func calledName(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.SelectorExpr:
		return value.Sel.Name
	case *ast.IndexExpr:
		return calledName(value.X)
	case *ast.IndexListExpr:
		return calledName(value.X)
	default:
		return ""
	}
}

func mappingEntries(mapping ast.Expr) []*ast.KeyValueExpr {
	call, ok := mapping.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return nil
	}
	literal, ok := call.Args[0].(*ast.CompositeLit)
	if !ok {
		return nil
	}
	entries := make([]*ast.KeyValueExpr, 0, len(literal.Elts))
	for _, item := range literal.Elts {
		if entry, ok := item.(*ast.KeyValueExpr); ok {
			entries = append(entries, entry)
		}
	}
	return entries
}

func validateProtoMapping(pass *analysis.Pass, file *ast.File, mapping ast.Expr, mapName string, entries []*ast.KeyValueExpr) {
	protoPath, protoType, ok := protoMappingType(file, mapping)
	if !ok {
		pass.Reportf(mapping.Pos(), "枚举规范: %s 必须使用 `enum.NewMapping[Internal, proto.Enum]`", mapName)
		return
	}
	protoValues, ok := protoEnumValues(protoPath, protoType)
	if !ok {
		pass.Reportf(mapping.Pos(), "枚举规范: 无法读取 %s 的 proto 枚举 %s", protoPath, protoType)
		return
	}
	if len(entries) != len(protoValues) {
		pass.Reportf(mapping.Pos(), "枚举规范: %s 的映射数量 %d 必须与 proto %s 的非 UNSPECIFIED 值数量 %d 一致", mapName, len(entries), protoType, len(protoValues))
		return
	}
	for i, entry := range entries {
		protoName, ok := protoEntryName(entry)
		if !ok || protoName != protoValues[i] {
			pass.Reportf(entry.Pos(), "枚举规范: %s 的 Proto 映射顺序必须与 proto %s 一致", mapName, protoType)
		}
	}
}

func protoMappingType(file *ast.File, mapping ast.Expr) (string, string, bool) {
	call, ok := mapping.(*ast.CallExpr)
	if !ok {
		return "", "", false
	}
	indices, ok := call.Fun.(*ast.IndexListExpr)
	if !ok || len(indices.Indices) != 2 {
		return "", "", false
	}
	selector, ok := indices.Indices[1].(*ast.SelectorExpr)
	if !ok {
		return "", "", false
	}
	alias, ok := selector.X.(*ast.Ident)
	if !ok {
		return "", "", false
	}
	for _, imp := range file.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}
		name := filepath.Base(path)
		if imp.Name != nil {
			name = imp.Name.Name
		}
		if name == alias.Name {
			return path, selector.Sel.Name, true
		}
	}
	return "", "", false
}

func protoEntryName(entry *ast.KeyValueExpr) (string, bool) {
	literal, ok := entry.Value.(*ast.CompositeLit)
	if !ok {
		return "", false
	}
	for _, item := range literal.Elts {
		field, ok := item.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		name, ok := field.Key.(*ast.Ident)
		if !ok || name.Name != "Proto" {
			continue
		}
		selector, ok := field.Value.(*ast.SelectorExpr)
		if ok {
			return selector.Sel.Name, true
		}
	}
	return "", false
}

func protoEnumValues(importPath, typeName string) ([]string, bool) {
	root := repoRoot()
	if root == "" || !strings.HasPrefix(importPath, "common/") {
		return nil, false
	}
	packages, err := parser.ParseDir(token.NewFileSet(), filepath.Join(root, filepath.FromSlash(importPath)), nil, 0)
	if err != nil {
		return nil, false
	}
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			if values := protoConsts(file, typeName); len(values) > 0 {
				return values, true
			}
		}
	}
	return nil, false
}

func protoConsts(file *ast.File, typeName string) []string {
	var values []string
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		for _, spec := range gen.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || !isTypedEnumValue(value, typeName) {
				continue
			}
			for _, name := range value.Names {
				if !isProtoZeroValue(name.Name) {
					values = append(values, name.Name)
				}
			}
		}
	}
	return values
}

func isProtoZeroValue(name string) bool {
	return strings.HasSuffix(name, "_UNSPECIFIED")
}

func isGeneratedFile(filename string) bool {
	return strings.HasSuffix(filename, ".pb.go") || strings.HasSuffix(filename, "wire_gen.go")
}

func isAllowedPackageFunc(name string) bool {
	return name == "main" || name == "newApp" ||
		strings.HasPrefix(name, "New") || strings.HasPrefix(name, "Provide") || strings.HasPrefix(name, "Provider") ||
		strings.HasPrefix(name, "wire") || strings.HasPrefix(name, "Wire")
}

var (
	modulesOnce sync.Once
	modulesMap  map[string]struct{}
)

// modules caches the app directories that define a Go module.
func modules() map[string]struct{} {
	modulesOnce.Do(func() {
		modulesMap = map[string]struct{}{}
		root := repoRoot()
		if root == "" {
			return
		}
		entries, err := os.ReadDir(filepath.Join(root, "app"))
		if err != nil {
			return
		}
		for _, e := range entries {
			if e.IsDir() && fileExists(filepath.Join(root, "app", e.Name(), "go.mod")) {
				modulesMap[e.Name()] = struct{}{}
			}
		}
	})
	return modulesMap
}

// repoRoot finds the configuration anchor from the invocation directory.
func repoRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for i := 0; i < 20 && dir != ""; i++ {
		if fileExists(filepath.Join(dir, ".custom-gcl.yml")) || fileExists(filepath.Join(dir, ".golangci.yml")) {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
	return ""
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
