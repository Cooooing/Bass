# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](TagService.md#create) | **POST** /v1/content/tag/create |  |
| [**list**](TagService.md#list) | **POST** /v1/content/tag/list |  |
| [**update**](TagService.md#update) | **POST** /v1/content/tag/update |  |



## create

> CreateTagReply create(createTagRequest)



创建标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // CreateTagRequest
    createTagRequest: ...,
  } satisfies CreateRequest;

  try {
    const data = await api.create(body);
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
| **createTagRequest** | [CreateTagRequest](CreateTagRequest.md) |  | |

### Return type

[**CreateTagReply**](CreateTagReply.md)

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


## list

> ListTagsReply list(listTagsRequest)



分页查询标签列表。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // ListTagsRequest
    listTagsRequest: ...,
  } satisfies ListRequest;

  try {
    const data = await api.list(body);
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


## update

> UpdateTagReply update(updateTagRequest)



更新标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // UpdateTagRequest
    updateTagRequest: ...,
  } satisfies UpdateRequest;

  try {
    const data = await api.update(body);
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
| **updateTagRequest** | [UpdateTagRequest](UpdateTagRequest.md) |  | |

### Return type

[**UpdateTagReply**](UpdateTagReply.md)

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

