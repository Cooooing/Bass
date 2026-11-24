# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_comment**](CommentService.md#create_comment) | **POST** /v1/content/comment/create | 
[**like_comment**](CommentService.md#like_comment) | **POST** /v1/content/comment/like | 
[**list_comment_replies**](CommentService.md#list_comment_replies) | **POST** /v1/content/comment/list-replies | 
[**list_comment_threads**](CommentService.md#list_comment_threads) | **POST** /v1/content/comment/list-threads | 
[**list_comment_timeline**](CommentService.md#list_comment_timeline) | **POST** /v1/content/comment/list-timeline | 
[**list_comments**](CommentService.md#list_comments) | **POST** /v1/content/comment/list | 
[**thank_comment**](CommentService.md#thank_comment) | **POST** /v1/content/comment/thank | 



## create_comment

> models::CreateCommentResp create_comment(create_comment_req)


创建评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_comment_req** | [**CreateCommentReq**](CreateCommentReq.md) |  | [required] |

### Return type

[**models::CreateCommentResp**](CreateComment_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## like_comment

> models::LikeCommentResp like_comment(like_comment_req)


点赞或取消点赞评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**like_comment_req** | [**LikeCommentReq**](LikeCommentReq.md) |  | [required] |

### Return type

[**models::LikeCommentResp**](LikeComment_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_comment_replies

> models::ListCommentRepliesResp list_comment_replies(list_comment_replies_req)


分页查询评论回复。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_replies_req** | [**ListCommentRepliesReq**](ListCommentRepliesReq.md) |  | [required] |

### Return type

[**models::ListCommentRepliesResp**](ListCommentReplies_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_comment_threads

> models::ListCommentThreadsResp list_comment_threads(list_comment_threads_req)


分页查询评论楼层。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_threads_req** | [**ListCommentThreadsReq**](ListCommentThreadsReq.md) |  | [required] |

### Return type

[**models::ListCommentThreadsResp**](ListCommentThreads_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_comment_timeline

> models::ListCommentTimelineResp list_comment_timeline(list_comment_timeline_req)


分页查询评论时间线。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_timeline_req** | [**ListCommentTimelineReq**](ListCommentTimelineReq.md) |  | [required] |

### Return type

[**models::ListCommentTimelineResp**](ListCommentTimeline_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_comments

> models::ListCommentsResp list_comments(list_comments_req)


分页查询评论列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comments_req** | [**ListCommentsReq**](ListCommentsReq.md) |  | [required] |

### Return type

[**models::ListCommentsResp**](ListComments_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## thank_comment

> models::ThankCommentResp thank_comment(thank_comment_req)


感谢或取消感谢评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**thank_comment_req** | [**ThankCommentReq**](ThankCommentReq.md) |  | [required] |

### Return type

[**models::ThankCommentResp**](ThankComment_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

