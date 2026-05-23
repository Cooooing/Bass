
# CommonApiAppBbsV1UserPageRequest

分页请求

## Properties

Name | Type
------------ | -------------
`current` | string
`pageSize` | string

## Example

```typescript
import type { CommonApiAppBbsV1UserPageRequest } from '@bass/bbs-sdk'

// TODO: Update the object below with actual values
const example = {
  "current": null,
  "pageSize": null,
} satisfies CommonApiAppBbsV1UserPageRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as CommonApiAppBbsV1UserPageRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


