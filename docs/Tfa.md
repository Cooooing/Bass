
# Tfa

二步验证状态

## Properties

Name | Type
------------ | -------------
`userId` | string
`enable` | boolean
`enableTime` | string

## Example

```typescript
import type { Tfa } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "enable": null,
  "enableTime": null,
} satisfies Tfa

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Tfa
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


