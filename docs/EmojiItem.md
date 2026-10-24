
# EmojiItem

EmojiItem 是可直接插入编辑器的表情图片。

## Properties

Name | Type
------------ | -------------
`id` | string
`assetId` | string
`name` | string
`imageUrl` | string

## Example

```typescript
import type { EmojiItem } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "assetId": null,
  "name": null,
  "imageUrl": null,
} satisfies EmojiItem

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as EmojiItem
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


