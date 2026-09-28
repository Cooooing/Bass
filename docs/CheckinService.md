# \CheckinService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**check_in**](CheckinService.md#check_in) | **POST** /v1/user/checkin/check-in | 
[**get_overview**](CheckinService.md#get_overview) | **POST** /v1/user/checkin/get-overview | 



## check_in

> models::CheckInResp check_in(body)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::CheckInResp**](CheckIn_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_overview

> models::GetCheckinOverviewResp get_overview(get_checkin_overview_req)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_checkin_overview_req** | [**GetCheckinOverviewReq**](GetCheckinOverviewReq.md) |  | [required] |

### Return type

[**models::GetCheckinOverviewResp**](GetCheckinOverview_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

