

# RespTag

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 标签 ID |  [optional] |
|**code** | **String** | 稳定编码 |  [optional] |
|**name** | **String** | 标签名称 |  [optional] |
|**description** | **String** | 描述 |  [optional] |
|**domainId** | **String** | 领域 ID |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 启停状态 |  [optional] |
|**icon** | **String** | 图标 |  [optional] |
|**sort** | **Integer** | 排序值，越小越靠前 |  [optional] |
|**articleCount** | **Integer** | 文章数量 |  [optional] |
|**createdBy** | **String** | 创建用户 ID |  [optional] |
|**updatedBy** | **String** | 更新用户 ID |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| TAG_STATUS_UNSPECIFIED | &quot;TAG_STATUS_UNSPECIFIED&quot; |
| TAG_STATUS_ENABLED | &quot;TAG_STATUS_ENABLED&quot; |
| TAG_STATUS_DISABLED | &quot;TAG_STATUS_DISABLED&quot; |



