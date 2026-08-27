
# GetCheckinOverviewResp


## Properties

Name | Type
------------ | -------------
`records` | [Array&lt;RespRecord&gt;](RespRecord.md)
`currentStreak` | number
`longestStreak` | number

## Example

```typescript
import type { GetCheckinOverviewResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "records": null,
  "currentStreak": null,
  "longestStreak": null,
} satisfies GetCheckinOverviewResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as GetCheckinOverviewResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


