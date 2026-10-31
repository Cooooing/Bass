
# TransferAccountPointsReq


## Properties

Name | Type
------------ | -------------
`targetName` | string
`amount` | string
`remark` | string
`idempotencyKey` | string

## Example

```typescript
import type { TransferAccountPointsReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "targetName": null,
  "amount": null,
  "remark": null,
  "idempotencyKey": null,
} satisfies TransferAccountPointsReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as TransferAccountPointsReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


