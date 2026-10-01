
# RespPrivacySetting


## Properties

Name | Type
------------ | -------------
`userId` | string
`publicPoints` | boolean
`publicFollowerList` | boolean
`publicFollowingList` | boolean
`publicArticleList` | boolean
`publicCommentList` | boolean
`publicOnlineStatus` | boolean
`publicLocation` | boolean
`publicMoonbreezeList` | boolean

## Example

```typescript
import type { RespPrivacySetting } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "userId": null,
  "publicPoints": null,
  "publicFollowerList": null,
  "publicFollowingList": null,
  "publicArticleList": null,
  "publicCommentList": null,
  "publicOnlineStatus": null,
  "publicLocation": null,
  "publicMoonbreezeList": null,
} satisfies RespPrivacySetting

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespPrivacySetting
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


