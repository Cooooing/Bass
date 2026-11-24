# CharacterService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createCharacter**](CharacterService.md#createCharacter) | **POST** /v1/game-idle/character/create |  |
| [**createCharacterWithHttpInfo**](CharacterService.md#createCharacterWithHttpInfo) | **POST** /v1/game-idle/character/create |  |
| [**listCharacter**](CharacterService.md#listCharacter) | **POST** /v1/game-idle/character/list |  |
| [**listCharacterWithHttpInfo**](CharacterService.md#listCharacterWithHttpInfo) | **POST** /v1/game-idle/character/list |  |



## createCharacter

> CreateCharacterResp createCharacter(createCharacterRequest)



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
            APIcreateCharacterRequest request = APIcreateCharacterRequest.newBuilder()
                .createCharacterReq(createCharacterReq)
                .build();
            CreateCharacterResp result = apiInstance.createCharacter(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#createCharacter");
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
| createCharacterRequest | [**APIcreateCharacterRequest**](CharacterService.md#APIcreateCharacterRequest)|-|-|

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

## createCharacterWithHttpInfo

> ApiResponse<CreateCharacterResp> createCharacterWithHttpInfo(createCharacterRequest)



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
            APIcreateCharacterRequest request = APIcreateCharacterRequest.newBuilder()
                .createCharacterReq(createCharacterReq)
                .build();
            ApiResponse<CreateCharacterResp> response = apiInstance.createCharacterWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#createCharacter");
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
| createCharacterRequest | [**APIcreateCharacterRequest**](CharacterService.md#APIcreateCharacterRequest)|-|-|

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


<a id="APIcreateCharacterRequest"></a>
## APIcreateCharacterRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createCharacterReq** | [**CreateCharacterReq**](CreateCharacterReq.md) |  | |



## listCharacter

> ListCharacterResp listCharacter(listCharacterRequest)



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
            APIlistCharacterRequest request = APIlistCharacterRequest.newBuilder()
                .body(body)
                .build();
            ListCharacterResp result = apiInstance.listCharacter(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#listCharacter");
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
| listCharacterRequest | [**APIlistCharacterRequest**](CharacterService.md#APIlistCharacterRequest)|-|-|

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

## listCharacterWithHttpInfo

> ApiResponse<ListCharacterResp> listCharacterWithHttpInfo(listCharacterRequest)



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
            APIlistCharacterRequest request = APIlistCharacterRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ListCharacterResp> response = apiInstance.listCharacterWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CharacterService#listCharacter");
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
| listCharacterRequest | [**APIlistCharacterRequest**](CharacterService.md#APIlistCharacterRequest)|-|-|

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


<a id="APIlistCharacterRequest"></a>
## APIlistCharacterRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |


