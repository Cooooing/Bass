# \ArticleServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticleServiceAcceptAnswer**](ArticleServiceAPI.md#ArticleServiceAcceptAnswer) | **Post** /v1/content/article/accept-answer | 
[**ArticleServiceCollect**](ArticleServiceAPI.md#ArticleServiceCollect) | **Post** /v1/content/article/collect | 
[**ArticleServiceCreate**](ArticleServiceAPI.md#ArticleServiceCreate) | **Post** /v1/content/article/create | 
[**ArticleServiceDelete**](ArticleServiceAPI.md#ArticleServiceDelete) | **Post** /v1/content/article/delete | 
[**ArticleServiceGet**](ArticleServiceAPI.md#ArticleServiceGet) | **Post** /v1/content/article/get | 
[**ArticleServiceLike**](ArticleServiceAPI.md#ArticleServiceLike) | **Post** /v1/content/article/like | 
[**ArticleServiceList**](ArticleServiceAPI.md#ArticleServiceList) | **Post** /v1/content/article/list | 
[**ArticleServicePublish**](ArticleServiceAPI.md#ArticleServicePublish) | **Post** /v1/content/article/publish | 
[**ArticleServiceReward**](ArticleServiceAPI.md#ArticleServiceReward) | **Post** /v1/content/article/reward | 
[**ArticleServiceThank**](ArticleServiceAPI.md#ArticleServiceThank) | **Post** /v1/content/article/thank | 
[**ArticleServiceUpdateDraft**](ArticleServiceAPI.md#ArticleServiceUpdateDraft) | **Post** /v1/content/article/update-draft | 
[**ArticleServiceWatch**](ArticleServiceAPI.md#ArticleServiceWatch) | **Post** /v1/content/article/watch | 



## ArticleServiceAcceptAnswer

> map[string]interface{} ArticleServiceAcceptAnswer(ctx).AcceptAnswerArticleRequest(acceptAnswerArticleRequest).Execute()



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
	acceptAnswerArticleRequest := *openapiclient.NewAcceptAnswerArticleRequest() // AcceptAnswerArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceAcceptAnswer(context.Background()).AcceptAnswerArticleRequest(acceptAnswerArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceAcceptAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceAcceptAnswer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceAcceptAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceAcceptAnswerRequest struct via the builder pattern


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


## ArticleServiceCollect

> map[string]interface{} ArticleServiceCollect(ctx).CollectArticleRequest(collectArticleRequest).Execute()



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
	collectArticleRequest := *openapiclient.NewCollectArticleRequest() // CollectArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceCollect(context.Background()).CollectArticleRequest(collectArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceCollect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceCollect`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceCollect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceCollectRequest struct via the builder pattern


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


## ArticleServiceCreate

> CreateArticleReply ArticleServiceCreate(ctx).CreateArticleRequest(createArticleRequest).Execute()



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
	createArticleRequest := *openapiclient.NewCreateArticleRequest() // CreateArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceCreate(context.Background()).CreateArticleRequest(createArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceCreate`: CreateArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceCreateRequest struct via the builder pattern


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


## ArticleServiceDelete

> map[string]interface{} ArticleServiceDelete(ctx).DeleteArticleRequest(deleteArticleRequest).Execute()



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
	deleteArticleRequest := *openapiclient.NewDeleteArticleRequest() // DeleteArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceDelete(context.Background()).DeleteArticleRequest(deleteArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceDeleteRequest struct via the builder pattern


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


## ArticleServiceGet

> GetArticleReply ArticleServiceGet(ctx).GetArticleRequest(getArticleRequest).Execute()



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
	getArticleRequest := *openapiclient.NewGetArticleRequest() // GetArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceGet(context.Background()).GetArticleRequest(getArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceGet`: GetArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceGetRequest struct via the builder pattern


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


## ArticleServiceLike

> map[string]interface{} ArticleServiceLike(ctx).LikeArticleRequest(likeArticleRequest).Execute()



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
	likeArticleRequest := *openapiclient.NewLikeArticleRequest() // LikeArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceLike(context.Background()).LikeArticleRequest(likeArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceLike``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceLike`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceLike`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceLikeRequest struct via the builder pattern


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


## ArticleServiceList

> ListArticlesReply ArticleServiceList(ctx).ListArticlesRequest(listArticlesRequest).Execute()



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
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceList(context.Background()).ListArticlesRequest(listArticlesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceList`: ListArticlesReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceListRequest struct via the builder pattern


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


## ArticleServicePublish

> map[string]interface{} ArticleServicePublish(ctx).PublishArticleRequest(publishArticleRequest).Execute()



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
	publishArticleRequest := *openapiclient.NewPublishArticleRequest() // PublishArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServicePublish(context.Background()).PublishArticleRequest(publishArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServicePublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServicePublish`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServicePublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServicePublishRequest struct via the builder pattern


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


## ArticleServiceReward

> map[string]interface{} ArticleServiceReward(ctx).RewardArticleRequest(rewardArticleRequest).Execute()



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
	rewardArticleRequest := *openapiclient.NewRewardArticleRequest() // RewardArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceReward(context.Background()).RewardArticleRequest(rewardArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceReward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceReward`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceReward`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceRewardRequest struct via the builder pattern


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


## ArticleServiceThank

> map[string]interface{} ArticleServiceThank(ctx).ThankArticleRequest(thankArticleRequest).Execute()



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
	thankArticleRequest := *openapiclient.NewThankArticleRequest() // ThankArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceThank(context.Background()).ThankArticleRequest(thankArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceThank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceThank`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceThank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceThankRequest struct via the builder pattern


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


## ArticleServiceUpdateDraft

> UpdateDraftArticleReply ArticleServiceUpdateDraft(ctx).UpdateDraftArticleRequest(updateDraftArticleRequest).Execute()



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
	updateDraftArticleRequest := *openapiclient.NewUpdateDraftArticleRequest() // UpdateDraftArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceUpdateDraft(context.Background()).UpdateDraftArticleRequest(updateDraftArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceUpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceUpdateDraft`: UpdateDraftArticleReply
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceUpdateDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceUpdateDraftRequest struct via the builder pattern


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


## ArticleServiceWatch

> map[string]interface{} ArticleServiceWatch(ctx).WatchArticleRequest(watchArticleRequest).Execute()



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
	watchArticleRequest := *openapiclient.NewWatchArticleRequest() // WatchArticleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleServiceAPI.ArticleServiceWatch(context.Background()).WatchArticleRequest(watchArticleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleServiceAPI.ArticleServiceWatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleServiceWatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ArticleServiceAPI.ArticleServiceWatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticleServiceWatchRequest struct via the builder pattern


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

