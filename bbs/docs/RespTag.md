
# RespTag


## Properties

Name | Type
------------ | -------------
`id` | string
`code` | string
`name` | string
`description` | string
`domainId` | string
`status` | string
`icon` | string
`sort` | number
`articleCount` | number
`createdBy` | string
`updatedBy` | string
`createdAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { RespTag } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "code": null,
  "name": null,
  "description": null,
  "domainId": null,
  "status": null,
  "icon": null,
  "sort": null,
  "articleCount": null,
  "createdBy": null,
  "updatedBy": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies RespTag

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespTag
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


