# RealtimeService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createTicket**](#createticket) | **POST** /v1/realtime/ticket | |

# **createTicket**
> CreateRealtimeTicketResp createTicket(body)


### Example

```typescript
import {
    RealtimeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RealtimeService(configuration);

let body: object; //

const { status, data } = await apiInstance.createTicket(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CreateRealtimeTicketResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

