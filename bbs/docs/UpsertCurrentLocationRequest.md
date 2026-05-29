
# UpsertCurrentLocationRequest


## Properties

Name | Type
------------ | -------------
`country` | string
`province` | string
`city` | string

## Example

```typescript
import type { UpsertCurrentLocationRequest } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "country": null,
  "province": null,
  "city": null,
} satisfies UpsertCurrentLocationRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UpsertCurrentLocationRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


