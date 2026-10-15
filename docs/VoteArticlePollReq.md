
# VoteArticlePollReq


## Properties

Name | Type
------------ | -------------
`articleId` | string
`optionIds` | Array&lt;string&gt;

## Example

```typescript
import type { VoteArticlePollReq } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "articleId": null,
  "optionIds": null,
} satisfies VoteArticlePollReq

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as VoteArticlePollReq
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


