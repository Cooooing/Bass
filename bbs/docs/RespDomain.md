
# RespDomain


## Properties

Name | Type
------------ | -------------
`id` | string
`code` | string
`name` | string
`description` | string
`status` | string
`url` | string
`icon` | string
`isNav` | boolean
`sort` | number
`createdBy` | string
`updatedBy` | string
`createdAt` | Date
`updatedAt` | Date

## Example

```typescript
import type { RespDomain } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "code": null,
  "name": null,
  "description": null,
  "status": null,
  "url": null,
  "icon": null,
  "isNav": null,
  "sort": null,
  "createdBy": null,
  "updatedBy": null,
  "createdAt": null,
  "updatedAt": null,
} satisfies RespDomain

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RespDomain
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


