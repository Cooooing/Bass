
# CommentViewerActionState

当前查看账号的评论行为状态。

## Properties

Name | Type
------------ | -------------
`liked` | boolean
`thanked` | boolean

## Example

```typescript
import type { CommentViewerActionState } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "liked": null,
  "thanked": null,
} satisfies CommentViewerActionState

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommentViewerActionState
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


