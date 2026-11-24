# AuthService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**loginAccount**](#loginaccount) | **POST** /v1/game-idle/auth/login | |
|[**registerAccount**](#registeraccount) | **POST** /v1/game-idle/auth/register | |

# **loginAccount**
> LoginAccountResp loginAccount(loginAccountReq)

登录账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    LoginAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let loginAccountReq: LoginAccountReq; //

const { status, data } = await apiInstance.loginAccount(
    loginAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loginAccountReq** | **LoginAccountReq**|  | |


### Return type

**LoginAccountResp**

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

# **registerAccount**
> object registerAccount(registerAccountReq)

注册账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    RegisterAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let registerAccountReq: RegisterAccountReq; //

const { status, data } = await apiInstance.registerAccount(
    registerAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **registerAccountReq** | **RegisterAccountReq**|  | |


### Return type

**object**

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

