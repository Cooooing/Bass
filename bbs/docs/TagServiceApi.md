# TagServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**tagServiceList**](TagServiceApi.md#tagservicelist) | **POST** /v1/content/tag/list |  |



## tagServiceList

> ListTagsReply tagServiceList(listTagsRequest)



### Example

```ts
import {
  Configuration,
  TagServiceApi,
} from '@bass/bbs-sdk';
import type { TagServiceListRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TagServiceApi();

  const body = {
    // ListTagsRequest
    listTagsRequest: ...,
  } satisfies TagServiceListRequest;

  try {
    const data = await api.tagServiceList(body);
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
| **listTagsRequest** | [ListTagsRequest](ListTagsRequest.md) |  | |

### Return type

[**ListTagsReply**](ListTagsReply.md)

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

