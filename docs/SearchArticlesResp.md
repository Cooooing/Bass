
# SearchArticlesResp


## Properties

Name | Type
------------ | -------------
`rows` | [Array&lt;ArticleListItem&gt;](ArticleListItem.md)
`nextCursor` | string

## Example

```typescript
import type { SearchArticlesResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "rows": null,
  "nextCursor": null,
} satisfies SearchArticlesResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SearchArticlesResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


