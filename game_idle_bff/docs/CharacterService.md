# CharacterService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/game-idle/character/create | |
|[**list**](#list) | **POST** /v1/game-idle/character/list | |

# **create**
> CreateCharacterResp create(createCharacterReq)

创建角色。

### Example

```typescript
import {
    CharacterService,
    Configuration,
    CreateCharacterReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CharacterService(configuration);

let createCharacterReq: CreateCharacterReq; //

const { status, data } = await apiInstance.create(
    createCharacterReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createCharacterReq** | **CreateCharacterReq**|  | |


### Return type

**CreateCharacterResp**

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

# **list**
> ListCharacterResp list(body)

查询当前账号角色列表。

### Example

```typescript
import {
    CharacterService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CharacterService(configuration);

let body: object; //

const { status, data } = await apiInstance.list(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**ListCharacterResp**

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

