
# UpdateCurrentPrivacySettingRequest


## Properties

Name | Type
------------ | -------------
`publicPoints` | boolean
`publicFollowers` | boolean
`publicArticles` | boolean
`publicComments` | boolean
`publicOnlineStatus` | boolean
`publicLocation` | boolean

## Example

```typescript
import type { UpdateCurrentPrivacySettingRequest } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "publicPoints": null,
  "publicFollowers": null,
  "publicArticles": null,
  "publicComments": null,
  "publicOnlineStatus": null,
  "publicLocation": null,
} satisfies UpdateCurrentPrivacySettingRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UpdateCurrentPrivacySettingRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


