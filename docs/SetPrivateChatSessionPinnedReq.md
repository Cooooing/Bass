
# SetPrivateChatSessionPinnedReq


## Properties

Name | Type
------------ | -------------
`peerUserId` | string
`pinned` | boolean

## Example

```typescript
import type { SetPrivateChatSessionPinnedReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "peerUserId": null,
  "pinned": null,
} satisfies SetPrivateChatSessionPinnedReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SetPrivateChatSessionPinnedReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


