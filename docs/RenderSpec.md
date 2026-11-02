# RenderSpec

RenderSpec 是客户端 Maker 传递的完整渲染参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image_asset_id** | **string** |  | [optional] [default to undefined]
**text** | **string** |  | [optional] [default to undefined]
**background_colors** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**text_colors** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**background_direction** | **string** |  | [optional] [default to undefined]
**text_direction** | **string** |  | [optional] [default to undefined]
**size** | **number** |  | [optional] [default to undefined]
**border** | **number** |  | [optional] [default to undefined]
**text_bar_length** | **number** |  | [optional] [default to undefined]
**text_size** | **number** |  | [optional] [default to undefined]
**text_bar_radius** | **number** |  | [optional] [default to undefined]
**scale** | **number** |  | [optional] [default to undefined]
**shadow** | **number** |  | [optional] [default to undefined]
**animation_seconds** | **number** |  | [optional] [default to undefined]
**font_asset_id** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { RenderSpec } from '@bass/bbs-sdk-axios';

const instance: RenderSpec = {
    image_asset_id,
    text,
    background_colors,
    text_colors,
    background_direction,
    text_direction,
    size,
    border,
    text_bar_length,
    text_size,
    text_bar_radius,
    scale,
    shadow,
    animation_seconds,
    font_asset_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
