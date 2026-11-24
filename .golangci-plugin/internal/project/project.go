// Package project 提供三个 lint 规则族共享的仓库路径与 Go AST 上下文。
package project

import (
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/tools/go/analysis"
)

const (
	LayerEnum       = "enum"
	LayerConfig     = "config"
	LayerBizModel   = "biz/model"
	LayerBizRepo    = "biz/repo"
	LayerBizUsecase = "biz/usecase"
	LayerBiz        = "biz"
	LayerData       = "data"
	LayerService    = "service"
	LayerServer     = "server"
)

// File 描述一份被分析的源码及其在服务模块中的位置。
type File struct {
	Pass      *analysis.Pass
	AST       *ast.File
	Path      string
	Module    string
	Layer     string
	Generated bool
}

// NewFile 从 analysis 输入推导服务、internal 层和生成文件状态。
func NewFile(pass *analysis.Pass, file *ast.File) File {
	path := filepath.ToSlash(pass.Fset.Position(file.Pos()).Filename)
	module := moduleName(path)
	return File{
		Pass:      pass,
		AST:       file,
		Path:      path,
		Module:    module,
		Layer:     InternalLayer(path, module),
		Generated: strings.HasSuffix(path, ".pb.go") || strings.HasSuffix(path, "wire_gen.go"),
	}
}

// IsAppSource 只接受 app 服务模块下的手写生产代码。
func (f File) IsAppSource() bool {
	return f.Module != "" && !f.Generated && !strings.HasSuffix(f.Path, "_test.go")
}

// IsInternal 表示文件处于本服务的 internal 目录。
func (f File) IsInternal() bool { return strings.Contains(f.Path, "/internal/") }

// HasPath 判断文件是否位于任一以斜杠分隔的目录片段中。
func (f File) HasPath(paths ...string) bool {
	for _, path := range paths {
		if strings.Contains(f.Path, path) {
			return true
		}
	}
	return false
}

// IsLayer 判断当前文件是否处于指定 internal 层。
func (f File) IsLayer(layer string) bool { return f.Layer == layer }

// ImportPath 返回 import 声明的未加引号路径。
func (f File) ImportPath(spec *ast.ImportSpec) string { return strings.Trim(spec.Path.Value, "\"") }

// ImportAlias 返回 Go 源码中引用 import 时使用的名称。
func (f File) ImportAlias(spec *ast.ImportSpec) string {
	if spec.Name != nil {
		return spec.Name.Name
	}
	return filepath.Base(f.ImportPath(spec))
}

// UsesImportAlias 判断标识符是否指向指定 import 路径。
func (f File) UsesImportAlias(name, importPath string) bool {
	for _, spec := range f.AST.Imports {
		if f.ImportPath(spec) == importPath && f.ImportAlias(spec) == name {
			return true
		}
	}
	return false
}

// Position 返回节点在当前 analysis 文件集中的源码位置。
func (f File) Position(node ast.Node) token.Pos { return node.Pos() }

// InternalLayer 根据路径判断模块内层，未知 internal 目录以 other 表示。
func InternalLayer(path, module string) string {
	if module == "" {
		return ""
	}
	prefix := "/app/" + module + "/internal/"
	index := strings.Index(path, prefix)
	if index < 0 {
		return ""
	}
	path = path[index+len(prefix):]
	switch {
	case path == LayerEnum || strings.HasPrefix(path, LayerEnum+"/"):
		return LayerEnum
	case path == LayerConfig || strings.HasPrefix(path, LayerConfig+"/"):
		return LayerConfig
	case path == LayerBizModel || strings.HasPrefix(path, LayerBizModel+"/"):
		return LayerBizModel
	case path == LayerBizRepo || strings.HasPrefix(path, LayerBizRepo+"/"):
		return LayerBizRepo
	case path == LayerBizUsecase || strings.HasPrefix(path, LayerBizUsecase+"/"):
		return LayerBizUsecase
	case path == LayerBiz || strings.HasPrefix(path, LayerBiz+"/"):
		return LayerBiz
	case path == LayerData || strings.HasPrefix(path, LayerData+"/"):
		return LayerData
	case path == LayerService || strings.HasPrefix(path, LayerService+"/"):
		return LayerService
	case path == LayerServer || strings.HasPrefix(path, LayerServer+"/"):
		return LayerServer
	default:
		return "other"
	}
}

var (
	modulesOnce sync.Once
	modulesMap  map[string]struct{}
	rootOnce    sync.Once
	rootPath    string
)

// Modules 返回 app 下定义了 go.mod 的模块名称集合。
func Modules() map[string]struct{} {
	modulesOnce.Do(func() {
		modulesMap = map[string]struct{}{}
		root := RepositoryRoot()
		entries, err := os.ReadDir(filepath.Join(root, "app"))
		if err != nil {
			return
		}
		for _, entry := range entries {
			if entry.IsDir() && fileExists(filepath.Join(root, "app", entry.Name(), "go.mod")) {
				modulesMap[entry.Name()] = struct{}{}
			}
		}
	})
	return modulesMap
}

// RepositoryRoot 从当前工作目录向上定位 linter 配置锚点。
func RepositoryRoot() string {
	rootOnce.Do(func() {
		directory, err := os.Getwd()
		if err != nil {
			return
		}
		for range 20 {
			if fileExists(filepath.Join(directory, ".custom-gcl.yml")) || fileExists(filepath.Join(directory, ".golangci.yml")) {
				rootPath = directory
				return
			}
			parent := filepath.Dir(directory)
			if parent == directory {
				return
			}
			directory = parent
		}
	})
	return rootPath
}

func moduleName(path string) string {
	_, after, found := strings.Cut(path, "/app/")
	if !found {
		return ""
	}
	name, _, found := strings.Cut(after, "/")
	if !found {
		return ""
	}
	return name
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
