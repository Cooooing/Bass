# \WebSocketService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_session**](WebSocketService.md#create_session) | **POST** /v1/game-idle/ws/create-session | 



## create_session

> models::CreateWebSocketSessionResp create_session(create_web_socket_session_req)


创建角色 WS session。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_web_socket_session_req** | [**CreateWebSocketSessionReq**](CreateWebSocketSessionReq.md) |  | [required] |

### Return type

[**models::CreateWebSocketSessionResp**](CreateWebSocketSession_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

