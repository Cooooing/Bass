
# CommentThread

评论楼层项。

## Properties

Name | Type
------------ | -------------
`root` | [CommentListItem](CommentListItem.md)
`previewReplies` | [Array&lt;CommentListItem&gt;](CommentListItem.md)
`replyCount` | number
`hasMoreReplies` | boolean

## Example

```typescript
import type { CommentThread } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "root": null,
  "previewReplies": null,
  "replyCount": null,
  "hasMoreReplies": null,
} satisfies CommentThread

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommentThread
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


