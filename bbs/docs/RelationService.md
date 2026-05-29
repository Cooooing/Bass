# \RelationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**block**](RelationService.md#block) | **POST** /v1/user/relation/block | 
[**follow**](RelationService.md#follow) | **POST** /v1/user/relation/follow | 
[**get_status**](RelationService.md#get_status) | **POST** /v1/user/relation/get-status | 
[**list_blocked**](RelationService.md#list_blocked) | **POST** /v1/user/relation/list-blocked | 
[**list_followers**](RelationService.md#list_followers) | **POST** /v1/user/relation/list-followers | 
[**list_following**](RelationService.md#list_following) | **POST** /v1/user/relation/list-following | 
[**unblock**](RelationService.md#unblock) | **POST** /v1/user/relation/unblock | 
[**unfollow**](RelationService.md#unfollow) | **POST** /v1/user/relation/unfollow | 



## block

> serde_json::Value block(block_relation_request)


当前账号拉黑目标账号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**block_relation_request** | [**BlockRelationRequest**](BlockRelationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## follow

> serde_json::Value follow(follow_relation_request)


当前账号关注目标账号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**follow_relation_request** | [**FollowRelationRequest**](FollowRelationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_status

> models::GetStatusRelationReply get_status(get_status_relation_request)


查询当前账号与目标账号之间的关系

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_status_relation_request** | [**GetStatusRelationRequest**](GetStatusRelationRequest.md) |  | [required] |

### Return type

[**models::GetStatusRelationReply**](GetStatusRelation_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_blocked

> models::ListBlockedRelationsReply list_blocked(list_blocked_relations_request)


分页查询当前账号拉黑的账号列表

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_blocked_relations_request** | [**ListBlockedRelationsRequest**](ListBlockedRelationsRequest.md) |  | [required] |

### Return type

[**models::ListBlockedRelationsReply**](ListBlockedRelations_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_followers

> models::ListFollowersRelationsReply list_followers(list_followers_relations_request)


分页查询当前账号的粉丝账号列表

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_followers_relations_request** | [**ListFollowersRelationsRequest**](ListFollowersRelationsRequest.md) |  | [required] |

### Return type

[**models::ListFollowersRelationsReply**](ListFollowersRelations_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_following

> models::ListFollowingRelationsReply list_following(list_following_relations_request)


分页查询当前账号关注的账号列表

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_following_relations_request** | [**ListFollowingRelationsRequest**](ListFollowingRelationsRequest.md) |  | [required] |

### Return type

[**models::ListFollowingRelationsReply**](ListFollowingRelations_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## unblock

> serde_json::Value unblock(unblock_relation_request)


当前账号取消拉黑目标账号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**unblock_relation_request** | [**UnblockRelationRequest**](UnblockRelationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## unfollow

> serde_json::Value unfollow(unfollow_relation_request)


当前账号取消关注目标账号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**unfollow_relation_request** | [**UnfollowRelationRequest**](UnfollowRelationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

