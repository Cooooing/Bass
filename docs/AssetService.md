# \AssetService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**complete_markdown_image_upload**](AssetService.md#complete_markdown_image_upload) | **POST** /v1/asset/markdown-image/complete | 
[**prepare_markdown_image_upload**](AssetService.md#prepare_markdown_image_upload) | **POST** /v1/asset/markdown-image/prepare | 



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

