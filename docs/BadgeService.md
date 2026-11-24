# \BadgeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBadgeDraft**](BadgeService.md#DeleteBadgeDraft) | **Post** /v1/badge/draft/delete | 
[**ListBadgeDrafts**](BadgeService.md#ListBadgeDrafts) | **Post** /v1/badge/draft/list | 
[**RenderMakerBadge**](BadgeService.md#RenderMakerBadge) | **Get** /v1/badge/maker/render | 
[**RenderPublishedBadge**](BadgeService.md#RenderPublishedBadge) | **Get** /v1/badge/render | 
[**UpsertBadgeDraft**](BadgeService.md#UpsertBadgeDraft) | **Post** /v1/badge/draft/upsert | 



## DeleteBadgeDraft

> map[string]interface{} DeleteBadgeDraft(ctx).DeleteBadgeDraftReq(deleteBadgeDraftReq).Execute()





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
	resp, r, err := apiClient.BadgeService.DeleteBadgeDraft(context.Background()).DeleteBadgeDraftReq(deleteBadgeDraftReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.DeleteBadgeDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBadgeDraft`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.DeleteBadgeDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBadgeDraftRequest struct via the builder pattern


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


## ListBadgeDrafts

> ListBadgeDraftsResp ListBadgeDrafts(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.BadgeService.ListBadgeDrafts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.ListBadgeDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBadgeDrafts`: ListBadgeDraftsResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.ListBadgeDrafts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBadgeDraftsRequest struct via the builder pattern


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


## RenderMakerBadge

> RenderMakerBadgeResp RenderMakerBadge(ctx).RenderSpecImageAssetId(renderSpecImageAssetId).RenderSpecText(renderSpecText).RenderSpecBackgroundColors(renderSpecBackgroundColors).RenderSpecTextColors(renderSpecTextColors).RenderSpecBackgroundDirection(renderSpecBackgroundDirection).RenderSpecTextDirection(renderSpecTextDirection).RenderSpecSize(renderSpecSize).RenderSpecBorder(renderSpecBorder).RenderSpecTextBarLength(renderSpecTextBarLength).RenderSpecTextSize(renderSpecTextSize).RenderSpecTextBarRadius(renderSpecTextBarRadius).RenderSpecScale(renderSpecScale).RenderSpecShadow(renderSpecShadow).RenderSpecAnimationSeconds(renderSpecAnimationSeconds).RenderSpecFontAssetId(renderSpecFontAssetId).Execute()





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
	renderSpecImageAssetId := "renderSpecImageAssetId_example" // string | Platform 已确认可用的徽章图标资源 ID。 (optional)
	renderSpecText := "renderSpecText_example" // string | 文字条内容。 (optional)
	renderSpecBackgroundColors := []string{"Inner_example"} // []string | 背景颜色列表。 (optional)
	renderSpecTextColors := []string{"Inner_example"} // []string | 文字颜色列表。 (optional)
	renderSpecBackgroundDirection := "renderSpecBackgroundDirection_example" // string | 背景渐变方向或角度。 (optional)
	renderSpecTextDirection := "renderSpecTextDirection_example" // string | 文字渐变方向或角度。 (optional)
	renderSpecSize := int32(56) // int32 | 图标边长。 (optional)
	renderSpecBorder := int32(56) // int32 | 图标边距。 (optional)
	renderSpecTextBarLength := int32(56) // int32 | 文字条长度。 (optional)
	renderSpecTextSize := int32(56) // int32 | 文字字号。 (optional)
	renderSpecTextBarRadius := int32(56) // int32 | 文字条圆角。 (optional)
	renderSpecScale := float64(1.2) // float64 | 整体缩放比例。 (optional)
	renderSpecShadow := float64(1.2) // float64 | 阴影强度。 (optional)
	renderSpecAnimationSeconds := float64(1.2) // float64 | 动画时长。 (optional)
	renderSpecFontAssetId := "renderSpecFontAssetId_example" // string | Platform 已确认可用的自定义字体资源 ID。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.RenderMakerBadge(context.Background()).RenderSpecImageAssetId(renderSpecImageAssetId).RenderSpecText(renderSpecText).RenderSpecBackgroundColors(renderSpecBackgroundColors).RenderSpecTextColors(renderSpecTextColors).RenderSpecBackgroundDirection(renderSpecBackgroundDirection).RenderSpecTextDirection(renderSpecTextDirection).RenderSpecSize(renderSpecSize).RenderSpecBorder(renderSpecBorder).RenderSpecTextBarLength(renderSpecTextBarLength).RenderSpecTextSize(renderSpecTextSize).RenderSpecTextBarRadius(renderSpecTextBarRadius).RenderSpecScale(renderSpecScale).RenderSpecShadow(renderSpecShadow).RenderSpecAnimationSeconds(renderSpecAnimationSeconds).RenderSpecFontAssetId(renderSpecFontAssetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.RenderMakerBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderMakerBadge`: RenderMakerBadgeResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.RenderMakerBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderMakerBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderSpecImageAssetId** | **string** | Platform 已确认可用的徽章图标资源 ID。 | 
 **renderSpecText** | **string** | 文字条内容。 | 
 **renderSpecBackgroundColors** | **[]string** | 背景颜色列表。 | 
 **renderSpecTextColors** | **[]string** | 文字颜色列表。 | 
 **renderSpecBackgroundDirection** | **string** | 背景渐变方向或角度。 | 
 **renderSpecTextDirection** | **string** | 文字渐变方向或角度。 | 
 **renderSpecSize** | **int32** | 图标边长。 | 
 **renderSpecBorder** | **int32** | 图标边距。 | 
 **renderSpecTextBarLength** | **int32** | 文字条长度。 | 
 **renderSpecTextSize** | **int32** | 文字字号。 | 
 **renderSpecTextBarRadius** | **int32** | 文字条圆角。 | 
 **renderSpecScale** | **float64** | 整体缩放比例。 | 
 **renderSpecShadow** | **float64** | 阴影强度。 | 
 **renderSpecAnimationSeconds** | **float64** | 动画时长。 | 
 **renderSpecFontAssetId** | **string** | Platform 已确认可用的自定义字体资源 ID。 | 

### Return type

[**RenderMakerBadgeResp**](RenderMakerBadgeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderPublishedBadge

> RenderPublishedBadgeResp RenderPublishedBadge(ctx).Code(code).Execute()





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
	code := "code_example" // string | 已发布徽章稳定编码。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BadgeService.RenderPublishedBadge(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.RenderPublishedBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderPublishedBadge`: RenderPublishedBadgeResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.RenderPublishedBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderPublishedBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | 已发布徽章稳定编码。 | 

### Return type

[**RenderPublishedBadgeResp**](RenderPublishedBadgeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBadgeDraft

> UpsertBadgeDraftResp UpsertBadgeDraft(ctx).UpsertBadgeDraftReq(upsertBadgeDraftReq).Execute()





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
	resp, r, err := apiClient.BadgeService.UpsertBadgeDraft(context.Background()).UpsertBadgeDraftReq(upsertBadgeDraftReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BadgeService.UpsertBadgeDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBadgeDraft`: UpsertBadgeDraftResp
	fmt.Fprintf(os.Stdout, "Response from `BadgeService.UpsertBadgeDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBadgeDraftRequest struct via the builder pattern


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

