
# ListCommentRepliesReq


## Properties

Name | Type
------------ | -------------
`page` | [PageReq](PageReq.md)
`articleId` | string
`parentId` | string
`order` | string

## Example

```typescript
import type { ListCommentRepliesReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "page": null,
  "articleId": null,
  "parentId": null,
  "order": null,
} satisfies ListCommentRepliesReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListCommentRepliesReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


