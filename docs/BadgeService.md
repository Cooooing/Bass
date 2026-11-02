# \BadgeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_draft**](BadgeService.md#delete_draft) | **POST** /v1/badge/draft/delete | 
[**list_drafts**](BadgeService.md#list_drafts) | **POST** /v1/badge/draft/list | 
[**render_maker**](BadgeService.md#render_maker) | **GET** /v1/badge/maker/render | 
[**render_published**](BadgeService.md#render_published) | **GET** /v1/badge/render | 
[**upsert_draft**](BadgeService.md#upsert_draft) | **POST** /v1/badge/draft/upsert | 



## delete_draft

> serde_json::Value delete_draft(delete_badge_draft_req)


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


## list_drafts

> models::ListBadgeDraftsResp list_drafts(body)


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


## render_maker

> models::ImageResp render_maker(render_spec_image_asset_id, render_spec_text, render_spec_background_colors, render_spec_text_colors, render_spec_background_direction, render_spec_text_direction, render_spec_size, render_spec_border, render_spec_text_bar_length, render_spec_text_size, render_spec_text_bar_radius, render_spec_scale, render_spec_shadow, render_spec_animation_seconds, render_spec_font_asset_id)


RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**render_spec_image_asset_id** | Option<**String**> |  |  |
**render_spec_text** | Option<**String**> |  |  |
**render_spec_background_colors** | Option<[**Vec<String>**](String.md)> |  |  |
**render_spec_text_colors** | Option<[**Vec<String>**](String.md)> |  |  |
**render_spec_background_direction** | Option<**String**> |  |  |
**render_spec_text_direction** | Option<**String**> |  |  |
**render_spec_size** | Option<**i32**> |  |  |
**render_spec_border** | Option<**i32**> |  |  |
**render_spec_text_bar_length** | Option<**i32**> |  |  |
**render_spec_text_size** | Option<**i32**> |  |  |
**render_spec_text_bar_radius** | Option<**i32**> |  |  |
**render_spec_scale** | Option<**f64**> |  |  |
**render_spec_shadow** | Option<**f64**> |  |  |
**render_spec_animation_seconds** | Option<**f64**> |  |  |
**render_spec_font_asset_id** | Option<**String**> |  |  |

### Return type

[**models::ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## render_published

> models::ImageResp render_published(code)


RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**code** | Option<**String**> |  |  |

### Return type

[**models::ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## upsert_draft

> models::UpsertBadgeDraftResp upsert_draft(upsert_badge_draft_req)


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

