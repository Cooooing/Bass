
# ReqDomainQuery


## Properties

Name | Type
------------ | -------------
`ids` | Array&lt;string&gt;
`code` | string
`name` | string
`description` | string
`status` | string
`url` | string
`icon` | string
`isNav` | boolean

## Example

```typescript
import type { ReqDomainQuery } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "ids": null,
  "code": null,
  "name": null,
  "description": null,
  "status": null,
  "url": null,
  "icon": null,
  "isNav": null,
} satisfies ReqDomainQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ReqDomainQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


