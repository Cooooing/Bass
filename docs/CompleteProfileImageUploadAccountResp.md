
# CompleteProfileImageUploadAccountResp


## Properties

Name | Type
------------ | -------------
`profile` | [AccountProfile](AccountProfile.md)
`imageUrl` | string

## Example

```typescript
import type { CompleteProfileImageUploadAccountResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "profile": null,
  "imageUrl": null,
} satisfies CompleteProfileImageUploadAccountResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CompleteProfileImageUploadAccountResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


