# RespDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 领域 ID | [optional]
**code** | Option<**String**> | 稳定编码 | [optional]
**name** | Option<**String**> | 领域名称 | [optional]
**description** | Option<**String**> | 描述 | [optional]
**status** | Option<**Status**> | 启停状态 (enum: DOMAIN_STATUS_UNSPECIFIED, DOMAIN_STATUS_ENABLED, DOMAIN_STATUS_DISABLED) | [optional]
**url** | Option<**String**> | 链接 | [optional]
**icon** | Option<**String**> | 图标 | [optional]
**is_nav** | Option<**bool**> | 是否在导航中展示 | [optional]
**sort** | Option<**i32**> | 排序值，越小越靠前 | [optional]
**created_by** | Option<**String**> | 创建用户 ID | [optional]
**updated_by** | Option<**String**> | 更新用户 ID | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


