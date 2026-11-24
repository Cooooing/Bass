# RelationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**blockRelation**](RelationService.md#blockrelation) | **POST** /v1/user/relation/block |  |
| [**followRelation**](RelationService.md#followrelation) | **POST** /v1/user/relation/follow |  |
| [**getStatusRelation**](RelationService.md#getstatusrelation) | **POST** /v1/user/relation/get-status |  |
| [**listBlockedRelations**](RelationService.md#listblockedrelations) | **POST** /v1/user/relation/list-blocked |  |
| [**listFollowersRelations**](RelationService.md#listfollowersrelations) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowingRelations**](RelationService.md#listfollowingrelations) | **POST** /v1/user/relation/list-following |  |
| [**unblockRelation**](RelationService.md#unblockrelation) | **POST** /v1/user/relation/unblock |  |
| [**unfollowRelation**](RelationService.md#unfollowrelation) | **POST** /v1/user/relation/unfollow |  |



## blockRelation

> object blockRelation(blockRelationReq)



当前账号拉黑目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { BlockRelationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // BlockRelationReq
    blockRelationReq: ...,
  } satisfies BlockRelationRequest;

  try {
    const data = await api.blockRelation(body);
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
| **blockRelationReq** | [BlockRelationReq](BlockRelationReq.md) |  | |

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


## followRelation

> object followRelation(followRelationReq)



当前账号关注目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { FollowRelationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // FollowRelationReq
    followRelationReq: ...,
  } satisfies FollowRelationRequest;

  try {
    const data = await api.followRelation(body);
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
| **followRelationReq** | [FollowRelationReq](FollowRelationReq.md) |  | |

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


## getStatusRelation

> GetStatusRelationResp getStatusRelation(getStatusRelationReq)



查询当前账号与目标账号之间的关系。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { GetStatusRelationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // GetStatusRelationReq
    getStatusRelationReq: ...,
  } satisfies GetStatusRelationRequest;

  try {
    const data = await api.getStatusRelation(body);
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
| **getStatusRelationReq** | [GetStatusRelationReq](GetStatusRelationReq.md) |  | |

### Return type

[**GetStatusRelationResp**](GetStatusRelationResp.md)

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


## listBlockedRelations

> ListBlockedRelationsResp listBlockedRelations(listBlockedRelationsReq)



分页查询当前账号拉黑的账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListBlockedRelationsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListBlockedRelationsReq
    listBlockedRelationsReq: ...,
  } satisfies ListBlockedRelationsRequest;

  try {
    const data = await api.listBlockedRelations(body);
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
| **listBlockedRelationsReq** | [ListBlockedRelationsReq](ListBlockedRelationsReq.md) |  | |

### Return type

[**ListBlockedRelationsResp**](ListBlockedRelationsResp.md)

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


## listFollowersRelations

> ListFollowersRelationsResp listFollowersRelations(listFollowersRelationsReq)



分页查询当前账号的粉丝账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowersRelationsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListFollowersRelationsReq
    listFollowersRelationsReq: ...,
  } satisfies ListFollowersRelationsRequest;

  try {
    const data = await api.listFollowersRelations(body);
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
| **listFollowersRelationsReq** | [ListFollowersRelationsReq](ListFollowersRelationsReq.md) |  | |

### Return type

[**ListFollowersRelationsResp**](ListFollowersRelationsResp.md)

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


## listFollowingRelations

> ListFollowingRelationsResp listFollowingRelations(listFollowingRelationsReq)



分页查询当前账号关注的账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowingRelationsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListFollowingRelationsReq
    listFollowingRelationsReq: ...,
  } satisfies ListFollowingRelationsRequest;

  try {
    const data = await api.listFollowingRelations(body);
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
| **listFollowingRelationsReq** | [ListFollowingRelationsReq](ListFollowingRelationsReq.md) |  | |

### Return type

[**ListFollowingRelationsResp**](ListFollowingRelationsResp.md)

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


## unblockRelation

> object unblockRelation(unblockRelationReq)



当前账号取消拉黑目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { UnblockRelationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // UnblockRelationReq
    unblockRelationReq: ...,
  } satisfies UnblockRelationRequest;

  try {
    const data = await api.unblockRelation(body);
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
| **unblockRelationReq** | [UnblockRelationReq](UnblockRelationReq.md) |  | |

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


## unfollowRelation

> object unfollowRelation(unfollowRelationReq)



当前账号取消关注目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { UnfollowRelationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // UnfollowRelationReq
    unfollowRelationReq: ...,
  } satisfies UnfollowRelationRequest;

  try {
    const data = await api.unfollowRelation(body);
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
| **unfollowRelationReq** | [UnfollowRelationReq](UnfollowRelationReq.md) |  | |

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

