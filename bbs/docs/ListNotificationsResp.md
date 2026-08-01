
# ListNotificationsResp


## Properties

Name | Type
------------ | -------------
`page` | [PageResp](PageResp.md)
`rows` | [Array&lt;RespNotification&gt;](RespNotification.md)

## Example

```typescript
import type { ListNotificationsResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "page": null,
  "rows": null,
} satisfies ListNotificationsResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListNotificationsResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


