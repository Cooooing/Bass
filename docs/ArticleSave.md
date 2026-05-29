
# ArticleSave


## Properties

Name | Type
------------ | -------------
`id` | string
`title` | string
`content` | string
`rewardContent` | string
`rewardPoints` | number
`status` | string
`type` | string
`bountyPoints` | number
`statement` | string
`commentable` | boolean
`anonymous` | boolean
`listable` | boolean
`tags` | [Array&lt;TagSave&gt;](TagSave.md)

## Example

```typescript
import type { ArticleSave } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "title": null,
  "content": null,
  "rewardContent": null,
  "rewardPoints": null,
  "status": null,
  "type": null,
  "bountyPoints": null,
  "statement": null,
  "commentable": null,
  "anonymous": null,
  "listable": null,
  "tags": null,
} satisfies ArticleSave

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleSave
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


