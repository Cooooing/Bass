# DomainService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createDomain**](DomainService.md#createDomain) | **POST** /v1/content/domain/create |  |
| [**createDomainWithHttpInfo**](DomainService.md#createDomainWithHttpInfo) | **POST** /v1/content/domain/create |  |
| [**listDomains**](DomainService.md#listDomains) | **POST** /v1/content/domain/list |  |
| [**listDomainsWithHttpInfo**](DomainService.md#listDomainsWithHttpInfo) | **POST** /v1/content/domain/list |  |
| [**updateDomain**](DomainService.md#updateDomain) | **POST** /v1/content/domain/update |  |
| [**updateDomainWithHttpInfo**](DomainService.md#updateDomainWithHttpInfo) | **POST** /v1/content/domain/update |  |



## createDomain

> CreateDomainResp createDomain(createDomainRequest)



创建领域。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        CreateDomainReq createDomainReq = new CreateDomainReq(); // CreateDomainReq | 
        try {
            APIcreateDomainRequest request = APIcreateDomainRequest.newBuilder()
                .createDomainReq(createDomainReq)
                .build();
            CreateDomainResp result = apiInstance.createDomain(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#createDomain");
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
| createDomainRequest | [**APIcreateDomainRequest**](DomainService.md#APIcreateDomainRequest)|-|-|

### Return type

[**CreateDomainResp**](CreateDomainResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createDomainWithHttpInfo

> ApiResponse<CreateDomainResp> createDomainWithHttpInfo(createDomainRequest)



创建领域。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        CreateDomainReq createDomainReq = new CreateDomainReq(); // CreateDomainReq | 
        try {
            APIcreateDomainRequest request = APIcreateDomainRequest.newBuilder()
                .createDomainReq(createDomainReq)
                .build();
            ApiResponse<CreateDomainResp> response = apiInstance.createDomainWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#createDomain");
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
| createDomainRequest | [**APIcreateDomainRequest**](DomainService.md#APIcreateDomainRequest)|-|-|

### Return type

ApiResponse<[**CreateDomainResp**](CreateDomainResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateDomainRequest"></a>
## APIcreateDomainRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createDomainReq** | [**CreateDomainReq**](CreateDomainReq.md) |  | |



## listDomains

> ListDomainsResp listDomains(listDomainsRequest)



查询领域列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        ListDomainsReq listDomainsReq = new ListDomainsReq(); // ListDomainsReq | 
        try {
            APIlistDomainsRequest request = APIlistDomainsRequest.newBuilder()
                .listDomainsReq(listDomainsReq)
                .build();
            ListDomainsResp result = apiInstance.listDomains(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#listDomains");
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
| listDomainsRequest | [**APIlistDomainsRequest**](DomainService.md#APIlistDomainsRequest)|-|-|

### Return type

[**ListDomainsResp**](ListDomainsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listDomainsWithHttpInfo

> ApiResponse<ListDomainsResp> listDomainsWithHttpInfo(listDomainsRequest)



查询领域列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        ListDomainsReq listDomainsReq = new ListDomainsReq(); // ListDomainsReq | 
        try {
            APIlistDomainsRequest request = APIlistDomainsRequest.newBuilder()
                .listDomainsReq(listDomainsReq)
                .build();
            ApiResponse<ListDomainsResp> response = apiInstance.listDomainsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#listDomains");
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
| listDomainsRequest | [**APIlistDomainsRequest**](DomainService.md#APIlistDomainsRequest)|-|-|

### Return type

ApiResponse<[**ListDomainsResp**](ListDomainsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistDomainsRequest"></a>
## APIlistDomainsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listDomainsReq** | [**ListDomainsReq**](ListDomainsReq.md) |  | |



## updateDomain

> UpdateDomainResp updateDomain(updateDomainRequest)



更新领域。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        UpdateDomainReq updateDomainReq = new UpdateDomainReq(); // UpdateDomainReq | 
        try {
            APIupdateDomainRequest request = APIupdateDomainRequest.newBuilder()
                .updateDomainReq(updateDomainReq)
                .build();
            UpdateDomainResp result = apiInstance.updateDomain(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#updateDomain");
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
| updateDomainRequest | [**APIupdateDomainRequest**](DomainService.md#APIupdateDomainRequest)|-|-|

### Return type

[**UpdateDomainResp**](UpdateDomainResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updateDomainWithHttpInfo

> ApiResponse<UpdateDomainResp> updateDomainWithHttpInfo(updateDomainRequest)



更新领域。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        UpdateDomainReq updateDomainReq = new UpdateDomainReq(); // UpdateDomainReq | 
        try {
            APIupdateDomainRequest request = APIupdateDomainRequest.newBuilder()
                .updateDomainReq(updateDomainReq)
                .build();
            ApiResponse<UpdateDomainResp> response = apiInstance.updateDomainWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#updateDomain");
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
| updateDomainRequest | [**APIupdateDomainRequest**](DomainService.md#APIupdateDomainRequest)|-|-|

### Return type

ApiResponse<[**UpdateDomainResp**](UpdateDomainResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdateDomainRequest"></a>
## APIupdateDomainRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateDomainReq** | [**UpdateDomainReq**](UpdateDomainReq.md) |  | |


