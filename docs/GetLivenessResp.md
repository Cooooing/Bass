# GetLivenessResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**score** | **number** | 当前活跃度总分。 | [optional] [default to undefined]
**max_score** | **number** | 每日活跃度上限。 | [optional] [default to undefined]
**checkin_score** | **number** | 自动签到所需的最低活跃度。 | [optional] [default to undefined]

## Example

```typescript
import { GetLivenessResp } from '@bass/bbs-sdk-axios';

const instance: GetLivenessResp = {
    score,
    max_score,
    checkin_score,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
