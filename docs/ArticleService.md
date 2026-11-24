# \ArticleService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**accept_article_answer**](ArticleService.md#accept_article_answer) | **POST** /v1/content/article/answer/accept | 
[**archive_article**](ArticleService.md#archive_article) | **POST** /v1/content/article/archive | 
[**cancel_publish_article**](ArticleService.md#cancel_publish_article) | **POST** /v1/content/article/publish/cancel | 
[**collect_article**](ArticleService.md#collect_article) | **POST** /v1/content/article/collect | 
[**count_collected_articles**](ArticleService.md#count_collected_articles) | **POST** /v1/content/article/collect/count | 
[**create_draft_article**](ArticleService.md#create_draft_article) | **POST** /v1/content/article/draft/create | 
[**discard_draft_article**](ArticleService.md#discard_draft_article) | **POST** /v1/content/article/draft/discard | 
[**enter_article_lottery**](ArticleService.md#enter_article_lottery) | **POST** /v1/content/article/lottery/entry | 
[**get_article**](ArticleService.md#get_article) | **POST** /v1/content/article/get | 
[**like_article**](ArticleService.md#like_article) | **POST** /v1/content/article/like | 
[**list_articles**](ArticleService.md#list_articles) | **POST** /v1/content/article/list | 
[**list_following_articles**](ArticleService.md#list_following_articles) | **POST** /v1/content/article/following/list | 
[**list_nearby_articles**](ArticleService.md#list_nearby_articles) | **POST** /v1/content/article/nearby/list | 
[**page_article_view_history**](ArticleService.md#page_article_view_history) | **POST** /v1/content/article/view-history/page | 
[**page_collected_articles**](ArticleService.md#page_collected_articles) | **POST** /v1/content/article/collected/page | 
[**page_watched_articles**](ArticleService.md#page_watched_articles) | **POST** /v1/content/article/watched/page | 
[**publish_article**](ArticleService.md#publish_article) | **POST** /v1/content/article/publish | 
[**reward_article**](ArticleService.md#reward_article) | **POST** /v1/content/article/reward | 
[**set_pinned_article_comment**](ArticleService.md#set_pinned_article_comment) | **POST** /v1/content/article/comment/pin | 
[**thank_article**](ArticleService.md#thank_article) | **POST** /v1/content/article/thank | 
[**update_draft_article**](ArticleService.md#update_draft_article) | **POST** /v1/content/article/draft/update | 
[**vote_article_poll**](ArticleService.md#vote_article_poll) | **POST** /v1/content/article/poll/vote | 
[**watch_article**](ArticleService.md#watch_article) | **POST** /v1/content/article/watch | 



## accept_article_answer

> serde_json::Value accept_article_answer(accept_article_answer_req)


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


## archive_article

> serde_json::Value archive_article(archive_article_req)


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


## cancel_publish_article

> serde_json::Value cancel_publish_article(cancel_publish_article_req)


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


## collect_article

> models::CollectArticleResp collect_article(collect_article_req)


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


## count_collected_articles

> models::CountCollectedArticlesResp count_collected_articles(body)


统计当前账号收藏的文章数量。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::CountCollectedArticlesResp**](CountCollectedArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_draft_article

> models::CreateDraftArticleResp create_draft_article(create_draft_article_req)


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


## discard_draft_article

> serde_json::Value discard_draft_article(discard_draft_article_req)


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


## enter_article_lottery

> serde_json::Value enter_article_lottery(enter_article_lottery_req)


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


## get_article

> models::GetArticleResp get_article(get_article_req)


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


## like_article

> models::LikeArticleResp like_article(like_article_req)


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


## list_articles

> models::ListArticlesResp list_articles(list_articles_req)


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


## list_following_articles

> models::ListFollowingArticlesResp list_following_articles(list_following_articles_req)


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


## list_nearby_articles

> models::ListNearbyArticlesResp list_nearby_articles(list_nearby_articles_req)


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


## page_article_view_history

> models::PageArticleViewHistoryResp page_article_view_history(page_article_view_history_req)


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


## page_collected_articles

> models::PageCollectedArticlesResp page_collected_articles(page_collected_articles_req)


查询当前账号收藏的文章。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_collected_articles_req** | [**PageCollectedArticlesReq**](PageCollectedArticlesReq.md) |  | [required] |

### Return type

[**models::PageCollectedArticlesResp**](PageCollectedArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## page_watched_articles

> models::PageWatchedArticlesResp page_watched_articles(page_watched_articles_req)


查询当前账号关注的文章。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_watched_articles_req** | [**PageWatchedArticlesReq**](PageWatchedArticlesReq.md) |  | [required] |

### Return type

[**models::PageWatchedArticlesResp**](PageWatchedArticles_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## publish_article

> serde_json::Value publish_article(publish_article_req)


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


## reward_article

> models::RewardArticleResp reward_article(reward_article_req)


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


## set_pinned_article_comment

> serde_json::Value set_pinned_article_comment(set_pinned_article_comment_req)


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


## thank_article

> models::ThankArticleResp thank_article(thank_article_req)


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


## update_draft_article

> models::UpdateDraftArticleResp update_draft_article(update_draft_article_req)


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


## vote_article_poll

> serde_json::Value vote_article_poll(vote_article_poll_req)


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


## watch_article

> models::WatchArticleResp watch_article(watch_article_req)


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

