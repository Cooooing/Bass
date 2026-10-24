# \PrivateChatService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**count_private_chat_unread**](PrivateChatService.md#count_private_chat_unread) | **POST** /v1/im/private-chat/unread-count | 
[**list_private_chat_messages**](PrivateChatService.md#list_private_chat_messages) | **POST** /v1/im/private-chat/messages/list | 
[**list_private_chat_sessions**](PrivateChatService.md#list_private_chat_sessions) | **GET** /v1/im/private-chat/sessions | 
[**mark_private_chat_session_read**](PrivateChatService.md#mark_private_chat_session_read) | **POST** /v1/im/private-chat/read | 
[**revoke_private_chat_message**](PrivateChatService.md#revoke_private_chat_message) | **POST** /v1/im/private-chat/messages/revoke | 
[**send_private_chat_message**](PrivateChatService.md#send_private_chat_message) | **POST** /v1/im/private-chat/messages | 
[**set_private_chat_session_muted**](PrivateChatService.md#set_private_chat_session_muted) | **POST** /v1/im/private-chat/muted | 
[**set_private_chat_session_pinned**](PrivateChatService.md#set_private_chat_session_pinned) | **POST** /v1/im/private-chat/pinned | 



## count_private_chat_unread

> models::CountPrivateChatUnreadResp count_private_chat_unread(body)


统计当前账号的私聊未读数。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::CountPrivateChatUnreadResp**](CountPrivateChatUnread_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_private_chat_messages

> models::ListPrivateChatMessagesResp list_private_chat_messages(list_private_chat_messages_req)


查询私聊消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_private_chat_messages_req** | [**ListPrivateChatMessagesReq**](ListPrivateChatMessagesReq.md) |  | [required] |

### Return type

[**models::ListPrivateChatMessagesResp**](ListPrivateChatMessages_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_private_chat_sessions

> models::ListPrivateChatSessionsResp list_private_chat_sessions(page_page, page_size)


查询当前账号的私聊会话。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page_page** | Option<**String**> | 页码 |  |
**page_size** | Option<**String**> | 页大小 |  |

### Return type

[**models::ListPrivateChatSessionsResp**](ListPrivateChatSessions_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## mark_private_chat_session_read

> serde_json::Value mark_private_chat_session_read(mark_private_chat_session_read_req)


标记私聊会话已读。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**mark_private_chat_session_read_req** | [**MarkPrivateChatSessionReadReq**](MarkPrivateChatSessionReadReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## revoke_private_chat_message

> models::RevokePrivateChatMessageResp revoke_private_chat_message(revoke_private_chat_message_req)


撤回私聊消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**revoke_private_chat_message_req** | [**RevokePrivateChatMessageReq**](RevokePrivateChatMessageReq.md) |  | [required] |

### Return type

[**models::RevokePrivateChatMessageResp**](RevokePrivateChatMessage_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## send_private_chat_message

> models::SendPrivateChatMessageResp send_private_chat_message(send_private_chat_message_req)


发送私聊消息。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**send_private_chat_message_req** | [**SendPrivateChatMessageReq**](SendPrivateChatMessageReq.md) |  | [required] |

### Return type

[**models::SendPrivateChatMessageResp**](SendPrivateChatMessage_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## set_private_chat_session_muted

> serde_json::Value set_private_chat_session_muted(set_private_chat_session_muted_req)


设置私聊会话免打扰。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**set_private_chat_session_muted_req** | [**SetPrivateChatSessionMutedReq**](SetPrivateChatSessionMutedReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## set_private_chat_session_pinned

> serde_json::Value set_private_chat_session_pinned(set_private_chat_session_pinned_req)


设置私聊会话置顶。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**set_private_chat_session_pinned_req** | [**SetPrivateChatSessionPinnedReq**](SetPrivateChatSessionPinnedReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

