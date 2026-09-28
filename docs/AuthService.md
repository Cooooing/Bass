# AuthService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**cancelAccount**](#cancelaccount) | **POST** /v1/user/auth/cancel-account | |
|[**login**](#login) | **POST** /v1/user/auth/login | |
|[**logout**](#logout) | **POST** /v1/user/auth/logout | |
|[**refreshToken**](#refreshtoken) | **POST** /v1/user/auth/refresh-token | |
|[**register**](#register) | **POST** /v1/user/auth/register | |

# **cancelAccount**
> object cancelAccount(cancelAccountReq)

注销账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    CancelAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let cancelAccountReq: CancelAccountReq; //

const { status, data } = await apiInstance.cancelAccount(
    cancelAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **cancelAccountReq** | **CancelAccountReq**|  | |


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

# **login**
> LoginResp login(loginReq)

登录账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    LoginReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let loginReq: LoginReq; //

const { status, data } = await apiInstance.login(
    loginReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loginReq** | **LoginReq**|  | |


### Return type

**LoginResp**

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

# **logout**
> object logout(body)

退出登录。

### Example

```typescript
import {
    AuthService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let body: object; //

const { status, data } = await apiInstance.logout(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


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

# **refreshToken**
> RefreshTokenResp refreshToken(refreshTokenReq)

刷新登录令牌。

### Example

```typescript
import {
    AuthService,
    Configuration,
    RefreshTokenReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let refreshTokenReq: RefreshTokenReq; //

const { status, data } = await apiInstance.refreshToken(
    refreshTokenReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **refreshTokenReq** | **RefreshTokenReq**|  | |


### Return type

**RefreshTokenResp**

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

# **register**
> object register(registerReq)

注册账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    RegisterReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let registerReq: RegisterReq; //

const { status, data } = await apiInstance.register(
    registerReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **registerReq** | **RegisterReq**|  | |


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

