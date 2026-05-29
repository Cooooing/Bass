# RelationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**relationServiceBlock**](RelationServiceApi.md#relationServiceBlock) | **POST** /v1/user/relation/block |  |
| [**relationServiceBlockWithHttpInfo**](RelationServiceApi.md#relationServiceBlockWithHttpInfo) | **POST** /v1/user/relation/block |  |
| [**relationServiceFollow**](RelationServiceApi.md#relationServiceFollow) | **POST** /v1/user/relation/follow |  |
| [**relationServiceFollowWithHttpInfo**](RelationServiceApi.md#relationServiceFollowWithHttpInfo) | **POST** /v1/user/relation/follow |  |
| [**relationServiceGetStatus**](RelationServiceApi.md#relationServiceGetStatus) | **POST** /v1/user/relation/get-status |  |
| [**relationServiceGetStatusWithHttpInfo**](RelationServiceApi.md#relationServiceGetStatusWithHttpInfo) | **POST** /v1/user/relation/get-status |  |
| [**relationServiceListBlocked**](RelationServiceApi.md#relationServiceListBlocked) | **POST** /v1/user/relation/list-blocked |  |
| [**relationServiceListBlockedWithHttpInfo**](RelationServiceApi.md#relationServiceListBlockedWithHttpInfo) | **POST** /v1/user/relation/list-blocked |  |
| [**relationServiceListFollowers**](RelationServiceApi.md#relationServiceListFollowers) | **POST** /v1/user/relation/list-followers |  |
| [**relationServiceListFollowersWithHttpInfo**](RelationServiceApi.md#relationServiceListFollowersWithHttpInfo) | **POST** /v1/user/relation/list-followers |  |
| [**relationServiceListFollowing**](RelationServiceApi.md#relationServiceListFollowing) | **POST** /v1/user/relation/list-following |  |
| [**relationServiceListFollowingWithHttpInfo**](RelationServiceApi.md#relationServiceListFollowingWithHttpInfo) | **POST** /v1/user/relation/list-following |  |
| [**relationServiceUnblock**](RelationServiceApi.md#relationServiceUnblock) | **POST** /v1/user/relation/unblock |  |
| [**relationServiceUnblockWithHttpInfo**](RelationServiceApi.md#relationServiceUnblockWithHttpInfo) | **POST** /v1/user/relation/unblock |  |
| [**relationServiceUnfollow**](RelationServiceApi.md#relationServiceUnfollow) | **POST** /v1/user/relation/unfollow |  |
| [**relationServiceUnfollowWithHttpInfo**](RelationServiceApi.md#relationServiceUnfollowWithHttpInfo) | **POST** /v1/user/relation/unfollow |  |



## relationServiceBlock

> Object relationServiceBlock(relationServiceBlockRequest)



当前账号拉黑目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        BlockRelationRequest blockRelationRequest = new BlockRelationRequest(); // BlockRelationRequest | 
        try {
            APIrelationServiceBlockRequest request = APIrelationServiceBlockRequest.newBuilder()
                .blockRelationRequest(blockRelationRequest)
                .build();
            Object result = apiInstance.relationServiceBlock(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceBlock");
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
| relationServiceBlockRequest | [**APIrelationServiceBlockRequest**](RelationServiceApi.md#APIrelationServiceBlockRequest)|-|-|

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

## relationServiceBlockWithHttpInfo

> ApiResponse<Object> relationServiceBlockWithHttpInfo(relationServiceBlockRequest)



当前账号拉黑目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        BlockRelationRequest blockRelationRequest = new BlockRelationRequest(); // BlockRelationRequest | 
        try {
            APIrelationServiceBlockRequest request = APIrelationServiceBlockRequest.newBuilder()
                .blockRelationRequest(blockRelationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.relationServiceBlockWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceBlock");
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
| relationServiceBlockRequest | [**APIrelationServiceBlockRequest**](RelationServiceApi.md#APIrelationServiceBlockRequest)|-|-|

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


<a id="APIrelationServiceBlockRequest"></a>
## APIrelationServiceBlockRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **blockRelationRequest** | [**BlockRelationRequest**](BlockRelationRequest.md) |  | |



## relationServiceFollow

> Object relationServiceFollow(relationServiceFollowRequest)



当前账号关注目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        FollowRelationRequest followRelationRequest = new FollowRelationRequest(); // FollowRelationRequest | 
        try {
            APIrelationServiceFollowRequest request = APIrelationServiceFollowRequest.newBuilder()
                .followRelationRequest(followRelationRequest)
                .build();
            Object result = apiInstance.relationServiceFollow(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceFollow");
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
| relationServiceFollowRequest | [**APIrelationServiceFollowRequest**](RelationServiceApi.md#APIrelationServiceFollowRequest)|-|-|

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

## relationServiceFollowWithHttpInfo

> ApiResponse<Object> relationServiceFollowWithHttpInfo(relationServiceFollowRequest)



当前账号关注目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        FollowRelationRequest followRelationRequest = new FollowRelationRequest(); // FollowRelationRequest | 
        try {
            APIrelationServiceFollowRequest request = APIrelationServiceFollowRequest.newBuilder()
                .followRelationRequest(followRelationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.relationServiceFollowWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceFollow");
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
| relationServiceFollowRequest | [**APIrelationServiceFollowRequest**](RelationServiceApi.md#APIrelationServiceFollowRequest)|-|-|

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


<a id="APIrelationServiceFollowRequest"></a>
## APIrelationServiceFollowRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **followRelationRequest** | [**FollowRelationRequest**](FollowRelationRequest.md) |  | |



## relationServiceGetStatus

> GetStatusRelationReply relationServiceGetStatus(relationServiceGetStatusRequest)



查询当前账号与目标账号之间的关系

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        GetStatusRelationRequest getStatusRelationRequest = new GetStatusRelationRequest(); // GetStatusRelationRequest | 
        try {
            APIrelationServiceGetStatusRequest request = APIrelationServiceGetStatusRequest.newBuilder()
                .getStatusRelationRequest(getStatusRelationRequest)
                .build();
            GetStatusRelationReply result = apiInstance.relationServiceGetStatus(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceGetStatus");
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
| relationServiceGetStatusRequest | [**APIrelationServiceGetStatusRequest**](RelationServiceApi.md#APIrelationServiceGetStatusRequest)|-|-|

### Return type

[**GetStatusRelationReply**](GetStatusRelationReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## relationServiceGetStatusWithHttpInfo

> ApiResponse<GetStatusRelationReply> relationServiceGetStatusWithHttpInfo(relationServiceGetStatusRequest)



查询当前账号与目标账号之间的关系

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        GetStatusRelationRequest getStatusRelationRequest = new GetStatusRelationRequest(); // GetStatusRelationRequest | 
        try {
            APIrelationServiceGetStatusRequest request = APIrelationServiceGetStatusRequest.newBuilder()
                .getStatusRelationRequest(getStatusRelationRequest)
                .build();
            ApiResponse<GetStatusRelationReply> response = apiInstance.relationServiceGetStatusWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceGetStatus");
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
| relationServiceGetStatusRequest | [**APIrelationServiceGetStatusRequest**](RelationServiceApi.md#APIrelationServiceGetStatusRequest)|-|-|

### Return type

ApiResponse<[**GetStatusRelationReply**](GetStatusRelationReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrelationServiceGetStatusRequest"></a>
## APIrelationServiceGetStatusRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getStatusRelationRequest** | [**GetStatusRelationRequest**](GetStatusRelationRequest.md) |  | |



## relationServiceListBlocked

> ListBlockedRelationsReply relationServiceListBlocked(relationServiceListBlockedRequest)



分页查询当前账号拉黑的账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListBlockedRelationsRequest listBlockedRelationsRequest = new ListBlockedRelationsRequest(); // ListBlockedRelationsRequest | 
        try {
            APIrelationServiceListBlockedRequest request = APIrelationServiceListBlockedRequest.newBuilder()
                .listBlockedRelationsRequest(listBlockedRelationsRequest)
                .build();
            ListBlockedRelationsReply result = apiInstance.relationServiceListBlocked(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListBlocked");
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
| relationServiceListBlockedRequest | [**APIrelationServiceListBlockedRequest**](RelationServiceApi.md#APIrelationServiceListBlockedRequest)|-|-|

### Return type

[**ListBlockedRelationsReply**](ListBlockedRelationsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## relationServiceListBlockedWithHttpInfo

> ApiResponse<ListBlockedRelationsReply> relationServiceListBlockedWithHttpInfo(relationServiceListBlockedRequest)



分页查询当前账号拉黑的账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListBlockedRelationsRequest listBlockedRelationsRequest = new ListBlockedRelationsRequest(); // ListBlockedRelationsRequest | 
        try {
            APIrelationServiceListBlockedRequest request = APIrelationServiceListBlockedRequest.newBuilder()
                .listBlockedRelationsRequest(listBlockedRelationsRequest)
                .build();
            ApiResponse<ListBlockedRelationsReply> response = apiInstance.relationServiceListBlockedWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListBlocked");
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
| relationServiceListBlockedRequest | [**APIrelationServiceListBlockedRequest**](RelationServiceApi.md#APIrelationServiceListBlockedRequest)|-|-|

### Return type

ApiResponse<[**ListBlockedRelationsReply**](ListBlockedRelationsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrelationServiceListBlockedRequest"></a>
## APIrelationServiceListBlockedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsRequest** | [**ListBlockedRelationsRequest**](ListBlockedRelationsRequest.md) |  | |



## relationServiceListFollowers

> ListFollowersRelationsReply relationServiceListFollowers(relationServiceListFollowersRequest)



分页查询当前账号的粉丝账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListFollowersRelationsRequest listFollowersRelationsRequest = new ListFollowersRelationsRequest(); // ListFollowersRelationsRequest | 
        try {
            APIrelationServiceListFollowersRequest request = APIrelationServiceListFollowersRequest.newBuilder()
                .listFollowersRelationsRequest(listFollowersRelationsRequest)
                .build();
            ListFollowersRelationsReply result = apiInstance.relationServiceListFollowers(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListFollowers");
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
| relationServiceListFollowersRequest | [**APIrelationServiceListFollowersRequest**](RelationServiceApi.md#APIrelationServiceListFollowersRequest)|-|-|

### Return type

[**ListFollowersRelationsReply**](ListFollowersRelationsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## relationServiceListFollowersWithHttpInfo

> ApiResponse<ListFollowersRelationsReply> relationServiceListFollowersWithHttpInfo(relationServiceListFollowersRequest)



分页查询当前账号的粉丝账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListFollowersRelationsRequest listFollowersRelationsRequest = new ListFollowersRelationsRequest(); // ListFollowersRelationsRequest | 
        try {
            APIrelationServiceListFollowersRequest request = APIrelationServiceListFollowersRequest.newBuilder()
                .listFollowersRelationsRequest(listFollowersRelationsRequest)
                .build();
            ApiResponse<ListFollowersRelationsReply> response = apiInstance.relationServiceListFollowersWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListFollowers");
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
| relationServiceListFollowersRequest | [**APIrelationServiceListFollowersRequest**](RelationServiceApi.md#APIrelationServiceListFollowersRequest)|-|-|

### Return type

ApiResponse<[**ListFollowersRelationsReply**](ListFollowersRelationsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrelationServiceListFollowersRequest"></a>
## APIrelationServiceListFollowersRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsRequest** | [**ListFollowersRelationsRequest**](ListFollowersRelationsRequest.md) |  | |



## relationServiceListFollowing

> ListFollowingRelationsReply relationServiceListFollowing(relationServiceListFollowingRequest)



分页查询当前账号关注的账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListFollowingRelationsRequest listFollowingRelationsRequest = new ListFollowingRelationsRequest(); // ListFollowingRelationsRequest | 
        try {
            APIrelationServiceListFollowingRequest request = APIrelationServiceListFollowingRequest.newBuilder()
                .listFollowingRelationsRequest(listFollowingRelationsRequest)
                .build();
            ListFollowingRelationsReply result = apiInstance.relationServiceListFollowing(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListFollowing");
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
| relationServiceListFollowingRequest | [**APIrelationServiceListFollowingRequest**](RelationServiceApi.md#APIrelationServiceListFollowingRequest)|-|-|

### Return type

[**ListFollowingRelationsReply**](ListFollowingRelationsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## relationServiceListFollowingWithHttpInfo

> ApiResponse<ListFollowingRelationsReply> relationServiceListFollowingWithHttpInfo(relationServiceListFollowingRequest)



分页查询当前账号关注的账号列表

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        ListFollowingRelationsRequest listFollowingRelationsRequest = new ListFollowingRelationsRequest(); // ListFollowingRelationsRequest | 
        try {
            APIrelationServiceListFollowingRequest request = APIrelationServiceListFollowingRequest.newBuilder()
                .listFollowingRelationsRequest(listFollowingRelationsRequest)
                .build();
            ApiResponse<ListFollowingRelationsReply> response = apiInstance.relationServiceListFollowingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceListFollowing");
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
| relationServiceListFollowingRequest | [**APIrelationServiceListFollowingRequest**](RelationServiceApi.md#APIrelationServiceListFollowingRequest)|-|-|

### Return type

ApiResponse<[**ListFollowingRelationsReply**](ListFollowingRelationsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrelationServiceListFollowingRequest"></a>
## APIrelationServiceListFollowingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsRequest** | [**ListFollowingRelationsRequest**](ListFollowingRelationsRequest.md) |  | |



## relationServiceUnblock

> Object relationServiceUnblock(relationServiceUnblockRequest)



当前账号取消拉黑目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        UnblockRelationRequest unblockRelationRequest = new UnblockRelationRequest(); // UnblockRelationRequest | 
        try {
            APIrelationServiceUnblockRequest request = APIrelationServiceUnblockRequest.newBuilder()
                .unblockRelationRequest(unblockRelationRequest)
                .build();
            Object result = apiInstance.relationServiceUnblock(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceUnblock");
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
| relationServiceUnblockRequest | [**APIrelationServiceUnblockRequest**](RelationServiceApi.md#APIrelationServiceUnblockRequest)|-|-|

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

## relationServiceUnblockWithHttpInfo

> ApiResponse<Object> relationServiceUnblockWithHttpInfo(relationServiceUnblockRequest)



当前账号取消拉黑目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        UnblockRelationRequest unblockRelationRequest = new UnblockRelationRequest(); // UnblockRelationRequest | 
        try {
            APIrelationServiceUnblockRequest request = APIrelationServiceUnblockRequest.newBuilder()
                .unblockRelationRequest(unblockRelationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.relationServiceUnblockWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceUnblock");
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
| relationServiceUnblockRequest | [**APIrelationServiceUnblockRequest**](RelationServiceApi.md#APIrelationServiceUnblockRequest)|-|-|

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


<a id="APIrelationServiceUnblockRequest"></a>
## APIrelationServiceUnblockRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unblockRelationRequest** | [**UnblockRelationRequest**](UnblockRelationRequest.md) |  | |



## relationServiceUnfollow

> Object relationServiceUnfollow(relationServiceUnfollowRequest)



当前账号取消关注目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        UnfollowRelationRequest unfollowRelationRequest = new UnfollowRelationRequest(); // UnfollowRelationRequest | 
        try {
            APIrelationServiceUnfollowRequest request = APIrelationServiceUnfollowRequest.newBuilder()
                .unfollowRelationRequest(unfollowRelationRequest)
                .build();
            Object result = apiInstance.relationServiceUnfollow(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceUnfollow");
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
| relationServiceUnfollowRequest | [**APIrelationServiceUnfollowRequest**](RelationServiceApi.md#APIrelationServiceUnfollowRequest)|-|-|

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

## relationServiceUnfollowWithHttpInfo

> ApiResponse<Object> relationServiceUnfollowWithHttpInfo(relationServiceUnfollowRequest)



当前账号取消关注目标账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RelationServiceApi;
import com.bass.bbs.api.RelationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RelationServiceApi apiInstance = new RelationServiceApi(defaultClient);
        UnfollowRelationRequest unfollowRelationRequest = new UnfollowRelationRequest(); // UnfollowRelationRequest | 
        try {
            APIrelationServiceUnfollowRequest request = APIrelationServiceUnfollowRequest.newBuilder()
                .unfollowRelationRequest(unfollowRelationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.relationServiceUnfollowWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RelationServiceApi#relationServiceUnfollow");
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
| relationServiceUnfollowRequest | [**APIrelationServiceUnfollowRequest**](RelationServiceApi.md#APIrelationServiceUnfollowRequest)|-|-|

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


<a id="APIrelationServiceUnfollowRequest"></a>
## APIrelationServiceUnfollowRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unfollowRelationRequest** | [**UnfollowRelationRequest**](UnfollowRelationRequest.md) |  | |


