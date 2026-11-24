

# ReqTag

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**code** | **String** | 稳定编码 |  |
|**name** | **String** | 标签名称 |  |
|**description** | **String** | 描述 |  [optional] |
|**domainId** | **String** | 领域 ID |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 启停状态 |  [optional] |
|**icon** | **String** | 图标 |  [optional] |
|**sort** | **Integer** | 排序值，越小越靠前 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| TAG_STATUS_UNSPECIFIED | &quot;TAG_STATUS_UNSPECIFIED&quot; |
| TAG_STATUS_ENABLED | &quot;TAG_STATUS_ENABLED&quot; |
| TAG_STATUS_DISABLED | &quot;TAG_STATUS_DISABLED&quot; |



