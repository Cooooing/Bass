
# GoogleProtobufAny

Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.

## Properties

Name | Type
------------ | -------------
`type` | string

## Example

```typescript
import type { GoogleProtobufAny } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "type": null,
} satisfies GoogleProtobufAny

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as GoogleProtobufAny
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


