# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**archive**](ArticleService.md#archive) | **POST** /v1/content/article/archive | 
[**cancel_publish**](ArticleService.md#cancel_publish) | **POST** /v1/content/article/publish/cancel | 
[**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect | 
[**create_draft**](ArticleService.md#create_draft) | **POST** /v1/content/article/draft/create | 
[**discard_draft**](ArticleService.md#discard_draft) | **POST** /v1/content/article/draft/discard | 
[**get**](ArticleService.md#get) | **POST** /v1/content/article/get | 
[**like**](ArticleService.md#like) | **POST** /v1/content/article/like | 
[**list**](ArticleService.md#list) | **POST** /v1/content/article/list | 
[**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish | 
[**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward | 
[**schedule_publish**](ArticleService.md#schedule_publish) | **POST** /v1/content/article/publish/schedule | 
[**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank | 
[**update_draft**](ArticleService.md#update_draft) | **POST** /v1/content/article/draft/update | 



## archive

> serde_json::Value archive(archive_article_req)


归档文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**archive_article_req** | [**ArchiveArticleReq**](ArchiveArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cancel_publish

> serde_json::Value cancel_publish(cancel_publish_article_req)


取消定时发布

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cancel_publish_article_req** | [**CancelPublishArticleReq**](CancelPublishArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## collect

> models::CollectArticleResp collect(collect_article_req)


收藏文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**collect_article_req** | [**CollectArticleReq**](CollectArticleReq.md) |  | [required] |

### Return type

[**models::CollectArticleResp**](CollectArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_draft

> models::CreateDraftArticleResp create_draft(create_draft_article_req)


创建文章草稿

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_draft_article_req** | [**CreateDraftArticleReq**](CreateDraftArticleReq.md) |  | [required] |

### Return type

[**models::CreateDraftArticleResp**](CreateDraftArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## discard_draft

> serde_json::Value discard_draft(discard_draft_article_req)


丢弃文章草稿

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**discard_draft_article_req** | [**DiscardDraftArticleReq**](DiscardDraftArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get

> models::GetArticleResp get(get_article_req)


查询文章详情

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_article_req** | [**GetArticleReq**](GetArticleReq.md) |  | [required] |

### Return type

[**models::GetArticleResp**](GetArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## like

> models::LikeArticleResp like(like_article_req)


点赞文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**like_article_req** | [**LikeArticleReq**](LikeArticleReq.md) |  | [required] |

### Return type

[**models::LikeArticleResp**](LikeArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListArticlesResp list(list_articles_req)


查询文章列表

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_articles_req** | [**ListArticlesReq**](ListArticlesReq.md) |  | [required] |

### Return type

[**models::ListArticlesResp**](ListArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## publish

> serde_json::Value publish(publish_article_req)


发布文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**publish_article_req** | [**PublishArticleReq**](PublishArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## reward

> serde_json::Value reward(reward_article_req)


打赏文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reward_article_req** | [**RewardArticleReq**](RewardArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## schedule_publish

> serde_json::Value schedule_publish(schedule_publish_article_req)


设置定时发布

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**schedule_publish_article_req** | [**SchedulePublishArticleReq**](SchedulePublishArticleReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## thank

> models::ThankArticleResp thank(thank_article_req)


感谢文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**thank_article_req** | [**ThankArticleReq**](ThankArticleReq.md) |  | [required] |

### Return type

[**models::ThankArticleResp**](ThankArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_draft

> models::UpdateDraftArticleResp update_draft(update_draft_article_req)


编辑文章草稿

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_draft_article_req** | [**UpdateDraftArticleReq**](UpdateDraftArticleReq.md) |  | [required] |

### Return type

[**models::UpdateDraftArticleResp**](UpdateDraftArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

