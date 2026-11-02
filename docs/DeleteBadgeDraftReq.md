
# DeleteBadgeDraftReq

Req 指定待删除草稿。

## Properties

Name | Type
------------ | -------------
`draftId` | string

## Example

```typescript
import type { DeleteBadgeDraftReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "draftId": null,
} satisfies DeleteBadgeDraftReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DeleteBadgeDraftReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


