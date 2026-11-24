# RenderSpec

RenderSpec 是客户端 Maker 传递的完整渲染参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image_asset_id** | **string** | Platform 已确认可用的徽章图标资源 ID。 | [optional] [default to undefined]
**text** | **string** | 文字条内容。 | [optional] [default to undefined]
**background_colors** | **Array&lt;string&gt;** | 背景颜色列表。 | [optional] [default to undefined]
**text_colors** | **Array&lt;string&gt;** | 文字颜色列表。 | [optional] [default to undefined]
**background_direction** | **string** | 背景渐变方向或角度。 | [optional] [default to undefined]
**text_direction** | **string** | 文字渐变方向或角度。 | [optional] [default to undefined]
**size** | **number** | 图标边长。 | [optional] [default to undefined]
**border** | **number** | 图标边距。 | [optional] [default to undefined]
**text_bar_length** | **number** | 文字条长度。 | [optional] [default to undefined]
**text_size** | **number** | 文字字号。 | [optional] [default to undefined]
**text_bar_radius** | **number** | 文字条圆角。 | [optional] [default to undefined]
**scale** | **number** | 整体缩放比例。 | [optional] [default to undefined]
**shadow** | **number** | 阴影强度。 | [optional] [default to undefined]
**animation_seconds** | **number** | 动画时长。 | [optional] [default to undefined]
**font_asset_id** | **string** | Platform 已确认可用的自定义字体资源 ID。 | [optional] [default to undefined]

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
