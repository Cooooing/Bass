# RelationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**block**](RelationService.md#block) | **POST** /v1/user/relation/block |  |
| [**blockWithHttpInfo**](RelationService.md#blockWithHttpInfo) | **POST** /v1/user/relation/block |  |
| [**follow**](RelationService.md#follow) | **POST** /v1/user/relation/follow |  |
| [**followWithHttpInfo**](RelationService.md#followWithHttpInfo) | **POST** /v1/user/relation/follow |  |
| [**getStatus**](RelationService.md#getStatus) | **POST** /v1/user/relation/get-status |  |
| [**getStatusWithHttpInfo**](RelationService.md#getStatusWithHttpInfo) | **POST** /v1/user/relation/get-status |  |
| [**listBlocked**](RelationService.md#listBlocked) | **POST** /v1/user/relation/list-blocked |  |
| [**listBlockedWithHttpInfo**](RelationService.md#listBlockedWithHttpInfo) | **POST** /v1/user/relation/list-blocked |  |
| [**listFollowers**](RelationService.md#listFollowers) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowersWithHttpInfo**](RelationService.md#listFollowersWithHttpInfo) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowing**](RelationService.md#listFollowing) | **POST** /v1/user/relation/list-following |  |
| [**listFollowingWithHttpInfo**](RelationService.md#listFollowingWithHttpInfo) | **POST** /v1/user/relation/list-following |  |
| [**unblock**](RelationService.md#unblock) | **POST** /v1/user/relation/unblock |  |
| [**unblockWithHttpInfo**](RelationService.md#unblockWithHttpInfo) | **POST** /v1/user/relation/unblock |  |
| [**unfollow**](RelationService.md#unfollow) | **POST** /v1/user/relation/unfollow |  |
| [**unfollowWithHttpInfo**](RelationService.md#unfollowWithHttpInfo) | **POST** /v1/user/relation/unfollow |  |



## block

> Object block(blockRequest)



当前账号拉黑目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        BlockRelationReq blockRelationReq = new BlockRelationReq(); // BlockRelationReq | 
        try {
            APIblockRequest request = APIblockRequest.newBuilder()
                .blockRelationReq(blockRelationReq)
                .build();
            Object result = apiInstance.block(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#block");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| blockRequest | [**APIblockRequest**](RelationService.md#APIblockRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## blockWithHttpInfo

> ApiResponse<Object> blockWithHttpInfo(blockRequest)



当前账号拉黑目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        BlockRelationReq blockRelationReq = new BlockRelationReq(); // BlockRelationReq | 
        try {
            APIblockRequest request = APIblockRequest.newBuilder()
                .blockRelationReq(blockRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.blockWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#block");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| blockRequest | [**APIblockRequest**](RelationService.md#APIblockRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIblockRequest"></a>
## APIblockRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **blockRelationReq** | [**BlockRelationReq**](BlockRelationReq.md) |  | |



## follow

> Object follow(followRequest)



当前账号关注目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        FollowRelationReq followRelationReq = new FollowRelationReq(); // FollowRelationReq | 
        try {
            APIfollowRequest request = APIfollowRequest.newBuilder()
                .followRelationReq(followRelationReq)
                .build();
            Object result = apiInstance.follow(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#follow");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| followRequest | [**APIfollowRequest**](RelationService.md#APIfollowRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## followWithHttpInfo

> ApiResponse<Object> followWithHttpInfo(followRequest)



当前账号关注目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        FollowRelationReq followRelationReq = new FollowRelationReq(); // FollowRelationReq | 
        try {
            APIfollowRequest request = APIfollowRequest.newBuilder()
                .followRelationReq(followRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.followWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#follow");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| followRequest | [**APIfollowRequest**](RelationService.md#APIfollowRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIfollowRequest"></a>
## APIfollowRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **followRelationReq** | [**FollowRelationReq**](FollowRelationReq.md) |  | |



## getStatus

> GetStatusRelationResp getStatus(getStatusRequest)



查询当前账号与目标账号之间的关系。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        GetStatusRelationReq getStatusRelationReq = new GetStatusRelationReq(); // GetStatusRelationReq | 
        try {
            APIgetStatusRequest request = APIgetStatusRequest.newBuilder()
                .getStatusRelationReq(getStatusRelationReq)
                .build();
            GetStatusRelationResp result = apiInstance.getStatus(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#getStatus");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| getStatusRequest | [**APIgetStatusRequest**](RelationService.md#APIgetStatusRequest)|-|-|

### Return type

[**GetStatusRelationResp**](GetStatusRelationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getStatusWithHttpInfo

> ApiResponse<GetStatusRelationResp> getStatusWithHttpInfo(getStatusRequest)



查询当前账号与目标账号之间的关系。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        GetStatusRelationReq getStatusRelationReq = new GetStatusRelationReq(); // GetStatusRelationReq | 
        try {
            APIgetStatusRequest request = APIgetStatusRequest.newBuilder()
                .getStatusRelationReq(getStatusRelationReq)
                .build();
            ApiResponse<GetStatusRelationResp> response = apiInstance.getStatusWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#getStatus");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| getStatusRequest | [**APIgetStatusRequest**](RelationService.md#APIgetStatusRequest)|-|-|

### Return type

ApiResponse<[**GetStatusRelationResp**](GetStatusRelationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetStatusRequest"></a>
## APIgetStatusRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getStatusRelationReq** | [**GetStatusRelationReq**](GetStatusRelationReq.md) |  | |



## listBlocked

> ListBlockedRelationsResp listBlocked(listBlockedRequest)



分页查询当前账号拉黑的账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListBlockedRelationsReq listBlockedRelationsReq = new ListBlockedRelationsReq(); // ListBlockedRelationsReq | 
        try {
            APIlistBlockedRequest request = APIlistBlockedRequest.newBuilder()
                .listBlockedRelationsReq(listBlockedRelationsReq)
                .build();
            ListBlockedRelationsResp result = apiInstance.listBlocked(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listBlocked");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listBlockedRequest | [**APIlistBlockedRequest**](RelationService.md#APIlistBlockedRequest)|-|-|

### Return type

[**ListBlockedRelationsResp**](ListBlockedRelationsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listBlockedWithHttpInfo

> ApiResponse<ListBlockedRelationsResp> listBlockedWithHttpInfo(listBlockedRequest)



分页查询当前账号拉黑的账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListBlockedRelationsReq listBlockedRelationsReq = new ListBlockedRelationsReq(); // ListBlockedRelationsReq | 
        try {
            APIlistBlockedRequest request = APIlistBlockedRequest.newBuilder()
                .listBlockedRelationsReq(listBlockedRelationsReq)
                .build();
            ApiResponse<ListBlockedRelationsResp> response = apiInstance.listBlockedWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listBlocked");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listBlockedRequest | [**APIlistBlockedRequest**](RelationService.md#APIlistBlockedRequest)|-|-|

### Return type

ApiResponse<[**ListBlockedRelationsResp**](ListBlockedRelationsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistBlockedRequest"></a>
## APIlistBlockedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsReq** | [**ListBlockedRelationsReq**](ListBlockedRelationsReq.md) |  | |



## listFollowers

> ListFollowersRelationsResp listFollowers(listFollowersRequest)



分页查询当前账号的粉丝账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListFollowersRelationsReq listFollowersRelationsReq = new ListFollowersRelationsReq(); // ListFollowersRelationsReq | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersRelationsReq(listFollowersRelationsReq)
                .build();
            ListFollowersRelationsResp result = apiInstance.listFollowers(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowers");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listFollowersRequest | [**APIlistFollowersRequest**](RelationService.md#APIlistFollowersRequest)|-|-|

### Return type

[**ListFollowersRelationsResp**](ListFollowersRelationsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listFollowersWithHttpInfo

> ApiResponse<ListFollowersRelationsResp> listFollowersWithHttpInfo(listFollowersRequest)



分页查询当前账号的粉丝账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListFollowersRelationsReq listFollowersRelationsReq = new ListFollowersRelationsReq(); // ListFollowersRelationsReq | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersRelationsReq(listFollowersRelationsReq)
                .build();
            ApiResponse<ListFollowersRelationsResp> response = apiInstance.listFollowersWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowers");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listFollowersRequest | [**APIlistFollowersRequest**](RelationService.md#APIlistFollowersRequest)|-|-|

### Return type

ApiResponse<[**ListFollowersRelationsResp**](ListFollowersRelationsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistFollowersRequest"></a>
## APIlistFollowersRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsReq** | [**ListFollowersRelationsReq**](ListFollowersRelationsReq.md) |  | |



## listFollowing

> ListFollowingRelationsResp listFollowing(listFollowingRequest)



分页查询当前账号关注的账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListFollowingRelationsReq listFollowingRelationsReq = new ListFollowingRelationsReq(); // ListFollowingRelationsReq | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingRelationsReq(listFollowingRelationsReq)
                .build();
            ListFollowingRelationsResp result = apiInstance.listFollowing(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowing");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listFollowingRequest | [**APIlistFollowingRequest**](RelationService.md#APIlistFollowingRequest)|-|-|

### Return type

[**ListFollowingRelationsResp**](ListFollowingRelationsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listFollowingWithHttpInfo

> ApiResponse<ListFollowingRelationsResp> listFollowingWithHttpInfo(listFollowingRequest)



分页查询当前账号关注的账号列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        ListFollowingRelationsReq listFollowingRelationsReq = new ListFollowingRelationsReq(); // ListFollowingRelationsReq | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingRelationsReq(listFollowingRelationsReq)
                .build();
            ApiResponse<ListFollowingRelationsResp> response = apiInstance.listFollowingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowing");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| listFollowingRequest | [**APIlistFollowingRequest**](RelationService.md#APIlistFollowingRequest)|-|-|

### Return type

ApiResponse<[**ListFollowingRelationsResp**](ListFollowingRelationsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistFollowingRequest"></a>
## APIlistFollowingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsReq** | [**ListFollowingRelationsReq**](ListFollowingRelationsReq.md) |  | |



## unblock

> Object unblock(unblockRequest)



当前账号取消拉黑目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        UnblockRelationReq unblockRelationReq = new UnblockRelationReq(); // UnblockRelationReq | 
        try {
            APIunblockRequest request = APIunblockRequest.newBuilder()
                .unblockRelationReq(unblockRelationReq)
                .build();
            Object result = apiInstance.unblock(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unblock");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| unblockRequest | [**APIunblockRequest**](RelationService.md#APIunblockRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## unblockWithHttpInfo

> ApiResponse<Object> unblockWithHttpInfo(unblockRequest)



当前账号取消拉黑目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        UnblockRelationReq unblockRelationReq = new UnblockRelationReq(); // UnblockRelationReq | 
        try {
            APIunblockRequest request = APIunblockRequest.newBuilder()
                .unblockRelationReq(unblockRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.unblockWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unblock");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| unblockRequest | [**APIunblockRequest**](RelationService.md#APIunblockRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIunblockRequest"></a>
## APIunblockRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unblockRelationReq** | [**UnblockRelationReq**](UnblockRelationReq.md) |  | |



## unfollow

> Object unfollow(unfollowRequest)



当前账号取消关注目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        UnfollowRelationReq unfollowRelationReq = new UnfollowRelationReq(); // UnfollowRelationReq | 
        try {
            APIunfollowRequest request = APIunfollowRequest.newBuilder()
                .unfollowRelationReq(unfollowRelationReq)
                .build();
            Object result = apiInstance.unfollow(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unfollow");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| unfollowRequest | [**APIunfollowRequest**](RelationService.md#APIunfollowRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## unfollowWithHttpInfo

> ApiResponse<Object> unfollowWithHttpInfo(unfollowRequest)



当前账号取消关注目标账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationService;
import com.bass.bbs.api.RelationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationService apiInstance = new RelationService(defaultClient);
        UnfollowRelationReq unfollowRelationReq = new UnfollowRelationReq(); // UnfollowRelationReq | 
        try {
            APIunfollowRequest request = APIunfollowRequest.newBuilder()
                .unfollowRelationReq(unfollowRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.unfollowWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unfollow");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| unfollowRequest | [**APIunfollowRequest**](RelationService.md#APIunfollowRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIunfollowRequest"></a>
## APIunfollowRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unfollowRelationReq** | [**UnfollowRelationReq**](UnfollowRelationReq.md) |  | |


