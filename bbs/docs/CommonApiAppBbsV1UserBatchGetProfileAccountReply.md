
# CommonApiAppBbsV1UserBatchGetProfileAccountReply


## Properties

Name | Type
------------ | -------------
`profiles` | [{ [key: string]: CommonApiAppBbsV1UserAccountProfile; }](CommonApiAppBbsV1UserAccountProfile.md)

## Example

```typescript
import type { CommonApiAppBbsV1UserBatchGetProfileAccountReply } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "profiles": null,
} satisfies CommonApiAppBbsV1UserBatchGetProfileAccountReply

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommonApiAppBbsV1UserBatchGetProfileAccountReply
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


