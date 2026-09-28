
# CancelAccountReq


## Properties

Name | Type
------------ | -------------
`password` | string
`code` | string

## Example

```typescript
import type { CancelAccountReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "password": null,
  "code": null,
} satisfies CancelAccountReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CancelAccountReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


