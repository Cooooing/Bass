
# ReqArticleQuery


## Properties

Name | Type
------------ | -------------
`tagId` | string
`domainId` | string
`type` | string
`order` | string
`keyword` | string
`authorId` | string
`publishStatus` | string
`publishStatuses` | Array&lt;string&gt;
`visibility` | string
`visibilities` | Array&lt;string&gt;
`restriction` | string
`restrictions` | Array&lt;string&gt;

## Example

```typescript
import type { ReqArticleQuery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "tagId": null,
  "domainId": null,
  "type": null,
  "order": null,
  "keyword": null,
  "authorId": null,
  "publishStatus": null,
  "publishStatuses": null,
  "visibility": null,
  "visibilities": null,
  "restriction": null,
  "restrictions": null,
} satisfies ReqArticleQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ReqArticleQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


