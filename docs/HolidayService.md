# \HolidayService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current_holiday**](HolidayService.md#get_current_holiday) | **POST** /v1/platform/holiday/get-current | 



## get_current_holiday

> models::GetCurrentHolidayResp get_current_holiday(body)


获取当前节假日状态。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentHolidayResp**](GetCurrentHoliday_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

