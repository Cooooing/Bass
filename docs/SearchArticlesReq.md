# SearchArticlesReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query** | **string** | 必填，搜索词 | [default to undefined]
**cursor** | **string** | 可选，不透明续页游标 | [optional] [default to undefined]
**size** | **number** | 可选，返回条数 | [optional] [default to undefined]

## Example

```typescript
import { SearchArticlesReq } from '@bass/bbs-sdk-axios';

const instance: SearchArticlesReq = {
    query,
    cursor,
    size,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
