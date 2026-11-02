
# CompleteBadgeFontUploadResp

Resp 返回已确认的字体资源与前端本地预览地址。

## Properties

Name | Type
------------ | -------------
`assetId` | string
`fontUrl` | string

## Example

```typescript
import type { CompleteBadgeFontUploadResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "assetId": null,
  "fontUrl": null,
} satisfies CompleteBadgeFontUploadResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CompleteBadgeFontUploadResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


