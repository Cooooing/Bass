
# RegisterReq


## Properties

Name | Type
------------ | -------------
`type` | string
`name` | string
`password` | string
`nickname` | string
`emailCredential` | [ReqEmailCredential](ReqEmailCredential.md)
`phoneCredential` | [ReqPhoneCredential](ReqPhoneCredential.md)

## Example

```typescript
import type { RegisterReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "type": null,
  "name": null,
  "password": null,
  "nickname": null,
  "emailCredential": null,
  "phoneCredential": null,
} satisfies RegisterReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegisterReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


