# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatarAccount**](AccountService.md#avataraccount) | **GET** /v1/user/account/avatar |  |
| [**completeProfileImageUploadAccount**](AccountService.md#completeprofileimageuploadaccount) | **POST** /v1/user/account/complete-profile-image-upload |  |
| [**getAccountEconomy**](AccountService.md#getaccounteconomy) | **POST** /v1/user/account/economy/get |  |
| [**getCurrentAccount**](AccountService.md#getcurrentaccount) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getprofile) | **POST** /v1/user/account/get-profile |  |
| [**listFollowers**](AccountService.md#listfollowers) | **POST** /v1/user/account/list-followers |  |
| [**listFollowing**](AccountService.md#listfollowing) | **POST** /v1/user/account/list-following |  |
| [**listPointRecords**](AccountService.md#listpointrecords) | **POST** /v1/user/account/list-point-records |  |
| [**prepareProfileImageUploadAccount**](AccountService.md#prepareprofileimageuploadaccount) | **POST** /v1/user/account/prepare-profile-image-upload |  |
| [**searchMentionAccounts**](AccountService.md#searchmentionaccounts) | **POST** /v1/user/account/search-mention |  |
| [**transferAccountPoints**](AccountService.md#transferaccountpoints) | **POST** /v1/user/account/economy/transfer |  |
| [**updateEmailAccount**](AccountService.md#updateemailaccount) | **POST** /v1/user/account/update-email |  |
| [**updatePasswordAccount**](AccountService.md#updatepasswordaccount) | **POST** /v1/user/account/update-password |  |
| [**updatePhoneAccount**](AccountService.md#updatephoneaccount) | **POST** /v1/user/account/update-phone |  |
| [**updateProfileAccount**](AccountService.md#updateprofileaccount) | **POST** /v1/user/account/update-profile |  |



## avatarAccount

> AvatarAccountResp avatarAccount(name)



生成默认账号头像

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { AvatarAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // string | 账号名，必填 (optional)
    name: name_example,
  } satisfies AvatarAccountRequest;

  try {
    const data = await api.avatarAccount(body);
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
| **name** | `string` | 账号名，必填 | [Optional] [Defaults to `undefined`] |

### Return type

[**AvatarAccountResp**](AvatarAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## completeProfileImageUploadAccount

> CompleteProfileImageUploadAccountResp completeProfileImageUploadAccount(completeProfileImageUploadAccountReq)



将已由 MinIO 回调确认的资源绑定到当前账号资料。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { CompleteProfileImageUploadAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // CompleteProfileImageUploadAccountReq
    completeProfileImageUploadAccountReq: ...,
  } satisfies CompleteProfileImageUploadAccountRequest;

  try {
    const data = await api.completeProfileImageUploadAccount(body);
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
| **completeProfileImageUploadAccountReq** | [CompleteProfileImageUploadAccountReq](CompleteProfileImageUploadAccountReq.md) |  | |

### Return type

[**CompleteProfileImageUploadAccountResp**](CompleteProfileImageUploadAccountResp.md)

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


## getAccountEconomy

> GetAccountEconomyResp getAccountEconomy(body)



查询当前账号的积分余额和累计收支。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { GetAccountEconomyRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // object
    body: Object,
  } satisfies GetAccountEconomyRequest;

  try {
    const data = await api.getAccountEconomy(body);
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

[**GetAccountEconomyResp**](GetAccountEconomyResp.md)

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


## getCurrentAccount

> GetCurrentAccountResp getCurrentAccount(body)



获取当前账号完整资料

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentAccountRequest;

  try {
    const data = await api.getCurrentAccount(body);
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

[**GetCurrentAccountResp**](GetCurrentAccountResp.md)

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


## getProfile

> GetProfileResp getProfile(getProfileReq)



按账号名获取个人主页资料。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { GetProfileRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // GetProfileReq
    getProfileReq: ...,
  } satisfies GetProfileRequest;

  try {
    const data = await api.getProfile(body);
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
| **getProfileReq** | [GetProfileReq](GetProfileReq.md) |  | |

### Return type

[**GetProfileResp**](GetProfileResp.md)

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


## listFollowers

> ListFollowersResp listFollowers(listFollowersReq)



查询账号公开的粉丝列表。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowersRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // ListFollowersReq
    listFollowersReq: ...,
  } satisfies ListFollowersRequest;

  try {
    const data = await api.listFollowers(body);
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
| **listFollowersReq** | [ListFollowersReq](ListFollowersReq.md) |  | |

### Return type

[**ListFollowersResp**](ListFollowersResp.md)

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


## listFollowing

> ListFollowingResp listFollowing(listFollowingReq)



查询账号公开的关注列表。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // ListFollowingReq
    listFollowingReq: ...,
  } satisfies ListFollowingRequest;

  try {
    const data = await api.listFollowing(body);
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
| **listFollowingReq** | [ListFollowingReq](ListFollowingReq.md) |  | |

### Return type

[**ListFollowingResp**](ListFollowingResp.md)

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


## listPointRecords

> ListPointRecordsResp listPointRecords(listPointRecordsReq)



查询账号公开的积分流水。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { ListPointRecordsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // ListPointRecordsReq
    listPointRecordsReq: ...,
  } satisfies ListPointRecordsRequest;

  try {
    const data = await api.listPointRecords(body);
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
| **listPointRecordsReq** | [ListPointRecordsReq](ListPointRecordsReq.md) |  | |

### Return type

[**ListPointRecordsResp**](ListPointRecordsResp.md)

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


## prepareProfileImageUploadAccount

> PrepareProfileImageUploadAccountResp prepareProfileImageUploadAccount(prepareProfileImageUploadAccountReq)



申请当前账号资料图片的内容寻址直传能力。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { PrepareProfileImageUploadAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // PrepareProfileImageUploadAccountReq
    prepareProfileImageUploadAccountReq: ...,
  } satisfies PrepareProfileImageUploadAccountRequest;

  try {
    const data = await api.prepareProfileImageUploadAccount(body);
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
| **prepareProfileImageUploadAccountReq** | [PrepareProfileImageUploadAccountReq](PrepareProfileImageUploadAccountReq.md) |  | |

### Return type

[**PrepareProfileImageUploadAccountResp**](PrepareProfileImageUploadAccountResp.md)

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


## searchMentionAccounts

> SearchMentionAccountsResp searchMentionAccounts(searchMentionAccountsReq)



检索可提及账号的候选项。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { SearchMentionAccountsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // SearchMentionAccountsReq
    searchMentionAccountsReq: ...,
  } satisfies SearchMentionAccountsRequest;

  try {
    const data = await api.searchMentionAccounts(body);
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
| **searchMentionAccountsReq** | [SearchMentionAccountsReq](SearchMentionAccountsReq.md) |  | |

### Return type

[**SearchMentionAccountsResp**](SearchMentionAccountsResp.md)

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


## transferAccountPoints

> TransferAccountPointsResp transferAccountPoints(transferAccountPointsReq)



向指定账号转移积分。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { TransferAccountPointsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // TransferAccountPointsReq
    transferAccountPointsReq: ...,
  } satisfies TransferAccountPointsRequest;

  try {
    const data = await api.transferAccountPoints(body);
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
| **transferAccountPointsReq** | [TransferAccountPointsReq](TransferAccountPointsReq.md) |  | |

### Return type

[**TransferAccountPointsResp**](TransferAccountPointsResp.md)

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


## updateEmailAccount

> object updateEmailAccount(updateEmailAccountReq)



更新当前账号邮箱

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateEmailAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // UpdateEmailAccountReq
    updateEmailAccountReq: ...,
  } satisfies UpdateEmailAccountRequest;

  try {
    const data = await api.updateEmailAccount(body);
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
| **updateEmailAccountReq** | [UpdateEmailAccountReq](UpdateEmailAccountReq.md) |  | |

### Return type

**object**

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


## updatePasswordAccount

> object updatePasswordAccount(updatePasswordAccountReq)



更新当前账号密码

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { UpdatePasswordAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // UpdatePasswordAccountReq
    updatePasswordAccountReq: ...,
  } satisfies UpdatePasswordAccountRequest;

  try {
    const data = await api.updatePasswordAccount(body);
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
| **updatePasswordAccountReq** | [UpdatePasswordAccountReq](UpdatePasswordAccountReq.md) |  | |

### Return type

**object**

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


## updatePhoneAccount

> object updatePhoneAccount(updatePhoneAccountReq)



更新当前账号手机号

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { UpdatePhoneAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // UpdatePhoneAccountReq
    updatePhoneAccountReq: ...,
  } satisfies UpdatePhoneAccountRequest;

  try {
    const data = await api.updatePhoneAccount(body);
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
| **updatePhoneAccountReq** | [UpdatePhoneAccountReq](UpdatePhoneAccountReq.md) |  | |

### Return type

**object**

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


## updateProfileAccount

> UpdateProfileAccountResp updateProfileAccount(updateProfileAccountReq)



更新当前账号展示资料

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateProfileAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // UpdateProfileAccountReq
    updateProfileAccountReq: ...,
  } satisfies UpdateProfileAccountRequest;

  try {
    const data = await api.updateProfileAccount(body);
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
| **updateProfileAccountReq** | [UpdateProfileAccountReq](UpdateProfileAccountReq.md) |  | |

### Return type

[**UpdateProfileAccountResp**](UpdateProfileAccountResp.md)

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

