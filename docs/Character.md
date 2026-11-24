# Character

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 角色 ID | [optional]
**name** | Option<**String**> | 角色名称 | [optional]
**status** | Option<**Status**> | 角色状态 (enum: CHARACTER_STATUS_UNSPECIFIED, CHARACTER_STATUS_ACTIVE, CHARACTER_STATUS_DISABLED) | [optional]
**slot** | Option<**i32**> | 角色槽位 | [optional]
**action_queue_capacity** | Option<**i32**> | 行动队列容量上限 | [optional]
**max_offline_seconds** | Option<**String**> | 最大离线收益结算秒数 | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]
**last_offline_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 最近离线时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


