# \WebSocketService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_ticket**](WebSocketService.md#create_ticket) | **POST** /v1/game-idle/ws/create-ticket | 



## create_ticket

> models::CreateWebSocketTicketResp create_ticket(create_web_socket_ticket_req)


创建角色 WS 连接凭证。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_web_socket_ticket_req** | [**CreateWebSocketTicketReq**](CreateWebSocketTicketReq.md) |  | [required] |

### Return type

[**models::CreateWebSocketTicketResp**](CreateWebSocketTicket_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

