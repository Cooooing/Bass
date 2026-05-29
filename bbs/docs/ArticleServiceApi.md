# \ArticleServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**article_service_accept_answer**](ArticleServiceApi.md#article_service_accept_answer) | **POST** /v1/content/article/accept-answer | 
[**article_service_collect**](ArticleServiceApi.md#article_service_collect) | **POST** /v1/content/article/collect | 
[**article_service_create**](ArticleServiceApi.md#article_service_create) | **POST** /v1/content/article/create | 
[**article_service_delete**](ArticleServiceApi.md#article_service_delete) | **POST** /v1/content/article/delete | 
[**article_service_get**](ArticleServiceApi.md#article_service_get) | **POST** /v1/content/article/get | 
[**article_service_like**](ArticleServiceApi.md#article_service_like) | **POST** /v1/content/article/like | 
[**article_service_list**](ArticleServiceApi.md#article_service_list) | **POST** /v1/content/article/list | 
[**article_service_publish**](ArticleServiceApi.md#article_service_publish) | **POST** /v1/content/article/publish | 
[**article_service_reward**](ArticleServiceApi.md#article_service_reward) | **POST** /v1/content/article/reward | 
[**article_service_thank**](ArticleServiceApi.md#article_service_thank) | **POST** /v1/content/article/thank | 
[**article_service_update_draft**](ArticleServiceApi.md#article_service_update_draft) | **POST** /v1/content/article/update-draft | 
[**article_service_watch**](ArticleServiceApi.md#article_service_watch) | **POST** /v1/content/article/watch | 



## article_service_accept_answer

> serde_json::Value article_service_accept_answer(accept_answer_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**accept_answer_article_request** | [**AcceptAnswerArticleRequest**](AcceptAnswerArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_collect

> serde_json::Value article_service_collect(collect_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**collect_article_request** | [**CollectArticleRequest**](CollectArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_create

> models::CreateArticleReply article_service_create(create_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_article_request** | [**CreateArticleRequest**](CreateArticleRequest.md) |  | [required] |

### Return type

[**models::CreateArticleReply**](CreateArticle_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_delete

> serde_json::Value article_service_delete(delete_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**delete_article_request** | [**DeleteArticleRequest**](DeleteArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_get

> models::GetArticleReply article_service_get(get_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_article_request** | [**GetArticleRequest**](GetArticleRequest.md) |  | [required] |

### Return type

[**models::GetArticleReply**](GetArticle_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_like

> serde_json::Value article_service_like(like_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**like_article_request** | [**LikeArticleRequest**](LikeArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_list

> models::ListArticlesReply article_service_list(list_articles_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_articles_request** | [**ListArticlesRequest**](ListArticlesRequest.md) |  | [required] |

### Return type

[**models::ListArticlesReply**](ListArticles_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_publish

> serde_json::Value article_service_publish(publish_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**publish_article_request** | [**PublishArticleRequest**](PublishArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_reward

> serde_json::Value article_service_reward(reward_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reward_article_request** | [**RewardArticleRequest**](RewardArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_thank

> serde_json::Value article_service_thank(thank_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**thank_article_request** | [**ThankArticleRequest**](ThankArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_update_draft

> models::UpdateDraftArticleReply article_service_update_draft(update_draft_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_draft_article_request** | [**UpdateDraftArticleRequest**](UpdateDraftArticleRequest.md) |  | [required] |

### Return type

[**models::UpdateDraftArticleReply**](UpdateDraftArticle_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## article_service_watch

> serde_json::Value article_service_watch(watch_article_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**watch_article_request** | [**WatchArticleRequest**](WatchArticleRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

