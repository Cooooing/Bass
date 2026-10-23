
# PrivateChatMessage


## Properties

Name | Type
------------ | -------------
`id` | string
`senderId` | string
`receiverId` | string
`content` | string
`messageType` | string
`revoked` | boolean
`createdAt` | Date
`revokedAt` | Date

## Example

```typescript
import type { PrivateChatMessage } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "senderId": null,
  "receiverId": null,
  "content": null,
  "messageType": null,
  "revoked": null,
  "createdAt": null,
  "revokedAt": null,
} satisfies PrivateChatMessage

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PrivateChatMessage
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


