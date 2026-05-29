
# Relation

账号关系记录

## Properties

Name | Type
------------ | -------------
`id` | string
`type` | number
`actorId` | string
`targetId` | string
`createdAt` | string
`updatedAt` | string

## Example

```typescript
import type { Relation } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "type": null,
  "actorId": null,
  "targetId": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies Relation

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Relation
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


