# RelationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**relationServiceBatchGetStatus**](RelationServiceApi.md#relationservicebatchgetstatus) | **POST** /v1/user/relation/batch-get-status |  |
| [**relationServiceBlock**](RelationServiceApi.md#relationserviceblock) | **POST** /v1/user/relation/block |  |
| [**relationServiceFollow**](RelationServiceApi.md#relationservicefollow) | **POST** /v1/user/relation/follow |  |
| [**relationServicePageBlocked**](RelationServiceApi.md#relationservicepageblocked) | **POST** /v1/user/relation/page-blocked |  |
| [**relationServicePageFollowers**](RelationServiceApi.md#relationservicepagefollowers) | **POST** /v1/user/relation/page-followers |  |
| [**relationServicePageFollowing**](RelationServiceApi.md#relationservicepagefollowing) | **POST** /v1/user/relation/page-following |  |
| [**relationServiceUnblock**](RelationServiceApi.md#relationserviceunblock) | **POST** /v1/user/relation/unblock |  |
| [**relationServiceUnfollow**](RelationServiceApi.md#relationserviceunfollow) | **POST** /v1/user/relation/unfollow |  |



## relationServiceBatchGetStatus

> CommonApiAppBbsV1UserBatchGetStatusRelationReply relationServiceBatchGetStatus(commonApiAppBbsV1UserBatchGetStatusRelationRequest)



批量查询当前账号与多个目标账号之间的关系状态

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceBatchGetStatusRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserBatchGetStatusRelationRequest
    commonApiAppBbsV1UserBatchGetStatusRelationRequest: ...,
  } satisfies RelationServiceBatchGetStatusRequest;

  try {
    const data = await api.relationServiceBatchGetStatus(body);
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
| **commonApiAppBbsV1UserBatchGetStatusRelationRequest** | [CommonApiAppBbsV1UserBatchGetStatusRelationRequest](CommonApiAppBbsV1UserBatchGetStatusRelationRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserBatchGetStatusRelationReply**](CommonApiAppBbsV1UserBatchGetStatusRelationReply.md)

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


## relationServiceBlock

> object relationServiceBlock(commonApiAppBbsV1UserBlockRelationRequest)



当前账号拉黑目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceBlockRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserBlockRelationRequest
    commonApiAppBbsV1UserBlockRelationRequest: ...,
  } satisfies RelationServiceBlockRequest;

  try {
    const data = await api.relationServiceBlock(body);
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
| **commonApiAppBbsV1UserBlockRelationRequest** | [CommonApiAppBbsV1UserBlockRelationRequest](CommonApiAppBbsV1UserBlockRelationRequest.md) |  | |

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


## relationServiceFollow

> object relationServiceFollow(commonApiAppBbsV1UserFollowRelationRequest)



当前账号关注目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceFollowRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserFollowRelationRequest
    commonApiAppBbsV1UserFollowRelationRequest: ...,
  } satisfies RelationServiceFollowRequest;

  try {
    const data = await api.relationServiceFollow(body);
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
| **commonApiAppBbsV1UserFollowRelationRequest** | [CommonApiAppBbsV1UserFollowRelationRequest](CommonApiAppBbsV1UserFollowRelationRequest.md) |  | |

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


## relationServicePageBlocked

> CommonApiAppBbsV1UserPageBlockedRelationReply relationServicePageBlocked(commonApiAppBbsV1UserPageBlockedRelationRequest)



分页查询当前账号拉黑的账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServicePageBlockedRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserPageBlockedRelationRequest
    commonApiAppBbsV1UserPageBlockedRelationRequest: ...,
  } satisfies RelationServicePageBlockedRequest;

  try {
    const data = await api.relationServicePageBlocked(body);
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
| **commonApiAppBbsV1UserPageBlockedRelationRequest** | [CommonApiAppBbsV1UserPageBlockedRelationRequest](CommonApiAppBbsV1UserPageBlockedRelationRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserPageBlockedRelationReply**](CommonApiAppBbsV1UserPageBlockedRelationReply.md)

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


## relationServicePageFollowers

> CommonApiAppBbsV1UserPageFollowersRelationReply relationServicePageFollowers(commonApiAppBbsV1UserPageFollowersRelationRequest)



分页查询当前账号的粉丝账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServicePageFollowersRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserPageFollowersRelationRequest
    commonApiAppBbsV1UserPageFollowersRelationRequest: ...,
  } satisfies RelationServicePageFollowersRequest;

  try {
    const data = await api.relationServicePageFollowers(body);
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
| **commonApiAppBbsV1UserPageFollowersRelationRequest** | [CommonApiAppBbsV1UserPageFollowersRelationRequest](CommonApiAppBbsV1UserPageFollowersRelationRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserPageFollowersRelationReply**](CommonApiAppBbsV1UserPageFollowersRelationReply.md)

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


## relationServicePageFollowing

> CommonApiAppBbsV1UserPageFollowingRelationReply relationServicePageFollowing(commonApiAppBbsV1UserPageFollowingRelationRequest)



分页查询当前账号关注的账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServicePageFollowingRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserPageFollowingRelationRequest
    commonApiAppBbsV1UserPageFollowingRelationRequest: ...,
  } satisfies RelationServicePageFollowingRequest;

  try {
    const data = await api.relationServicePageFollowing(body);
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
| **commonApiAppBbsV1UserPageFollowingRelationRequest** | [CommonApiAppBbsV1UserPageFollowingRelationRequest](CommonApiAppBbsV1UserPageFollowingRelationRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserPageFollowingRelationReply**](CommonApiAppBbsV1UserPageFollowingRelationReply.md)

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


## relationServiceUnblock

> object relationServiceUnblock(commonApiAppBbsV1UserUnblockRelationRequest)



当前账号取消拉黑目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceUnblockRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserUnblockRelationRequest
    commonApiAppBbsV1UserUnblockRelationRequest: ...,
  } satisfies RelationServiceUnblockRequest;

  try {
    const data = await api.relationServiceUnblock(body);
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
| **commonApiAppBbsV1UserUnblockRelationRequest** | [CommonApiAppBbsV1UserUnblockRelationRequest](CommonApiAppBbsV1UserUnblockRelationRequest.md) |  | |

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


## relationServiceUnfollow

> object relationServiceUnfollow(commonApiAppBbsV1UserUnfollowRelationRequest)



当前账号取消关注目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceUnfollowRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserUnfollowRelationRequest
    commonApiAppBbsV1UserUnfollowRelationRequest: ...,
  } satisfies RelationServiceUnfollowRequest;

  try {
    const data = await api.relationServiceUnfollow(body);
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
| **commonApiAppBbsV1UserUnfollowRelationRequest** | [CommonApiAppBbsV1UserUnfollowRelationRequest](CommonApiAppBbsV1UserUnfollowRelationRequest.md) |  | |

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

