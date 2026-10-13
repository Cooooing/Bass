
# SearchArticlesReq


## Properties

Name | Type
------------ | -------------
`query` | string
`cursor` | string
`size` | number

## Example

```typescript
import type { SearchArticlesReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "query": null,
  "cursor": null,
  "size": null,
} satisfies SearchArticlesReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SearchArticlesReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


