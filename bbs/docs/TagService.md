# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list**](TagService.md#list) | **POST** /v1/content/tag/list | 



## list

> models::ListTagsReply list(list_tags_request)


分页查询标签列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_tags_request** | [**ListTagsRequest**](ListTagsRequest.md) |  | [required] |

### Return type

[**models::ListTagsReply**](ListTags_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

