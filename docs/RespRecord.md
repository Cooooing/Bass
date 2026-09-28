
# RespRecord


## Properties

Name | Type
------------ | -------------
`id` | string
`transactionNo` | string
`recordType` | string
`direction` | string
`amount` | string
`balanceBefore` | string
`balanceAfter` | string
`remark` | string
`createdAt` | Date

## Example

```typescript
import type { RespRecord } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "transactionNo": null,
  "recordType": null,
  "direction": null,
  "amount": null,
  "balanceBefore": null,
  "balanceAfter": null,
  "remark": null,
  "createdAt": null,
} satisfies RespRecord

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespRecord
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


