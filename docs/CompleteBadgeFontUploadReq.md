
# CompleteBadgeFontUploadReq

Req 指定已上传字体的内容哈希。

## Properties

Name | Type
------------ | -------------
`hash` | string

## Example

```typescript
import type { CompleteBadgeFontUploadReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "hash": null,
} satisfies CompleteBadgeFontUploadReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CompleteBadgeFontUploadReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


