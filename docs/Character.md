

# Character

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 角色 ID |  [optional] |
|**name** | **String** | 角色名称 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 角色状态 |  [optional] |
|**slot** | **Integer** | 角色槽位 |  [optional] |
|**actionQueueCapacity** | **Integer** | 行动队列容量上限 |  [optional] |
|**maxOfflineSeconds** | **String** | 最大离线收益结算秒数 |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |
|**lastOfflineAt** | **OffsetDateTime** | 最近离线时间 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| CHARACTER_STATUS_UNSPECIFIED | &quot;CHARACTER_STATUS_UNSPECIFIED&quot; |
| CHARACTER_STATUS_ACTIVE | &quot;CHARACTER_STATUS_ACTIVE&quot; |
| CHARACTER_STATUS_DISABLED | &quot;CHARACTER_STATUS_DISABLED&quot; |



