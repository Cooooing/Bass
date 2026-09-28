
# RespCurrentAccount


## Properties

Name | Type
------------ | -------------
`profile` | [AccountProfile](AccountProfile.md)
`contact` | [AccountContact](AccountContact.md)

## Example

```typescript
import type { RespCurrentAccount } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "profile": null,
  "contact": null,
} satisfies RespCurrentAccount

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespCurrentAccount
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


