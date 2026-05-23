# \PrivacySettingServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivacySettingServiceGetCurrent**](PrivacySettingServiceAPI.md#PrivacySettingServiceGetCurrent) | **Post** /v1/user/privacy-setting/get-current | 
[**PrivacySettingServiceUpdate**](PrivacySettingServiceAPI.md#PrivacySettingServiceUpdate) | **Post** /v1/user/privacy-setting/update-current | 



## PrivacySettingServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentPrivacySettingReply PrivacySettingServiceGetCurrent(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivacySettingServiceGetCurrent`: CommonApiAppBbsV1UserGetCurrentPrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySettingServiceGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonApiAppBbsV1UserGetCurrentPrivacySettingReply**](CommonApiAppBbsV1UserGetCurrentPrivacySettingReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySettingServiceUpdate

> CommonApiAppBbsV1UserUpdatePrivacySettingReply PrivacySettingServiceUpdate(ctx).CommonApiAppBbsV1UserUpdatePrivacySettingRequest(commonApiAppBbsV1UserUpdatePrivacySettingRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserUpdatePrivacySettingRequest := *openapiclient.NewCommonApiAppBbsV1UserUpdatePrivacySettingRequest() // CommonApiAppBbsV1UserUpdatePrivacySettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivacySettingServiceAPI.PrivacySettingServiceUpdate(context.Background()).CommonApiAppBbsV1UserUpdatePrivacySettingRequest(commonApiAppBbsV1UserUpdatePrivacySettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingServiceAPI.PrivacySettingServiceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivacySettingServiceUpdate`: CommonApiAppBbsV1UserUpdatePrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingServiceAPI.PrivacySettingServiceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySettingServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserUpdatePrivacySettingRequest** | [**CommonApiAppBbsV1UserUpdatePrivacySettingRequest**](CommonApiAppBbsV1UserUpdatePrivacySettingRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserUpdatePrivacySettingReply**](CommonApiAppBbsV1UserUpdatePrivacySettingReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

