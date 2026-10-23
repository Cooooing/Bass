
# SendPrivateChatMessageReq


## Properties

Name | Type
------------ | -------------
`receiverId` | string
`clientMessageId` | string
`content` | string

## Example

```typescript
import type { SendPrivateChatMessageReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "receiverId": null,
  "clientMessageId": null,
  "content": null,
} satisfies SendPrivateChatMessageReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SendPrivateChatMessageReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


