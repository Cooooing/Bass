
# BadgeDraft

BadgeDraft 是 Maker 可编辑的徽章草稿。

## Properties

Name | Type
------------ | -------------
`id` | string
`name` | string
`description` | string
`renderSpec` | [RenderSpec](RenderSpec.md)
`designerUserId` | string

## Example

```typescript
import type { BadgeDraft } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "name": null,
  "description": null,
  "renderSpec": null,
  "designerUserId": null,
} satisfies BadgeDraft

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as BadgeDraft
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


