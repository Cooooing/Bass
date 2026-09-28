
# Character


## Properties

Name | Type
------------ | -------------
`id` | string
`name` | string
`status` | string
`slot` | number
`actionQueueCapacity` | number
`maxOfflineSeconds` | string
`createdAt` | Date
`updatedAt` | Date
`lastOfflineAt` | Date

## Example

```typescript
import type { Character } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "name": null,
  "status": null,
  "slot": null,
  "actionQueueCapacity": null,
  "maxOfflineSeconds": null,
  "createdAt": null,
  "updatedAt": null,
  "lastOfflineAt": null,
} satisfies Character

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as Character
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


