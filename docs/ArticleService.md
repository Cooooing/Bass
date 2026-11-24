# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptArticleAnswer**](ArticleService.md#AcceptArticleAnswer) | **Post** /v1/content/article/answer/accept | 
[**ArchiveArticle**](ArticleService.md#ArchiveArticle) | **Post** /v1/content/article/archive | 
[**CancelPublishArticle**](ArticleService.md#CancelPublishArticle) | **Post** /v1/content/article/publish/cancel | 
[**CollectArticle**](ArticleService.md#CollectArticle) | **Post** /v1/content/article/collect | 
[**CountCollectedArticles**](ArticleService.md#CountCollectedArticles) | **Post** /v1/content/article/collect/count | 
[**CreateDraftArticle**](ArticleService.md#CreateDraftArticle) | **Post** /v1/content/article/draft/create | 
[**DiscardDraftArticle**](ArticleService.md#DiscardDraftArticle) | **Post** /v1/content/article/draft/discard | 
[**EnterArticleLottery**](ArticleService.md#EnterArticleLottery) | **Post** /v1/content/article/lottery/entry | 
[**GetArticle**](ArticleService.md#GetArticle) | **Post** /v1/content/article/get | 
[**LikeArticle**](ArticleService.md#LikeArticle) | **Post** /v1/content/article/like | 
[**ListArticles**](ArticleService.md#ListArticles) | **Post** /v1/content/article/list | 
[**ListFollowingArticles**](ArticleService.md#ListFollowingArticles) | **Post** /v1/content/article/following/list | 
[**ListNearbyArticles**](ArticleService.md#ListNearbyArticles) | **Post** /v1/content/article/nearby/list | 
[**PageArticleViewHistory**](ArticleService.md#PageArticleViewHistory) | **Post** /v1/content/article/view-history/page | 
[**PageCollectedArticles**](ArticleService.md#PageCollectedArticles) | **Post** /v1/content/article/collected/page | 
[**PageWatchedArticles**](ArticleService.md#PageWatchedArticles) | **Post** /v1/content/article/watched/page | 
[**PublishArticle**](ArticleService.md#PublishArticle) | **Post** /v1/content/article/publish | 
[**RewardArticle**](ArticleService.md#RewardArticle) | **Post** /v1/content/article/reward | 
[**SetPinnedArticleComment**](ArticleService.md#SetPinnedArticleComment) | **Post** /v1/content/article/comment/pin | 
[**ThankArticle**](ArticleService.md#ThankArticle) | **Post** /v1/content/article/thank | 
[**UpdateDraftArticle**](ArticleService.md#UpdateDraftArticle) | **Post** /v1/content/article/draft/update | 
[**VoteArticlePoll**](ArticleService.md#VoteArticlePoll) | **Post** /v1/content/article/poll/vote | 
[**WatchArticle**](ArticleService.md#WatchArticle) | **Post** /v1/content/article/watch | 



## AcceptArticleAnswer

> map[string]interface{} AcceptArticleAnswer(ctx).AcceptArticleAnswerReq(acceptArticleAnswerReq).Execute()





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
	resp, r, err := apiClient.ArticleService.AcceptArticleAnswer(context.Background()).AcceptArticleAnswerReq(acceptArticleAnswerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.AcceptArticleAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptArticleAnswer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.AcceptArticleAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptArticleAnswerRequest struct via the builder pattern


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


## ArchiveArticle

> map[string]interface{} ArchiveArticle(ctx).ArchiveArticleReq(archiveArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.ArchiveArticle(context.Background()).ArchiveArticleReq(archiveArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ArchiveArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ArchiveArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveArticleRequest struct via the builder pattern


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


## CancelPublishArticle

> map[string]interface{} CancelPublishArticle(ctx).CancelPublishArticleReq(cancelPublishArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.CancelPublishArticle(context.Background()).CancelPublishArticleReq(cancelPublishArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CancelPublishArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPublishArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CancelPublishArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPublishArticleRequest struct via the builder pattern


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


## CollectArticle

> CollectArticleResp CollectArticle(ctx).CollectArticleReq(collectArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.CollectArticle(context.Background()).CollectArticleReq(collectArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CollectArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectArticle`: CollectArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CollectArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectArticleRequest struct via the builder pattern


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


## CountCollectedArticles

> CountCollectedArticlesResp CountCollectedArticles(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.ArticleService.CountCollectedArticles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CountCollectedArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCollectedArticles`: CountCollectedArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CountCollectedArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCollectedArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CountCollectedArticlesResp**](CountCollectedArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDraftArticle

> CreateDraftArticleResp CreateDraftArticle(ctx).CreateDraftArticleReq(createDraftArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.CreateDraftArticle(context.Background()).CreateDraftArticleReq(createDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.CreateDraftArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDraftArticle`: CreateDraftArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.CreateDraftArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDraftArticleRequest struct via the builder pattern


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


## DiscardDraftArticle

> map[string]interface{} DiscardDraftArticle(ctx).DiscardDraftArticleReq(discardDraftArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.DiscardDraftArticle(context.Background()).DiscardDraftArticleReq(discardDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.DiscardDraftArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscardDraftArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.DiscardDraftArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscardDraftArticleRequest struct via the builder pattern


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


## EnterArticleLottery

> map[string]interface{} EnterArticleLottery(ctx).EnterArticleLotteryReq(enterArticleLotteryReq).Execute()





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
	resp, r, err := apiClient.ArticleService.EnterArticleLottery(context.Background()).EnterArticleLotteryReq(enterArticleLotteryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.EnterArticleLottery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnterArticleLottery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.EnterArticleLottery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterArticleLotteryRequest struct via the builder pattern


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


## GetArticle

> GetArticleResp GetArticle(ctx).GetArticleReq(getArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.GetArticle(context.Background()).GetArticleReq(getArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.GetArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArticle`: GetArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.GetArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArticleRequest struct via the builder pattern


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


## LikeArticle

> LikeArticleResp LikeArticle(ctx).LikeArticleReq(likeArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.LikeArticle(context.Background()).LikeArticleReq(likeArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.LikeArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LikeArticle`: LikeArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.LikeArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeArticleRequest struct via the builder pattern


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


## ListArticles

> ListArticlesResp ListArticles(ctx).ListArticlesReq(listArticlesReq).Execute()





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
	resp, r, err := apiClient.ArticleService.ListArticles(context.Background()).ListArticlesReq(listArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ListArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArticles`: ListArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ListArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArticlesRequest struct via the builder pattern


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


## ListFollowingArticles

> ListFollowingArticlesResp ListFollowingArticles(ctx).ListFollowingArticlesReq(listFollowingArticlesReq).Execute()





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
	resp, r, err := apiClient.ArticleService.ListFollowingArticles(context.Background()).ListFollowingArticlesReq(listFollowingArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ListFollowingArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowingArticles`: ListFollowingArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ListFollowingArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingArticlesRequest struct via the builder pattern


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


## ListNearbyArticles

> ListNearbyArticlesResp ListNearbyArticles(ctx).ListNearbyArticlesReq(listNearbyArticlesReq).Execute()





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
	resp, r, err := apiClient.ArticleService.ListNearbyArticles(context.Background()).ListNearbyArticlesReq(listNearbyArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ListNearbyArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNearbyArticles`: ListNearbyArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ListNearbyArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNearbyArticlesRequest struct via the builder pattern


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


## PageArticleViewHistory

> PageArticleViewHistoryResp PageArticleViewHistory(ctx).PageArticleViewHistoryReq(pageArticleViewHistoryReq).Execute()





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
	resp, r, err := apiClient.ArticleService.PageArticleViewHistory(context.Background()).PageArticleViewHistoryReq(pageArticleViewHistoryReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.PageArticleViewHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageArticleViewHistory`: PageArticleViewHistoryResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.PageArticleViewHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageArticleViewHistoryRequest struct via the builder pattern


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


## PageCollectedArticles

> PageCollectedArticlesResp PageCollectedArticles(ctx).PageCollectedArticlesReq(pageCollectedArticlesReq).Execute()





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
	pageCollectedArticlesReq := *openapiclient.NewPageCollectedArticlesReq() // PageCollectedArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.PageCollectedArticles(context.Background()).PageCollectedArticlesReq(pageCollectedArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.PageCollectedArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageCollectedArticles`: PageCollectedArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.PageCollectedArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageCollectedArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageCollectedArticlesReq** | [**PageCollectedArticlesReq**](PageCollectedArticlesReq.md) |  | 

### Return type

[**PageCollectedArticlesResp**](PageCollectedArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageWatchedArticles

> PageWatchedArticlesResp PageWatchedArticles(ctx).PageWatchedArticlesReq(pageWatchedArticlesReq).Execute()





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
	pageWatchedArticlesReq := *openapiclient.NewPageWatchedArticlesReq() // PageWatchedArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.PageWatchedArticles(context.Background()).PageWatchedArticlesReq(pageWatchedArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.PageWatchedArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageWatchedArticles`: PageWatchedArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.PageWatchedArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageWatchedArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageWatchedArticlesReq** | [**PageWatchedArticlesReq**](PageWatchedArticlesReq.md) |  | 

### Return type

[**PageWatchedArticlesResp**](PageWatchedArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishArticle

> map[string]interface{} PublishArticle(ctx).PublishArticleReq(publishArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.PublishArticle(context.Background()).PublishArticleReq(publishArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.PublishArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.PublishArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishArticleRequest struct via the builder pattern


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


## RewardArticle

> RewardArticleResp RewardArticle(ctx).RewardArticleReq(rewardArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.RewardArticle(context.Background()).RewardArticleReq(rewardArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.RewardArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RewardArticle`: RewardArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.RewardArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRewardArticleRequest struct via the builder pattern


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


## SetPinnedArticleComment

> map[string]interface{} SetPinnedArticleComment(ctx).SetPinnedArticleCommentReq(setPinnedArticleCommentReq).Execute()





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
	resp, r, err := apiClient.ArticleService.SetPinnedArticleComment(context.Background()).SetPinnedArticleCommentReq(setPinnedArticleCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.SetPinnedArticleComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPinnedArticleComment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.SetPinnedArticleComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPinnedArticleCommentRequest struct via the builder pattern


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


## ThankArticle

> ThankArticleResp ThankArticle(ctx).ThankArticleReq(thankArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.ThankArticle(context.Background()).ThankArticleReq(thankArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.ThankArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThankArticle`: ThankArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.ThankArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankArticleRequest struct via the builder pattern


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


## UpdateDraftArticle

> UpdateDraftArticleResp UpdateDraftArticle(ctx).UpdateDraftArticleReq(updateDraftArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.UpdateDraftArticle(context.Background()).UpdateDraftArticleReq(updateDraftArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.UpdateDraftArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraftArticle`: UpdateDraftArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.UpdateDraftArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftArticleRequest struct via the builder pattern


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


## VoteArticlePoll

> map[string]interface{} VoteArticlePoll(ctx).VoteArticlePollReq(voteArticlePollReq).Execute()





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
	resp, r, err := apiClient.ArticleService.VoteArticlePoll(context.Background()).VoteArticlePollReq(voteArticlePollReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.VoteArticlePoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoteArticlePoll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.VoteArticlePoll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoteArticlePollRequest struct via the builder pattern


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


## WatchArticle

> WatchArticleResp WatchArticle(ctx).WatchArticleReq(watchArticleReq).Execute()





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
	resp, r, err := apiClient.ArticleService.WatchArticle(context.Background()).WatchArticleReq(watchArticleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.WatchArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchArticle`: WatchArticleResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.WatchArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchArticleRequest struct via the builder pattern


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

