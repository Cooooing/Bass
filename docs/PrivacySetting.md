
# PrivacySetting

账号隐私设置

## Properties

Name | Type
------------ | -------------
`userId` | string
`publicPoints` | boolean
`publicFollowers` | boolean
`publicArticles` | boolean
`publicComments` | boolean
`publicOnlineStatus` | boolean
`publicLocation` | boolean

## Example

```typescript
import type { PrivacySetting } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "publicPoints": null,
  "publicFollowers": null,
  "publicArticles": null,
  "publicComments": null,
  "publicOnlineStatus": null,
  "publicLocation": null,
} satisfies PrivacySetting

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PrivacySetting
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


