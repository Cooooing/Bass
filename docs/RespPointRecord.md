

# RespPointRecord

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**createdAt** | **OffsetDateTime** | 记录创建时间 |  [optional] |
|**typeLabel** | **String** | 流水类型展示文案 |  [optional] |
|**direction** | [**DirectionEnum**](#DirectionEnum) | 收入或支出方向 |  [optional] |
|**amount** | **String** | 绝对变动积分 |  [optional] |
|**balanceAfter** | **String** | 变动后的积分余额 |  [optional] |



## Enum: DirectionEnum

| Name | Value |
|---- | -----|
| ECONOMY_RECORD_DIRECTION_UNSPECIFIED | &quot;ECONOMY_RECORD_DIRECTION_UNSPECIFIED&quot; |
| ECONOMY_RECORD_DIRECTION_INCOME | &quot;ECONOMY_RECORD_DIRECTION_INCOME&quot; |
| ECONOMY_RECORD_DIRECTION_EXPENSE | &quot;ECONOMY_RECORD_DIRECTION_EXPENSE&quot; |



