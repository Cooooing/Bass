
# RespArticleBrief


## Properties

Name | Type
------------ | -------------
`id` | string
`title` | string
`type` | string
`authorUser` | [RespAccountProfile](RespAccountProfile.md)
`coverImageUrl` | string
`publishStatus` | string
`visibility` | string
`restriction` | string
`createdBy` | string
`updatedBy` | string
`createdAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { RespArticleBrief } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "title": null,
  "type": null,
  "authorUser": null,
  "coverImageUrl": null,
  "publishStatus": null,
  "visibility": null,
  "restriction": null,
  "createdBy": null,
  "updatedBy": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies RespArticleBrief

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespArticleBrief
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


