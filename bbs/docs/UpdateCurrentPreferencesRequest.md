
# UpdateCurrentPreferencesRequest


## Properties

Name | Type
------------ | -------------
`timezone` | string
`theme` | string
`mobileTheme` | string
`language` | string

## Example

```typescript
import type { UpdateCurrentPreferencesRequest } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "timezone": null,
  "theme": null,
  "mobileTheme": null,
  "language": null,
} satisfies UpdateCurrentPreferencesRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UpdateCurrentPreferencesRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


