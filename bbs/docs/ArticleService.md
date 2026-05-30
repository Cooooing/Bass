# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAnswer**](ArticleService.md#AcceptAnswer) | **Post** /v1/content/article/accept-answer | 
[**Collect**](ArticleService.md#Collect) | **Post** /v1/content/article/collect | 
[**Create**](ArticleService.md#Create) | **Post** /v1/content/article/create | 
[**Delete**](ArticleService.md#Delete) | **Post** /v1/content/article/delete | 
[**Get**](ArticleService.md#Get) | **Post** /v1/content/article/get | 
[**Like**](ArticleService.md#Like) | **Post** /v1/content/article/like | 
[**List**](ArticleService.md#List) | **Post** /v1/content/article/list | 
[**Publish**](ArticleService.md#Publish) | **Post** /v1/content/article/publish | 
[**Reward**](ArticleService.md#Reward) | **Post** /v1/content/article/reward | 
[**Thank**](ArticleService.md#Thank) | **Post** /v1/content/article/thank | 
[**UpdateDraft**](ArticleService.md#UpdateDraft) | **Post** /v1/content/article/update-draft | 
[**Watch**](ArticleService.md#Watch) | **Post** /v1/content/article/watch | 



## AcceptAnswer

> map[string]interface{} AcceptAnswer(ctx).AcceptAnswerArticleRequest(acceptAnswerArticleRequest).Execute()





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
	acceptAnswerArticleRequest := *openapiclient.NewAcceptAnswerArticleRequest("ArticleId_example", "CommentId_example") // AcceptAnswerArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.AcceptAnswer(context.Background()).AcceptAnswerArticleRequest(acceptAnswerArticleRequest).Execute()
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
 **acceptAnswerArticleRequest** | [**AcceptAnswerArticleRequest**](AcceptAnswerArticleRequest.md) |  | 

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

> map[string]interface{} Collect(ctx).CollectArticleRequest(collectArticleRequest).Execute()





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
	collectArticleRequest := *openapiclient.NewCollectArticleRequest("ArticleId_example", false) // CollectArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Collect(context.Background()).CollectArticleRequest(collectArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Collect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Collect`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Collect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectArticleRequest** | [**CollectArticleRequest**](CollectArticleRequest.md) |  | 

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


## Create

> CreateArticleReply Create(ctx).CreateArticleRequest(createArticleRequest).Execute()





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
	createArticleRequest := *openapiclient.NewCreateArticleRequest(*openapiclient.NewArticleSave("Title_example", "Content_example", "Status_example", "Type_example")) // CreateArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Create(context.Background()).CreateArticleRequest(createArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createArticleRequest** | [**CreateArticleRequest**](CreateArticleRequest.md) |  | 

### Return type

[**CreateArticleReply**](CreateArticleReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> map[string]interface{} Delete(ctx).DeleteArticleRequest(deleteArticleRequest).Execute()





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
	deleteArticleRequest := *openapiclient.NewDeleteArticleRequest("ArticleId_example") // DeleteArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Delete(context.Background()).DeleteArticleRequest(deleteArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteArticleRequest** | [**DeleteArticleRequest**](DeleteArticleRequest.md) |  | 

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

> GetArticleReply Get(ctx).GetArticleRequest(getArticleRequest).Execute()





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
	getArticleRequest := *openapiclient.NewGetArticleRequest("ArticleId_example") // GetArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Get(context.Background()).GetArticleRequest(getArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: GetArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getArticleRequest** | [**GetArticleRequest**](GetArticleRequest.md) |  | 

### Return type

[**GetArticleReply**](GetArticleReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Like

> map[string]interface{} Like(ctx).LikeArticleRequest(likeArticleRequest).Execute()





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
	likeArticleRequest := *openapiclient.NewLikeArticleRequest("ArticleId_example", false) // LikeArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Like(context.Background()).LikeArticleRequest(likeArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Like``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Like`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Like`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeArticleRequest** | [**LikeArticleRequest**](LikeArticleRequest.md) |  | 

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


## List

> ListArticlesReply List(ctx).ListArticlesRequest(listArticlesRequest).Execute()





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
	listArticlesRequest := *openapiclient.NewListArticlesRequest() // ListArticlesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.List(context.Background()).ListArticlesRequest(listArticlesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListArticlesReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listArticlesRequest** | [**ListArticlesRequest**](ListArticlesRequest.md) |  | 

### Return type

[**ListArticlesReply**](ListArticlesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Publish

> map[string]interface{} Publish(ctx).PublishArticleRequest(publishArticleRequest).Execute()





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
	publishArticleRequest := *openapiclient.NewPublishArticleRequest("ArticleId_example") // PublishArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Publish(context.Background()).PublishArticleRequest(publishArticleRequest).Execute()
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
 **publishArticleRequest** | [**PublishArticleRequest**](PublishArticleRequest.md) |  | 

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

> map[string]interface{} Reward(ctx).RewardArticleRequest(rewardArticleRequest).Execute()





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
	rewardArticleRequest := *openapiclient.NewRewardArticleRequest("ArticleId_example") // RewardArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Reward(context.Background()).RewardArticleRequest(rewardArticleRequest).Execute()
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
 **rewardArticleRequest** | [**RewardArticleRequest**](RewardArticleRequest.md) |  | 

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

> map[string]interface{} Thank(ctx).ThankArticleRequest(thankArticleRequest).Execute()





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
	thankArticleRequest := *openapiclient.NewThankArticleRequest("ArticleId_example", false) // ThankArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Thank(context.Background()).ThankArticleRequest(thankArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Thank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Thank`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Thank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thankArticleRequest** | [**ThankArticleRequest**](ThankArticleRequest.md) |  | 

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


## UpdateDraft

> UpdateDraftArticleReply UpdateDraft(ctx).UpdateDraftArticleRequest(updateDraftArticleRequest).Execute()





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
	updateDraftArticleRequest := *openapiclient.NewUpdateDraftArticleRequest(*openapiclient.NewArticleSave("Title_example", "Content_example", "Status_example", "Type_example")) // UpdateDraftArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.UpdateDraft(context.Background()).UpdateDraftArticleRequest(updateDraftArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.UpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraft`: UpdateDraftArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.UpdateDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDraftArticleRequest** | [**UpdateDraftArticleRequest**](UpdateDraftArticleRequest.md) |  | 

### Return type

[**UpdateDraftArticleReply**](UpdateDraftArticleReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Watch

> map[string]interface{} Watch(ctx).WatchArticleRequest(watchArticleRequest).Execute()





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
	watchArticleRequest := *openapiclient.NewWatchArticleRequest("ArticleId_example", false) // WatchArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleService.Watch(context.Background()).WatchArticleRequest(watchArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleService.Watch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Watch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleService.Watch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchArticleRequest** | [**WatchArticleRequest**](WatchArticleRequest.md) |  | 

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

