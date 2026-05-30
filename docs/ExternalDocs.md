
# ExternalDocs

Allows referencing an external resource for extended documentation.

## Properties

Name | Type
------------ | -------------
`description` | string
`url` | string
`specificationExtension` | [Array&lt;NamedAny&gt;](NamedAny.md)

## Example

```typescript
import type { ExternalDocs } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "description": null,
  "url": null,
  "specificationExtension": null,
} satisfies ExternalDocs

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ExternalDocs
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


