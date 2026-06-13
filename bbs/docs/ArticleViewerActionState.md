
# ArticleViewerActionState

当前查看账号的文章行为状态。

## Properties

Name | Type
------------ | -------------
`liked` | boolean
`thanked` | boolean
`collected` | boolean
`watched` | boolean

## Example

```typescript
import type { ArticleViewerActionState } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "liked": null,
  "thanked": null,
  "collected": null,
  "watched": null,
} satisfies ArticleViewerActionState

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ArticleViewerActionState
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


