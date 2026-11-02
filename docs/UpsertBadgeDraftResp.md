
# UpsertBadgeDraftResp

Resp 返回保存后的草稿。

## Properties

Name | Type
------------ | -------------
`draft` | [BadgeDraft](BadgeDraft.md)

## Example

```typescript
import type { UpsertBadgeDraftResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "draft": null,
} satisfies UpsertBadgeDraftResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UpsertBadgeDraftResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


