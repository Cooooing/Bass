
# AccountContact

当前账号联系方式

## Properties

Name | Type
------------ | -------------
`userId` | string
`email` | string
`phone` | string

## Example

```typescript
import type { AccountContact } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "email": null,
  "phone": null,
} satisfies AccountContact

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as AccountContact
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


