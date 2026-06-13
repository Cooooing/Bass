
# RequestArticle


## Properties

Name | Type
------------ | -------------
`title` | string
`content` | string
`rewardContent` | string
`rewardPoints` | number
`type` | string
`bountyPoints` | number
`statement` | string
`commentable` | boolean
`anonymous` | boolean
`tagIds` | Array&lt;string&gt;

## Example

```typescript
import type { RequestArticle } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "title": null,
  "content": null,
  "rewardContent": null,
  "rewardPoints": null,
  "type": null,
  "bountyPoints": null,
  "statement": null,
  "commentable": null,
  "anonymous": null,
  "tagIds": null,
} satisfies RequestArticle

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RequestArticle
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


