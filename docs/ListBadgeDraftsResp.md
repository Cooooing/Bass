
# ListBadgeDraftsResp

Resp 返回草稿行。

## Properties

Name | Type
------------ | -------------
`rows` | [Array&lt;BadgeDraft&gt;](BadgeDraft.md)

## Example

```typescript
import type { ListBadgeDraftsResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "rows": null,
} satisfies ListBadgeDraftsResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListBadgeDraftsResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


