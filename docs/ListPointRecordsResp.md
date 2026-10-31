
# ListPointRecordsResp


## Properties

Name | Type
------------ | -------------
`rows` | [Array&lt;RespPointRecord&gt;](RespPointRecord.md)
`page` | [PageResp](PageResp.md)

## Example

```typescript
import type { ListPointRecordsResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "rows": null,
  "page": null,
} satisfies ListPointRecordsResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ListPointRecordsResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


