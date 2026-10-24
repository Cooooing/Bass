
# GlobalChatTopic


## Properties

Name | Type
------------ | -------------
`id` | string
`userId` | string
`content` | string
`pointsCost` | string
`createdAt` | Date

## Example

```typescript
import type { GlobalChatTopic } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "userId": null,
  "content": null,
  "pointsCost": null,
  "createdAt": null,
} satisfies GlobalChatTopic

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as GlobalChatTopic
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


