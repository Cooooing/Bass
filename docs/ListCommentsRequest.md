# ListCommentsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageRequest**](PageRequest.md) | 分页参数。 | [optional] [default to undefined]
**query** | [**CommentQuery**](CommentQuery.md) | 查询条件。 | [optional] [default to undefined]

## Example

```typescript
import { ListCommentsRequest } from '@bass/bbs-sdk-axios';

const instance: ListCommentsRequest = {
    page,
    query,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
