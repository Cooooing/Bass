

# ReqDomainQuery

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**ids** | **List&lt;String&gt;** | 领域 ID 列表 |  [optional] |
|**code** | **String** | 稳定编码 |  [optional] |
|**name** | **String** | 领域名称 |  [optional] |
|**description** | **String** | 描述 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 启停状态 |  [optional] |
|**url** | **String** | 链接 |  [optional] |
|**icon** | **String** | 图标 |  [optional] |
|**isNav** | **Boolean** | 是否在导航中展示 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| DOMAIN_STATUS_UNSPECIFIED | &quot;DOMAIN_STATUS_UNSPECIFIED&quot; |
| DOMAIN_STATUS_ENABLED | &quot;DOMAIN_STATUS_ENABLED&quot; |
| DOMAIN_STATUS_DISABLED | &quot;DOMAIN_STATUS_DISABLED&quot; |



