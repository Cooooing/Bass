# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**accept_answer**](ArticleService.md#accept_answer) | **POST** /v1/content/article/accept-answer | 
[**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect | 
[**create**](ArticleService.md#create) | **POST** /v1/content/article/create | 
[**delete**](ArticleService.md#delete) | **POST** /v1/content/article/delete | 
[**get**](ArticleService.md#get) | **POST** /v1/content/article/get | 
[**like**](ArticleService.md#like) | **POST** /v1/content/article/like | 
[**list**](ArticleService.md#list) | **POST** /v1/content/article/list | 
[**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish | 
[**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward | 
[**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank | 
[**update_draft**](ArticleService.md#update_draft) | **POST** /v1/content/article/update-draft | 
[**watch**](ArticleService.md#watch) | **POST** /v1/content/article/watch | 



## accept_answer

> serde_json::Value accept_answer(accept_answer_article_request)


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


## collect

> serde_json::Value collect(collect_article_request)


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


## create

> models::CreateArticleReply create(create_article_request)


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


## delete

> serde_json::Value delete(delete_article_request)


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


## get

> models::GetArticleReply get(get_article_request)


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


## like

> serde_json::Value like(like_article_request)


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


## list

> models::ListArticlesReply list(list_articles_request)


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


## publish

> serde_json::Value publish(publish_article_request)


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


## reward

> serde_json::Value reward(reward_article_request)


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


## thank

> serde_json::Value thank(thank_article_request)


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


## update_draft

> models::UpdateDraftArticleReply update_draft(update_draft_article_request)


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


## watch

> serde_json::Value watch(watch_article_request)


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

