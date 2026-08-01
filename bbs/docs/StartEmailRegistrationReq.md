
# StartEmailRegistrationReq


## Properties

Name | Type
------------ | -------------
`email` | string
`password` | string
`name` | string
`nickname` | string

## Example

```typescript
import type { StartEmailRegistrationReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "email": null,
  "password": null,
  "name": null,
  "nickname": null,
} satisfies StartEmailRegistrationReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as StartEmailRegistrationReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


