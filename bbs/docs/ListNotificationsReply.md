
# ListNotificationsReply


## Properties

Name | Type
------------ | -------------
`page` | [PageReply](PageReply.md)
`rows` | [Array&lt;Notification&gt;](Notification.md)

## Example

```typescript
import type { ListNotificationsReply } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "page": null,
  "rows": null,
} satisfies ListNotificationsReply

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListNotificationsReply
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


