
# RespRelationStatus


## Properties

Name | Type
------------ | -------------
`targetId` | string
`following` | boolean
`followedBy` | boolean
`blocking` | boolean
`blockedBy` | boolean

## Example

```typescript
import type { RespRelationStatus } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "targetId": null,
  "following": null,
  "followedBy": null,
  "blocking": null,
  "blockedBy": null,
} satisfies RespRelationStatus

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespRelationStatus
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


