# AccountService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/account/get-current | |
|[**getProfile**](#getprofile) | **POST** /v1/user/account/get-profile | |
|[**updateProfile**](#updateprofile) | **POST** /v1/user/account/update-profile | |

# **getCurrent**
> GetCurrentAccountReply getCurrent(body)

获取当前登录账号的完整资料

### Example

```typescript
import {
    AccountService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let body: object; //

const { status, data } = await apiInstance.getCurrent(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentAccountReply**

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

# **getProfile**
> GetProfileAccountReply getProfile(getProfileAccountRequest)

按账号 ID 获取账号展示资料

### Example

```typescript
import {
    AccountService,
    Configuration,
    GetProfileAccountRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let getProfileAccountRequest: GetProfileAccountRequest; //

const { status, data } = await apiInstance.getProfile(
    getProfileAccountRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getProfileAccountRequest** | **GetProfileAccountRequest**|  | |


### Return type

**GetProfileAccountReply**

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

# **updateProfile**
> UpdateProfileAccountReply updateProfile(updateProfileAccountRequest)

更新当前登录账号的展示资料

### Example

```typescript
import {
    AccountService,
    Configuration,
    UpdateProfileAccountRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let updateProfileAccountRequest: UpdateProfileAccountRequest; //

const { status, data } = await apiInstance.updateProfile(
    updateProfileAccountRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateProfileAccountRequest** | **UpdateProfileAccountRequest**|  | |


### Return type

**UpdateProfileAccountReply**

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

