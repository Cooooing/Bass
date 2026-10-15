
# ArticlePoll


## Properties

Name | Type
------------ | -------------
`endsAt` | Date
`maxChoices` | number
`options` | [Array&lt;ArticlePollOption&gt;](ArticlePollOption.md)
`selectedOptionIds` | Array&lt;string&gt;
`resultVisible` | boolean
`participantCount` | string
`totalVoteCount` | string

## Example

```typescript
import type { ArticlePoll } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "endsAt": null,
  "maxChoices": null,
  "options": null,
  "selectedOptionIds": null,
  "resultVisible": null,
  "participantCount": null,
  "totalVoteCount": null,
} satisfies ArticlePoll

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticlePoll
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


