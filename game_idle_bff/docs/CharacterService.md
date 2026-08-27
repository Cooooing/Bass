# CharacterService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](CharacterService.md#callList) | **POST** /v1/game-idle/character/list |  |
| [**callListWithHttpInfo**](CharacterService.md#callListWithHttpInfo) | **POST** /v1/game-idle/character/list |  |
| [**create**](CharacterService.md#create) | **POST** /v1/game-idle/character/create |  |
| [**createWithHttpInfo**](CharacterService.md#createWithHttpInfo) | **POST** /v1/game-idle/character/create |  |



## callList

> ListCharacterResp callList(callListRequest)



查询当前账号角色列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CharacterService;
import com.bass.bbs.api.CharacterService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CharacterService apiInstance = new CharacterService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .body(body)
                .build();
            ListCharacterResp result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#callList");
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
| callListRequest | [**APIcallListRequest**](CharacterService.md#APIcallListRequest)|-|-|

### Return type

[**ListCharacterResp**](ListCharacterResp.md)


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

> ApiResponse<ListCharacterResp> callListWithHttpInfo(callListRequest)



查询当前账号角色列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CharacterService;
import com.bass.bbs.api.CharacterService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CharacterService apiInstance = new CharacterService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ListCharacterResp> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#callList");
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
| callListRequest | [**APIcallListRequest**](CharacterService.md#APIcallListRequest)|-|-|

### Return type

ApiResponse<[**ListCharacterResp**](ListCharacterResp.md)>


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
| **body** | **Object** |  | |



## create

> CreateCharacterResp create(createRequest)



创建角色。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CharacterService;
import com.bass.bbs.api.CharacterService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CharacterService apiInstance = new CharacterService(defaultClient);
        CreateCharacterReq createCharacterReq = new CreateCharacterReq(); // CreateCharacterReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCharacterReq(createCharacterReq)
                .build();
            CreateCharacterResp result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#create");
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
| createRequest | [**APIcreateRequest**](CharacterService.md#APIcreateRequest)|-|-|

### Return type

[**CreateCharacterResp**](CreateCharacterResp.md)


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

> ApiResponse<CreateCharacterResp> createWithHttpInfo(createRequest)



创建角色。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CharacterService;
import com.bass.bbs.api.CharacterService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CharacterService apiInstance = new CharacterService(defaultClient);
        CreateCharacterReq createCharacterReq = new CreateCharacterReq(); // CreateCharacterReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCharacterReq(createCharacterReq)
                .build();
            ApiResponse<CreateCharacterResp> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#create");
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
| createRequest | [**APIcreateRequest**](CharacterService.md#APIcreateRequest)|-|-|

### Return type

ApiResponse<[**CreateCharacterResp**](CreateCharacterResp.md)>


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
| **createCharacterReq** | [**CreateCharacterReq**](CreateCharacterReq.md) |  | |


