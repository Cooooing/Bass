# \LocationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrent**](LocationService.md#GetCurrent) | **Post** /v1/user/location/get-current | 
[**UpsertCurrent**](LocationService.md#UpsertCurrent) | **Post** /v1/user/location/upsert-current | 



## GetCurrent

> GetCurrentLocationResp GetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.LocationService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentLocationResp
	fmt.Fprintf(os.Stdout, "Response from `LocationService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentLocationResp**](GetCurrentLocationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCurrent

> UpsertCurrentLocationResp UpsertCurrent(ctx).UpsertCurrentLocationReq(upsertCurrentLocationReq).Execute()





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
	upsertCurrentLocationReq := *openapiclient.NewUpsertCurrentLocationReq() // UpsertCurrentLocationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationService.UpsertCurrent(context.Background()).UpsertCurrentLocationReq(upsertCurrentLocationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationService.UpsertCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCurrent`: UpsertCurrentLocationResp
	fmt.Fprintf(os.Stdout, "Response from `LocationService.UpsertCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertCurrentLocationReq** | [**UpsertCurrentLocationReq**](UpsertCurrentLocationReq.md) |  | 

### Return type

[**UpsertCurrentLocationResp**](UpsertCurrentLocationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

