
# ArticleLotteryPrize


## Properties

Name | Type
------------ | -------------
`id` | string
`name` | string
`description` | string
`winnerLimit` | number
`winnerUsers` | [Array&lt;AccountProfile&gt;](AccountProfile.md)

## Example

```typescript
import type { ArticleLotteryPrize } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "name": null,
  "description": null,
  "winnerLimit": null,
  "winnerUsers": null,
} satisfies ArticleLotteryPrize

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleLotteryPrize
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


