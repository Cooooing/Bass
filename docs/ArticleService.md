# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**accept_answer**](ArticleService.md#accept_answer) | **POST** /v1/content/article/answer/accept | 
[**archive**](ArticleService.md#archive) | **POST** /v1/content/article/archive | 
[**cancel_publish**](ArticleService.md#cancel_publish) | **POST** /v1/content/article/publish/cancel | 
[**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect | 
[**create_draft**](ArticleService.md#create_draft) | **POST** /v1/content/article/draft/create | 
[**discard_draft**](ArticleService.md#discard_draft) | **POST** /v1/content/article/draft/discard | 
[**enter_lottery**](ArticleService.md#enter_lottery) | **POST** /v1/content/article/lottery/entry | 
[**get**](ArticleService.md#get) | **POST** /v1/content/article/get | 
[**like**](ArticleService.md#like) | **POST** /v1/content/article/like | 
[**list**](ArticleService.md#list) | **POST** /v1/content/article/list | 
[**list_following**](ArticleService.md#list_following) | **POST** /v1/content/article/following/list | 
[**list_nearby**](ArticleService.md#list_nearby) | **POST** /v1/content/article/nearby/list | 
[**page_view_history**](ArticleService.md#page_view_history) | **POST** /v1/content/article/view-history/page | 
[**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish | 
[**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward | 
[**set_pinned_comment**](ArticleService.md#set_pinned_comment) | **POST** /v1/content/article/comment/pin | 
[**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank | 
[**update_draft**](ArticleService.md#update_draft) | **POST** /v1/content/article/draft/update | 
[**vote_poll**](ArticleService.md#vote_poll) | **POST** /v1/content/article/poll/vote | 
[**watch**](ArticleService.md#watch) | **POST** /v1/content/article/watch | 



## accept_answer

> serde_json::Value accept_answer(accept_article_answer_req)


采纳文章回答

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**accept_article_answer_req** | [**AcceptArticleAnswerReq**](AcceptArticleAnswerReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


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


## enter_lottery

> serde_json::Value enter_lottery(enter_article_lottery_req)


报名或撤销抽奖报名。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**enter_article_lottery_req** | [**EnterArticleLotteryReq**](EnterArticleLotteryReq.md) |  | [required] |

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


## list_following

> models::ListFollowingArticlesResp list_following(list_following_articles_req)


查询当前账号关注作者发布的文章。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_following_articles_req** | [**ListFollowingArticlesReq**](ListFollowingArticlesReq.md) |  | [required] |

### Return type

[**models::ListFollowingArticlesResp**](ListFollowingArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_nearby

> models::ListNearbyArticlesResp list_nearby(list_nearby_articles_req)


查询当前账号所在城市发布的文章。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_nearby_articles_req** | [**ListNearbyArticlesReq**](ListNearbyArticlesReq.md) |  | [required] |

### Return type

[**models::ListNearbyArticlesResp**](ListNearbyArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_view_history

> models::PageArticleViewHistoryResp page_view_history(page_article_view_history_req)


查询当前账号的浏览历史。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_article_view_history_req** | [**PageArticleViewHistoryReq**](PageArticleViewHistoryReq.md) |  | [required] |

### Return type

[**models::PageArticleViewHistoryResp**](PageArticleViewHistory_Resp.md)

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

> models::RewardArticleResp reward(reward_article_req)


打赏文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reward_article_req** | [**RewardArticleReq**](RewardArticleReq.md) |  | [required] |

### Return type

[**models::RewardArticleResp**](RewardArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## set_pinned_comment

> serde_json::Value set_pinned_comment(set_pinned_article_comment_req)


设置或取消当前账号文章的置顶评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**set_pinned_article_comment_req** | [**SetPinnedArticleCommentReq**](SetPinnedArticleCommentReq.md) |  | [required] |

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


感谢文章，确认后不可撤销。

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


## vote_poll

> serde_json::Value vote_poll(vote_article_poll_req)


提交一次不可修改的文章投票。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**vote_article_poll_req** | [**VoteArticlePollReq**](VoteArticlePollReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## watch

> models::WatchArticleResp watch(watch_article_req)


关注文章

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**watch_article_req** | [**WatchArticleReq**](WatchArticleReq.md) |  | [required] |

### Return type

[**models::WatchArticleResp**](WatchArticle_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

