# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](ArticleService.md#Archive) | **Post** /v1/content/article/archive | 
[**CancelPublish**](ArticleService.md#CancelPublish) | **Post** /v1/content/article/publish/cancel | 
[**Collect**](ArticleService.md#Collect) | **Post** /v1/content/article/collect | 
[**CreateDraft**](ArticleService.md#CreateDraft) | **Post** /v1/content/article/draft/create | 
[**DiscardDraft**](ArticleService.md#DiscardDraft) | **Post** /v1/content/article/draft/discard | 
[**Get**](ArticleService.md#Get) | **Post** /v1/content/article/get | 
[**Like**](ArticleService.md#Like) | **Post** /v1/content/article/like | 
[**List**](ArticleService.md#List) | **Post** /v1/content/article/list | 
[**Publish**](ArticleService.md#Publish) | **Post** /v1/content/article/publish | 
[**Reward**](ArticleService.md#Reward) | **Post** /v1/content/article/reward | 
[**SchedulePublish**](ArticleService.md#SchedulePublish) | **Post** /v1/content/article/publish/schedule | 
[**Thank**](ArticleService.md#Thank) | **Post** /v1/content/article/thank | 
[**UpdateDraft**](ArticleService.md#UpdateDraft) | **Post** /v1/content/article/draft/update | 



## Archive

> map[string]interface{} Archive(ctx).ArchiveArticleReq(archiveArticleReq).Execute()





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
	archiveArticleReq := *openapiclient.NewArchiveArticleReq("ArticleId_example") // ArchiveArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Archive(context.Background()).ArchiveArticleReq(archiveArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Archive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Archive`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Archive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveArticleReq** | [**ArchiveArticleReq**](ArchiveArticleReq.md) |  | 

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


## CancelPublish

> map[string]interface{} CancelPublish(ctx).CancelPublishArticleReq(cancelPublishArticleReq).Execute()





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
	cancelPublishArticleReq := *openapiclient.NewCancelPublishArticleReq("ArticleId_example") // CancelPublishArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.CancelPublish(context.Background()).CancelPublishArticleReq(cancelPublishArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CancelPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPublish`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CancelPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelPublishArticleReq** | [**CancelPublishArticleReq**](CancelPublishArticleReq.md) |  | 

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


## Collect

> CollectArticleResp Collect(ctx).CollectArticleReq(collectArticleReq).Execute()





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
	collectArticleReq := *openapiclient.NewCollectArticleReq("ArticleId_example", false) // CollectArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Collect(context.Background()).CollectArticleReq(collectArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Collect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Collect`: CollectArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Collect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectArticleReq** | [**CollectArticleReq**](CollectArticleReq.md) |  | 

### Return type

[**CollectArticleResp**](CollectArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDraft

> CreateDraftArticleResp CreateDraft(ctx).CreateDraftArticleReq(createDraftArticleReq).Execute()





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
	createDraftArticleReq := *openapiclient.NewCreateDraftArticleReq(*openapiclient.NewReqArticle("Title_example", "Content_example", "Type_example")) // CreateDraftArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.CreateDraft(context.Background()).CreateDraftArticleReq(createDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CreateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDraft`: CreateDraftArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CreateDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDraftArticleReq** | [**CreateDraftArticleReq**](CreateDraftArticleReq.md) |  | 

### Return type

[**CreateDraftArticleResp**](CreateDraftArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscardDraft

> map[string]interface{} DiscardDraft(ctx).DiscardDraftArticleReq(discardDraftArticleReq).Execute()





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
	discardDraftArticleReq := *openapiclient.NewDiscardDraftArticleReq("ArticleId_example") // DiscardDraftArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.DiscardDraft(context.Background()).DiscardDraftArticleReq(discardDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.DiscardDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscardDraft`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.DiscardDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscardDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discardDraftArticleReq** | [**DiscardDraftArticleReq**](DiscardDraftArticleReq.md) |  | 

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


## Get

> GetArticleResp Get(ctx).GetArticleReq(getArticleReq).Execute()





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
	getArticleReq := *openapiclient.NewGetArticleReq("ArticleId_example") // GetArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Get(context.Background()).GetArticleReq(getArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: GetArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getArticleReq** | [**GetArticleReq**](GetArticleReq.md) |  | 

### Return type

[**GetArticleResp**](GetArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Like

> LikeArticleResp Like(ctx).LikeArticleReq(likeArticleReq).Execute()





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
	likeArticleReq := *openapiclient.NewLikeArticleReq("ArticleId_example", false) // LikeArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Like(context.Background()).LikeArticleReq(likeArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Like``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Like`: LikeArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Like`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeArticleReq** | [**LikeArticleReq**](LikeArticleReq.md) |  | 

### Return type

[**LikeArticleResp**](LikeArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListArticlesResp List(ctx).ListArticlesReq(listArticlesReq).Execute()





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
	listArticlesReq := *openapiclient.NewListArticlesReq() // ListArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.List(context.Background()).ListArticlesReq(listArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listArticlesReq** | [**ListArticlesReq**](ListArticlesReq.md) |  | 

### Return type

[**ListArticlesResp**](ListArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Publish

> map[string]interface{} Publish(ctx).PublishArticleReq(publishArticleReq).Execute()





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
	publishArticleReq := *openapiclient.NewPublishArticleReq("ArticleId_example") // PublishArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Publish(context.Background()).PublishArticleReq(publishArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Publish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Publish`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Publish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishArticleReq** | [**PublishArticleReq**](PublishArticleReq.md) |  | 

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


## Reward

> map[string]interface{} Reward(ctx).RewardArticleReq(rewardArticleReq).Execute()





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
	rewardArticleReq := *openapiclient.NewRewardArticleReq("ArticleId_example", int32(123)) // RewardArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Reward(context.Background()).RewardArticleReq(rewardArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Reward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reward`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Reward`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRewardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardArticleReq** | [**RewardArticleReq**](RewardArticleReq.md) |  | 

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


## SchedulePublish

> map[string]interface{} SchedulePublish(ctx).SchedulePublishArticleReq(schedulePublishArticleReq).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	schedulePublishArticleReq := *openapiclient.NewSchedulePublishArticleReq("ArticleId_example", time.Now()) // SchedulePublishArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.SchedulePublish(context.Background()).SchedulePublishArticleReq(schedulePublishArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.SchedulePublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulePublish`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.SchedulePublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulePublishArticleReq** | [**SchedulePublishArticleReq**](SchedulePublishArticleReq.md) |  | 

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


## Thank

> ThankArticleResp Thank(ctx).ThankArticleReq(thankArticleReq).Execute()





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
	thankArticleReq := *openapiclient.NewThankArticleReq("ArticleId_example", false) // ThankArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Thank(context.Background()).ThankArticleReq(thankArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Thank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Thank`: ThankArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Thank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thankArticleReq** | [**ThankArticleReq**](ThankArticleReq.md) |  | 

### Return type

[**ThankArticleResp**](ThankArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDraft

> UpdateDraftArticleResp UpdateDraft(ctx).UpdateDraftArticleReq(updateDraftArticleReq).Execute()





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
	updateDraftArticleReq := *openapiclient.NewUpdateDraftArticleReq("ArticleId_example", *openapiclient.NewReqArticle("Title_example", "Content_example", "Type_example")) // UpdateDraftArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.UpdateDraft(context.Background()).UpdateDraftArticleReq(updateDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.UpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraft`: UpdateDraftArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.UpdateDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDraftArticleReq** | [**UpdateDraftArticleReq**](UpdateDraftArticleReq.md) |  | 

### Return type

[**UpdateDraftArticleResp**](UpdateDraftArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

