
# ArticleLotteryDraft


## Properties

Name | Type
------------ | -------------
`prizes` | [Array&lt;ArticleLotteryPrizeDraft&gt;](ArticleLotteryPrizeDraft.md)
`drawAt` | Date

## Example

```typescript
import type { ArticleLotteryDraft } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "prizes": null,
  "drawAt": null,
} satisfies ArticleLotteryDraft

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleLotteryDraft
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


