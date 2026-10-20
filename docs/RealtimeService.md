# RealtimeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTicket**](RealtimeService.md#createticket) | **POST** /v1/realtime/ticket |  |



## createTicket

> CreateRealtimeTicketResp createTicket(body)



### Example

```ts
import {
  Configuration,
  RealtimeService,
} from '@bass/bbs-sdk-fetch';
import type { CreateTicketRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RealtimeService();

  const body = {
    // object
    body: Object,
  } satisfies CreateTicketRequest;

  try {
    const data = await api.createTicket(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | `object` |  | |

### Return type

[**CreateRealtimeTicketResp**](CreateRealtimeTicketResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

