// Package bassenum 校验服务内部 Go enum 与 Proto enum 的映射契约。
package bassenum

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/cooooing/bass-lint-plugin/internal/project"
	"github.com/cooooing/bass-lint-plugin/internal/rules"
	"golang.org/x/tools/go/analysis"
)

// Analyzer 执行内部 enum 的文件、值与 Mapping 契约检查。
var Analyzer = &analysis.Analyzer{
	Name: "bassenum",
	Doc:  "Bass 内部枚举与 Proto 映射约束",
	Run: func(pass *analysis.Pass) (any, error) {
		checker := checker{pass: pass}
		checker.check()
		return nil, nil
	},
}

type checker struct {
	pass *analysis.Pass
}

type enumValue struct {
	name    string
	literal string
	pos     token.Pos
}

func (c checker) check() {
	for _, astFile := range c.pass.Files {
		file := project.NewFile(c.pass, astFile)
		if file.IsAppSource() && file.IsLayer(project.LayerEnum) {
			c.checkFile(file)
		}
	}
}

func (c checker) checkFile(file project.File) {
	types := c.enumTypes(file)
	if len(types) != 1 {
		rules.EnumSingle.Reportf(c.pass, file.AST.Package, len(types))
		return
	}
	typeName := types[0].Name.Name
	c.checkFilename(file, types[0], typeName)
	values := c.enumValues(file, typeName)
	c.checkMapping(file, typeName, values)
}

func (c checker) enumTypes(file project.File) []*ast.TypeSpec {
	var types []*ast.TypeSpec
	for _, declaration := range file.AST.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.TYPE {
			continue
		}
		for _, spec := range generic.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok && c.isStringEnum(typeSpec) {
				types = append(types, typeSpec)
			}
		}
	}
	return types
}

func (c checker) isStringEnum(spec *ast.TypeSpec) bool {
	ident, ok := spec.Type.(*ast.Ident)
	return ok && ident.Name == "string"
}

func (c checker) checkFilename(file project.File, typeSpec *ast.TypeSpec, typeName string) {
	got := strings.TrimSuffix(filepath.Base(file.Path), ".go")
	want := c.snakeCase(typeName)
	if got != want {
		rules.EnumFilename.Reportf(c.pass, typeSpec.Pos(), typeName, want, got)
	}
}

func (c checker) enumValues(file project.File, typeName string) []enumValue {
	values := make([]enumValue, 0)
	seen := map[string]enumValue{}
	for _, declaration := range file.AST.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.CONST {
			continue
		}
		for _, spec := range generic.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || !c.isTypedValue(value, typeName) {
				continue
			}
			item, valid := c.enumValue(file, value, typeName)
			if !valid {
				continue
			}
			if prior, duplicated := seen[item.literal]; duplicated {
				rules.EnumDuplicate.Reportf(c.pass, item.pos, typeName, strconv.Quote(item.literal), prior.name)
			} else {
				seen[item.literal] = item
			}
			values = append(values, item)
		}
	}
	return values
}

func (c checker) isTypedValue(value *ast.ValueSpec, typeName string) bool {
	ident, ok := value.Type.(*ast.Ident)
	return ok && ident.Name == typeName
}

func (c checker) enumValue(file project.File, value *ast.ValueSpec, typeName string) (enumValue, bool) {
	if len(value.Names) != 1 || len(value.Values) != 1 {
		rules.EnumValue.Reportf(c.pass, value.Pos(), typeName, typeName)
		return enumValue{}, false
	}
	literal, ok := value.Values[0].(*ast.BasicLit)
	if !ok || literal.Kind != token.STRING {
		rules.EnumValue.Reportf(c.pass, value.Pos(), typeName, typeName)
		return enumValue{}, false
	}
	if !c.hasChineseTrailingComment(file, value) {
		rules.EnumComment.Reportf(c.pass, value.Pos(), typeName, value.Names[0].Name)
	}
	decoded, err := strconv.Unquote(literal.Value)
	if err != nil {
		return enumValue{}, false
	}
	return enumValue{name: value.Names[0].Name, literal: decoded, pos: value.Pos()}, true
}

func (c checker) hasChineseTrailingComment(file project.File, value *ast.ValueSpec) bool {
	line := c.pass.Fset.Position(value.End()).Line
	for _, group := range file.AST.Comments {
		for _, comment := range group.List {
			if c.pass.Fset.Position(comment.Slash).Line == line && comment.Slash > value.End() && c.hasChinese(comment.Text) {
				return true
			}
		}
	}
	return false
}

func (c checker) checkMapping(file project.File, typeName string, values []enumValue) {
	mapName := typeName + "Map"
	mapping := c.findMapping(file, mapName)
	if mapping == nil {
		rules.EnumMapping.Reportf(c.pass, file.AST.Package, typeName, mapName)
		return
	}
	entries := c.mappingEntries(mapping)
	if len(entries) != len(values) {
		rules.EnumMappingSize.Reportf(c.pass, mapping.Pos(), mapName, len(entries), len(values))
	}
	for index, entry := range entries {
		if index >= len(values) {
			break
		}
		key, ok := entry.Key.(*ast.Ident)
		if !ok || key.Name != values[index].name {
			rules.EnumMappingKey.Reportf(c.pass, entry.Pos(), mapName, typeName)
		}
	}
	c.checkProtoMapping(file, mapping, mapName, entries)
}

