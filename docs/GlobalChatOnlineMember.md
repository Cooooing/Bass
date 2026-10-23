
# GlobalChatOnlineMember


## Properties

Name | Type
------------ | -------------
`userId` | string
`name` | string
`nickname` | string
`avatarUrl` | string

## Example

```typescript
import type { GlobalChatOnlineMember } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "name": null,
  "nickname": null,
  "avatarUrl": null,
} satisfies GlobalChatOnlineMember

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as GlobalChatOnlineMember
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


