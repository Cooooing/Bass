
# RespRelation


## Properties

Name | Type
------------ | -------------
`id` | string
`type` | string
`actorId` | string
`targetId` | string
`createdAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { RespRelation } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "type": null,
  "actorId": null,
  "targetId": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies RespRelation

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespRelation
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