func (c checker) findMapping(file project.File, mapName string) ast.Expr {
	for _, declaration := range file.AST.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.VAR {
			continue
		}
		for _, spec := range generic.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || len(value.Names) != 1 || len(value.Values) != 1 || value.Names[0].Name != mapName {
				continue
			}
			if c.isNewMapping(value.Values[0]) {
				return value.Values[0]
			}
		}
	}
	return nil
}

func (c checker) isNewMapping(expression ast.Expr) bool {
	call, ok := expression.(*ast.CallExpr)
	if !ok {
		return false
	}
	indices, ok := call.Fun.(*ast.IndexListExpr)
	return ok && len(indices.Indices) == 2 && c.calledName(indices.X) == "NewMapping"
}

func (c checker) calledName(expression ast.Expr) string {
	switch value := expression.(type) {
	case *ast.SelectorExpr:
		return value.Sel.Name
	case *ast.IndexExpr:
		return c.calledName(value.X)
	case *ast.IndexListExpr:
		return c.calledName(value.X)
	default:
		return ""
	}
}

func (c checker) mappingEntries(mapping ast.Expr) []*ast.KeyValueExpr {
	call, ok := mapping.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return nil
	}
	literal, ok := call.Args[0].(*ast.CompositeLit)
	if !ok {
		return nil
	}
	entries := make([]*ast.KeyValueExpr, 0, len(literal.Elts))
	for _, element := range literal.Elts {
		if entry, ok := element.(*ast.KeyValueExpr); ok {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (c checker) checkProtoMapping(file project.File, mapping ast.Expr, mapName string, entries []*ast.KeyValueExpr) {
	path, typeName, ok := c.protoMappingType(file, mapping)
	if !ok {
		rules.EnumMapping.Reportf(c.pass, mapping.Pos(), "内部枚举", mapName)
		return
	}
	values, ok := c.protoValues(path, typeName)
	if !ok || len(values) != len(entries) {
		rules.EnumProto.Reportf(c.pass, mapping.Pos(), mapName, typeName)
		return
	}
	for index, entry := range entries {
		if name, valid := c.protoEntryName(entry); !valid || name != values[index] {
			rules.EnumProto.Reportf(c.pass, entry.Pos(), mapName, typeName)
		}
	}
}

func (c checker) protoMappingType(file project.File, mapping ast.Expr) (string, string, bool) {
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
	for _, spec := range file.AST.Imports {
		path := file.ImportPath(spec)
		if spec.Name != nil && spec.Name.Name == alias.Name {
			return path, selector.Sel.Name, true
		}
		if spec.Name == nil && c.importPackageName(path) == alias.Name {
			return path, selector.Sel.Name, true
		}
	}
	return "", "", false
}

func (c checker) importPackageName(importPath string) string {
	directory := filepath.Join(project.RepositoryRoot(), filepath.FromSlash(importPath))
	entries, err := os.ReadDir(directory)
	if err != nil {
		return filepath.Base(importPath)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), filepath.Join(directory, entry.Name()), nil, parser.PackageClauseOnly)
		if parseErr == nil {
			return file.Name.Name
		}
	}
	return filepath.Base(importPath)
}

func (c checker) protoValues(importPath, typeName string) ([]string, bool) {
	if !strings.HasPrefix(importPath, "common/") {
		return nil, false
	}
	directory := filepath.Join(project.RepositoryRoot(), filepath.FromSlash(importPath))
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, false
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), filepath.Join(directory, entry.Name()), nil, 0)
		if parseErr != nil {
			continue
		}
		if values := c.protoConstants(file, typeName); len(values) > 0 {
			return values, true
		}
	}
	return nil, false
}

func (c checker) protoConstants(file *ast.File, typeName string) []string {
	values := make([]string, 0)
	for _, declaration := range file.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.CONST {
			continue
		}
		for _, spec := range generic.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok || !c.isTypedValue(value, typeName) {
				continue
			}
			for _, name := range value.Names {
				if !strings.HasSuffix(name.Name, "_UNSPECIFIED") {
					values = append(values, name.Name)
				}
			}
		}
	}
	return values
}

func (c checker) protoEntryName(entry *ast.KeyValueExpr) (string, bool) {
	literal, ok := entry.Value.(*ast.CompositeLit)
	if !ok {
		return "", false
	}
	for _, item := range literal.Elts {
		field, ok := item.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := field.Key.(*ast.Ident)
		if !ok || key.Name != "Proto" {
			continue
		}
		selector, ok := field.Value.(*ast.SelectorExpr)
		if ok {
			return selector.Sel.Name, true
		}
	}
	return "", false
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

func (c checker) hasChinese(value string) bool {
	for _, runeValue := range value {
		if unicode.Is(unicode.Han, runeValue) {
			return true
		}
	}
	return false
}
