# MoonbreezeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](MoonbreezeService.md#create) | **POST** /v1/content/moonbreeze/create |  |
| [**createWithHttpInfo**](MoonbreezeService.md#createWithHttpInfo) | **POST** /v1/content/moonbreeze/create |  |
| [**pageMember**](MoonbreezeService.md#pageMember) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pageMemberWithHttpInfo**](MoonbreezeService.md#pageMemberWithHttpInfo) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pagePublic**](MoonbreezeService.md#pagePublic) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pagePublicWithHttpInfo**](MoonbreezeService.md#pagePublicWithHttpInfo) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pageWatching**](MoonbreezeService.md#pageWatching) | **POST** /v1/content/moonbreeze/page-watching |  |
| [**pageWatchingWithHttpInfo**](MoonbreezeService.md#pageWatchingWithHttpInfo) | **POST** /v1/content/moonbreeze/page-watching |  |



## create

> CreateMoonbreezeResp create(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        CreateMoonbreezeReq createMoonbreezeReq = new CreateMoonbreezeReq(); // CreateMoonbreezeReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createMoonbreezeReq(createMoonbreezeReq)
                .build();
            CreateMoonbreezeResp result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#create");
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
| createRequest | [**APIcreateRequest**](MoonbreezeService.md#APIcreateRequest)|-|-|

### Return type

[**CreateMoonbreezeResp**](CreateMoonbreezeResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createWithHttpInfo

> ApiResponse<CreateMoonbreezeResp> createWithHttpInfo(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        CreateMoonbreezeReq createMoonbreezeReq = new CreateMoonbreezeReq(); // CreateMoonbreezeReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createMoonbreezeReq(createMoonbreezeReq)
                .build();
            ApiResponse<CreateMoonbreezeResp> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#create");
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
| createRequest | [**APIcreateRequest**](MoonbreezeService.md#APIcreateRequest)|-|-|

### Return type

ApiResponse<[**CreateMoonbreezeResp**](CreateMoonbreezeResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateRequest"></a>
## APIcreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createMoonbreezeReq** | [**CreateMoonbreezeReq**](CreateMoonbreezeReq.md) |  | |



## pageMember

> PageMoonbreezesResp pageMember(pageMemberRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PageMemberMoonbreezesReq pageMemberMoonbreezesReq = new PageMemberMoonbreezesReq(); // PageMemberMoonbreezesReq | 
        try {
            APIpageMemberRequest request = APIpageMemberRequest.newBuilder()
                .pageMemberMoonbreezesReq(pageMemberMoonbreezesReq)
                .build();
            PageMoonbreezesResp result = apiInstance.pageMember(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageMember");
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
| pageMemberRequest | [**APIpageMemberRequest**](MoonbreezeService.md#APIpageMemberRequest)|-|-|

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageMemberWithHttpInfo

> ApiResponse<PageMoonbreezesResp> pageMemberWithHttpInfo(pageMemberRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PageMemberMoonbreezesReq pageMemberMoonbreezesReq = new PageMemberMoonbreezesReq(); // PageMemberMoonbreezesReq | 
        try {
            APIpageMemberRequest request = APIpageMemberRequest.newBuilder()
                .pageMemberMoonbreezesReq(pageMemberMoonbreezesReq)
                .build();
            ApiResponse<PageMoonbreezesResp> response = apiInstance.pageMemberWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageMember");
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
| pageMemberRequest | [**APIpageMemberRequest**](MoonbreezeService.md#APIpageMemberRequest)|-|-|

### Return type

ApiResponse<[**PageMoonbreezesResp**](PageMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageMemberRequest"></a>
## APIpageMemberRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageMemberMoonbreezesReq** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | |



## pagePublic

> PageMoonbreezesResp pagePublic(pagePublicRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PagePublicMoonbreezesReq pagePublicMoonbreezesReq = new PagePublicMoonbreezesReq(); // PagePublicMoonbreezesReq | 
        try {
            APIpagePublicRequest request = APIpagePublicRequest.newBuilder()
                .pagePublicMoonbreezesReq(pagePublicMoonbreezesReq)
                .build();
            PageMoonbreezesResp result = apiInstance.pagePublic(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pagePublic");
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
| pagePublicRequest | [**APIpagePublicRequest**](MoonbreezeService.md#APIpagePublicRequest)|-|-|

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pagePublicWithHttpInfo

> ApiResponse<PageMoonbreezesResp> pagePublicWithHttpInfo(pagePublicRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PagePublicMoonbreezesReq pagePublicMoonbreezesReq = new PagePublicMoonbreezesReq(); // PagePublicMoonbreezesReq | 
        try {
            APIpagePublicRequest request = APIpagePublicRequest.newBuilder()
                .pagePublicMoonbreezesReq(pagePublicMoonbreezesReq)
                .build();
            ApiResponse<PageMoonbreezesResp> response = apiInstance.pagePublicWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pagePublic");
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
| pagePublicRequest | [**APIpagePublicRequest**](MoonbreezeService.md#APIpagePublicRequest)|-|-|

### Return type

ApiResponse<[**PageMoonbreezesResp**](PageMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpagePublicRequest"></a>
## APIpagePublicRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pagePublicMoonbreezesReq** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | |



## pageWatching

> PageMoonbreezesResp pageWatching(pageWatchingRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PageWatchingMoonbreezesReq pageWatchingMoonbreezesReq = new PageWatchingMoonbreezesReq(); // PageWatchingMoonbreezesReq | 
        try {
            APIpageWatchingRequest request = APIpageWatchingRequest.newBuilder()
                .pageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq)
                .build();
            PageMoonbreezesResp result = apiInstance.pageWatching(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageWatching");
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
| pageWatchingRequest | [**APIpageWatchingRequest**](MoonbreezeService.md#APIpageWatchingRequest)|-|-|

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageWatchingWithHttpInfo

> ApiResponse<PageMoonbreezesResp> pageWatchingWithHttpInfo(pageWatchingRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.MoonbreezeService;
import com.bass.bbs.api.MoonbreezeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        MoonbreezeService apiInstance = new MoonbreezeService(defaultClient);
        PageWatchingMoonbreezesReq pageWatchingMoonbreezesReq = new PageWatchingMoonbreezesReq(); // PageWatchingMoonbreezesReq | 
        try {
            APIpageWatchingRequest request = APIpageWatchingRequest.newBuilder()
                .pageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq)
                .build();
            ApiResponse<PageMoonbreezesResp> response = apiInstance.pageWatchingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageWatching");
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
| pageWatchingRequest | [**APIpageWatchingRequest**](MoonbreezeService.md#APIpageWatchingRequest)|-|-|

### Return type

ApiResponse<[**PageMoonbreezesResp**](PageMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageWatchingRequest"></a>
## APIpageWatchingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageWatchingMoonbreezesReq** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | |


