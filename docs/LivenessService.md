# \LivenessService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**claim_yesterday_liveness_reward**](LivenessService.md#claim_yesterday_liveness_reward) | **POST** /v1/user/liveness/yesterday-reward/claim | 
[**get_liveness**](LivenessService.md#get_liveness) | **GET** /v1/user/liveness | 



## claim_yesterday_liveness_reward

> models::ClaimYesterdayLivenessRewardResp claim_yesterday_liveness_reward(body)


服务方法。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::ClaimYesterdayLivenessRewardResp**](ClaimYesterdayLivenessReward_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_liveness

> models::GetLivenessResp get_liveness()


服务方法。

### Parameters

This endpoint does not need any parameter.

### Return type

[**models::GetLivenessResp**](GetLiveness_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

