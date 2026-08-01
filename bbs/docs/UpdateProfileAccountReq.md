
# UpdateProfileAccountReq


## Properties

Name | Type
------------ | -------------
`avatarUrl` | string
`nickname` | string
`url` | string
`introduction` | string
`mbti` | string

## Example

```typescript
import type { UpdateProfileAccountReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "avatarUrl": null,
  "nickname": null,
  "url": null,
  "introduction": null,
  "mbti": null,
} satisfies UpdateProfileAccountReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UpdateProfileAccountReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


