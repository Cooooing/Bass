
# ProfileVisibility


## Properties

Name | Type
------------ | -------------
`articles` | boolean
`comments` | boolean
`followers` | boolean
`following` | boolean

## Example

```typescript
import type { ProfileVisibility } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "articles": null,
  "comments": null,
  "followers": null,
  "following": null,
} satisfies ProfileVisibility

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ProfileVisibility
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


