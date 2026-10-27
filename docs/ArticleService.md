# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAnswer**](ArticleService.md#AcceptAnswer) | **Post** /v1/content/article/answer/accept | 
[**Archive**](ArticleService.md#Archive) | **Post** /v1/content/article/archive | 
[**CancelPublish**](ArticleService.md#CancelPublish) | **Post** /v1/content/article/publish/cancel | 
[**Collect**](ArticleService.md#Collect) | **Post** /v1/content/article/collect | 
[**CreateDraft**](ArticleService.md#CreateDraft) | **Post** /v1/content/article/draft/create | 
[**DiscardDraft**](ArticleService.md#DiscardDraft) | **Post** /v1/content/article/draft/discard | 
[**EnterLottery**](ArticleService.md#EnterLottery) | **Post** /v1/content/article/lottery/entry | 
[**Get**](ArticleService.md#Get) | **Post** /v1/content/article/get | 
[**Like**](ArticleService.md#Like) | **Post** /v1/content/article/like | 
[**List**](ArticleService.md#List) | **Post** /v1/content/article/list | 
[**ListFollowing**](ArticleService.md#ListFollowing) | **Post** /v1/content/article/following/list | 
[**ListNearby**](ArticleService.md#ListNearby) | **Post** /v1/content/article/nearby/list | 
[**PageViewHistory**](ArticleService.md#PageViewHistory) | **Post** /v1/content/article/view-history/page | 
[**Publish**](ArticleService.md#Publish) | **Post** /v1/content/article/publish | 
[**Reward**](ArticleService.md#Reward) | **Post** /v1/content/article/reward | 
[**SetPinnedComment**](ArticleService.md#SetPinnedComment) | **Post** /v1/content/article/comment/pin | 
[**Thank**](ArticleService.md#Thank) | **Post** /v1/content/article/thank | 
[**UpdateDraft**](ArticleService.md#UpdateDraft) | **Post** /v1/content/article/draft/update | 
[**VotePoll**](ArticleService.md#VotePoll) | **Post** /v1/content/article/poll/vote | 
[**Watch**](ArticleService.md#Watch) | **Post** /v1/content/article/watch | 



## AcceptAnswer

> map[string]interface{} AcceptAnswer(ctx).AcceptArticleAnswerReq(acceptArticleAnswerReq).Execute()





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
	acceptArticleAnswerReq := *openapiclient.NewAcceptArticleAnswerReq("ArticleId_example", "CommentId_example") // AcceptArticleAnswerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.AcceptAnswer(context.Background()).AcceptArticleAnswerReq(acceptArticleAnswerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.AcceptAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptAnswer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.AcceptAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptArticleAnswerReq** | [**AcceptArticleAnswerReq**](AcceptArticleAnswerReq.md) |  | 

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


## EnterLottery

> map[string]interface{} EnterLottery(ctx).EnterArticleLotteryReq(enterArticleLotteryReq).Execute()





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
	enterArticleLotteryReq := *openapiclient.NewEnterArticleLotteryReq("ArticleId_example") // EnterArticleLotteryReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.EnterLottery(context.Background()).EnterArticleLotteryReq(enterArticleLotteryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.EnterLottery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnterLottery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.EnterLottery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterLotteryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enterArticleLotteryReq** | [**EnterArticleLotteryReq**](EnterArticleLotteryReq.md) |  | 

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


## ListFollowing

> ListFollowingArticlesResp ListFollowing(ctx).ListFollowingArticlesReq(listFollowingArticlesReq).Execute()





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
	listFollowingArticlesReq := *openapiclient.NewListFollowingArticlesReq() // ListFollowingArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.ListFollowing(context.Background()).ListFollowingArticlesReq(listFollowingArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ListFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowing`: ListFollowingArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowingArticlesReq** | [**ListFollowingArticlesReq**](ListFollowingArticlesReq.md) |  | 

### Return type

[**ListFollowingArticlesResp**](ListFollowingArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNearby

> ListNearbyArticlesResp ListNearby(ctx).ListNearbyArticlesReq(listNearbyArticlesReq).Execute()





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
	listNearbyArticlesReq := *openapiclient.NewListNearbyArticlesReq() // ListNearbyArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.ListNearby(context.Background()).ListNearbyArticlesReq(listNearbyArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ListNearby``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNearby`: ListNearbyArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ListNearby`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNearbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listNearbyArticlesReq** | [**ListNearbyArticlesReq**](ListNearbyArticlesReq.md) |  | 

### Return type

[**ListNearbyArticlesResp**](ListNearbyArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageViewHistory

> PageArticleViewHistoryResp PageViewHistory(ctx).PageArticleViewHistoryReq(pageArticleViewHistoryReq).Execute()





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
	pageArticleViewHistoryReq := *openapiclient.NewPageArticleViewHistoryReq() // PageArticleViewHistoryReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.PageViewHistory(context.Background()).PageArticleViewHistoryReq(pageArticleViewHistoryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.PageViewHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageViewHistory`: PageArticleViewHistoryResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.PageViewHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageViewHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageArticleViewHistoryReq** | [**PageArticleViewHistoryReq**](PageArticleViewHistoryReq.md) |  | 

### Return type

[**PageArticleViewHistoryResp**](PageArticleViewHistoryResp.md)

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

> RewardArticleResp Reward(ctx).RewardArticleReq(rewardArticleReq).Execute()





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
	rewardArticleReq := *openapiclient.NewRewardArticleReq("ArticleId_example") // RewardArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Reward(context.Background()).RewardArticleReq(rewardArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Reward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reward`: RewardArticleResp
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

[**RewardArticleResp**](RewardArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPinnedComment

> map[string]interface{} SetPinnedComment(ctx).SetPinnedArticleCommentReq(setPinnedArticleCommentReq).Execute()





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
	setPinnedArticleCommentReq := *openapiclient.NewSetPinnedArticleCommentReq("ArticleId_example") // SetPinnedArticleCommentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.SetPinnedComment(context.Background()).SetPinnedArticleCommentReq(setPinnedArticleCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.SetPinnedComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPinnedComment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.SetPinnedComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPinnedCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPinnedArticleCommentReq** | [**SetPinnedArticleCommentReq**](SetPinnedArticleCommentReq.md) |  | 

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
	thankArticleReq := *openapiclient.NewThankArticleReq("ArticleId_example") // ThankArticleReq | 

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


## VotePoll

> map[string]interface{} VotePoll(ctx).VoteArticlePollReq(voteArticlePollReq).Execute()





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
	voteArticlePollReq := *openapiclient.NewVoteArticlePollReq("ArticleId_example", []string{"OptionIds_example"}) // VoteArticlePollReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.VotePoll(context.Background()).VoteArticlePollReq(voteArticlePollReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.VotePoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VotePoll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.VotePoll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVotePollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voteArticlePollReq** | [**VoteArticlePollReq**](VoteArticlePollReq.md) |  | 

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


## Watch

> WatchArticleResp Watch(ctx).WatchArticleReq(watchArticleReq).Execute()





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
	watchArticleReq := *openapiclient.NewWatchArticleReq("ArticleId_example", false) // WatchArticleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Watch(context.Background()).WatchArticleReq(watchArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Watch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Watch`: WatchArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Watch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchArticleReq** | [**WatchArticleReq**](WatchArticleReq.md) |  | 

### Return type

[**WatchArticleResp**](WatchArticleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

