# \DomainService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainService.md#CreateDomain) | **Post** /v1/content/domain/create | 
[**ListDomains**](DomainService.md#ListDomains) | **Post** /v1/content/domain/list | 
[**UpdateDomain**](DomainService.md#UpdateDomain) | **Post** /v1/content/domain/update | 



## CreateDomain

> CreateDomainResp CreateDomain(ctx).CreateDomainReq(createDomainReq).Execute()





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
	createDomainReq := *openapiclient.NewCreateDomainReq(*openapiclient.NewReqDomain("Code_example", "Name_example")) // CreateDomainReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainService.CreateDomain(context.Background()).CreateDomainReq(createDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: CreateDomainResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainReq** | [**CreateDomainReq**](CreateDomainReq.md) |  | 

### Return type

[**CreateDomainResp**](CreateDomainResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> ListDomainsResp ListDomains(ctx).ListDomainsReq(listDomainsReq).Execute()





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
	listDomainsReq := *openapiclient.NewListDomainsReq() // ListDomainsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainService.ListDomains(context.Background()).ListDomainsReq(listDomainsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.ListDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomains`: ListDomainsResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.ListDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listDomainsReq** | [**ListDomainsReq**](ListDomainsReq.md) |  | 

### Return type

[**ListDomainsResp**](ListDomainsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomain

> UpdateDomainResp UpdateDomain(ctx).UpdateDomainReq(updateDomainReq).Execute()





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
	updateDomainReq := *openapiclient.NewUpdateDomainReq("DomainId_example", *openapiclient.NewReqDomain("Code_example", "Name_example")) // UpdateDomainReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainService.UpdateDomain(context.Background()).UpdateDomainReq(updateDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.UpdateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomain`: UpdateDomainResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDomainReq** | [**UpdateDomainReq**](UpdateDomainReq.md) |  | 

### Return type

[**UpdateDomainResp**](UpdateDomainResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

