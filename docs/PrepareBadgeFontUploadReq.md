
# PrepareBadgeFontUploadReq

Req 指定待上传字体的内容寻址信息。

## Properties

Name | Type
------------ | -------------
`hash` | string
`mimeType` | string
`size` | string

## Example

```typescript
import type { PrepareBadgeFontUploadReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "hash": null,
  "mimeType": null,
  "size": null,
} satisfies PrepareBadgeFontUploadReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PrepareBadgeFontUploadReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


