
# Notification

通知记录。

## Properties

Name | Type
------------ | -------------
`id` | string
`eventId` | string
`receiverId` | string
`eventType` | string
`title` | string
`content` | string
`readAt` | string
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { Notification } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "eventId": null,
  "receiverId": null,
  "eventType": null,
  "title": null,
  "content": null,
  "readAt": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies Notification

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Notification
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


