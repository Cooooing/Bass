

# ReqDomain

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**code** | **String** | 稳定编码 |  |
|**name** | **String** | 领域名称 |  |
|**description** | **String** | 描述 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 启停状态 |  [optional] |
|**url** | **String** | 链接 |  [optional] |
|**icon** | **String** | 图标 |  [optional] |
|**isNav** | **Boolean** | 是否在导航中展示 |  [optional] |
|**sort** | **Integer** | 排序值，越小越靠前 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| DOMAIN_STATUS_UNSPECIFIED | &quot;DOMAIN_STATUS_UNSPECIFIED&quot; |
| DOMAIN_STATUS_ENABLED | &quot;DOMAIN_STATUS_ENABLED&quot; |
| DOMAIN_STATUS_DISABLED | &quot;DOMAIN_STATUS_DISABLED&quot; |



