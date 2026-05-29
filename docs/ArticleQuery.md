
# ArticleQuery


## Properties

Name | Type
------------ | -------------
`tagId` | string
`domainId` | string
`status` | string
`type` | string
`order` | string
`keyword` | string
`authorId` | string
`listable` | boolean

## Example

```typescript
import type { ArticleQuery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "tagId": null,
  "domainId": null,
  "status": null,
  "type": null,
  "order": null,
  "keyword": null,
  "authorId": null,
  "listable": null,
} satisfies ArticleQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


