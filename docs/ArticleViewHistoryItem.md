
# ArticleViewHistoryItem


## Properties

Name | Type
------------ | -------------
`article` | [ArticleListItem](ArticleListItem.md)
`viewedAt` | Date

## Example

```typescript
import type { ArticleViewHistoryItem } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "article": null,
  "viewedAt": null,
} satisfies ArticleViewHistoryItem

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleViewHistoryItem
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


