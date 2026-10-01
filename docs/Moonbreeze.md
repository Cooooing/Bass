
# Moonbreeze


## Properties

Name | Type
------------ | -------------
`id` | string
`content` | string
`author` | [AccountProfile](AccountProfile.md)
`city` | string
`createdAt` | Date

## Example

```typescript
import type { Moonbreeze } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "content": null,
  "author": null,
  "city": null,
  "createdAt": null,
} satisfies Moonbreeze

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Moonbreeze
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


