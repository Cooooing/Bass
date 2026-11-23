// protocontract 包校验 Bass 的 Proto 合约约定。
package protocontract

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/bufbuild/protocompile/ast"
	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
)

// Violation 表示 Proto 源文件中的一处合约规则违规。
type Violation struct {
	Path    string
	Line    int
	Rule    string
	Message string
}

// Check 校验 root 下的全部 .proto 文件，并返回排序后的违规项。
func Check(root string) ([]Violation, error) {
	var violations []Violation
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(path) != ".proto" {
			return nil
		}
		//nolint:gosec // path 由受 root 限制的 WalkDir 生成。
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		file, err := parser.Parse(path, bytes.NewReader(data), reporter.NewHandler(nil))
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		checker := fileChecker{
			root:     root,
			path:     path,
			file:     file,
			messages: map[string]*ast.MessageNode{},
		}
		checker.check()
		violations = append(violations, checker.violations...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(violations, func(i, j int) bool {
		if violations[i].Path != violations[j].Path {
			return violations[i].Path < violations[j].Path
		}
		if violations[i].Line != violations[j].Line {
			return violations[i].Line < violations[j].Line
		}
		return violations[i].Rule < violations[j].Rule
	})
	return violations, nil
}

type fileChecker struct {
	root       string
	path       string
	file       *ast.FileNode
	messages   map[string]*ast.MessageNode
	services   []*ast.ServiceNode
	violations []Violation
}

func (c *fileChecker) check() {
	for _, decl := range c.file.Decls {
		switch value := decl.(type) {
		case *ast.MessageNode:
			c.messages[value.Name.Val] = value
		case *ast.ServiceNode:
			c.services = append(c.services, value)
		}
	}
	if len(c.services) > 1 {
		c.add(c.services[1], ruleServiceCount, len(c.services))
	}
	for _, service := range c.services {
		c.checkService(service)
	}
}

func (c *fileChecker) add(node ast.Node, rule lintRule, args ...any) {
	path, err := filepath.Rel(c.root, c.path)
	if err != nil {
		path = c.path
	}
	c.violations = append(c.violations, Violation{
		Path:    filepath.ToSlash(path),
		Line:    c.file.NodeInfo(node).Start().Line,
		Rule:    rule.code,
		Message: fmt.Sprintf(rule.message, args...),
	})
}

// 普通一元 RPC 独占自己的请求和响应容器，避免无关接口复用通用消息。
func (c *fileChecker) checkService(service *ast.ServiceNode) {
	if strings.HasPrefix(service.Name.Val, "TCC") {
		c.checkTCCService(service)
		return
	}
	if c.isTCCFile() {
		c.add(service, ruleTCCFileService, service.Name.Val)
	}
	for _, decl := range service.Decls {
		rpc, ok := decl.(*ast.RPCNode)
		if !ok || rpc.Input.Stream != nil || rpc.Output.Stream != nil {
			continue
		}
		c.checkUnaryRPC(rpc)
	}
}

func (c *fileChecker) checkUnaryRPC(rpc *ast.RPCNode) {
	input := string(rpc.Input.MessageType.AsIdentifier())
	output := string(rpc.Output.MessageType.AsIdentifier())
	operation, inputKind, inputOK := c.operationType(input)
	outputOperation, outputKind, outputOK := c.operationType(output)
	if !inputOK || inputKind != operationRequest || operation != rpc.Name.Val {
		c.add(rpc.Input.MessageType, ruleRPCRequest, rpc.Name.Val, rpc.Name.Val, input)
	}
	if !outputOK || outputKind != operationResponse || outputOperation != rpc.Name.Val {
		c.add(rpc.Output.MessageType, ruleRPCResponse, rpc.Name.Val, rpc.Name.Val, output)
		return
	}
	if !inputOK || inputKind != operationRequest || operation != rpc.Name.Val {
		return
	}
	container := c.messages[operation]
	if container == nil {
		c.add(rpc.Input.MessageType, ruleRPCContainer, rpc.Name.Val, operation)
		return
	}
	c.checkOperationContainer(container)
}

// 参数容器是接口的边界。嵌套数据结构应定义在 Req 或 Resp 内，
// 参数容器本身只能直接暴露这两个消息。
func (c *fileChecker) checkOperationContainer(container *ast.MessageNode) {
	found := map[string]bool{}
	for _, decl := range container.Decls {
		nested, ok := decl.(*ast.MessageNode)
		if !ok {
			if _, empty := decl.(*ast.EmptyDeclNode); !empty {
				c.add(decl, ruleDirectMember, container.Name.Val)
			}
			continue
		}
		found[nested.Name.Val] = true
		if nested.Name.Val != operationRequest && nested.Name.Val != operationResponse {
			c.add(nested, ruleExtraMember, container.Name.Val, nested.Name.Val)
		}
	}
	for _, required := range []string{operationRequest, operationResponse} {
		if !found[required] {
			c.add(container, ruleMissingMember, container.Name.Val, required)
		}
	}
}

// TCC 的 Try、Confirm、Cancel 共享同一个业务参数，
// 因此其命名规则有意区别于普通 RPC 按方法名匹配参数的规则。
func (c *fileChecker) checkTCCService(service *ast.ServiceNode) {
	base := strings.TrimSuffix(filepath.Base(c.path), ".proto")
	serviceCore, serviceOK := c.tccServiceCore(service.Name.Val)
	if !c.isTCCFile() {
		c.add(service, ruleTCCFileName, service.Name.Val)
	} else if !serviceOK || base != "tcc_"+c.snakeCase(serviceCore) {
		c.add(service, ruleTCCFileMatch, service.Name.Val, c.snakeCase(serviceCore))
	}
	if !serviceOK {
		c.add(service, ruleTCCServiceName, service.Name.Val)
		return
	}
	methods := map[string]*ast.RPCNode{}
	for _, decl := range service.Decls {
		rpc, ok := decl.(*ast.RPCNode)
		if !ok {
			continue
		}
		methods[rpc.Name.Val] = rpc
		if rpc.Name.Val != "Try" && rpc.Name.Val != "Confirm" && rpc.Name.Val != "Cancel" {
			c.add(rpc, ruleTCCMethodName, service.Name.Val, rpc.Name.Val)
		}
		if rpc.Input.Stream != nil || rpc.Output.Stream != nil {
			c.add(rpc, ruleTCCStreaming, rpc.Name.Val)
			continue
		}
		operation, inputKind, inputOK := c.operationType(string(rpc.Input.MessageType.AsIdentifier()))
		outputOperation, outputKind, outputOK := c.operationType(string(rpc.Output.MessageType.AsIdentifier()))
		if !inputOK || inputKind != operationRequest || !outputOK || outputKind != operationResponse || operation != outputOperation {
			c.add(rpc, ruleTCCParameter, rpc.Name.Val)
			continue
		}
		if container := c.messages[operation]; container == nil {
			c.add(rpc, ruleTCCContainer, rpc.Name.Val, operation)
		} else {
			c.checkOperationContainer(container)
		}
		expectedOperation := "TCC" + serviceCore
		if operation != expectedOperation {
			c.add(rpc, ruleTCCParameterID, expectedOperation, operation)
		}
	}
	for _, required := range []string{"Try", "Confirm", "Cancel"} {
		if methods[required] == nil {
			c.add(service, ruleTCCMethodCount, service.Name.Val, required)
		}
	}
}

func (c *fileChecker) operationType(value string) (string, string, bool) {
	operation, kind, ok := strings.Cut(value, ".")
	return operation, kind, ok && operation != "" && (kind == operationRequest || kind == operationResponse) && !strings.Contains(kind, ".")
}

func (c *fileChecker) isTCCFile() bool { return strings.HasPrefix(filepath.Base(c.path), "tcc_") }

func (c *fileChecker) tccServiceCore(service string) (string, bool) {
	if !strings.HasPrefix(service, "TCC") || !strings.HasSuffix(service, "Service") {
		return "", false
	}
	core := strings.TrimSuffix(strings.TrimPrefix(service, "TCC"), "Service")
	return core, core != ""
}

func (c *fileChecker) snakeCase(value string) string {
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
