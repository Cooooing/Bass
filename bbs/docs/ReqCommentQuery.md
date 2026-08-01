
# ReqCommentQuery


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
`restriction` | string
`restrictions` | Array&lt;string&gt;

## Example

```typescript
import type { ReqCommentQuery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "commentId": null,
  "articleId": null,
  "parentId": null,
  "replyId": null,
  "order": null,
  "userId": null,
  "level": null,
  "restriction": null,
  "restrictions": null,
} satisfies ReqCommentQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ReqCommentQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


