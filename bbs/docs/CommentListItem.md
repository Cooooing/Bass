
# CommentListItem

评论列表项。

## Properties

Name | Type
------------ | -------------
`id` | string
`articleId` | string
`content` | string
`contentRender` | string
`level` | number
`parentId` | string
`replyId` | string
`replyCount` | number
`likeCount` | number
`thankCount` | number
`user` | [AccountProfile](AccountProfile.md)
`replyUser` | [AccountProfile](AccountProfile.md)
`article` | [ArticleBrief](ArticleBrief.md)
`viewerActionState` | [CommentViewerActionState](CommentViewerActionState.md)
`restriction` | string
`deletedAt` | string
`createdBy` | string
`updatedBy` | string
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { CommentListItem } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "articleId": null,
  "content": null,
  "contentRender": null,
  "level": null,
  "parentId": null,
  "replyId": null,
  "replyCount": null,
  "likeCount": null,
  "thankCount": null,
  "user": null,
  "replyUser": null,
  "article": null,
  "viewerActionState": null,
  "restriction": null,
  "deletedAt": null,
  "createdBy": null,
  "updatedBy": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies CommentListItem

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommentListItem
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


