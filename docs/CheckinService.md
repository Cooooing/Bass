# \CheckinService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckIn**](CheckinService.md#CheckIn) | **Post** /v1/user/checkin/check-in | 
[**GetOverview**](CheckinService.md#GetOverview) | **Post** /v1/user/checkin/get-overview | 



## CheckIn

> CheckInResp CheckIn(ctx).Body(body).Execute()



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
	resp, r, err := apiClient.CheckinService.CheckIn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckinService.CheckIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIn`: CheckInResp
	fmt.Fprintf(os.Stdout, "Response from `CheckinService.CheckIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CheckInResp**](CheckInResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverview

> GetCheckinOverviewResp GetOverview(ctx).GetCheckinOverviewReq(getCheckinOverviewReq).Execute()



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
	getCheckinOverviewReq := *openapiclient.NewGetCheckinOverviewReq() // GetCheckinOverviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckinService.GetOverview(context.Background()).GetCheckinOverviewReq(getCheckinOverviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckinService.GetOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverview`: GetCheckinOverviewResp
	fmt.Fprintf(os.Stdout, "Response from `CheckinService.GetOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCheckinOverviewReq** | [**GetCheckinOverviewReq**](GetCheckinOverviewReq.md) |  | 

### Return type

[**GetCheckinOverviewResp**](GetCheckinOverviewResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

