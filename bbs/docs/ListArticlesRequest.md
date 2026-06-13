
# ListArticlesRequest


## Properties

Name | Type
------------ | -------------
`page` | [PageRequest](PageRequest.md)
`query` | [ArticleQuery](ArticleQuery.md)

## Example

```typescript
import type { ListArticlesRequest } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "page": null,
  "query": null,
} satisfies ListArticlesRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListArticlesRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


