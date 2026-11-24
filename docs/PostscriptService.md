# \PostscriptService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_postscript**](PostscriptService.md#add_postscript) | **POST** /v1/content/postscript/add | 
[**list_postscripts**](PostscriptService.md#list_postscripts) | **POST** /v1/content/postscript/list | 



## add_postscript

> models::AddPostscriptResp add_postscript(add_postscript_req)


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


## list_postscripts

> models::ListPostscriptsResp list_postscripts(list_postscripts_req)


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

