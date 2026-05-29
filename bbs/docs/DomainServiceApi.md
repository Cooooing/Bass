# DomainServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**domainServiceList**](DomainServiceApi.md#domainservicelist) | **POST** /v1/content/domain/list |  |



## domainServiceList

> ListDomainsReply domainServiceList(listDomainsRequest)



### Example

```ts
import {
  Configuration,
  DomainServiceApi,
} from '@bass/bbs-sdk';
import type { DomainServiceListRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new DomainServiceApi();

  const body = {
    // ListDomainsRequest
    listDomainsRequest: ...,
  } satisfies DomainServiceListRequest;

  try {
    const data = await api.domainServiceList(body);
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
| **listDomainsRequest** | [ListDomainsRequest](ListDomainsRequest.md) |  | |

### Return type

[**ListDomainsReply**](ListDomainsReply.md)

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

