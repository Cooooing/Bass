
# CreateCommentReq


## Properties

Name | Type
------------ | -------------
`articleId` | string
`content` | string
`replyId` | string

## Example

```typescript
import type { CreateCommentReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "articleId": null,
  "content": null,
  "replyId": null,
} satisfies CreateCommentReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CreateCommentReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


