
# CommentQuery

评论查询条件。

## Properties

Name | Type
------------ | -------------
`commentId` | string
`articleId` | string
`parentId` | string
`replyId` | string
`order` | string
`userId` | string
`level` | number
`status` | string

## Example

```typescript
import type { CommentQuery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "commentId": null,
  "articleId": null,
  "parentId": null,
  "replyId": null,
  "order": null,
  "userId": null,
  "level": null,
  "status": null,
} satisfies CommentQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommentQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


