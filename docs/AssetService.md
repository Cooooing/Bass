# \AssetService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**complete_badge_font_upload**](AssetService.md#complete_badge_font_upload) | **POST** /v1/asset/badge-font/complete | 
[**complete_badge_image_upload**](AssetService.md#complete_badge_image_upload) | **POST** /v1/asset/badge-image/complete | 
[**complete_markdown_image_upload**](AssetService.md#complete_markdown_image_upload) | **POST** /v1/asset/markdown-image/complete | 
[**prepare_badge_font_upload**](AssetService.md#prepare_badge_font_upload) | **POST** /v1/asset/badge-font/prepare | 
[**prepare_badge_image_upload**](AssetService.md#prepare_badge_image_upload) | **POST** /v1/asset/badge-image/prepare | 
[**prepare_markdown_image_upload**](AssetService.md#prepare_markdown_image_upload) | **POST** /v1/asset/markdown-image/prepare | 



## complete_badge_font_upload

> models::CompleteBadgeFontUploadResp complete_badge_font_upload(complete_badge_font_upload_req)


查询已由对象存储回调确认的徽章 Maker 自定义字体。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**complete_badge_font_upload_req** | [**CompleteBadgeFontUploadReq**](CompleteBadgeFontUploadReq.md) |  | [required] |

### Return type

[**models::CompleteBadgeFontUploadResp**](CompleteBadgeFontUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## complete_badge_image_upload

> models::CompleteBadgeImageUploadResp complete_badge_image_upload(complete_badge_image_upload_req)


查询已由对象存储回调确认的徽章 Maker 图片。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**complete_badge_image_upload_req** | [**CompleteBadgeImageUploadReq**](CompleteBadgeImageUploadReq.md) |  | [required] |

### Return type

[**models::CompleteBadgeImageUploadResp**](CompleteBadgeImageUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## complete_markdown_image_upload

> models::CompleteMarkdownImageUploadResp complete_markdown_image_upload(complete_markdown_image_upload_req)


查询已由对象存储回调确认的 Markdown 图片。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**complete_markdown_image_upload_req** | [**CompleteMarkdownImageUploadReq**](CompleteMarkdownImageUploadReq.md) |  | [required] |

### Return type

[**models::CompleteMarkdownImageUploadResp**](CompleteMarkdownImageUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## prepare_badge_font_upload

> models::PrepareBadgeFontUploadResp prepare_badge_font_upload(prepare_badge_font_upload_req)


申请徽章 Maker 自定义字体的内容寻址直传凭证。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**prepare_badge_font_upload_req** | [**PrepareBadgeFontUploadReq**](PrepareBadgeFontUploadReq.md) |  | [required] |

### Return type

[**models::PrepareBadgeFontUploadResp**](PrepareBadgeFontUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## prepare_badge_image_upload

> models::PrepareBadgeImageUploadResp prepare_badge_image_upload(prepare_badge_image_upload_req)


申请徽章 Maker 图片的内容寻址直传凭证。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**prepare_badge_image_upload_req** | [**PrepareBadgeImageUploadReq**](PrepareBadgeImageUploadReq.md) |  | [required] |

### Return type

[**models::PrepareBadgeImageUploadResp**](PrepareBadgeImageUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## prepare_markdown_image_upload

> models::PrepareMarkdownImageUploadResp prepare_markdown_image_upload(prepare_markdown_image_upload_req)


申请 Markdown 图片的内容寻址直传凭证。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**prepare_markdown_image_upload_req** | [**PrepareMarkdownImageUploadReq**](PrepareMarkdownImageUploadReq.md) |  | [required] |

### Return type

[**models::PrepareMarkdownImageUploadResp**](PrepareMarkdownImageUpload_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

