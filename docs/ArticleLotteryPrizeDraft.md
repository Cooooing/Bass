
# ArticleLotteryPrizeDraft


## Properties

Name | Type
------------ | -------------
`name` | string
`description` | string
`winnerLimit` | number

## Example

```typescript
import type { ArticleLotteryPrizeDraft } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "name": null,
  "description": null,
  "winnerLimit": null,
} satisfies ArticleLotteryPrizeDraft

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleLotteryPrizeDraft
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


