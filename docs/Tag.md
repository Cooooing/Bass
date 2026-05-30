
# Tag

Adds metadata to a single tag that is used by the Operation Object. It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.

## Properties

Name | Type
------------ | -------------
`name` | string
`description` | string
`externalDocs` | [ExternalDocs](ExternalDocs.md)
`specificationExtension` | [Array&lt;NamedAny&gt;](NamedAny.md)

## Example

```typescript
import type { Tag } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "name": null,
  "description": null,
  "externalDocs": null,
  "specificationExtension": null,
} satisfies Tag

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Tag
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


