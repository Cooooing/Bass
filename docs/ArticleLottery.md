
# ArticleLottery


## Properties

Name | Type
------------ | -------------
`drawAt` | Date
`drawnAt` | Date
`entryCount` | string
`entered` | boolean
`prizes` | [Array&lt;ArticleLotteryPrize&gt;](ArticleLotteryPrize.md)

## Example

```typescript
import type { ArticleLottery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "drawAt": null,
  "drawnAt": null,
  "entryCount": null,
  "entered": null,
  "prizes": null,
} satisfies ArticleLottery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleLottery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


