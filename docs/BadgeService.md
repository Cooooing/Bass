# \BadgeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDraft**](BadgeService.md#DeleteDraft) | **Post** /v1/badge/draft/delete | 
[**ListDrafts**](BadgeService.md#ListDrafts) | **Post** /v1/badge/draft/list | 
[**RenderMaker**](BadgeService.md#RenderMaker) | **Get** /v1/badge/maker/render | 
[**RenderPublished**](BadgeService.md#RenderPublished) | **Get** /v1/badge/render | 
[**UpsertDraft**](BadgeService.md#UpsertDraft) | **Post** /v1/badge/draft/upsert | 



## DeleteDraft

> map[string]interface{} DeleteDraft(ctx).DeleteBadgeDraftReq(deleteBadgeDraftReq).Execute()





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
	deleteBadgeDraftReq := *openapiclient.NewDeleteBadgeDraftReq() // DeleteBadgeDraftReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.DeleteDraft(context.Background()).DeleteBadgeDraftReq(deleteBadgeDraftReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.DeleteDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDraft`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.DeleteDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBadgeDraftReq** | [**DeleteBadgeDraftReq**](DeleteBadgeDraftReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDrafts

> ListBadgeDraftsResp ListDrafts(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.BadgeService.ListDrafts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.ListDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDrafts`: ListBadgeDraftsResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.ListDrafts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDraftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListBadgeDraftsResp**](ListBadgeDraftsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderMaker

> ImageResp RenderMaker(ctx).RenderSpecImageAssetId(renderSpecImageAssetId).RenderSpecText(renderSpecText).RenderSpecBackgroundColors(renderSpecBackgroundColors).RenderSpecTextColors(renderSpecTextColors).RenderSpecBackgroundDirection(renderSpecBackgroundDirection).RenderSpecTextDirection(renderSpecTextDirection).RenderSpecSize(renderSpecSize).RenderSpecBorder(renderSpecBorder).RenderSpecTextBarLength(renderSpecTextBarLength).RenderSpecTextSize(renderSpecTextSize).RenderSpecTextBarRadius(renderSpecTextBarRadius).RenderSpecScale(renderSpecScale).RenderSpecShadow(renderSpecShadow).RenderSpecAnimationSeconds(renderSpecAnimationSeconds).RenderSpecFontAssetId(renderSpecFontAssetId).Execute()





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
	renderSpecImageAssetId := "renderSpecImageAssetId_example" // string |  (optional)
	renderSpecText := "renderSpecText_example" // string |  (optional)
	renderSpecBackgroundColors := []string{"Inner_example"} // []string |  (optional)
	renderSpecTextColors := []string{"Inner_example"} // []string |  (optional)
	renderSpecBackgroundDirection := "renderSpecBackgroundDirection_example" // string |  (optional)
	renderSpecTextDirection := "renderSpecTextDirection_example" // string |  (optional)
	renderSpecSize := int32(56) // int32 |  (optional)
	renderSpecBorder := int32(56) // int32 |  (optional)
	renderSpecTextBarLength := int32(56) // int32 |  (optional)
	renderSpecTextSize := int32(56) // int32 |  (optional)
	renderSpecTextBarRadius := int32(56) // int32 |  (optional)
	renderSpecScale := float64(1.2) // float64 |  (optional)
	renderSpecShadow := float64(1.2) // float64 |  (optional)
	renderSpecAnimationSeconds := float64(1.2) // float64 |  (optional)
	renderSpecFontAssetId := "renderSpecFontAssetId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.RenderMaker(context.Background()).RenderSpecImageAssetId(renderSpecImageAssetId).RenderSpecText(renderSpecText).RenderSpecBackgroundColors(renderSpecBackgroundColors).RenderSpecTextColors(renderSpecTextColors).RenderSpecBackgroundDirection(renderSpecBackgroundDirection).RenderSpecTextDirection(renderSpecTextDirection).RenderSpecSize(renderSpecSize).RenderSpecBorder(renderSpecBorder).RenderSpecTextBarLength(renderSpecTextBarLength).RenderSpecTextSize(renderSpecTextSize).RenderSpecTextBarRadius(renderSpecTextBarRadius).RenderSpecScale(renderSpecScale).RenderSpecShadow(renderSpecShadow).RenderSpecAnimationSeconds(renderSpecAnimationSeconds).RenderSpecFontAssetId(renderSpecFontAssetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.RenderMaker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderMaker`: ImageResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.RenderMaker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderMakerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderSpecImageAssetId** | **string** |  | 
 **renderSpecText** | **string** |  | 
 **renderSpecBackgroundColors** | **[]string** |  | 
 **renderSpecTextColors** | **[]string** |  | 
 **renderSpecBackgroundDirection** | **string** |  | 
 **renderSpecTextDirection** | **string** |  | 
 **renderSpecSize** | **int32** |  | 
 **renderSpecBorder** | **int32** |  | 
 **renderSpecTextBarLength** | **int32** |  | 
 **renderSpecTextSize** | **int32** |  | 
 **renderSpecTextBarRadius** | **int32** |  | 
 **renderSpecScale** | **float64** |  | 
 **renderSpecShadow** | **float64** |  | 
 **renderSpecAnimationSeconds** | **float64** |  | 
 **renderSpecFontAssetId** | **string** |  | 

### Return type

[**ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderPublished

> ImageResp RenderPublished(ctx).Code(code).Execute()





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
	code := "code_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.RenderPublished(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.RenderPublished``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderPublished`: ImageResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.RenderPublished`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 

### Return type

[**ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertDraft

> UpsertBadgeDraftResp UpsertDraft(ctx).UpsertBadgeDraftReq(upsertBadgeDraftReq).Execute()





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
	upsertBadgeDraftReq := *openapiclient.NewUpsertBadgeDraftReq() // UpsertBadgeDraftReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.UpsertDraft(context.Background()).UpsertBadgeDraftReq(upsertBadgeDraftReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.UpsertDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertDraft`: UpsertBadgeDraftResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.UpsertDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertBadgeDraftReq** | [**UpsertBadgeDraftReq**](UpsertBadgeDraftReq.md) |  | 

### Return type

[**UpsertBadgeDraftResp**](UpsertBadgeDraftResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

