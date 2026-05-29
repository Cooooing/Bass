
# TagQuery


## Properties

Name | Type
------------ | -------------
`ids` | Array&lt;string&gt;
`name` | string
`names` | Array&lt;string&gt;
`description` | string
`status` | string
`domainId` | string

## Example

```typescript
import type { TagQuery } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "ids": null,
  "name": null,
  "names": null,
  "description": null,
  "status": null,
  "domainId": null,
} satisfies TagQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as TagQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


