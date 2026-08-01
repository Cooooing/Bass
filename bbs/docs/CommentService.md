# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](CommentService.md#create) | **POST** /v1/content/comment/create | 
[**like**](CommentService.md#like) | **POST** /v1/content/comment/like | 
[**list**](CommentService.md#list) | **POST** /v1/content/comment/list | 
[**list_replies**](CommentService.md#list_replies) | **POST** /v1/content/comment/list-replies | 
[**list_threads**](CommentService.md#list_threads) | **POST** /v1/content/comment/list-threads | 
[**list_timeline**](CommentService.md#list_timeline) | **POST** /v1/content/comment/list-timeline | 
[**thank**](CommentService.md#thank) | **POST** /v1/content/comment/thank | 



## create

> models::CreateCommentResp create(create_comment_req)


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


## like

> models::LikeCommentResp like(like_comment_req)


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


## list

> models::ListCommentsResp list(list_comments_req)


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


## list_replies

> models::ListCommentRepliesResp list_replies(list_comment_replies_req)


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


## list_threads

> models::ListCommentThreadsResp list_threads(list_comment_threads_req)


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


## list_timeline

> models::ListCommentTimelineResp list_timeline(list_comment_timeline_req)


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


## thank

> models::ThankCommentResp thank(thank_comment_req)


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

