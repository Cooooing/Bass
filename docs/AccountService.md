# AccountService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**avatarAccount**](#avataraccount) | **GET** /v1/user/account/avatar | |
|[**completeProfileImageUploadAccount**](#completeprofileimageuploadaccount) | **POST** /v1/user/account/complete-profile-image-upload | |
|[**getAccountEconomy**](#getaccounteconomy) | **POST** /v1/user/account/economy/get | |
|[**getCurrentAccount**](#getcurrentaccount) | **POST** /v1/user/account/get-current | |
|[**getProfile**](#getprofile) | **POST** /v1/user/account/get-profile | |
|[**listFollowers**](#listfollowers) | **POST** /v1/user/account/list-followers | |
|[**listFollowing**](#listfollowing) | **POST** /v1/user/account/list-following | |
|[**listPointRecords**](#listpointrecords) | **POST** /v1/user/account/list-point-records | |
|[**prepareProfileImageUploadAccount**](#prepareprofileimageuploadaccount) | **POST** /v1/user/account/prepare-profile-image-upload | |
|[**searchMentionAccounts**](#searchmentionaccounts) | **POST** /v1/user/account/search-mention | |
|[**transferAccountPoints**](#transferaccountpoints) | **POST** /v1/user/account/economy/transfer | |
|[**updateEmailAccount**](#updateemailaccount) | **POST** /v1/user/account/update-email | |
|[**updatePasswordAccount**](#updatepasswordaccount) | **POST** /v1/user/account/update-password | |
|[**updatePhoneAccount**](#updatephoneaccount) | **POST** /v1/user/account/update-phone | |
|[**updateProfileAccount**](#updateprofileaccount) | **POST** /v1/user/account/update-profile | |

# **avatarAccount**
> AvatarAccountResp avatarAccount()

生成默认账号头像

### Example

```typescript
import {
    AccountService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let name: string; //账号名，必填 (optional) (default to undefined)

const { status, data } = await apiInstance.avatarAccount(
    name
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | [**string**] | 账号名，必填 | (optional) defaults to undefined|


### Return type

**AvatarAccountResp**

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

# **completeProfileImageUploadAccount**
> CompleteProfileImageUploadAccountResp completeProfileImageUploadAccount(completeProfileImageUploadAccountReq)

将已由 MinIO 回调确认的资源绑定到当前账号资料。

### Example

```typescript
import {
    AccountService,
    Configuration,
    CompleteProfileImageUploadAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let completeProfileImageUploadAccountReq: CompleteProfileImageUploadAccountReq; //

const { status, data } = await apiInstance.completeProfileImageUploadAccount(
    completeProfileImageUploadAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **completeProfileImageUploadAccountReq** | **CompleteProfileImageUploadAccountReq**|  | |


### Return type

**CompleteProfileImageUploadAccountResp**

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

# **getAccountEconomy**
> GetAccountEconomyResp getAccountEconomy(body)

查询当前账号的积分余额和累计收支。

### Example

```typescript
import {
    AccountService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let body: object; //

const { status, data } = await apiInstance.getAccountEconomy(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetAccountEconomyResp**

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

# **getCurrentAccount**
> GetCurrentAccountResp getCurrentAccount(body)

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

const { status, data } = await apiInstance.getCurrentAccount(
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

# **listPointRecords**
> ListPointRecordsResp listPointRecords(listPointRecordsReq)

查询账号公开的积分流水。

### Example

```typescript
import {
    AccountService,
    Configuration,
    ListPointRecordsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let listPointRecordsReq: ListPointRecordsReq; //

const { status, data } = await apiInstance.listPointRecords(
    listPointRecordsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listPointRecordsReq** | **ListPointRecordsReq**|  | |


### Return type

**ListPointRecordsResp**

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

# **prepareProfileImageUploadAccount**
> PrepareProfileImageUploadAccountResp prepareProfileImageUploadAccount(prepareProfileImageUploadAccountReq)

申请当前账号资料图片的内容寻址直传能力。

### Example

```typescript
import {
    AccountService,
    Configuration,
    PrepareProfileImageUploadAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let prepareProfileImageUploadAccountReq: PrepareProfileImageUploadAccountReq; //

const { status, data } = await apiInstance.prepareProfileImageUploadAccount(
    prepareProfileImageUploadAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **prepareProfileImageUploadAccountReq** | **PrepareProfileImageUploadAccountReq**|  | |


### Return type

**PrepareProfileImageUploadAccountResp**

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

# **searchMentionAccounts**
> SearchMentionAccountsResp searchMentionAccounts(searchMentionAccountsReq)

检索可提及账号的候选项。

### Example

```typescript
import {
    AccountService,
    Configuration,
    SearchMentionAccountsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let searchMentionAccountsReq: SearchMentionAccountsReq; //

const { status, data } = await apiInstance.searchMentionAccounts(
    searchMentionAccountsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **searchMentionAccountsReq** | **SearchMentionAccountsReq**|  | |


### Return type

**SearchMentionAccountsResp**

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

# **transferAccountPoints**
> TransferAccountPointsResp transferAccountPoints(transferAccountPointsReq)

向指定账号转移积分。

### Example

```typescript
import {
    AccountService,
    Configuration,
    TransferAccountPointsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AccountService(configuration);

let transferAccountPointsReq: TransferAccountPointsReq; //

const { status, data } = await apiInstance.transferAccountPoints(
    transferAccountPointsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **transferAccountPointsReq** | **TransferAccountPointsReq**|  | |


### Return type

**TransferAccountPointsResp**

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

# **updateEmailAccount**
> object updateEmailAccount(updateEmailAccountReq)

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

const { status, data } = await apiInstance.updateEmailAccount(
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

# **updatePasswordAccount**
> object updatePasswordAccount(updatePasswordAccountReq)

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

const { status, data } = await apiInstance.updatePasswordAccount(
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

# **updatePhoneAccount**
> object updatePhoneAccount(updatePhoneAccountReq)

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

const { status, data } = await apiInstance.updatePhoneAccount(
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

# **updateProfileAccount**
> UpdateProfileAccountResp updateProfileAccount(updateProfileAccountReq)

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

const { status, data } = await apiInstance.updateProfileAccount(
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

