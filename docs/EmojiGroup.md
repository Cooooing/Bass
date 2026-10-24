
# EmojiGroup

EmojiGroup 是当前账号可操作的表情分组。

## Properties

Name | Type
------------ | -------------
`id` | string
`name` | string
`items` | [Array&lt;EmojiItem&gt;](EmojiItem.md)

## Example

```typescript
import type { EmojiGroup } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "name": null,
  "items": null,
} satisfies EmojiGroup

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as EmojiGroup
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


