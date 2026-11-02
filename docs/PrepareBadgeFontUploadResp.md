
# PrepareBadgeFontUploadResp

Resp 返回已存在字体或一次直传凭证。

## Properties

Name | Type
------------ | -------------
`assetId` | string
`uploadUrl` | string
`formFields` | { [key: string]: string; }

## Example

```typescript
import type { PrepareBadgeFontUploadResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "assetId": null,
  "uploadUrl": null,
  "formFields": null,
} satisfies PrepareBadgeFontUploadResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PrepareBadgeFontUploadResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


