
# PrivateChatSession


## Properties

Name | Type
------------ | -------------
`peerUserId` | string
`lastReadMessageId` | string
`unreadCount` | number
`muted` | boolean
`pinned` | boolean
`lastMessage` | [PrivateChatMessage](PrivateChatMessage.md)
`lastMessageAt` | Date
`peer` | [PrivateChatMember](PrivateChatMember.md)

## Example

```typescript
import type { PrivateChatSession } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "peerUserId": null,
  "lastReadMessageId": null,
  "unreadCount": null,
  "muted": null,
  "pinned": null,
  "lastMessage": null,
  "lastMessageAt": null,
  "peer": null,
} satisfies PrivateChatSession

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PrivateChatSession
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


