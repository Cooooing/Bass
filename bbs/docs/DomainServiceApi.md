# DomainServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**domainServiceList**](DomainServiceApi.md#domainServiceList) | **POST** /v1/content/domain/list |  |
| [**domainServiceListWithHttpInfo**](DomainServiceApi.md#domainServiceListWithHttpInfo) | **POST** /v1/content/domain/list |  |



## domainServiceList

> ListDomainsReply domainServiceList(domainServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainServiceApi;
import com.bass.bbs.api.DomainServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainServiceApi apiInstance = new DomainServiceApi(defaultClient);
        ListDomainsRequest listDomainsRequest = new ListDomainsRequest(); // ListDomainsRequest | 
        try {
            APIdomainServiceListRequest request = APIdomainServiceListRequest.newBuilder()
                .listDomainsRequest(listDomainsRequest)
                .build();
            ListDomainsReply result = apiInstance.domainServiceList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainServiceApi#domainServiceList");
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
| domainServiceListRequest | [**APIdomainServiceListRequest**](DomainServiceApi.md#APIdomainServiceListRequest)|-|-|

### Return type

[**ListDomainsReply**](ListDomainsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## domainServiceListWithHttpInfo

> ApiResponse<ListDomainsReply> domainServiceListWithHttpInfo(domainServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainServiceApi;
import com.bass.bbs.api.DomainServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainServiceApi apiInstance = new DomainServiceApi(defaultClient);
        ListDomainsRequest listDomainsRequest = new ListDomainsRequest(); // ListDomainsRequest | 
        try {
            APIdomainServiceListRequest request = APIdomainServiceListRequest.newBuilder()
                .listDomainsRequest(listDomainsRequest)
                .build();
            ApiResponse<ListDomainsReply> response = apiInstance.domainServiceListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainServiceApi#domainServiceList");
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
| domainServiceListRequest | [**APIdomainServiceListRequest**](DomainServiceApi.md#APIdomainServiceListRequest)|-|-|

### Return type

ApiResponse<[**ListDomainsReply**](ListDomainsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIdomainServiceListRequest"></a>
## APIdomainServiceListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listDomainsRequest** | [**ListDomainsRequest**](ListDomainsRequest.md) |  | |


