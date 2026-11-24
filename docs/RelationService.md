# RelationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**blockRelation**](RelationService.md#blockRelation) | **POST** /v1/user/relation/block |  |
| [**blockRelationWithHttpInfo**](RelationService.md#blockRelationWithHttpInfo) | **POST** /v1/user/relation/block |  |
| [**followRelation**](RelationService.md#followRelation) | **POST** /v1/user/relation/follow |  |
| [**followRelationWithHttpInfo**](RelationService.md#followRelationWithHttpInfo) | **POST** /v1/user/relation/follow |  |
| [**getStatusRelation**](RelationService.md#getStatusRelation) | **POST** /v1/user/relation/get-status |  |
| [**getStatusRelationWithHttpInfo**](RelationService.md#getStatusRelationWithHttpInfo) | **POST** /v1/user/relation/get-status |  |
| [**listBlockedRelations**](RelationService.md#listBlockedRelations) | **POST** /v1/user/relation/list-blocked |  |
| [**listBlockedRelationsWithHttpInfo**](RelationService.md#listBlockedRelationsWithHttpInfo) | **POST** /v1/user/relation/list-blocked |  |
| [**listFollowersRelations**](RelationService.md#listFollowersRelations) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowersRelationsWithHttpInfo**](RelationService.md#listFollowersRelationsWithHttpInfo) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowingRelations**](RelationService.md#listFollowingRelations) | **POST** /v1/user/relation/list-following |  |
| [**listFollowingRelationsWithHttpInfo**](RelationService.md#listFollowingRelationsWithHttpInfo) | **POST** /v1/user/relation/list-following |  |
| [**unblockRelation**](RelationService.md#unblockRelation) | **POST** /v1/user/relation/unblock |  |
| [**unblockRelationWithHttpInfo**](RelationService.md#unblockRelationWithHttpInfo) | **POST** /v1/user/relation/unblock |  |
| [**unfollowRelation**](RelationService.md#unfollowRelation) | **POST** /v1/user/relation/unfollow |  |
| [**unfollowRelationWithHttpInfo**](RelationService.md#unfollowRelationWithHttpInfo) | **POST** /v1/user/relation/unfollow |  |



## blockRelation

> Object blockRelation(blockRelationRequest)



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
            APIblockRelationRequest request = APIblockRelationRequest.newBuilder()
                .blockRelationReq(blockRelationReq)
                .build();
            Object result = apiInstance.blockRelation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#blockRelation");
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
| blockRelationRequest | [**APIblockRelationRequest**](RelationService.md#APIblockRelationRequest)|-|-|

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

## blockRelationWithHttpInfo

> ApiResponse<Object> blockRelationWithHttpInfo(blockRelationRequest)



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
            APIblockRelationRequest request = APIblockRelationRequest.newBuilder()
                .blockRelationReq(blockRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.blockRelationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#blockRelation");
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
| blockRelationRequest | [**APIblockRelationRequest**](RelationService.md#APIblockRelationRequest)|-|-|

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


<a id="APIblockRelationRequest"></a>
## APIblockRelationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **blockRelationReq** | [**BlockRelationReq**](BlockRelationReq.md) |  | |



## followRelation

> Object followRelation(followRelationRequest)



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
            APIfollowRelationRequest request = APIfollowRelationRequest.newBuilder()
                .followRelationReq(followRelationReq)
                .build();
            Object result = apiInstance.followRelation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#followRelation");
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
| followRelationRequest | [**APIfollowRelationRequest**](RelationService.md#APIfollowRelationRequest)|-|-|

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

## followRelationWithHttpInfo

> ApiResponse<Object> followRelationWithHttpInfo(followRelationRequest)



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
            APIfollowRelationRequest request = APIfollowRelationRequest.newBuilder()
                .followRelationReq(followRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.followRelationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#followRelation");
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
| followRelationRequest | [**APIfollowRelationRequest**](RelationService.md#APIfollowRelationRequest)|-|-|

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


<a id="APIfollowRelationRequest"></a>
## APIfollowRelationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **followRelationReq** | [**FollowRelationReq**](FollowRelationReq.md) |  | |



## getStatusRelation

> GetStatusRelationResp getStatusRelation(getStatusRelationRequest)



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
            APIgetStatusRelationRequest request = APIgetStatusRelationRequest.newBuilder()
                .getStatusRelationReq(getStatusRelationReq)
                .build();
            GetStatusRelationResp result = apiInstance.getStatusRelation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#getStatusRelation");
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
| getStatusRelationRequest | [**APIgetStatusRelationRequest**](RelationService.md#APIgetStatusRelationRequest)|-|-|

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

## getStatusRelationWithHttpInfo

> ApiResponse<GetStatusRelationResp> getStatusRelationWithHttpInfo(getStatusRelationRequest)



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
            APIgetStatusRelationRequest request = APIgetStatusRelationRequest.newBuilder()
                .getStatusRelationReq(getStatusRelationReq)
                .build();
            ApiResponse<GetStatusRelationResp> response = apiInstance.getStatusRelationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#getStatusRelation");
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
| getStatusRelationRequest | [**APIgetStatusRelationRequest**](RelationService.md#APIgetStatusRelationRequest)|-|-|

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


<a id="APIgetStatusRelationRequest"></a>
## APIgetStatusRelationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getStatusRelationReq** | [**GetStatusRelationReq**](GetStatusRelationReq.md) |  | |



## listBlockedRelations

> ListBlockedRelationsResp listBlockedRelations(listBlockedRelationsRequest)



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
            APIlistBlockedRelationsRequest request = APIlistBlockedRelationsRequest.newBuilder()
                .listBlockedRelationsReq(listBlockedRelationsReq)
                .build();
            ListBlockedRelationsResp result = apiInstance.listBlockedRelations(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listBlockedRelations");
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
| listBlockedRelationsRequest | [**APIlistBlockedRelationsRequest**](RelationService.md#APIlistBlockedRelationsRequest)|-|-|

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

## listBlockedRelationsWithHttpInfo

> ApiResponse<ListBlockedRelationsResp> listBlockedRelationsWithHttpInfo(listBlockedRelationsRequest)



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
            APIlistBlockedRelationsRequest request = APIlistBlockedRelationsRequest.newBuilder()
                .listBlockedRelationsReq(listBlockedRelationsReq)
                .build();
            ApiResponse<ListBlockedRelationsResp> response = apiInstance.listBlockedRelationsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listBlockedRelations");
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
| listBlockedRelationsRequest | [**APIlistBlockedRelationsRequest**](RelationService.md#APIlistBlockedRelationsRequest)|-|-|

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


<a id="APIlistBlockedRelationsRequest"></a>
## APIlistBlockedRelationsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsReq** | [**ListBlockedRelationsReq**](ListBlockedRelationsReq.md) |  | |



## listFollowersRelations

> ListFollowersRelationsResp listFollowersRelations(listFollowersRelationsRequest)



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
            APIlistFollowersRelationsRequest request = APIlistFollowersRelationsRequest.newBuilder()
                .listFollowersRelationsReq(listFollowersRelationsReq)
                .build();
            ListFollowersRelationsResp result = apiInstance.listFollowersRelations(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowersRelations");
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
| listFollowersRelationsRequest | [**APIlistFollowersRelationsRequest**](RelationService.md#APIlistFollowersRelationsRequest)|-|-|

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

## listFollowersRelationsWithHttpInfo

> ApiResponse<ListFollowersRelationsResp> listFollowersRelationsWithHttpInfo(listFollowersRelationsRequest)



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
            APIlistFollowersRelationsRequest request = APIlistFollowersRelationsRequest.newBuilder()
                .listFollowersRelationsReq(listFollowersRelationsReq)
                .build();
            ApiResponse<ListFollowersRelationsResp> response = apiInstance.listFollowersRelationsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowersRelations");
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
| listFollowersRelationsRequest | [**APIlistFollowersRelationsRequest**](RelationService.md#APIlistFollowersRelationsRequest)|-|-|

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


<a id="APIlistFollowersRelationsRequest"></a>
## APIlistFollowersRelationsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsReq** | [**ListFollowersRelationsReq**](ListFollowersRelationsReq.md) |  | |



## listFollowingRelations

> ListFollowingRelationsResp listFollowingRelations(listFollowingRelationsRequest)



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
            APIlistFollowingRelationsRequest request = APIlistFollowingRelationsRequest.newBuilder()
                .listFollowingRelationsReq(listFollowingRelationsReq)
                .build();
            ListFollowingRelationsResp result = apiInstance.listFollowingRelations(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowingRelations");
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
| listFollowingRelationsRequest | [**APIlistFollowingRelationsRequest**](RelationService.md#APIlistFollowingRelationsRequest)|-|-|

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

## listFollowingRelationsWithHttpInfo

> ApiResponse<ListFollowingRelationsResp> listFollowingRelationsWithHttpInfo(listFollowingRelationsRequest)



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
            APIlistFollowingRelationsRequest request = APIlistFollowingRelationsRequest.newBuilder()
                .listFollowingRelationsReq(listFollowingRelationsReq)
                .build();
            ApiResponse<ListFollowingRelationsResp> response = apiInstance.listFollowingRelationsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#listFollowingRelations");
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
| listFollowingRelationsRequest | [**APIlistFollowingRelationsRequest**](RelationService.md#APIlistFollowingRelationsRequest)|-|-|

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


<a id="APIlistFollowingRelationsRequest"></a>
## APIlistFollowingRelationsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsReq** | [**ListFollowingRelationsReq**](ListFollowingRelationsReq.md) |  | |



## unblockRelation

> Object unblockRelation(unblockRelationRequest)



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
            APIunblockRelationRequest request = APIunblockRelationRequest.newBuilder()
                .unblockRelationReq(unblockRelationReq)
                .build();
            Object result = apiInstance.unblockRelation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unblockRelation");
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
| unblockRelationRequest | [**APIunblockRelationRequest**](RelationService.md#APIunblockRelationRequest)|-|-|

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

## unblockRelationWithHttpInfo

> ApiResponse<Object> unblockRelationWithHttpInfo(unblockRelationRequest)



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
            APIunblockRelationRequest request = APIunblockRelationRequest.newBuilder()
                .unblockRelationReq(unblockRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.unblockRelationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unblockRelation");
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
| unblockRelationRequest | [**APIunblockRelationRequest**](RelationService.md#APIunblockRelationRequest)|-|-|

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


<a id="APIunblockRelationRequest"></a>
## APIunblockRelationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unblockRelationReq** | [**UnblockRelationReq**](UnblockRelationReq.md) |  | |



## unfollowRelation

> Object unfollowRelation(unfollowRelationRequest)



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
            APIunfollowRelationRequest request = APIunfollowRelationRequest.newBuilder()
                .unfollowRelationReq(unfollowRelationReq)
                .build();
            Object result = apiInstance.unfollowRelation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unfollowRelation");
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
| unfollowRelationRequest | [**APIunfollowRelationRequest**](RelationService.md#APIunfollowRelationRequest)|-|-|

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

## unfollowRelationWithHttpInfo

> ApiResponse<Object> unfollowRelationWithHttpInfo(unfollowRelationRequest)



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
            APIunfollowRelationRequest request = APIunfollowRelationRequest.newBuilder()
                .unfollowRelationReq(unfollowRelationReq)
                .build();
            ApiResponse<Object> response = apiInstance.unfollowRelationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationService#unfollowRelation");
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
| unfollowRelationRequest | [**APIunfollowRelationRequest**](RelationService.md#APIunfollowRelationRequest)|-|-|

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


<a id="APIunfollowRelationRequest"></a>
## APIunfollowRelationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unfollowRelationReq** | [**UnfollowRelationReq**](UnfollowRelationReq.md) |  | |


