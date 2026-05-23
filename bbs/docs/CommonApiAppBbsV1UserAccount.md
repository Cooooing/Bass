
# CommonApiAppBbsV1UserAccount

当前账号完整资料

## Properties

Name | Type
------------ | -------------
`profile` | [CommonApiAppBbsV1UserAccountProfile](CommonApiAppBbsV1UserAccountProfile.md)
`contact` | [CommonApiAppBbsV1UserAccountContact](CommonApiAppBbsV1UserAccountContact.md)

## Example

```typescript
import type { CommonApiAppBbsV1UserAccount } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "profile": null,
  "contact": null,
} satisfies CommonApiAppBbsV1UserAccount

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommonApiAppBbsV1UserAccount
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


