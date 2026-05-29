# \LocationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current**](LocationService.md#get_current) | **POST** /v1/user/location/get-current | 
[**upsert_current**](LocationService.md#upsert_current) | **POST** /v1/user/location/upsert-current | 



## get_current

> models::GetCurrentLocationReply get_current(body)


获取当前登录账号的地理资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentLocationReply**](GetCurrentLocation_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## upsert_current

> models::UpsertCurrentLocationReply upsert_current(upsert_current_location_request)


更新当前登录账号的地理资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**upsert_current_location_request** | [**UpsertCurrentLocationRequest**](UpsertCurrentLocationRequest.md) |  | [required] |

### Return type

[**models::UpsertCurrentLocationReply**](UpsertCurrentLocation_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

