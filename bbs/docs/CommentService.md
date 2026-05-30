# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create**](CommentService.md#create) | **POST** /v1/content/comment/create | 
[**like**](CommentService.md#like) | **POST** /v1/content/comment/like | 
[**list**](CommentService.md#list) | **POST** /v1/content/comment/list | 
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

> serde_json::Value like(like_comment_request)


点赞或取消点赞评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**like_comment_request** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

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


## thank

> serde_json::Value thank(thank_comment_request)


感谢或取消感谢评论。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**thank_comment_request** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

