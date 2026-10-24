# EmojiService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addItem**](EmojiService.md#addItem) | **POST** /v1/im/emoji/items |  |
| [**addItemWithHttpInfo**](EmojiService.md#addItemWithHttpInfo) | **POST** /v1/im/emoji/items |  |
| [**callList**](EmojiService.md#callList) | **POST** /v1/im/emoji/groups/list |  |
| [**callListWithHttpInfo**](EmojiService.md#callListWithHttpInfo) | **POST** /v1/im/emoji/groups/list |  |
| [**createGroup**](EmojiService.md#createGroup) | **POST** /v1/im/emoji/groups |  |
| [**createGroupWithHttpInfo**](EmojiService.md#createGroupWithHttpInfo) | **POST** /v1/im/emoji/groups |  |
| [**createShare**](EmojiService.md#createShare) | **POST** /v1/im/emoji/shares |  |
| [**createShareWithHttpInfo**](EmojiService.md#createShareWithHttpInfo) | **POST** /v1/im/emoji/shares |  |
| [**deleteGroup**](EmojiService.md#deleteGroup) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteGroupWithHttpInfo**](EmojiService.md#deleteGroupWithHttpInfo) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteItem**](EmojiService.md#deleteItem) | **POST** /v1/im/emoji/items/delete |  |
| [**deleteItemWithHttpInfo**](EmojiService.md#deleteItemWithHttpInfo) | **POST** /v1/im/emoji/items/delete |  |
| [**importShare**](EmojiService.md#importShare) | **POST** /v1/im/emoji/shares/import |  |
| [**importShareWithHttpInfo**](EmojiService.md#importShareWithHttpInfo) | **POST** /v1/im/emoji/shares/import |  |
| [**moveItem**](EmojiService.md#moveItem) | **POST** /v1/im/emoji/items/move |  |
| [**moveItemWithHttpInfo**](EmojiService.md#moveItemWithHttpInfo) | **POST** /v1/im/emoji/items/move |  |
| [**renameGroup**](EmojiService.md#renameGroup) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameGroupWithHttpInfo**](EmojiService.md#renameGroupWithHttpInfo) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameItem**](EmojiService.md#renameItem) | **POST** /v1/im/emoji/items/rename |  |
| [**renameItemWithHttpInfo**](EmojiService.md#renameItemWithHttpInfo) | **POST** /v1/im/emoji/items/rename |  |
| [**reorderGroups**](EmojiService.md#reorderGroups) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderGroupsWithHttpInfo**](EmojiService.md#reorderGroupsWithHttpInfo) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderItems**](EmojiService.md#reorderItems) | **POST** /v1/im/emoji/items/reorder |  |
| [**reorderItemsWithHttpInfo**](EmojiService.md#reorderItemsWithHttpInfo) | **POST** /v1/im/emoji/items/reorder |  |



## addItem

> AddEmojiItemResp addItem(addItemRequest)



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
            APIaddItemRequest request = APIaddItemRequest.newBuilder()
                .addEmojiItemReq(addEmojiItemReq)
                .build();
            AddEmojiItemResp result = apiInstance.addItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#addItem");
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
| addItemRequest | [**APIaddItemRequest**](EmojiService.md#APIaddItemRequest)|-|-|

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

## addItemWithHttpInfo

> ApiResponse<AddEmojiItemResp> addItemWithHttpInfo(addItemRequest)



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
            APIaddItemRequest request = APIaddItemRequest.newBuilder()
                .addEmojiItemReq(addEmojiItemReq)
                .build();
            ApiResponse<AddEmojiItemResp> response = apiInstance.addItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#addItem");
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
| addItemRequest | [**APIaddItemRequest**](EmojiService.md#APIaddItemRequest)|-|-|

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


<a id="APIaddItemRequest"></a>
## APIaddItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **addEmojiItemReq** | [**AddEmojiItemReq**](AddEmojiItemReq.md) |  | |



## callList

> ListEmojiGroupsResp callList(callListRequest)



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
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listEmojiGroupsReq(listEmojiGroupsReq)
                .build();
            ListEmojiGroupsResp result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#callList");
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
| callListRequest | [**APIcallListRequest**](EmojiService.md#APIcallListRequest)|-|-|

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

## callListWithHttpInfo

> ApiResponse<ListEmojiGroupsResp> callListWithHttpInfo(callListRequest)



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
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listEmojiGroupsReq(listEmojiGroupsReq)
                .build();
            ApiResponse<ListEmojiGroupsResp> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#callList");
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
| callListRequest | [**APIcallListRequest**](EmojiService.md#APIcallListRequest)|-|-|

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


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listEmojiGroupsReq** | [**ListEmojiGroupsReq**](ListEmojiGroupsReq.md) |  | |



## createGroup

> CreateEmojiGroupResp createGroup(createGroupRequest)



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
            APIcreateGroupRequest request = APIcreateGroupRequest.newBuilder()
                .createEmojiGroupReq(createEmojiGroupReq)
                .build();
            CreateEmojiGroupResp result = apiInstance.createGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createGroup");
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
| createGroupRequest | [**APIcreateGroupRequest**](EmojiService.md#APIcreateGroupRequest)|-|-|

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

## createGroupWithHttpInfo

> ApiResponse<CreateEmojiGroupResp> createGroupWithHttpInfo(createGroupRequest)



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
            APIcreateGroupRequest request = APIcreateGroupRequest.newBuilder()
                .createEmojiGroupReq(createEmojiGroupReq)
                .build();
            ApiResponse<CreateEmojiGroupResp> response = apiInstance.createGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createGroup");
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
| createGroupRequest | [**APIcreateGroupRequest**](EmojiService.md#APIcreateGroupRequest)|-|-|

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


<a id="APIcreateGroupRequest"></a>
## APIcreateGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createEmojiGroupReq** | [**CreateEmojiGroupReq**](CreateEmojiGroupReq.md) |  | |



## createShare

> CreateEmojiShareResp createShare(createShareRequest)



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
            APIcreateShareRequest request = APIcreateShareRequest.newBuilder()
                .createEmojiShareReq(createEmojiShareReq)
                .build();
            CreateEmojiShareResp result = apiInstance.createShare(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createShare");
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
| createShareRequest | [**APIcreateShareRequest**](EmojiService.md#APIcreateShareRequest)|-|-|

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

## createShareWithHttpInfo

> ApiResponse<CreateEmojiShareResp> createShareWithHttpInfo(createShareRequest)



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
            APIcreateShareRequest request = APIcreateShareRequest.newBuilder()
                .createEmojiShareReq(createEmojiShareReq)
                .build();
            ApiResponse<CreateEmojiShareResp> response = apiInstance.createShareWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#createShare");
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
| createShareRequest | [**APIcreateShareRequest**](EmojiService.md#APIcreateShareRequest)|-|-|

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


<a id="APIcreateShareRequest"></a>
## APIcreateShareRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createEmojiShareReq** | [**CreateEmojiShareReq**](CreateEmojiShareReq.md) |  | |



## deleteGroup

> Object deleteGroup(deleteGroupRequest)



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
            APIdeleteGroupRequest request = APIdeleteGroupRequest.newBuilder()
                .deleteEmojiGroupReq(deleteEmojiGroupReq)
                .build();
            Object result = apiInstance.deleteGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteGroup");
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
| deleteGroupRequest | [**APIdeleteGroupRequest**](EmojiService.md#APIdeleteGroupRequest)|-|-|

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

## deleteGroupWithHttpInfo

> ApiResponse<Object> deleteGroupWithHttpInfo(deleteGroupRequest)



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
            APIdeleteGroupRequest request = APIdeleteGroupRequest.newBuilder()
                .deleteEmojiGroupReq(deleteEmojiGroupReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteGroup");
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
| deleteGroupRequest | [**APIdeleteGroupRequest**](EmojiService.md#APIdeleteGroupRequest)|-|-|

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


<a id="APIdeleteGroupRequest"></a>
## APIdeleteGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteEmojiGroupReq** | [**DeleteEmojiGroupReq**](DeleteEmojiGroupReq.md) |  | |



## deleteItem

> Object deleteItem(deleteItemRequest)



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
            APIdeleteItemRequest request = APIdeleteItemRequest.newBuilder()
                .deleteEmojiItemReq(deleteEmojiItemReq)
                .build();
            Object result = apiInstance.deleteItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteItem");
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
| deleteItemRequest | [**APIdeleteItemRequest**](EmojiService.md#APIdeleteItemRequest)|-|-|

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

## deleteItemWithHttpInfo

> ApiResponse<Object> deleteItemWithHttpInfo(deleteItemRequest)



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
            APIdeleteItemRequest request = APIdeleteItemRequest.newBuilder()
                .deleteEmojiItemReq(deleteEmojiItemReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#deleteItem");
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
| deleteItemRequest | [**APIdeleteItemRequest**](EmojiService.md#APIdeleteItemRequest)|-|-|

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


<a id="APIdeleteItemRequest"></a>
## APIdeleteItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteEmojiItemReq** | [**DeleteEmojiItemReq**](DeleteEmojiItemReq.md) |  | |



## importShare

> ImportEmojiShareResp importShare(importShareRequest)



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
            APIimportShareRequest request = APIimportShareRequest.newBuilder()
                .importEmojiShareReq(importEmojiShareReq)
                .build();
            ImportEmojiShareResp result = apiInstance.importShare(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#importShare");
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
| importShareRequest | [**APIimportShareRequest**](EmojiService.md#APIimportShareRequest)|-|-|

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

## importShareWithHttpInfo

> ApiResponse<ImportEmojiShareResp> importShareWithHttpInfo(importShareRequest)



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
            APIimportShareRequest request = APIimportShareRequest.newBuilder()
                .importEmojiShareReq(importEmojiShareReq)
                .build();
            ApiResponse<ImportEmojiShareResp> response = apiInstance.importShareWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#importShare");
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
| importShareRequest | [**APIimportShareRequest**](EmojiService.md#APIimportShareRequest)|-|-|

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


<a id="APIimportShareRequest"></a>
## APIimportShareRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **importEmojiShareReq** | [**ImportEmojiShareReq**](ImportEmojiShareReq.md) |  | |



## moveItem

> MoveEmojiItemResp moveItem(moveItemRequest)



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
            APImoveItemRequest request = APImoveItemRequest.newBuilder()
                .moveEmojiItemReq(moveEmojiItemReq)
                .build();
            MoveEmojiItemResp result = apiInstance.moveItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#moveItem");
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
| moveItemRequest | [**APImoveItemRequest**](EmojiService.md#APImoveItemRequest)|-|-|

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

## moveItemWithHttpInfo

> ApiResponse<MoveEmojiItemResp> moveItemWithHttpInfo(moveItemRequest)



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
            APImoveItemRequest request = APImoveItemRequest.newBuilder()
                .moveEmojiItemReq(moveEmojiItemReq)
                .build();
            ApiResponse<MoveEmojiItemResp> response = apiInstance.moveItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#moveItem");
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
| moveItemRequest | [**APImoveItemRequest**](EmojiService.md#APImoveItemRequest)|-|-|

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


<a id="APImoveItemRequest"></a>
## APImoveItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **moveEmojiItemReq** | [**MoveEmojiItemReq**](MoveEmojiItemReq.md) |  | |



## renameGroup

> RenameEmojiGroupResp renameGroup(renameGroupRequest)



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
            APIrenameGroupRequest request = APIrenameGroupRequest.newBuilder()
                .renameEmojiGroupReq(renameEmojiGroupReq)
                .build();
            RenameEmojiGroupResp result = apiInstance.renameGroup(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameGroup");
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
| renameGroupRequest | [**APIrenameGroupRequest**](EmojiService.md#APIrenameGroupRequest)|-|-|

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

## renameGroupWithHttpInfo

> ApiResponse<RenameEmojiGroupResp> renameGroupWithHttpInfo(renameGroupRequest)



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
            APIrenameGroupRequest request = APIrenameGroupRequest.newBuilder()
                .renameEmojiGroupReq(renameEmojiGroupReq)
                .build();
            ApiResponse<RenameEmojiGroupResp> response = apiInstance.renameGroupWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameGroup");
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
| renameGroupRequest | [**APIrenameGroupRequest**](EmojiService.md#APIrenameGroupRequest)|-|-|

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


<a id="APIrenameGroupRequest"></a>
## APIrenameGroupRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renameEmojiGroupReq** | [**RenameEmojiGroupReq**](RenameEmojiGroupReq.md) |  | |



## renameItem

> RenameEmojiItemResp renameItem(renameItemRequest)



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
            APIrenameItemRequest request = APIrenameItemRequest.newBuilder()
                .renameEmojiItemReq(renameEmojiItemReq)
                .build();
            RenameEmojiItemResp result = apiInstance.renameItem(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameItem");
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
| renameItemRequest | [**APIrenameItemRequest**](EmojiService.md#APIrenameItemRequest)|-|-|

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

## renameItemWithHttpInfo

> ApiResponse<RenameEmojiItemResp> renameItemWithHttpInfo(renameItemRequest)



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
            APIrenameItemRequest request = APIrenameItemRequest.newBuilder()
                .renameEmojiItemReq(renameEmojiItemReq)
                .build();
            ApiResponse<RenameEmojiItemResp> response = apiInstance.renameItemWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#renameItem");
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
| renameItemRequest | [**APIrenameItemRequest**](EmojiService.md#APIrenameItemRequest)|-|-|

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


<a id="APIrenameItemRequest"></a>
## APIrenameItemRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renameEmojiItemReq** | [**RenameEmojiItemReq**](RenameEmojiItemReq.md) |  | |



## reorderGroups

> Object reorderGroups(reorderGroupsRequest)



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
            APIreorderGroupsRequest request = APIreorderGroupsRequest.newBuilder()
                .reorderEmojiGroupsReq(reorderEmojiGroupsReq)
                .build();
            Object result = apiInstance.reorderGroups(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderGroups");
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
| reorderGroupsRequest | [**APIreorderGroupsRequest**](EmojiService.md#APIreorderGroupsRequest)|-|-|

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

## reorderGroupsWithHttpInfo

> ApiResponse<Object> reorderGroupsWithHttpInfo(reorderGroupsRequest)



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
            APIreorderGroupsRequest request = APIreorderGroupsRequest.newBuilder()
                .reorderEmojiGroupsReq(reorderEmojiGroupsReq)
                .build();
            ApiResponse<Object> response = apiInstance.reorderGroupsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderGroups");
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
| reorderGroupsRequest | [**APIreorderGroupsRequest**](EmojiService.md#APIreorderGroupsRequest)|-|-|

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


<a id="APIreorderGroupsRequest"></a>
## APIreorderGroupsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **reorderEmojiGroupsReq** | [**ReorderEmojiGroupsReq**](ReorderEmojiGroupsReq.md) |  | |



## reorderItems

> Object reorderItems(reorderItemsRequest)



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
            APIreorderItemsRequest request = APIreorderItemsRequest.newBuilder()
                .reorderEmojiItemsReq(reorderEmojiItemsReq)
                .build();
            Object result = apiInstance.reorderItems(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderItems");
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
| reorderItemsRequest | [**APIreorderItemsRequest**](EmojiService.md#APIreorderItemsRequest)|-|-|

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

## reorderItemsWithHttpInfo

> ApiResponse<Object> reorderItemsWithHttpInfo(reorderItemsRequest)



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
            APIreorderItemsRequest request = APIreorderItemsRequest.newBuilder()
                .reorderEmojiItemsReq(reorderEmojiItemsReq)
                .build();
            ApiResponse<Object> response = apiInstance.reorderItemsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling EmojiService#reorderItems");
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
| reorderItemsRequest | [**APIreorderItemsRequest**](EmojiService.md#APIreorderItemsRequest)|-|-|

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


<a id="APIreorderItemsRequest"></a>
## APIreorderItemsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **reorderEmojiItemsReq** | [**ReorderEmojiItemsReq**](ReorderEmojiItemsReq.md) |  | |


