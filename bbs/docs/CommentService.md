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

> models::CreateCommentReply create(create_comment_request)


创建评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_comment_request** | [**CreateCommentRequest**](CreateCommentRequest.md) |  | [required] |

### Return type

[**models::CreateCommentReply**](CreateComment_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## like

> models::LikeCommentReply like(like_comment_request)


点赞或取消点赞评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**like_comment_request** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | [required] |

### Return type

[**models::LikeCommentReply**](LikeComment_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListCommentsReply list(list_comments_request)


分页查询评论列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comments_request** | [**ListCommentsRequest**](ListCommentsRequest.md) |  | [required] |

### Return type

[**models::ListCommentsReply**](ListComments_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_replies

> models::ListCommentRepliesReply list_replies(list_comment_replies_request)


分页查询评论回复。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_replies_request** | [**ListCommentRepliesRequest**](ListCommentRepliesRequest.md) |  | [required] |

### Return type

[**models::ListCommentRepliesReply**](ListCommentReplies_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_threads

> models::ListCommentThreadsReply list_threads(list_comment_threads_request)


分页查询评论楼层。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_threads_request** | [**ListCommentThreadsRequest**](ListCommentThreadsRequest.md) |  | [required] |

### Return type

[**models::ListCommentThreadsReply**](ListCommentThreads_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_timeline

> models::ListCommentTimelineReply list_timeline(list_comment_timeline_request)


分页查询评论时间线。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_comment_timeline_request** | [**ListCommentTimelineRequest**](ListCommentTimelineRequest.md) |  | [required] |

### Return type

[**models::ListCommentTimelineReply**](ListCommentTimeline_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## thank

> models::ThankCommentReply thank(thank_comment_request)


感谢或取消感谢评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**thank_comment_request** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | [required] |

### Return type

[**models::ThankCommentReply**](ThankComment_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

