
# CommonApiAppBbsV1UserPreference

账号偏好设置

## Properties

Name | Type
------------ | -------------
`userId` | string
`language` | string
`timezone` | string
`theme` | string
`mobileTheme` | string
`enableWebNotify` | boolean
`enableEmailSubscribe` | boolean

## Example

```typescript
import type { CommonApiAppBbsV1UserPreference } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "language": null,
  "timezone": null,
  "theme": null,
  "mobileTheme": null,
  "enableWebNotify": null,
  "enableEmailSubscribe": null,
} satisfies CommonApiAppBbsV1UserPreference

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommonApiAppBbsV1UserPreference
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


