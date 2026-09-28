# \DomainService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DomainService.md#Create) | **Post** /v1/content/domain/create | 
[**List**](DomainService.md#List) | **Post** /v1/content/domain/list | 
[**Update**](DomainService.md#Update) | **Post** /v1/content/domain/update | 



## Create

> CreateDomainResp Create(ctx).CreateDomainReq(createDomainReq).Execute()





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
	resp, r, err := apiClient.DomainService.Create(context.Background()).CreateDomainReq(createDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDomainResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


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


## List

> ListDomainsResp List(ctx).ListDomainsReq(listDomainsReq).Execute()





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
	resp, r, err := apiClient.DomainService.List(context.Background()).ListDomainsReq(listDomainsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDomainsResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


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


## Update

> UpdateDomainResp Update(ctx).UpdateDomainReq(updateDomainReq).Execute()





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
	resp, r, err := apiClient.DomainService.Update(context.Background()).UpdateDomainReq(updateDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainService.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDomainResp
	fmt.Fprintf(os.Stdout, "Response from `DomainService.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


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

