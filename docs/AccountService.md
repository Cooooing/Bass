# AccountService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**avatar**](#avatar) | **GET** /v1/user/account/avatar | |
|[**getCurrent**](#getcurrent) | **POST** /v1/user/account/get-current | |
|[**getProfile**](#getprofile) | **POST** /v1/user/account/get-profile | |
|[**listFollowers**](#listfollowers) | **POST** /v1/user/account/list-followers | |
|[**listFollowing**](#listfollowing) | **POST** /v1/user/account/list-following | |
|[**updateEmail**](#updateemail) | **POST** /v1/user/account/update-email | |
|[**updatePassword**](#updatepassword) | **POST** /v1/user/account/update-password | |
|[**updatePhone**](#updatephone) | **POST** /v1/user/account/update-phone | |
|[**updateProfile**](#updateprofile) | **POST** /v1/user/account/update-profile | |

# **avatar**
> ImageResp avatar()

生成默认账号头像

### Example

```typescript
import {
    AccountService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let name: string; // (optional) (default to undefined)

const { status, data } = await apiInstance.avatar(
    name
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | [**string**] |  | (optional) defaults to undefined|


### Return type

**ImageResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCurrent**
> GetCurrentAccountResp getCurrent(body)

获取当前账号完整资料

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

**GetCurrentAccountResp**

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
> GetProfileResp getProfile(getProfileReq)

按账号名获取个人主页资料。

### Example

```typescript
import {
    AccountService,
    Configuration,
    GetProfileReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let getProfileReq: GetProfileReq; //

const { status, data } = await apiInstance.getProfile(
    getProfileReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getProfileReq** | **GetProfileReq**|  | |


### Return type

**GetProfileResp**

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

# **listFollowers**
> ListFollowersResp listFollowers(listFollowersReq)

查询账号公开的粉丝列表。

### Example

```typescript
import {
    AccountService,
    Configuration,
    ListFollowersReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let listFollowersReq: ListFollowersReq; //

const { status, data } = await apiInstance.listFollowers(
    listFollowersReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowersReq** | **ListFollowersReq**|  | |


### Return type

**ListFollowersResp**

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

# **listFollowing**
> ListFollowingResp listFollowing(listFollowingReq)

查询账号公开的关注列表。

### Example

```typescript
import {
    AccountService,
    Configuration,
    ListFollowingReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let listFollowingReq: ListFollowingReq; //

const { status, data } = await apiInstance.listFollowing(
    listFollowingReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowingReq** | **ListFollowingReq**|  | |


### Return type

**ListFollowingResp**

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

# **updateEmail**
> object updateEmail(updateEmailAccountReq)

更新当前账号邮箱

### Example

```typescript
import {
    AccountService,
    Configuration,
    UpdateEmailAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let updateEmailAccountReq: UpdateEmailAccountReq; //

const { status, data } = await apiInstance.updateEmail(
    updateEmailAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateEmailAccountReq** | **UpdateEmailAccountReq**|  | |


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

# **updatePassword**
> object updatePassword(updatePasswordAccountReq)

更新当前账号密码

### Example

```typescript
import {
    AccountService,
    Configuration,
    UpdatePasswordAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let updatePasswordAccountReq: UpdatePasswordAccountReq; //

const { status, data } = await apiInstance.updatePassword(
    updatePasswordAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updatePasswordAccountReq** | **UpdatePasswordAccountReq**|  | |


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

# **updatePhone**
> object updatePhone(updatePhoneAccountReq)

更新当前账号手机号

### Example

```typescript
import {
    AccountService,
    Configuration,
    UpdatePhoneAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let updatePhoneAccountReq: UpdatePhoneAccountReq; //

const { status, data } = await apiInstance.updatePhone(
    updatePhoneAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updatePhoneAccountReq** | **UpdatePhoneAccountReq**|  | |


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

# **updateProfile**
> UpdateProfileAccountResp updateProfile(updateProfileAccountReq)

更新当前账号展示资料

### Example

```typescript
import {
    AccountService,
    Configuration,
    UpdateProfileAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let updateProfileAccountReq: UpdateProfileAccountReq; //

const { status, data } = await apiInstance.updateProfile(
    updateProfileAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateProfileAccountReq** | **UpdateProfileAccountReq**|  | |


### Return type

**UpdateProfileAccountResp**

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

