
# ListCommentThreadsReply


## Properties

Name | Type
------------ | -------------
`page` | [PageReply](PageReply.md)
`rows` | [Array&lt;CommentThread&gt;](CommentThread.md)

## Example

```typescript
import type { ListCommentThreadsReply } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "page": null,
  "rows": null,
} satisfies ListCommentThreadsReply

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListCommentThreadsReply
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


