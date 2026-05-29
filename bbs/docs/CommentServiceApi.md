# \CommentServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**comment_service_create**](CommentServiceApi.md#comment_service_create) | **POST** /v1/content/comment/create | 
[**comment_service_like**](CommentServiceApi.md#comment_service_like) | **POST** /v1/content/comment/like | 
[**comment_service_list**](CommentServiceApi.md#comment_service_list) | **POST** /v1/content/comment/list | 
[**comment_service_thank**](CommentServiceApi.md#comment_service_thank) | **POST** /v1/content/comment/thank | 



## comment_service_create

> models::CreateCommentReply comment_service_create(create_comment_request)


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


## comment_service_like

> serde_json::Value comment_service_like(like_comment_request)


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


## comment_service_list

> models::ListCommentsReply comment_service_list(list_comments_request)


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


## comment_service_thank

> serde_json::Value comment_service_thank(thank_comment_request)


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

