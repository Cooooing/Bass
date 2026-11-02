
# RenderSpec

RenderSpec 是客户端 Maker 传递的完整渲染参数。

## Properties

Name | Type
------------ | -------------
`imageAssetId` | string
`text` | string
`backgroundColors` | Array&lt;string&gt;
`textColors` | Array&lt;string&gt;
`backgroundDirection` | string
`textDirection` | string
`size` | number
`border` | number
`textBarLength` | number
`textSize` | number
`textBarRadius` | number
`scale` | number
`shadow` | number
`animationSeconds` | number
`fontAssetId` | string

## Example

```typescript
import type { RenderSpec } from '@bass/bbs-sdk-fetch'

// TODO: Update the object below with actual values
const example = {
  "imageAssetId": null,
  "text": null,
  "backgroundColors": null,
  "textColors": null,
  "backgroundDirection": null,
  "textDirection": null,
  "size": null,
  "border": null,
  "textBarLength": null,
  "textSize": null,
  "textBarRadius": null,
  "scale": null,
  "shadow": null,
  "animationSeconds": null,
  "fontAssetId": null,
} satisfies RenderSpec

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RenderSpec
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


