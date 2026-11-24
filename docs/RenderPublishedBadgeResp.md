
# RenderPublishedBadgeResp

Resp 通过 raw_body 返回 SVG 原始字节及内容类型。

## Properties

Name | Type
------------ | -------------
`rawBody` | [ImageResp](ImageResp.md)

## Example

```typescript
import type { RenderPublishedBadgeResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "rawBody": null,
} satisfies RenderPublishedBadgeResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RenderPublishedBadgeResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


