
# CommonApiAppBbsV1UserAccountProfile

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
`status` | number
`groupName` | string
`followCount` | number
`followerCount` | number
`blockCount` | number
`blockedCount` | number
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { CommonApiAppBbsV1UserAccountProfile } from '@bass/bbs-sdk'

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
  "groupName": null,
  "followCount": null,
  "followerCount": null,
  "blockCount": null,
  "blockedCount": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies CommonApiAppBbsV1UserAccountProfile

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommonApiAppBbsV1UserAccountProfile
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


