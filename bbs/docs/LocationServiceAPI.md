# \LocationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationServiceGetCurrent**](LocationServiceAPI.md#LocationServiceGetCurrent) | **Post** /v1/user/location/get-current | 
[**LocationServiceUpsert**](LocationServiceAPI.md#LocationServiceUpsert) | **Post** /v1/user/location/upsert-current | 



## LocationServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentLocationReply LocationServiceGetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.LocationServiceAPI.LocationServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationServiceAPI.LocationServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationServiceGetCurrent`: CommonApiAppBbsV1UserGetCurrentLocationReply
	fmt.Fprintf(os.Stdout, "Response from `LocationServiceAPI.LocationServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationServiceGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonApiAppBbsV1UserGetCurrentLocationReply**](CommonApiAppBbsV1UserGetCurrentLocationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationServiceUpsert

> CommonApiAppBbsV1UserUpsertLocationReply LocationServiceUpsert(ctx).CommonApiAppBbsV1UserUpsertLocationRequest(commonApiAppBbsV1UserUpsertLocationRequest).Execute()





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
	commonApiAppBbsV1UserUpsertLocationRequest := *openapiclient.NewCommonApiAppBbsV1UserUpsertLocationRequest() // CommonApiAppBbsV1UserUpsertLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationServiceAPI.LocationServiceUpsert(context.Background()).CommonApiAppBbsV1UserUpsertLocationRequest(commonApiAppBbsV1UserUpsertLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationServiceAPI.LocationServiceUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationServiceUpsert`: CommonApiAppBbsV1UserUpsertLocationReply
	fmt.Fprintf(os.Stdout, "Response from `LocationServiceAPI.LocationServiceUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationServiceUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserUpsertLocationRequest** | [**CommonApiAppBbsV1UserUpsertLocationRequest**](CommonApiAppBbsV1UserUpsertLocationRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserUpsertLocationReply**](CommonApiAppBbsV1UserUpsertLocationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

