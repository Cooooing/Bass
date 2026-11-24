// Package bassschema 校验 Ent schema 的文件结构与业务说明。
package bassschema

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/cooooing/bass-lint-plugin/internal/project"
	"github.com/cooooing/bass-lint-plugin/internal/rules"
	"golang.org/x/tools/go/analysis"
)

const (
	methodFields = "Fields"
)

// Analyzer 执行 Ent schema 的文件与注释规范检查。
var Analyzer = &analysis.Analyzer{
	Name: "bassschema",
	Doc:  "Bass Ent schema 文件结构和业务说明约束",
	Run: func(pass *analysis.Pass) (any, error) {
		checker := checker{pass: pass}
		checker.check()
		return nil, nil
	},
}

type checker struct {
	pass *analysis.Pass
}

type entity struct {
	name string
	spec *ast.TypeSpec
	doc  *ast.CommentGroup
}

func (c checker) check() {
	for _, astFile := range c.pass.Files {
		file := project.NewFile(c.pass, astFile)
		if file.IsAppSource() && file.HasPath("/internal/data/schema/") {
			c.checkFile(file)
		}
	}
}

func (c checker) checkFile(file project.File) {
	entities := c.entities(file)
	for index, item := range entities {
		if index > 0 {
			rules.SchemaSingle.Reportf(c.pass, item.spec.Pos(), item.name)
		}
		c.checkEntity(file, item)
	}
}

func (c checker) entities(file project.File) []entity {
	entities := make([]entity, 0, 1)
	for _, declaration := range file.AST.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.TYPE {
			continue
		}
		for _, spec := range generic.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok && c.isEntity(typeSpec) {
				entities = append(entities, entity{name: typeSpec.Name.Name, spec: typeSpec, doc: generic.Doc})
			}
		}
	}
	return entities
}

func (c checker) isEntity(spec *ast.TypeSpec) bool {
	structure, ok := spec.Type.(*ast.StructType)
	if !ok {
		return false
	}
	for _, field := range structure.Fields.List {
		selector, ok := field.Type.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "Schema" {
			continue
		}
		if packageName, ok := selector.X.(*ast.Ident); ok && packageName.Name == "ent" {
			return true
		}
	}
	return false
}

func (c checker) checkEntity(file project.File, item entity) {
	c.checkFilename(file, item)
	c.checkEntityComment(item)
	methods := c.entityMethods(file, item.name)
	fields := methods[methodFields]
	if fields == nil {
		if !c.allowsMixinOnlyEntity(file, item) {
			rules.SchemaFields.Reportf(c.pass, item.spec.Pos(), item.name)
		}
		return
	}
	c.checkFields(item, fields)
}

// allowsMixinOnlyEntity 仅豁免字段完全由 DTM 屏障 mixin 提供的两张基础设施表。
func (c checker) allowsMixinOnlyEntity(file project.File, item entity) bool {
	if item.name != "Barrier" {
		return false
	}
	return file.HasPath(
		"/app/content/internal/data/schema/barrier.go",
		"/app/economy/internal/data/schema/barrier.go",
	)
}

func (c checker) checkFilename(file project.File, item entity) {
	got := strings.TrimSuffix(filepath.Base(file.Path), ".go")
	want := c.snakeCase(item.name)
	if got != want {
		rules.SchemaFilename.Reportf(c.pass, item.spec.Pos(), item.name, want, got)
	}
}

func (c checker) checkEntityComment(item entity) {
	if !c.hasChineseDoc(item.doc) {
		rules.SchemaComment.Reportf(c.pass, item.spec.Pos(), "实体", item.name)
	}
}

func (c checker) entityMethods(file project.File, entityName string) map[string]*ast.FuncDecl {
	methods := map[string]*ast.FuncDecl{}
	for _, declaration := range file.AST.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Recv == nil || len(function.Recv.List) != 1 {
			continue
		}
		if c.receiverName(function.Recv.List[0].Type) == entityName {
			methods[function.Name.Name] = function
		}
	}
	return methods
}

