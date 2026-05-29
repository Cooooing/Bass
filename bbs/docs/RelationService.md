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



当前账号拉黑目标账号

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
        BlockRelationRequest blockRelationRequest = new BlockRelationRequest(); // BlockRelationRequest | 
        try {
            APIblockRequest request = APIblockRequest.newBuilder()
                .blockRelationRequest(blockRelationRequest)
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



当前账号拉黑目标账号

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
        BlockRelationRequest blockRelationRequest = new BlockRelationRequest(); // BlockRelationRequest | 
        try {
            APIblockRequest request = APIblockRequest.newBuilder()
                .blockRelationRequest(blockRelationRequest)
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
| **blockRelationRequest** | [**BlockRelationRequest**](BlockRelationRequest.md) |  | |



## follow

> Object follow(followRequest)



当前账号关注目标账号

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
        FollowRelationRequest followRelationRequest = new FollowRelationRequest(); // FollowRelationRequest | 
        try {
            APIfollowRequest request = APIfollowRequest.newBuilder()
                .followRelationRequest(followRelationRequest)
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



当前账号关注目标账号

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
        FollowRelationRequest followRelationRequest = new FollowRelationRequest(); // FollowRelationRequest | 
        try {
            APIfollowRequest request = APIfollowRequest.newBuilder()
                .followRelationRequest(followRelationRequest)
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
| **followRelationRequest** | [**FollowRelationRequest**](FollowRelationRequest.md) |  | |



## getStatus

> GetStatusRelationReply getStatus(getStatusRequest)



查询当前账号与目标账号之间的关系

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
        GetStatusRelationRequest getStatusRelationRequest = new GetStatusRelationRequest(); // GetStatusRelationRequest | 
        try {
            APIgetStatusRequest request = APIgetStatusRequest.newBuilder()
                .getStatusRelationRequest(getStatusRelationRequest)
                .build();
            GetStatusRelationReply result = apiInstance.getStatus(request);
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

## getStatusWithHttpInfo

> ApiResponse<GetStatusRelationReply> getStatusWithHttpInfo(getStatusRequest)



查询当前账号与目标账号之间的关系

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
        GetStatusRelationRequest getStatusRelationRequest = new GetStatusRelationRequest(); // GetStatusRelationRequest | 
        try {
            APIgetStatusRequest request = APIgetStatusRequest.newBuilder()
                .getStatusRelationRequest(getStatusRelationRequest)
                .build();
            ApiResponse<GetStatusRelationReply> response = apiInstance.getStatusWithHttpInfo(request);
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


<a id="APIgetStatusRequest"></a>
## APIgetStatusRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getStatusRelationRequest** | [**GetStatusRelationRequest**](GetStatusRelationRequest.md) |  | |



## listBlocked

> ListBlockedRelationsReply listBlocked(listBlockedRequest)



分页查询当前账号拉黑的账号列表

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
        ListBlockedRelationsRequest listBlockedRelationsRequest = new ListBlockedRelationsRequest(); // ListBlockedRelationsRequest | 
        try {
            APIlistBlockedRequest request = APIlistBlockedRequest.newBuilder()
                .listBlockedRelationsRequest(listBlockedRelationsRequest)
                .build();
            ListBlockedRelationsReply result = apiInstance.listBlocked(request);
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

## listBlockedWithHttpInfo

> ApiResponse<ListBlockedRelationsReply> listBlockedWithHttpInfo(listBlockedRequest)



分页查询当前账号拉黑的账号列表

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
        ListBlockedRelationsRequest listBlockedRelationsRequest = new ListBlockedRelationsRequest(); // ListBlockedRelationsRequest | 
        try {
            APIlistBlockedRequest request = APIlistBlockedRequest.newBuilder()
                .listBlockedRelationsRequest(listBlockedRelationsRequest)
                .build();
            ApiResponse<ListBlockedRelationsReply> response = apiInstance.listBlockedWithHttpInfo(request);
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


<a id="APIlistBlockedRequest"></a>
## APIlistBlockedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsRequest** | [**ListBlockedRelationsRequest**](ListBlockedRelationsRequest.md) |  | |



## listFollowers

> ListFollowersRelationsReply listFollowers(listFollowersRequest)



分页查询当前账号的粉丝账号列表

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
        ListFollowersRelationsRequest listFollowersRelationsRequest = new ListFollowersRelationsRequest(); // ListFollowersRelationsRequest | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersRelationsRequest(listFollowersRelationsRequest)
                .build();
            ListFollowersRelationsReply result = apiInstance.listFollowers(request);
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

## listFollowersWithHttpInfo

> ApiResponse<ListFollowersRelationsReply> listFollowersWithHttpInfo(listFollowersRequest)



分页查询当前账号的粉丝账号列表

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
        ListFollowersRelationsRequest listFollowersRelationsRequest = new ListFollowersRelationsRequest(); // ListFollowersRelationsRequest | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersRelationsRequest(listFollowersRelationsRequest)
                .build();
            ApiResponse<ListFollowersRelationsReply> response = apiInstance.listFollowersWithHttpInfo(request);
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


<a id="APIlistFollowersRequest"></a>
## APIlistFollowersRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsRequest** | [**ListFollowersRelationsRequest**](ListFollowersRelationsRequest.md) |  | |



## listFollowing

> ListFollowingRelationsReply listFollowing(listFollowingRequest)



分页查询当前账号关注的账号列表

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
        ListFollowingRelationsRequest listFollowingRelationsRequest = new ListFollowingRelationsRequest(); // ListFollowingRelationsRequest | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingRelationsRequest(listFollowingRelationsRequest)
                .build();
            ListFollowingRelationsReply result = apiInstance.listFollowing(request);
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

## listFollowingWithHttpInfo

> ApiResponse<ListFollowingRelationsReply> listFollowingWithHttpInfo(listFollowingRequest)



分页查询当前账号关注的账号列表

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
        ListFollowingRelationsRequest listFollowingRelationsRequest = new ListFollowingRelationsRequest(); // ListFollowingRelationsRequest | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingRelationsRequest(listFollowingRelationsRequest)
                .build();
            ApiResponse<ListFollowingRelationsReply> response = apiInstance.listFollowingWithHttpInfo(request);
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


<a id="APIlistFollowingRequest"></a>
## APIlistFollowingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsRequest** | [**ListFollowingRelationsRequest**](ListFollowingRelationsRequest.md) |  | |



## unblock

> Object unblock(unblockRequest)



当前账号取消拉黑目标账号

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
        UnblockRelationRequest unblockRelationRequest = new UnblockRelationRequest(); // UnblockRelationRequest | 
        try {
            APIunblockRequest request = APIunblockRequest.newBuilder()
                .unblockRelationRequest(unblockRelationRequest)
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



当前账号取消拉黑目标账号

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
        UnblockRelationRequest unblockRelationRequest = new UnblockRelationRequest(); // UnblockRelationRequest | 
        try {
            APIunblockRequest request = APIunblockRequest.newBuilder()
                .unblockRelationRequest(unblockRelationRequest)
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
| **unblockRelationRequest** | [**UnblockRelationRequest**](UnblockRelationRequest.md) |  | |



## unfollow

> Object unfollow(unfollowRequest)



当前账号取消关注目标账号

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
        UnfollowRelationRequest unfollowRelationRequest = new UnfollowRelationRequest(); // UnfollowRelationRequest | 
        try {
            APIunfollowRequest request = APIunfollowRequest.newBuilder()
                .unfollowRelationRequest(unfollowRelationRequest)
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



当前账号取消关注目标账号

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
        UnfollowRelationRequest unfollowRelationRequest = new UnfollowRelationRequest(); // UnfollowRelationRequest | 
        try {
            APIunfollowRequest request = APIunfollowRequest.newBuilder()
                .unfollowRelationRequest(unfollowRelationRequest)
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
| **unfollowRelationRequest** | [**UnfollowRelationRequest**](UnfollowRelationRequest.md) |  | |


