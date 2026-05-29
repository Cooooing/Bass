# \DomainServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainServiceList**](DomainServiceAPI.md#DomainServiceList) | **Post** /v1/content/domain/list | 



## DomainServiceList

> ListDomainsReply DomainServiceList(ctx).ListDomainsRequest(listDomainsRequest).Execute()



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
	listDomainsRequest := *openapiclient.NewListDomainsRequest() // ListDomainsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainServiceAPI.DomainServiceList(context.Background()).ListDomainsRequest(listDomainsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainServiceAPI.DomainServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainServiceList`: ListDomainsReply
	fmt.Fprintf(os.Stdout, "Response from `DomainServiceAPI.DomainServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listDomainsRequest** | [**ListDomainsRequest**](ListDomainsRequest.md) |  | 

### Return type

[**ListDomainsReply**](ListDomainsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

