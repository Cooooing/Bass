
# AccountProfile

账号展示资料

## Properties

Name | Type
------------ | -------------
`id` | string
`name` | string
`nickname` | string
`url` | string
`avatarUrl` | string
`introduction` | string
`mbti` | string
`status` | string
`followCount` | number
`followerCount` | number
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { AccountProfile } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "name": null,
  "nickname": null,
  "url": null,
  "avatarUrl": null,
  "introduction": null,
  "mbti": null,
  "status": null,
  "followCount": null,
  "followerCount": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies AccountProfile

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as AccountProfile
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


