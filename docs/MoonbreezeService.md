# MoonbreezeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createMoonbreeze**](MoonbreezeService.md#createMoonbreeze) | **POST** /v1/content/moonbreeze/create |  |
| [**createMoonbreezeWithHttpInfo**](MoonbreezeService.md#createMoonbreezeWithHttpInfo) | **POST** /v1/content/moonbreeze/create |  |
| [**pageMemberMoonbreezes**](MoonbreezeService.md#pageMemberMoonbreezes) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pageMemberMoonbreezesWithHttpInfo**](MoonbreezeService.md#pageMemberMoonbreezesWithHttpInfo) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pagePublicMoonbreezes**](MoonbreezeService.md#pagePublicMoonbreezes) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pagePublicMoonbreezesWithHttpInfo**](MoonbreezeService.md#pagePublicMoonbreezesWithHttpInfo) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pageWatchingMoonbreezes**](MoonbreezeService.md#pageWatchingMoonbreezes) | **POST** /v1/content/moonbreeze/page-watching |  |
| [**pageWatchingMoonbreezesWithHttpInfo**](MoonbreezeService.md#pageWatchingMoonbreezesWithHttpInfo) | **POST** /v1/content/moonbreeze/page-watching |  |



## createMoonbreeze

> CreateMoonbreezeResp createMoonbreeze(createMoonbreezeRequest)



服务方法。

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
            APIcreateMoonbreezeRequest request = APIcreateMoonbreezeRequest.newBuilder()
                .createMoonbreezeReq(createMoonbreezeReq)
                .build();
            CreateMoonbreezeResp result = apiInstance.createMoonbreeze(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#createMoonbreeze");
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
| createMoonbreezeRequest | [**APIcreateMoonbreezeRequest**](MoonbreezeService.md#APIcreateMoonbreezeRequest)|-|-|

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

## createMoonbreezeWithHttpInfo

> ApiResponse<CreateMoonbreezeResp> createMoonbreezeWithHttpInfo(createMoonbreezeRequest)



服务方法。

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
            APIcreateMoonbreezeRequest request = APIcreateMoonbreezeRequest.newBuilder()
                .createMoonbreezeReq(createMoonbreezeReq)
                .build();
            ApiResponse<CreateMoonbreezeResp> response = apiInstance.createMoonbreezeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#createMoonbreeze");
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
| createMoonbreezeRequest | [**APIcreateMoonbreezeRequest**](MoonbreezeService.md#APIcreateMoonbreezeRequest)|-|-|

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


<a id="APIcreateMoonbreezeRequest"></a>
## APIcreateMoonbreezeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createMoonbreezeReq** | [**CreateMoonbreezeReq**](CreateMoonbreezeReq.md) |  | |



## pageMemberMoonbreezes

> PageMemberMoonbreezesResp pageMemberMoonbreezes(pageMemberMoonbreezesRequest)



服务方法。

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
            APIpageMemberMoonbreezesRequest request = APIpageMemberMoonbreezesRequest.newBuilder()
                .pageMemberMoonbreezesReq(pageMemberMoonbreezesReq)
                .build();
            PageMemberMoonbreezesResp result = apiInstance.pageMemberMoonbreezes(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageMemberMoonbreezes");
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
| pageMemberMoonbreezesRequest | [**APIpageMemberMoonbreezesRequest**](MoonbreezeService.md#APIpageMemberMoonbreezesRequest)|-|-|

### Return type

[**PageMemberMoonbreezesResp**](PageMemberMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageMemberMoonbreezesWithHttpInfo

> ApiResponse<PageMemberMoonbreezesResp> pageMemberMoonbreezesWithHttpInfo(pageMemberMoonbreezesRequest)



服务方法。

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
            APIpageMemberMoonbreezesRequest request = APIpageMemberMoonbreezesRequest.newBuilder()
                .pageMemberMoonbreezesReq(pageMemberMoonbreezesReq)
                .build();
            ApiResponse<PageMemberMoonbreezesResp> response = apiInstance.pageMemberMoonbreezesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageMemberMoonbreezes");
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
| pageMemberMoonbreezesRequest | [**APIpageMemberMoonbreezesRequest**](MoonbreezeService.md#APIpageMemberMoonbreezesRequest)|-|-|

### Return type

ApiResponse<[**PageMemberMoonbreezesResp**](PageMemberMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageMemberMoonbreezesRequest"></a>
## APIpageMemberMoonbreezesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageMemberMoonbreezesReq** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | |



## pagePublicMoonbreezes

> PagePublicMoonbreezesResp pagePublicMoonbreezes(pagePublicMoonbreezesRequest)



服务方法。

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
            APIpagePublicMoonbreezesRequest request = APIpagePublicMoonbreezesRequest.newBuilder()
                .pagePublicMoonbreezesReq(pagePublicMoonbreezesReq)
                .build();
            PagePublicMoonbreezesResp result = apiInstance.pagePublicMoonbreezes(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pagePublicMoonbreezes");
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
| pagePublicMoonbreezesRequest | [**APIpagePublicMoonbreezesRequest**](MoonbreezeService.md#APIpagePublicMoonbreezesRequest)|-|-|

### Return type

[**PagePublicMoonbreezesResp**](PagePublicMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pagePublicMoonbreezesWithHttpInfo

> ApiResponse<PagePublicMoonbreezesResp> pagePublicMoonbreezesWithHttpInfo(pagePublicMoonbreezesRequest)



服务方法。

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
            APIpagePublicMoonbreezesRequest request = APIpagePublicMoonbreezesRequest.newBuilder()
                .pagePublicMoonbreezesReq(pagePublicMoonbreezesReq)
                .build();
            ApiResponse<PagePublicMoonbreezesResp> response = apiInstance.pagePublicMoonbreezesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pagePublicMoonbreezes");
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
| pagePublicMoonbreezesRequest | [**APIpagePublicMoonbreezesRequest**](MoonbreezeService.md#APIpagePublicMoonbreezesRequest)|-|-|

### Return type

ApiResponse<[**PagePublicMoonbreezesResp**](PagePublicMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpagePublicMoonbreezesRequest"></a>
## APIpagePublicMoonbreezesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pagePublicMoonbreezesReq** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | |



## pageWatchingMoonbreezes

> PageWatchingMoonbreezesResp pageWatchingMoonbreezes(pageWatchingMoonbreezesRequest)



服务方法。

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
            APIpageWatchingMoonbreezesRequest request = APIpageWatchingMoonbreezesRequest.newBuilder()
                .pageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq)
                .build();
            PageWatchingMoonbreezesResp result = apiInstance.pageWatchingMoonbreezes(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageWatchingMoonbreezes");
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
| pageWatchingMoonbreezesRequest | [**APIpageWatchingMoonbreezesRequest**](MoonbreezeService.md#APIpageWatchingMoonbreezesRequest)|-|-|

### Return type

[**PageWatchingMoonbreezesResp**](PageWatchingMoonbreezesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageWatchingMoonbreezesWithHttpInfo

> ApiResponse<PageWatchingMoonbreezesResp> pageWatchingMoonbreezesWithHttpInfo(pageWatchingMoonbreezesRequest)



服务方法。

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
            APIpageWatchingMoonbreezesRequest request = APIpageWatchingMoonbreezesRequest.newBuilder()
                .pageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq)
                .build();
            ApiResponse<PageWatchingMoonbreezesResp> response = apiInstance.pageWatchingMoonbreezesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling MoonbreezeService#pageWatchingMoonbreezes");
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
| pageWatchingMoonbreezesRequest | [**APIpageWatchingMoonbreezesRequest**](MoonbreezeService.md#APIpageWatchingMoonbreezesRequest)|-|-|

### Return type

ApiResponse<[**PageWatchingMoonbreezesResp**](PageWatchingMoonbreezesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageWatchingMoonbreezesRequest"></a>
## APIpageWatchingMoonbreezesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageWatchingMoonbreezesReq** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | |


