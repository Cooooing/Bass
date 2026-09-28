
# RespCommentThread


## Properties

Name | Type
------------ | -------------
`root` | [RespCommentListItem](RespCommentListItem.md)
`previewReplies` | [Array&lt;RespCommentListItem&gt;](RespCommentListItem.md)
`replyCount` | number
`hasMoreReplies` | boolean

## Example

```typescript
import type { RespCommentThread } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "root": null,
  "previewReplies": null,
  "replyCount": null,
  "hasMoreReplies": null,
} satisfies RespCommentThread

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespCommentThread
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


