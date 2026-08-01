# \PostscriptService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add**](PostscriptService.md#add) | **POST** /v1/content/postscript/add | 
[**list**](PostscriptService.md#list) | **POST** /v1/content/postscript/list | 



## add

> models::AddPostscriptResp add(add_postscript_req)


添加文章附言。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**add_postscript_req** | [**AddPostscriptReq**](AddPostscriptReq.md) |  | [required] |

### Return type

[**models::AddPostscriptResp**](AddPostscript_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListPostscriptsResp list(list_postscripts_req)


查询文章附言列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_postscripts_req** | [**ListPostscriptsReq**](ListPostscriptsReq.md) |  | [required] |

### Return type

[**models::ListPostscriptsResp**](ListPostscripts_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