func (c checker) receiverName(expression ast.Expr) string {
	switch value := expression.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.StarExpr:
		return c.receiverName(value.X)
	default:
		return ""
	}
}

func (c checker) checkFields(item entity, method *ast.FuncDecl) {
	for _, call := range c.collectionCalls(method) {
		if c.calledPackage(call) != "field" {
			continue
		}
		name := c.fieldName(call)
		if name == "id" {
			continue
		}
		if !c.hasChineseCommentCall(call) {
			rules.SchemaField.Reportf(c.pass, call.Pos(), item.name, name)
		}
	}
}

func (c checker) collectionCalls(method *ast.FuncDecl) []*ast.CallExpr {
	if method.Body == nil {
		return nil
	}
	calls := make([]*ast.CallExpr, 0)
	ast.Inspect(method.Body, func(node ast.Node) bool {
		literal, ok := node.(*ast.CompositeLit)
		if !ok {
			return true
		}
		for _, element := range literal.Elts {
			if call, ok := element.(*ast.CallExpr); ok {
				calls = append(calls, call)
			}
		}
		return true
	})
	return calls
}

func (c checker) calledPackage(call *ast.CallExpr) string {
	expression := call.Fun
	for {
		switch value := expression.(type) {
		case *ast.SelectorExpr:
			if ident, ok := value.X.(*ast.Ident); ok {
				return ident.Name
			}
			expression = value.X
		case *ast.CallExpr:
			expression = value.Fun
		default:
			return ""
		}
	}
}

func (c checker) fieldName(call *ast.CallExpr) string {
	constructor := c.constructorCall(call, "field")
	if constructor == nil || len(constructor.Args) == 0 {
		return "字段"
	}
	literal, ok := constructor.Args[0].(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		return "动态字段"
	}
	return strings.Trim(literal.Value, "\"")
}

// constructorCall 沿链式调用回溯到指定包的字段、边或索引构造入口。
func (c checker) constructorCall(call *ast.CallExpr, packageName string) *ast.CallExpr {
	for expression := ast.Expr(call); expression != nil; {
		nested, ok := expression.(*ast.CallExpr)
		if !ok {
			return nil
		}
		selector, ok := nested.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil
		}
		if ident, ok := selector.X.(*ast.Ident); ok && ident.Name == packageName {
			return nested
		}
		inner, ok := selector.X.(*ast.CallExpr)
		if !ok {
			return nil
		}
		expression = inner
	}
	return nil
}

func (c checker) hasChineseCommentCall(call *ast.CallExpr) bool {
	matched := false
	ast.Inspect(call, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "Comment" {
			return true
		}
		parent, ok := c.parentCall(call, selector)
		if !ok || len(parent.Args) != 1 {
			return true
		}
		literal, ok := parent.Args[0].(*ast.BasicLit)
		if ok && literal.Kind == token.STRING && c.hasChinese(literal.Value) {
			matched = true
		}
		return true
	})
	return matched
}

func (c checker) parentCall(root ast.Node, selector *ast.SelectorExpr) (*ast.CallExpr, bool) {
	var parent *ast.CallExpr
	ast.Inspect(root, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if ok && call.Fun == selector {
			parent = call
			return false
		}
		return parent == nil
	})
	return parent, parent != nil
}

func (c checker) hasChineseDoc(group *ast.CommentGroup) bool {
	if group == nil {
		return false
	}
	return c.hasChinese(group.Text())
}

func (c checker) hasChinese(value string) bool {
	for _, runeValue := range value {
		if unicode.Is(unicode.Han, runeValue) {
			return true
		}
	}
	return false
}

func (c checker) snakeCase(value string) string {
	runes := []rune(value)
	var result strings.Builder
	for index, runeValue := range runes {
		if unicode.IsUpper(runeValue) {
			previousLower := index > 0 && unicode.IsLower(runes[index-1])
			nextLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
			if index > 0 && (previousLower || nextLower) {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(runeValue))
			continue
		}
		result.WriteRune(runeValue)
	}
	return result.String()
}
