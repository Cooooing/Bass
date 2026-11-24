# EnterArticleLotteryReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**entered** | **boolean** | true 报名，false 撤销报名 | [optional] [default to undefined]

## Example

```typescript
import { EnterArticleLotteryReq } from '@bass/bbs-sdk-axios';

const instance: EnterArticleLotteryReq = {
    article_id,
    entered,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
