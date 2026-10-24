# \GlobalChatService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_global_chat_topic**](GlobalChatService.md#get_global_chat_topic) | **POST** /v1/im/global-chat/topic/get | 
[**list_global_chat_barrages**](GlobalChatService.md#list_global_chat_barrages) | **POST** /v1/im/global-chat/barrages/list | 
[**list_global_chat_messages**](GlobalChatService.md#list_global_chat_messages) | **POST** /v1/im/global-chat/messages/list | 
[**list_global_chat_online_members**](GlobalChatService.md#list_global_chat_online_members) | **POST** /v1/im/global-chat/online-members/list | 
[**revoke_global_chat_message**](GlobalChatService.md#revoke_global_chat_message) | **POST** /v1/im/global-chat/messages/revoke | 
[**send_global_chat_barrage**](GlobalChatService.md#send_global_chat_barrage) | **POST** /v1/im/global-chat/barrages | 
[**send_global_chat_message**](GlobalChatService.md#send_global_chat_message) | **POST** /v1/im/global-chat/messages | 
[**set_global_chat_topic**](GlobalChatService.md#set_global_chat_topic) | **POST** /v1/im/global-chat/topic/set | 



## get_global_chat_topic

> models::GetGlobalChatTopicResp get_global_chat_topic(body)


查询当前全局聊天室话题。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetGlobalChatTopicResp**](GetGlobalChatTopic_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_global_chat_barrages

> models::ListGlobalChatBarragesResp list_global_chat_barrages(list_global_chat_barrages_req)


分页查询已发布的全局聊天室弹幕。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_global_chat_barrages_req** | [**ListGlobalChatBarragesReq**](ListGlobalChatBarragesReq.md) |  | [required] |

### Return type

[**models::ListGlobalChatBarragesResp**](ListGlobalChatBarrages_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_global_chat_messages

> models::ListGlobalChatMessagesResp list_global_chat_messages(list_global_chat_messages_req)


查询全局聊天室消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_global_chat_messages_req** | [**ListGlobalChatMessagesReq**](ListGlobalChatMessagesReq.md) |  | [required] |

### Return type

[**models::ListGlobalChatMessagesResp**](ListGlobalChatMessages_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_global_chat_online_members

> models::ListGlobalChatOnlineMembersResp list_global_chat_online_members(body)


查询全部在线的聊天室成员。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::ListGlobalChatOnlineMembersResp**](ListGlobalChatOnlineMembers_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## revoke_global_chat_message

> models::RevokeGlobalChatMessageResp revoke_global_chat_message(revoke_global_chat_message_req)


撤回自己的全局聊天室消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**revoke_global_chat_message_req** | [**RevokeGlobalChatMessageReq**](RevokeGlobalChatMessageReq.md) |  | [required] |

### Return type

[**models::RevokeGlobalChatMessageResp**](RevokeGlobalChatMessage_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## send_global_chat_barrage

> models::SendGlobalChatBarrageResp send_global_chat_barrage(send_global_chat_barrage_req)


发送全局聊天室弹幕。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**send_global_chat_barrage_req** | [**SendGlobalChatBarrageReq**](SendGlobalChatBarrageReq.md) |  | [required] |

### Return type

[**models::SendGlobalChatBarrageResp**](SendGlobalChatBarrage_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## send_global_chat_message

> models::SendGlobalChatMessageResp send_global_chat_message(send_global_chat_message_req)


发送全局聊天室消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**send_global_chat_message_req** | [**SendGlobalChatMessageReq**](SendGlobalChatMessageReq.md) |  | [required] |

### Return type

[**models::SendGlobalChatMessageResp**](SendGlobalChatMessage_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## set_global_chat_topic

> models::SetGlobalChatTopicResp set_global_chat_topic(set_global_chat_topic_req)


设置全局聊天室话题。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**set_global_chat_topic_req** | [**SetGlobalChatTopicReq**](SetGlobalChatTopicReq.md) |  | [required] |

### Return type

[**models::SetGlobalChatTopicResp**](SetGlobalChatTopic_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

