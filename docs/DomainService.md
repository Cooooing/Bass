# DomainService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](DomainService.md#callList) | **POST** /v1/content/domain/list |  |
| [**callListWithHttpInfo**](DomainService.md#callListWithHttpInfo) | **POST** /v1/content/domain/list |  |
| [**create**](DomainService.md#create) | **POST** /v1/content/domain/create |  |
| [**createWithHttpInfo**](DomainService.md#createWithHttpInfo) | **POST** /v1/content/domain/create |  |
| [**update**](DomainService.md#update) | **POST** /v1/content/domain/update |  |
| [**updateWithHttpInfo**](DomainService.md#updateWithHttpInfo) | **POST** /v1/content/domain/update |  |



## callList

> ListDomainsResp callList(callListRequest)



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
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listDomainsReq(listDomainsReq)
                .build();
            ListDomainsResp result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#callList");
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
| callListRequest | [**APIcallListRequest**](DomainService.md#APIcallListRequest)|-|-|

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

## callListWithHttpInfo

> ApiResponse<ListDomainsResp> callListWithHttpInfo(callListRequest)



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
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listDomainsReq(listDomainsReq)
                .build();
            ApiResponse<ListDomainsResp> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#callList");
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
| callListRequest | [**APIcallListRequest**](DomainService.md#APIcallListRequest)|-|-|

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


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listDomainsReq** | [**ListDomainsReq**](ListDomainsReq.md) |  | |



## create

> CreateDomainResp create(createRequest)



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
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createDomainReq(createDomainReq)
                .build();
            CreateDomainResp result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#create");
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
| createRequest | [**APIcreateRequest**](DomainService.md#APIcreateRequest)|-|-|

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

## createWithHttpInfo

> ApiResponse<CreateDomainResp> createWithHttpInfo(createRequest)



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
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createDomainReq(createDomainReq)
                .build();
            ApiResponse<CreateDomainResp> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#create");
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
| createRequest | [**APIcreateRequest**](DomainService.md#APIcreateRequest)|-|-|

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


<a id="APIcreateRequest"></a>
## APIcreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createDomainReq** | [**CreateDomainReq**](CreateDomainReq.md) |  | |



## update

> UpdateDomainResp update(updateRequest)



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
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateDomainReq(updateDomainReq)
                .build();
            UpdateDomainResp result = apiInstance.update(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#update");
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
| updateRequest | [**APIupdateRequest**](DomainService.md#APIupdateRequest)|-|-|

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

## updateWithHttpInfo

> ApiResponse<UpdateDomainResp> updateWithHttpInfo(updateRequest)



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
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateDomainReq(updateDomainReq)
                .build();
            ApiResponse<UpdateDomainResp> response = apiInstance.updateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#update");
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
| updateRequest | [**APIupdateRequest**](DomainService.md#APIupdateRequest)|-|-|

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


<a id="APIupdateRequest"></a>
## APIupdateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateDomainReq** | [**UpdateDomainReq**](UpdateDomainReq.md) |  | |


