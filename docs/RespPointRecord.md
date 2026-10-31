
# RespPointRecord


## Properties

Name | Type
------------ | -------------
`createdAt` | Date
`typeLabel` | string
`direction` | string
`amount` | string
`balanceAfter` | string

## Example

```typescript
import type { RespPointRecord } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "typeLabel": null,
  "direction": null,
  "amount": null,
  "balanceAfter": null,
} satisfies RespPointRecord

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespPointRecord
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


