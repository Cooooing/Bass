# \RelationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**relation_service_block**](RelationServiceApi.md#relation_service_block) | **POST** /v1/user/relation/block | 
[**relation_service_follow**](RelationServiceApi.md#relation_service_follow) | **POST** /v1/user/relation/follow | 
[**relation_service_get_status**](RelationServiceApi.md#relation_service_get_status) | **POST** /v1/user/relation/get-status | 
[**relation_service_list_blocked**](RelationServiceApi.md#relation_service_list_blocked) | **POST** /v1/user/relation/list-blocked | 
[**relation_service_list_followers**](RelationServiceApi.md#relation_service_list_followers) | **POST** /v1/user/relation/list-followers | 
[**relation_service_list_following**](RelationServiceApi.md#relation_service_list_following) | **POST** /v1/user/relation/list-following | 
[**relation_service_unblock**](RelationServiceApi.md#relation_service_unblock) | **POST** /v1/user/relation/unblock | 
[**relation_service_unfollow**](RelationServiceApi.md#relation_service_unfollow) | **POST** /v1/user/relation/unfollow | 



## relation_service_block

> serde_json::Value relation_service_block(block_relation_request)


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


## relation_service_follow

> serde_json::Value relation_service_follow(follow_relation_request)


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


## relation_service_get_status

> models::GetStatusRelationReply relation_service_get_status(get_status_relation_request)


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


## relation_service_list_blocked

> models::ListBlockedRelationsReply relation_service_list_blocked(list_blocked_relations_request)


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


## relation_service_list_followers

> models::ListFollowersRelationsReply relation_service_list_followers(list_followers_relations_request)


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


## relation_service_list_following

> models::ListFollowingRelationsReply relation_service_list_following(list_following_relations_request)


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


## relation_service_unblock

> serde_json::Value relation_service_unblock(unblock_relation_request)


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


## relation_service_unfollow

> serde_json::Value relation_service_unfollow(unfollow_relation_request)


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

