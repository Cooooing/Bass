
# RespAccount


## Properties

Name | Type
------------ | -------------
`basic` | [RespAccountBasic](RespAccountBasic.md)
`contact` | [RespAccountContact](RespAccountContact.md)

## Example

```typescript
import type { RespAccount } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "basic": null,
  "contact": null,
} satisfies RespAccount

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespAccount
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


