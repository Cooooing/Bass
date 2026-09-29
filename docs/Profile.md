
# Profile


## Properties

Name | Type
------------ | -------------
`account` | [AccountProfile](AccountProfile.md)
`backgroundUrl` | string
`location` | [ProfileLocation](ProfileLocation.md)
`lastSuccessLoginAt` | Date
`visibility` | [ProfileVisibility](ProfileVisibility.md)
`viewerRelation` | [ProfileRelation](ProfileRelation.md)

## Example

```typescript
import type { Profile } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "account": null,
  "backgroundUrl": null,
  "location": null,
  "lastSuccessLoginAt": null,
  "visibility": null,
  "viewerRelation": null,
} satisfies Profile

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Profile
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


