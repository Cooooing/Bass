
# ImportEmojiShareResp


## Properties

Name | Type
------------ | -------------
`group` | [EmojiGroup](EmojiGroup.md)
`skippedCount` | number

## Example

```typescript
import type { ImportEmojiShareResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "group": null,
  "skippedCount": null,
} satisfies ImportEmojiShareResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ImportEmojiShareResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


