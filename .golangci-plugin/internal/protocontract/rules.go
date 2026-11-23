package protocontract

// lintRule 将规则编号和中文提示模板绑定，便于集中维护全部检查结果。
type lintRule struct {
	code    string
	message string
}

// 所有 Proto 合约规则的编号和提示模板均在此处声明。
var (
	ruleServiceCount = lintRule{"SERVICE_COUNT", "每个 Proto 文件只能定义一个 service，当前定义 %d 个"}

	ruleTCCFileService = lintRule{"TCC_SERVICE", "tcc_xxx.proto 只能定义 TCCXxxService，当前为 %s"}
	ruleTCCFileName    = lintRule{"TCC_FILE", "TCC service %s 必须定义在 tcc_xxx.proto 文件中"}
	ruleTCCFileMatch   = lintRule{"TCC_FILE", "TCC service %s 必须定义在 tcc_%s.proto"}
	ruleTCCServiceName = lintRule{"TCC_SERVICE", "TCC service 必须命名为 TCCXxxService，当前为 %s"}
	ruleTCCMethodName  = lintRule{"TCC_METHOD", "TCC service %s 只能定义 Try、Confirm、Cancel，发现 %s"}
	ruleTCCStreaming   = lintRule{"TCC_METHOD", "TCC rpc %s 不得使用 streaming"}
	ruleTCCMethodCount = lintRule{"TCC_METHOD", "TCC service %s 必须定义 %s rpc"}
	ruleTCCParameter   = lintRule{"TCC_PARAMETER", "TCC rpc %s 必须使用同一参数的 .Req 和 .Resp"}
	ruleTCCContainer   = lintRule{"RPC_CONTAINER", "TCC rpc %s 缺少顶层参数 message %s"}
	ruleTCCParameterID = lintRule{"TCC_PARAMETER", "TCC 参数必须为 %s，当前为 %s"}

	ruleRPCRequest    = lintRule{"RPC_REQUEST", "rpc %s 的请求必须为 %s.Req，当前为 %s"}
	ruleRPCResponse   = lintRule{"RPC_RESPONSE", "rpc %s 的响应必须为 %s.Resp，当前为 %s"}
	ruleRPCContainer  = lintRule{"RPC_CONTAINER", "rpc %s 缺少顶层参数 message %s"}
	ruleDirectMember  = lintRule{"RPC_CONTAINER", "参数 message %s 的直接成员只能是 Req 和 Resp"}
	ruleExtraMember   = lintRule{"RPC_CONTAINER", "参数 message %s 只能直接定义 Req 和 Resp，发现 %s"}
	ruleMissingMember = lintRule{"RPC_CONTAINER", "参数 message %s 必须直接定义 %s"}
)
