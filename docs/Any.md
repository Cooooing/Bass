
# Any


## Properties

Name | Type
------------ | -------------
`value` | [GoogleProtobufAny](GoogleProtobufAny.md)
`yaml` | string

## Example

```typescript
import type { Any } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "value": null,
  "yaml": null,
} satisfies Any

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Any
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


