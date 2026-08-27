# RespRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**transaction_no** | Option<**String**> |  | [optional]
**record_type** | Option<**RecordType**> |  (enum: ECONOMY_RECORD_TYPE_UNSPECIFIED, ECONOMY_RECORD_TYPE_SIGN_IN_REWARD, ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD, ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD, ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT, ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN, ECONOMY_RECORD_TYPE_ADMIN_ADD, ECONOMY_RECORD_TYPE_ADMIN_DEDUCT) | [optional]
**direction** | Option<**Direction**> |  (enum: ECONOMY_RECORD_DIRECTION_UNSPECIFIED, ECONOMY_RECORD_DIRECTION_INCOME, ECONOMY_RECORD_DIRECTION_EXPENSE) | [optional]
**amount** | Option<**String**> |  | [optional]
**balance_before** | Option<**String**> |  | [optional]
**balance_after** | Option<**String**> |  | [optional]
**remark** | Option<**String**> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


