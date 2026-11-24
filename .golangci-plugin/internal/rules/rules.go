// Package rules 集中定义 Bass 自定义 lint 的规则编号和中文提示。
package rules

import (
	"fmt"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// Rule 绑定稳定规则编号与诊断模板，所有自定义提示必须从此处引用。
type Rule struct {
	Code    string
	Message string
}

// Reportf 统一在诊断中输出规则编号，便于门禁、审计与精确豁免定位。
func (r Rule) Reportf(pass *analysis.Pass, pos token.Pos, args ...any) {
	pass.Reportf(pos, "[%s] %s", r.Code, fmt.Sprintf(r.Message, args...))
}

// 所有项目级 Go 规则的编号与中文提示均在本文件维护。
var (
	DepCrossModule = Rule{"BASSDEP001", "禁止跨模块导入 internal 包：%s；请通过 common 或 RPC client 复用"}
	DepLayer       = Rule{"BASSDEP002", "%s 层不得依赖 %s（目标层：%s）"}
	DepConstruct   = Rule{"BASSDEP003", "基础设施客户端只能在 cmd、module 或 data 层构造，当前调用 %s"}
	DepBusiness    = Rule{"BASSDEP004", "biz 层不得依赖基础设施实现包：%s"}
	ConfigEnv      = Rule{"BASSCFG001", "环境变量只能在 cmd、module 或配置加载实现中读取，当前调用 %s"}
	RepoLeak       = Rule{"BASSREPO001", "biz/repo 不得泄漏数据实现类型或包：%s"}
	APIType        = Rule{"BASSAPI001", "service 边界不得使用数据实现类型：%s"}
	Database       = Rule{"BASSDB001", "SQL 只能在 data、migration 或装配层访问：%s"}
	ContextRoot    = Rule{"BASSCTX001", "internal 业务代码不得创建根 context，当前调用 %s"}
	GoRoutine      = Rule{"BASSCONC001", "业务代码不得直接启动 goroutine；请交由 server 生命周期或受控并发组件管理"}
	GoCancel       = Rule{"BASSCONC002", "goroutine 必须引用可取消 context、停止信号或等待组"}
	TimeSleep      = Rule{"BASSTIME001", "业务代码不得调用 time.Sleep；请使用受控调度、重试或定时任务"}

	EnumSingle      = Rule{"BASSENUM001", "每个 internal/enum 文件有且仅有一个 type X string 枚举，当前为 %d 个"}
	EnumFilename    = Rule{"BASSENUM002", "%s 必须位于 %s.go，当前文件为 %s.go"}
	EnumValue       = Rule{"BASSENUM003", "%s 的每个值必须写为 Name %s = \"value\" // 中文注释"}
	EnumComment     = Rule{"BASSENUM004", "%s.%s 必须在定义同行追加中文业务注释"}
	EnumMapping     = Rule{"BASSENUM005", "%s 必须定义 var %s = enum.NewMapping[...]"}
	EnumMappingSize = Rule{"BASSENUM006", "%s 的映射数量 %d 必须与枚举值数量 %d 一致"}
	EnumProto       = Rule{"BASSENUM007", "%s 的 Proto 映射必须与 proto %s 的非 UNSPECIFIED 值顺序一致"}
	EnumDuplicate   = Rule{"BASSENUM008", "%s 的字符串值 %s 与 %s 重复"}
	EnumMappingKey  = Rule{"BASSENUM009", "%s 的映射键必须是本文件声明的 %s 常量"}

	SchemaSingle   = Rule{"BASSSCHEMA001", "每个 internal/data/schema 文件只能声明一张 Ent 表；请将 %s 拆分到独立文件"}
	SchemaField    = Rule{"BASSSCHEMA002", "%s.%s 字段必须使用中文 Comment 说明业务语义"}
	SchemaFields   = Rule{"BASSSCHEMA003", "Ent 实体 %s 必须声明直接 Fields() 方法或使用精确路径豁免"}
	SchemaFilename = Rule{"BASSSCHEMA004", "Ent 实体 %s 必须位于 %s.go，当前文件为 %s.go"}
	SchemaComment  = Rule{"BASSSCHEMA005", "Ent %s %s 必须使用中文注释说明业务语义"}
)
