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

> serde_json::Value block(block_relation_req)


当前账号拉黑目标账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**block_relation_req** | [**BlockRelationReq**](BlockRelationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## follow

> serde_json::Value follow(follow_relation_req)


当前账号关注目标账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**follow_relation_req** | [**FollowRelationReq**](FollowRelationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_status

> models::GetStatusRelationResp get_status(get_status_relation_req)


查询当前账号与目标账号之间的关系。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_status_relation_req** | [**GetStatusRelationReq**](GetStatusRelationReq.md) |  | [required] |

### Return type

[**models::GetStatusRelationResp**](GetStatusRelation_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_blocked

> models::ListBlockedRelationsResp list_blocked(list_blocked_relations_req)


分页查询当前账号拉黑的账号列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_blocked_relations_req** | [**ListBlockedRelationsReq**](ListBlockedRelationsReq.md) |  | [required] |

### Return type

[**models::ListBlockedRelationsResp**](ListBlockedRelations_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_followers

> models::ListFollowersRelationsResp list_followers(list_followers_relations_req)


分页查询当前账号的粉丝账号列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_followers_relations_req** | [**ListFollowersRelationsReq**](ListFollowersRelationsReq.md) |  | [required] |

### Return type

[**models::ListFollowersRelationsResp**](ListFollowersRelations_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_following

> models::ListFollowingRelationsResp list_following(list_following_relations_req)


分页查询当前账号关注的账号列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_following_relations_req** | [**ListFollowingRelationsReq**](ListFollowingRelationsReq.md) |  | [required] |

### Return type

[**models::ListFollowingRelationsResp**](ListFollowingRelations_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## unblock

> serde_json::Value unblock(unblock_relation_req)


当前账号取消拉黑目标账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**unblock_relation_req** | [**UnblockRelationReq**](UnblockRelationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## unfollow

> serde_json::Value unfollow(unfollow_relation_req)


当前账号取消关注目标账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**unfollow_relation_req** | [**UnfollowRelationReq**](UnfollowRelationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

