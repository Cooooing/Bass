
# RefreshTokenResp


## Properties

Name | Type
------------ | -------------
`accessToken` | string
`refreshToken` | string
`accessTokenExpiresAt` | Date
`refreshTokenExpiresAt` | Date
`sessionExpiresAt` | Date

## Example

```typescript
import type { RefreshTokenResp } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "accessToken": null,
  "refreshToken": null,
  "accessTokenExpiresAt": null,
  "refreshTokenExpiresAt": null,
  "sessionExpiresAt": null,
} satisfies RefreshTokenResp

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RefreshTokenResp
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


