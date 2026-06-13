# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](TagService.md#create) | **POST** /v1/content/tag/create | 
[**list**](TagService.md#list) | **POST** /v1/content/tag/list | 
[**update**](TagService.md#update) | **POST** /v1/content/tag/update | 



## create

> models::CreateTagReply create(create_tag_request)


创建标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_tag_request** | [**CreateTagRequest**](CreateTagRequest.md) |  | [required] |

### Return type

[**models::CreateTagReply**](CreateTag_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


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


## update

> models::UpdateTagReply update(update_tag_request)


更新标签。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_tag_request** | [**UpdateTagRequest**](UpdateTagRequest.md) |  | [required] |

### Return type

[**models::UpdateTagReply**](UpdateTag_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

