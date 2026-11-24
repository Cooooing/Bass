# \BadgeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_badge_draft**](BadgeService.md#delete_badge_draft) | **POST** /v1/badge/draft/delete | 
[**list_badge_drafts**](BadgeService.md#list_badge_drafts) | **POST** /v1/badge/draft/list | 
[**render_maker_badge**](BadgeService.md#render_maker_badge) | **GET** /v1/badge/maker/render | 
[**render_published_badge**](BadgeService.md#render_published_badge) | **GET** /v1/badge/render | 
[**upsert_badge_draft**](BadgeService.md#upsert_badge_draft) | **POST** /v1/badge/draft/upsert | 



## delete_badge_draft

> serde_json::Value delete_badge_draft(delete_badge_draft_req)


DeleteDraft 删除一个未发布草稿。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**delete_badge_draft_req** | [**DeleteBadgeDraftReq**](DeleteBadgeDraftReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_badge_drafts

> models::ListBadgeDraftsResp list_badge_drafts(body)


ListDrafts 返回当前登录用户的徽章草稿。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::ListBadgeDraftsResp**](ListBadgeDrafts_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## render_maker_badge

> models::RenderMakerBadgeResp render_maker_badge(render_spec_image_asset_id, render_spec_text, render_spec_background_colors, render_spec_text_colors, render_spec_background_direction, render_spec_text_direction, render_spec_size, render_spec_border, render_spec_text_bar_length, render_spec_text_size, render_spec_text_bar_radius, render_spec_scale, render_spec_shadow, render_spec_animation_seconds, render_spec_font_asset_id)


RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**render_spec_image_asset_id** | Option<**String**> | Platform 已确认可用的徽章图标资源 ID。 |  |
**render_spec_text** | Option<**String**> | 文字条内容。 |  |
**render_spec_background_colors** | Option<[**Vec<String>**](String.md)> | 背景颜色列表。 |  |
**render_spec_text_colors** | Option<[**Vec<String>**](String.md)> | 文字颜色列表。 |  |
**render_spec_background_direction** | Option<**String**> | 背景渐变方向或角度。 |  |
**render_spec_text_direction** | Option<**String**> | 文字渐变方向或角度。 |  |
**render_spec_size** | Option<**i32**> | 图标边长。 |  |
**render_spec_border** | Option<**i32**> | 图标边距。 |  |
**render_spec_text_bar_length** | Option<**i32**> | 文字条长度。 |  |
**render_spec_text_size** | Option<**i32**> | 文字字号。 |  |
**render_spec_text_bar_radius** | Option<**i32**> | 文字条圆角。 |  |
**render_spec_scale** | Option<**f64**> | 整体缩放比例。 |  |
**render_spec_shadow** | Option<**f64**> | 阴影强度。 |  |
**render_spec_animation_seconds** | Option<**f64**> | 动画时长。 |  |
**render_spec_font_asset_id** | Option<**String**> | Platform 已确认可用的自定义字体资源 ID。 |  |

### Return type

[**models::RenderMakerBadgeResp**](RenderMakerBadge_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## render_published_badge

> models::RenderPublishedBadgeResp render_published_badge(code)


RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**code** | Option<**String**> | 已发布徽章稳定编码。 |  |

### Return type

[**models::RenderPublishedBadgeResp**](RenderPublishedBadge_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## upsert_badge_draft

> models::UpsertBadgeDraftResp upsert_badge_draft(upsert_badge_draft_req)


UpsertDraft 保存当前登录用户的徽章草稿。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**upsert_badge_draft_req** | [**UpsertBadgeDraftReq**](UpsertBadgeDraftReq.md) |  | [required] |

### Return type

[**models::UpsertBadgeDraftResp**](UpsertBadgeDraft_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

