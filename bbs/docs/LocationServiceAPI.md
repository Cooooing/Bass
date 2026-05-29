# \LocationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationServiceGetCurrent**](LocationServiceAPI.md#LocationServiceGetCurrent) | **Post** /v1/user/location/get-current | 
[**LocationServiceUpsertCurrent**](LocationServiceAPI.md#LocationServiceUpsertCurrent) | **Post** /v1/user/location/upsert-current | 



## LocationServiceGetCurrent

> GetCurrentLocationReply LocationServiceGetCurrent(ctx).Body(body).Execute()





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
	// response from `LocationServiceGetCurrent`: GetCurrentLocationReply
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

[**GetCurrentLocationReply**](GetCurrentLocationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationServiceUpsertCurrent

> UpsertCurrentLocationReply LocationServiceUpsertCurrent(ctx).UpsertCurrentLocationRequest(upsertCurrentLocationRequest).Execute()





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
	upsertCurrentLocationRequest := *openapiclient.NewUpsertCurrentLocationRequest() // UpsertCurrentLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationServiceAPI.LocationServiceUpsertCurrent(context.Background()).UpsertCurrentLocationRequest(upsertCurrentLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationServiceAPI.LocationServiceUpsertCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationServiceUpsertCurrent`: UpsertCurrentLocationReply
	fmt.Fprintf(os.Stdout, "Response from `LocationServiceAPI.LocationServiceUpsertCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationServiceUpsertCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertCurrentLocationRequest** | [**UpsertCurrentLocationRequest**](UpsertCurrentLocationRequest.md) |  | 

### Return type

[**UpsertCurrentLocationReply**](UpsertCurrentLocationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

