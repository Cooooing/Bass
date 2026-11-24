# EmojiService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addEmojiItem**](EmojiService.md#addEmojiItem) | **POST** /v1/im/emoji/items |  |
| [**addEmojiItemWithHttpInfo**](EmojiService.md#addEmojiItemWithHttpInfo) | **POST** /v1/im/emoji/items |  |
| [**createEmojiGroup**](EmojiService.md#createEmojiGroup) | **POST** /v1/im/emoji/groups |  |
| [**createEmojiGroupWithHttpInfo**](EmojiService.md#createEmojiGroupWithHttpInfo) | **POST** /v1/im/emoji/groups |  |
| [**createEmojiShare**](EmojiService.md#createEmojiShare) | **POST** /v1/im/emoji/shares |  |
| [**createEmojiShareWithHttpInfo**](EmojiService.md#createEmojiShareWithHttpInfo) | **POST** /v1/im/emoji/shares |  |
| [**deleteEmojiGroup**](EmojiService.md#deleteEmojiGroup) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteEmojiGroupWithHttpInfo**](EmojiService.md#deleteEmojiGroupWithHttpInfo) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteEmojiItem**](EmojiService.md#deleteEmojiItem) | **POST** /v1/im/emoji/items/delete |  |
| [**deleteEmojiItemWithHttpInfo**](EmojiService.md#deleteEmojiItemWithHttpInfo) | **POST** /v1/im/emoji/items/delete |  |
| [**importEmojiShare**](EmojiService.md#importEmojiShare) | **POST** /v1/im/emoji/shares/import |  |
| [**importEmojiShareWithHttpInfo**](EmojiService.md#importEmojiShareWithHttpInfo) | **POST** /v1/im/emoji/shares/import |  |
| [**listEmojiGroups**](EmojiService.md#listEmojiGroups) | **POST** /v1/im/emoji/groups/list |  |
| [**listEmojiGroupsWithHttpInfo**](EmojiService.md#listEmojiGroupsWithHttpInfo) | **POST** /v1/im/emoji/groups/list |  |
| [**moveEmojiItem**](EmojiService.md#moveEmojiItem) | **POST** /v1/im/emoji/items/move |  |
| [**moveEmojiItemWithHttpInfo**](EmojiService.md#moveEmojiItemWithHttpInfo) | **POST** /v1/im/emoji/items/move |  |
| [**renameEmojiGroup**](EmojiService.md#renameEmojiGroup) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameEmojiGroupWithHttpInfo**](EmojiService.md#renameEmojiGroupWithHttpInfo) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameEmojiItem**](EmojiService.md#renameEmojiItem) | **POST** /v1/im/emoji/items/rename |  |
| [**renameEmojiItemWithHttpInfo**](EmojiService.md#renameEmojiItemWithHttpInfo) | **POST** /v1/im/emoji/items/rename |  |
| [**reorderEmojiGroups**](EmojiService.md#reorderEmojiGroups) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderEmojiGroupsWithHttpInfo**](EmojiService.md#reorderEmojiGroupsWithHttpInfo) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderEmojiItems**](EmojiService.md#reorderEmojiItems) | **POST** /v1/im/emoji/items/reorder |  |
| [**reorderEmojiItemsWithHttpInfo**](EmojiService.md#reorderEmojiItemsWithHttpInfo) | **POST** /v1/im/emoji/items/reorder |  |



## addEmojiItem

> AddEmojiItemResp addEmojiItem(addEmojiItemRequest)



向分组加入已上传图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        AddEmojiItemReq addEmojiItemReq = new AddEmojiItemReq(); // AddEmojiItemReq | 
        try {
            APIaddEmojiItemRequest request = APIaddEmojiItemRequest.newBuilder()
                .addEmojiItemReq(addEmojiItemReq)
                .build();
            AddEmojiItemResp result = apiInstance.addEmojiItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#addEmojiItem");
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
| addEmojiItemRequest | [**APIaddEmojiItemRequest**](EmojiService.md#APIaddEmojiItemRequest)|-|-|

### Return type

[**AddEmojiItemResp**](AddEmojiItemResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## addEmojiItemWithHttpInfo

> ApiResponse<AddEmojiItemResp> addEmojiItemWithHttpInfo(addEmojiItemRequest)



向分组加入已上传图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        AddEmojiItemReq addEmojiItemReq = new AddEmojiItemReq(); // AddEmojiItemReq | 
        try {
            APIaddEmojiItemRequest request = APIaddEmojiItemRequest.newBuilder()
                .addEmojiItemReq(addEmojiItemReq)
                .build();
            ApiResponse<AddEmojiItemResp> response = apiInstance.addEmojiItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#addEmojiItem");
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
| addEmojiItemRequest | [**APIaddEmojiItemRequest**](EmojiService.md#APIaddEmojiItemRequest)|-|-|

### Return type

ApiResponse<[**AddEmojiItemResp**](AddEmojiItemResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIaddEmojiItemRequest"></a>
## APIaddEmojiItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **addEmojiItemReq** | [**AddEmojiItemReq**](AddEmojiItemReq.md) |  | |



## createEmojiGroup

> CreateEmojiGroupResp createEmojiGroup(createEmojiGroupRequest)



创建一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        CreateEmojiGroupReq createEmojiGroupReq = new CreateEmojiGroupReq(); // CreateEmojiGroupReq | 
        try {
            APIcreateEmojiGroupRequest request = APIcreateEmojiGroupRequest.newBuilder()
                .createEmojiGroupReq(createEmojiGroupReq)
                .build();
            CreateEmojiGroupResp result = apiInstance.createEmojiGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createEmojiGroup");
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
| createEmojiGroupRequest | [**APIcreateEmojiGroupRequest**](EmojiService.md#APIcreateEmojiGroupRequest)|-|-|

### Return type

[**CreateEmojiGroupResp**](CreateEmojiGroupResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createEmojiGroupWithHttpInfo

> ApiResponse<CreateEmojiGroupResp> createEmojiGroupWithHttpInfo(createEmojiGroupRequest)



创建一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        CreateEmojiGroupReq createEmojiGroupReq = new CreateEmojiGroupReq(); // CreateEmojiGroupReq | 
        try {
            APIcreateEmojiGroupRequest request = APIcreateEmojiGroupRequest.newBuilder()
                .createEmojiGroupReq(createEmojiGroupReq)
                .build();
            ApiResponse<CreateEmojiGroupResp> response = apiInstance.createEmojiGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createEmojiGroup");
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
| createEmojiGroupRequest | [**APIcreateEmojiGroupRequest**](EmojiService.md#APIcreateEmojiGroupRequest)|-|-|

### Return type

ApiResponse<[**CreateEmojiGroupResp**](CreateEmojiGroupResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateEmojiGroupRequest"></a>
## APIcreateEmojiGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createEmojiGroupReq** | [**CreateEmojiGroupReq**](CreateEmojiGroupReq.md) |  | |



## createEmojiShare

> CreateEmojiShareResp createEmojiShare(createEmojiShareRequest)



生成七天有效的分组分享码。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        CreateEmojiShareReq createEmojiShareReq = new CreateEmojiShareReq(); // CreateEmojiShareReq | 
        try {
            APIcreateEmojiShareRequest request = APIcreateEmojiShareRequest.newBuilder()
                .createEmojiShareReq(createEmojiShareReq)
                .build();
            CreateEmojiShareResp result = apiInstance.createEmojiShare(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createEmojiShare");
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
| createEmojiShareRequest | [**APIcreateEmojiShareRequest**](EmojiService.md#APIcreateEmojiShareRequest)|-|-|

### Return type

[**CreateEmojiShareResp**](CreateEmojiShareResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createEmojiShareWithHttpInfo

> ApiResponse<CreateEmojiShareResp> createEmojiShareWithHttpInfo(createEmojiShareRequest)



生成七天有效的分组分享码。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        CreateEmojiShareReq createEmojiShareReq = new CreateEmojiShareReq(); // CreateEmojiShareReq | 
        try {
            APIcreateEmojiShareRequest request = APIcreateEmojiShareRequest.newBuilder()
                .createEmojiShareReq(createEmojiShareReq)
                .build();
            ApiResponse<CreateEmojiShareResp> response = apiInstance.createEmojiShareWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createEmojiShare");
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
| createEmojiShareRequest | [**APIcreateEmojiShareRequest**](EmojiService.md#APIcreateEmojiShareRequest)|-|-|

### Return type

ApiResponse<[**CreateEmojiShareResp**](CreateEmojiShareResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateEmojiShareRequest"></a>
## APIcreateEmojiShareRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createEmojiShareReq** | [**CreateEmojiShareReq**](CreateEmojiShareReq.md) |  | |



## deleteEmojiGroup

> Object deleteEmojiGroup(deleteEmojiGroupRequest)



删除一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        DeleteEmojiGroupReq deleteEmojiGroupReq = new DeleteEmojiGroupReq(); // DeleteEmojiGroupReq | 
        try {
            APIdeleteEmojiGroupRequest request = APIdeleteEmojiGroupRequest.newBuilder()
                .deleteEmojiGroupReq(deleteEmojiGroupReq)
                .build();
            Object result = apiInstance.deleteEmojiGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteEmojiGroup");
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
| deleteEmojiGroupRequest | [**APIdeleteEmojiGroupRequest**](EmojiService.md#APIdeleteEmojiGroupRequest)|-|-|

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

## deleteEmojiGroupWithHttpInfo

> ApiResponse<Object> deleteEmojiGroupWithHttpInfo(deleteEmojiGroupRequest)



删除一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        DeleteEmojiGroupReq deleteEmojiGroupReq = new DeleteEmojiGroupReq(); // DeleteEmojiGroupReq | 
        try {
            APIdeleteEmojiGroupRequest request = APIdeleteEmojiGroupRequest.newBuilder()
                .deleteEmojiGroupReq(deleteEmojiGroupReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteEmojiGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteEmojiGroup");
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
| deleteEmojiGroupRequest | [**APIdeleteEmojiGroupRequest**](EmojiService.md#APIdeleteEmojiGroupRequest)|-|-|

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


<a id="APIdeleteEmojiGroupRequest"></a>
## APIdeleteEmojiGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteEmojiGroupReq** | [**DeleteEmojiGroupReq**](DeleteEmojiGroupReq.md) |  | |



## deleteEmojiItem

> Object deleteEmojiItem(deleteEmojiItemRequest)



从分组移除表情。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        DeleteEmojiItemReq deleteEmojiItemReq = new DeleteEmojiItemReq(); // DeleteEmojiItemReq | 
        try {
            APIdeleteEmojiItemRequest request = APIdeleteEmojiItemRequest.newBuilder()
                .deleteEmojiItemReq(deleteEmojiItemReq)
                .build();
            Object result = apiInstance.deleteEmojiItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteEmojiItem");
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
| deleteEmojiItemRequest | [**APIdeleteEmojiItemRequest**](EmojiService.md#APIdeleteEmojiItemRequest)|-|-|

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

## deleteEmojiItemWithHttpInfo

> ApiResponse<Object> deleteEmojiItemWithHttpInfo(deleteEmojiItemRequest)



从分组移除表情。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        DeleteEmojiItemReq deleteEmojiItemReq = new DeleteEmojiItemReq(); // DeleteEmojiItemReq | 
        try {
            APIdeleteEmojiItemRequest request = APIdeleteEmojiItemRequest.newBuilder()
                .deleteEmojiItemReq(deleteEmojiItemReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteEmojiItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteEmojiItem");
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
| deleteEmojiItemRequest | [**APIdeleteEmojiItemRequest**](EmojiService.md#APIdeleteEmojiItemRequest)|-|-|

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


<a id="APIdeleteEmojiItemRequest"></a>
## APIdeleteEmojiItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteEmojiItemReq** | [**DeleteEmojiItemReq**](DeleteEmojiItemReq.md) |  | |



## importEmojiShare

> ImportEmojiShareResp importEmojiShare(importEmojiShareRequest)



把分享码导入为一个新分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ImportEmojiShareReq importEmojiShareReq = new ImportEmojiShareReq(); // ImportEmojiShareReq | 
        try {
            APIimportEmojiShareRequest request = APIimportEmojiShareRequest.newBuilder()
                .importEmojiShareReq(importEmojiShareReq)
                .build();
            ImportEmojiShareResp result = apiInstance.importEmojiShare(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#importEmojiShare");
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
| importEmojiShareRequest | [**APIimportEmojiShareRequest**](EmojiService.md#APIimportEmojiShareRequest)|-|-|

### Return type

[**ImportEmojiShareResp**](ImportEmojiShareResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## importEmojiShareWithHttpInfo

> ApiResponse<ImportEmojiShareResp> importEmojiShareWithHttpInfo(importEmojiShareRequest)



把分享码导入为一个新分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ImportEmojiShareReq importEmojiShareReq = new ImportEmojiShareReq(); // ImportEmojiShareReq | 
        try {
            APIimportEmojiShareRequest request = APIimportEmojiShareRequest.newBuilder()
                .importEmojiShareReq(importEmojiShareReq)
                .build();
            ApiResponse<ImportEmojiShareResp> response = apiInstance.importEmojiShareWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#importEmojiShare");
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
| importEmojiShareRequest | [**APIimportEmojiShareRequest**](EmojiService.md#APIimportEmojiShareRequest)|-|-|

### Return type

ApiResponse<[**ImportEmojiShareResp**](ImportEmojiShareResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIimportEmojiShareRequest"></a>
## APIimportEmojiShareRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **importEmojiShareReq** | [**ImportEmojiShareReq**](ImportEmojiShareReq.md) |  | |



## listEmojiGroups

> ListEmojiGroupsResp listEmojiGroups(listEmojiGroupsRequest)



查询当前账号的表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ListEmojiGroupsReq listEmojiGroupsReq = new ListEmojiGroupsReq(); // ListEmojiGroupsReq | 
        try {
            APIlistEmojiGroupsRequest request = APIlistEmojiGroupsRequest.newBuilder()
                .listEmojiGroupsReq(listEmojiGroupsReq)
                .build();
            ListEmojiGroupsResp result = apiInstance.listEmojiGroups(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#listEmojiGroups");
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
| listEmojiGroupsRequest | [**APIlistEmojiGroupsRequest**](EmojiService.md#APIlistEmojiGroupsRequest)|-|-|

### Return type

[**ListEmojiGroupsResp**](ListEmojiGroupsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listEmojiGroupsWithHttpInfo

> ApiResponse<ListEmojiGroupsResp> listEmojiGroupsWithHttpInfo(listEmojiGroupsRequest)



查询当前账号的表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ListEmojiGroupsReq listEmojiGroupsReq = new ListEmojiGroupsReq(); // ListEmojiGroupsReq | 
        try {
            APIlistEmojiGroupsRequest request = APIlistEmojiGroupsRequest.newBuilder()
                .listEmojiGroupsReq(listEmojiGroupsReq)
                .build();
            ApiResponse<ListEmojiGroupsResp> response = apiInstance.listEmojiGroupsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#listEmojiGroups");
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
| listEmojiGroupsRequest | [**APIlistEmojiGroupsRequest**](EmojiService.md#APIlistEmojiGroupsRequest)|-|-|

### Return type

ApiResponse<[**ListEmojiGroupsResp**](ListEmojiGroupsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistEmojiGroupsRequest"></a>
## APIlistEmojiGroupsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listEmojiGroupsReq** | [**ListEmojiGroupsReq**](ListEmojiGroupsReq.md) |  | |



## moveEmojiItem

> MoveEmojiItemResp moveEmojiItem(moveEmojiItemRequest)



将表情移动到另一个当前账号分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        MoveEmojiItemReq moveEmojiItemReq = new MoveEmojiItemReq(); // MoveEmojiItemReq | 
        try {
            APImoveEmojiItemRequest request = APImoveEmojiItemRequest.newBuilder()
                .moveEmojiItemReq(moveEmojiItemReq)
                .build();
            MoveEmojiItemResp result = apiInstance.moveEmojiItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#moveEmojiItem");
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
| moveEmojiItemRequest | [**APImoveEmojiItemRequest**](EmojiService.md#APImoveEmojiItemRequest)|-|-|

### Return type

[**MoveEmojiItemResp**](MoveEmojiItemResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## moveEmojiItemWithHttpInfo

> ApiResponse<MoveEmojiItemResp> moveEmojiItemWithHttpInfo(moveEmojiItemRequest)



将表情移动到另一个当前账号分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        MoveEmojiItemReq moveEmojiItemReq = new MoveEmojiItemReq(); // MoveEmojiItemReq | 
        try {
            APImoveEmojiItemRequest request = APImoveEmojiItemRequest.newBuilder()
                .moveEmojiItemReq(moveEmojiItemReq)
                .build();
            ApiResponse<MoveEmojiItemResp> response = apiInstance.moveEmojiItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#moveEmojiItem");
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
| moveEmojiItemRequest | [**APImoveEmojiItemRequest**](EmojiService.md#APImoveEmojiItemRequest)|-|-|

### Return type

ApiResponse<[**MoveEmojiItemResp**](MoveEmojiItemResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APImoveEmojiItemRequest"></a>
## APImoveEmojiItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **moveEmojiItemReq** | [**MoveEmojiItemReq**](MoveEmojiItemReq.md) |  | |



## renameEmojiGroup

> RenameEmojiGroupResp renameEmojiGroup(renameEmojiGroupRequest)



重命名一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        RenameEmojiGroupReq renameEmojiGroupReq = new RenameEmojiGroupReq(); // RenameEmojiGroupReq | 
        try {
            APIrenameEmojiGroupRequest request = APIrenameEmojiGroupRequest.newBuilder()
                .renameEmojiGroupReq(renameEmojiGroupReq)
                .build();
            RenameEmojiGroupResp result = apiInstance.renameEmojiGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameEmojiGroup");
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
| renameEmojiGroupRequest | [**APIrenameEmojiGroupRequest**](EmojiService.md#APIrenameEmojiGroupRequest)|-|-|

### Return type

[**RenameEmojiGroupResp**](RenameEmojiGroupResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renameEmojiGroupWithHttpInfo

> ApiResponse<RenameEmojiGroupResp> renameEmojiGroupWithHttpInfo(renameEmojiGroupRequest)



重命名一个表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        RenameEmojiGroupReq renameEmojiGroupReq = new RenameEmojiGroupReq(); // RenameEmojiGroupReq | 
        try {
            APIrenameEmojiGroupRequest request = APIrenameEmojiGroupRequest.newBuilder()
                .renameEmojiGroupReq(renameEmojiGroupReq)
                .build();
            ApiResponse<RenameEmojiGroupResp> response = apiInstance.renameEmojiGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameEmojiGroup");
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
| renameEmojiGroupRequest | [**APIrenameEmojiGroupRequest**](EmojiService.md#APIrenameEmojiGroupRequest)|-|-|

### Return type

ApiResponse<[**RenameEmojiGroupResp**](RenameEmojiGroupResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenameEmojiGroupRequest"></a>
## APIrenameEmojiGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renameEmojiGroupReq** | [**RenameEmojiGroupReq**](RenameEmojiGroupReq.md) |  | |



## renameEmojiItem

> RenameEmojiItemResp renameEmojiItem(renameEmojiItemRequest)



重命名一个表情别名。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        RenameEmojiItemReq renameEmojiItemReq = new RenameEmojiItemReq(); // RenameEmojiItemReq | 
        try {
            APIrenameEmojiItemRequest request = APIrenameEmojiItemRequest.newBuilder()
                .renameEmojiItemReq(renameEmojiItemReq)
                .build();
            RenameEmojiItemResp result = apiInstance.renameEmojiItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameEmojiItem");
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
| renameEmojiItemRequest | [**APIrenameEmojiItemRequest**](EmojiService.md#APIrenameEmojiItemRequest)|-|-|

### Return type

[**RenameEmojiItemResp**](RenameEmojiItemResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renameEmojiItemWithHttpInfo

> ApiResponse<RenameEmojiItemResp> renameEmojiItemWithHttpInfo(renameEmojiItemRequest)



重命名一个表情别名。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        RenameEmojiItemReq renameEmojiItemReq = new RenameEmojiItemReq(); // RenameEmojiItemReq | 
        try {
            APIrenameEmojiItemRequest request = APIrenameEmojiItemRequest.newBuilder()
                .renameEmojiItemReq(renameEmojiItemReq)
                .build();
            ApiResponse<RenameEmojiItemResp> response = apiInstance.renameEmojiItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameEmojiItem");
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
| renameEmojiItemRequest | [**APIrenameEmojiItemRequest**](EmojiService.md#APIrenameEmojiItemRequest)|-|-|

### Return type

ApiResponse<[**RenameEmojiItemResp**](RenameEmojiItemResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenameEmojiItemRequest"></a>
## APIrenameEmojiItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renameEmojiItemReq** | [**RenameEmojiItemReq**](RenameEmojiItemReq.md) |  | |



## reorderEmojiGroups

> Object reorderEmojiGroups(reorderEmojiGroupsRequest)



重排表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ReorderEmojiGroupsReq reorderEmojiGroupsReq = new ReorderEmojiGroupsReq(); // ReorderEmojiGroupsReq | 
        try {
            APIreorderEmojiGroupsRequest request = APIreorderEmojiGroupsRequest.newBuilder()
                .reorderEmojiGroupsReq(reorderEmojiGroupsReq)
                .build();
            Object result = apiInstance.reorderEmojiGroups(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderEmojiGroups");
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
| reorderEmojiGroupsRequest | [**APIreorderEmojiGroupsRequest**](EmojiService.md#APIreorderEmojiGroupsRequest)|-|-|

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

## reorderEmojiGroupsWithHttpInfo

> ApiResponse<Object> reorderEmojiGroupsWithHttpInfo(reorderEmojiGroupsRequest)



重排表情分组。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ReorderEmojiGroupsReq reorderEmojiGroupsReq = new ReorderEmojiGroupsReq(); // ReorderEmojiGroupsReq | 
        try {
            APIreorderEmojiGroupsRequest request = APIreorderEmojiGroupsRequest.newBuilder()
                .reorderEmojiGroupsReq(reorderEmojiGroupsReq)
                .build();
            ApiResponse<Object> response = apiInstance.reorderEmojiGroupsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderEmojiGroups");
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
| reorderEmojiGroupsRequest | [**APIreorderEmojiGroupsRequest**](EmojiService.md#APIreorderEmojiGroupsRequest)|-|-|

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


<a id="APIreorderEmojiGroupsRequest"></a>
## APIreorderEmojiGroupsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **reorderEmojiGroupsReq** | [**ReorderEmojiGroupsReq**](ReorderEmojiGroupsReq.md) |  | |



## reorderEmojiItems

> Object reorderEmojiItems(reorderEmojiItemsRequest)



重排同一分组内的表情。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ReorderEmojiItemsReq reorderEmojiItemsReq = new ReorderEmojiItemsReq(); // ReorderEmojiItemsReq | 
        try {
            APIreorderEmojiItemsRequest request = APIreorderEmojiItemsRequest.newBuilder()
                .reorderEmojiItemsReq(reorderEmojiItemsReq)
                .build();
            Object result = apiInstance.reorderEmojiItems(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderEmojiItems");
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
| reorderEmojiItemsRequest | [**APIreorderEmojiItemsRequest**](EmojiService.md#APIreorderEmojiItemsRequest)|-|-|

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

## reorderEmojiItemsWithHttpInfo

> ApiResponse<Object> reorderEmojiItemsWithHttpInfo(reorderEmojiItemsRequest)



重排同一分组内的表情。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.EmojiService;
import com.bass.bbs.api.EmojiService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        EmojiService apiInstance = new EmojiService(defaultClient);
        ReorderEmojiItemsReq reorderEmojiItemsReq = new ReorderEmojiItemsReq(); // ReorderEmojiItemsReq | 
        try {
            APIreorderEmojiItemsRequest request = APIreorderEmojiItemsRequest.newBuilder()
                .reorderEmojiItemsReq(reorderEmojiItemsReq)
                .build();
            ApiResponse<Object> response = apiInstance.reorderEmojiItemsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderEmojiItems");
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
| reorderEmojiItemsRequest | [**APIreorderEmojiItemsRequest**](EmojiService.md#APIreorderEmojiItemsRequest)|-|-|

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


<a id="APIreorderEmojiItemsRequest"></a>
## APIreorderEmojiItemsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **reorderEmojiItemsReq** | [**ReorderEmojiItemsReq**](ReorderEmojiItemsReq.md) |  | |


