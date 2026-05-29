
# Location

账号地理资料

## Properties

Name | Type
------------ | -------------
`userId` | string
`country` | string
`province` | string
`city` | string

## Example

```typescript
import type { Location } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "country": null,
  "province": null,
  "city": null,
} satisfies Location

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Location
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


