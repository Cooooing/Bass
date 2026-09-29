
# AccountProfileListItem


## Properties

Name | Type
------------ | -------------
`account` | [AccountProfile](AccountProfile.md)
`viewerRelation` | [ProfileRelation](ProfileRelation.md)

## Example

```typescript
import type { AccountProfileListItem } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "account": null,
  "viewerRelation": null,
} satisfies AccountProfileListItem

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as AccountProfileListItem
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


