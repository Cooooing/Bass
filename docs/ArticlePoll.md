# ArticlePoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ends_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 截止时间 | [optional]
**max_choices** | Option<**i32**> | 每人最多选择数 | [optional]
**options** | Option<[**Vec<models::ArticlePollOption>**](ArticlePollOption.md)> | 投票选项 | [optional]
**selected_option_ids** | Option<**Vec<String>**> | 当前账号已选选项 ID | [optional]
**result_visible** | Option<**bool**> | 当前账号是否可查看实时结果 | [optional]
**participant_count** | Option<**String**> | 已投票账号数 | [optional]
**total_vote_count** | Option<**String**> | 所有选项累计票数 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


