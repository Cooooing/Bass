
# Article


## Properties

Name | Type
------------ | -------------
`id` | string
`title` | string
`content` | string
`contentRender` | string
`hasPostscript` | boolean
`hasReward` | boolean
`rewardContent` | string
`rewardContentRender` | string
`rewardPoints` | number
`status` | string
`type` | string
`statement` | string
`commentable` | boolean
`anonymous` | boolean
`listable` | boolean
`viewCount` | number
`thankCount` | number
`likeCount` | number
`collectCount` | number
`watchCount` | number
`replyCount` | number
`bountyPoints` | number
`acceptedAnswerId` | string
`authorUser` | [AccountProfile](AccountProfile.md)
`lastReplyUser` | [AccountProfile](AccountProfile.md)
`lastReplyAt` | string
`coverImageUrl` | string
`createdBy` | string
`updatedBy` | string
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { Article } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "title": null,
  "content": null,
  "contentRender": null,
  "hasPostscript": null,
  "hasReward": null,
  "rewardContent": null,
  "rewardContentRender": null,
  "rewardPoints": null,
  "status": null,
  "type": null,
  "statement": null,
  "commentable": null,
  "anonymous": null,
  "listable": null,
  "viewCount": null,
  "thankCount": null,
  "likeCount": null,
  "collectCount": null,
  "watchCount": null,
  "replyCount": null,
  "bountyPoints": null,
  "acceptedAnswerId": null,
  "authorUser": null,
  "lastReplyUser": null,
  "lastReplyAt": null,
  "coverImageUrl": null,
  "createdBy": null,
  "updatedBy": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies Article

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Article
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


