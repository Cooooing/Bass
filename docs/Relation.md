

# Relation

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 关系 ID |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 关系类型 |  [optional] |
|**actorId** | **String** | 发起账号 ID |  [optional] |
|**targetId** | **String** | 目标账号 ID |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| RELATION_TYPE_UNSPECIFIED | &quot;RELATION_TYPE_UNSPECIFIED&quot; |
| RELATION_TYPE_FOLLOW | &quot;RELATION_TYPE_FOLLOW&quot; |
| RELATION_TYPE_BLOCK | &quot;RELATION_TYPE_BLOCK&quot; |



