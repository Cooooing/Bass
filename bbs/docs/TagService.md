# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**bindArticle**](TagService.md#bindArticle) | **POST** /v1/content/tag/bind-article |  |
| [**bindArticleWithHttpInfo**](TagService.md#bindArticleWithHttpInfo) | **POST** /v1/content/tag/bind-article |  |
| [**callList**](TagService.md#callList) | **POST** /v1/content/tag/list |  |
| [**callListWithHttpInfo**](TagService.md#callListWithHttpInfo) | **POST** /v1/content/tag/list |  |
| [**create**](TagService.md#create) | **POST** /v1/content/tag/create |  |
| [**createWithHttpInfo**](TagService.md#createWithHttpInfo) | **POST** /v1/content/tag/create |  |
| [**listArticleTags**](TagService.md#listArticleTags) | **POST** /v1/content/tag/list-article-tags |  |
| [**listArticleTagsWithHttpInfo**](TagService.md#listArticleTagsWithHttpInfo) | **POST** /v1/content/tag/list-article-tags |  |
| [**unbindArticle**](TagService.md#unbindArticle) | **POST** /v1/content/tag/unbind-article |  |
| [**unbindArticleWithHttpInfo**](TagService.md#unbindArticleWithHttpInfo) | **POST** /v1/content/tag/unbind-article |  |
| [**update**](TagService.md#update) | **POST** /v1/content/tag/update |  |
| [**updateWithHttpInfo**](TagService.md#updateWithHttpInfo) | **POST** /v1/content/tag/update |  |



## bindArticle

> Object bindArticle(bindArticleRequest)



绑定文章标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        BindArticleTagsReq bindArticleTagsReq = new BindArticleTagsReq(); // BindArticleTagsReq | 
        try {
            APIbindArticleRequest request = APIbindArticleRequest.newBuilder()
                .bindArticleTagsReq(bindArticleTagsReq)
                .build();
            Object result = apiInstance.bindArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#bindArticle");
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
| bindArticleRequest | [**APIbindArticleRequest**](TagService.md#APIbindArticleRequest)|-|-|

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

## bindArticleWithHttpInfo

> ApiResponse<Object> bindArticleWithHttpInfo(bindArticleRequest)



绑定文章标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        BindArticleTagsReq bindArticleTagsReq = new BindArticleTagsReq(); // BindArticleTagsReq | 
        try {
            APIbindArticleRequest request = APIbindArticleRequest.newBuilder()
                .bindArticleTagsReq(bindArticleTagsReq)
                .build();
            ApiResponse<Object> response = apiInstance.bindArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#bindArticle");
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
| bindArticleRequest | [**APIbindArticleRequest**](TagService.md#APIbindArticleRequest)|-|-|

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


<a id="APIbindArticleRequest"></a>
## APIbindArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **bindArticleTagsReq** | [**BindArticleTagsReq**](BindArticleTagsReq.md) |  | |



## callList

> ListTagsResp callList(callListRequest)



查询标签列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListTagsReq listTagsReq = new ListTagsReq(); // ListTagsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listTagsReq(listTagsReq)
                .build();
            ListTagsResp result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#callList");
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
| callListRequest | [**APIcallListRequest**](TagService.md#APIcallListRequest)|-|-|

### Return type

[**ListTagsResp**](ListTagsResp.md)


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

> ApiResponse<ListTagsResp> callListWithHttpInfo(callListRequest)



查询标签列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListTagsReq listTagsReq = new ListTagsReq(); // ListTagsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listTagsReq(listTagsReq)
                .build();
            ApiResponse<ListTagsResp> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#callList");
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
| callListRequest | [**APIcallListRequest**](TagService.md#APIcallListRequest)|-|-|

### Return type

ApiResponse<[**ListTagsResp**](ListTagsResp.md)>


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
| **listTagsReq** | [**ListTagsReq**](ListTagsReq.md) |  | |



## create

> CreateTagResp create(createRequest)



创建标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        CreateTagReq createTagReq = new CreateTagReq(); // CreateTagReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createTagReq(createTagReq)
                .build();
            CreateTagResp result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#create");
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
| createRequest | [**APIcreateRequest**](TagService.md#APIcreateRequest)|-|-|

### Return type

[**CreateTagResp**](CreateTagResp.md)


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

> ApiResponse<CreateTagResp> createWithHttpInfo(createRequest)



创建标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        CreateTagReq createTagReq = new CreateTagReq(); // CreateTagReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createTagReq(createTagReq)
                .build();
            ApiResponse<CreateTagResp> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#create");
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
| createRequest | [**APIcreateRequest**](TagService.md#APIcreateRequest)|-|-|

### Return type

ApiResponse<[**CreateTagResp**](CreateTagResp.md)>


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
| **createTagReq** | [**CreateTagReq**](CreateTagReq.md) |  | |



## listArticleTags

> ListArticleTagsResp listArticleTags(listArticleTagsRequest)



查询文章标签列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListArticleTagsReq listArticleTagsReq = new ListArticleTagsReq(); // ListArticleTagsReq | 
        try {
            APIlistArticleTagsRequest request = APIlistArticleTagsRequest.newBuilder()
                .listArticleTagsReq(listArticleTagsReq)
                .build();
            ListArticleTagsResp result = apiInstance.listArticleTags(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#listArticleTags");
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
| listArticleTagsRequest | [**APIlistArticleTagsRequest**](TagService.md#APIlistArticleTagsRequest)|-|-|

### Return type

[**ListArticleTagsResp**](ListArticleTagsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listArticleTagsWithHttpInfo

> ApiResponse<ListArticleTagsResp> listArticleTagsWithHttpInfo(listArticleTagsRequest)



查询文章标签列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListArticleTagsReq listArticleTagsReq = new ListArticleTagsReq(); // ListArticleTagsReq | 
        try {
            APIlistArticleTagsRequest request = APIlistArticleTagsRequest.newBuilder()
                .listArticleTagsReq(listArticleTagsReq)
                .build();
            ApiResponse<ListArticleTagsResp> response = apiInstance.listArticleTagsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#listArticleTags");
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
| listArticleTagsRequest | [**APIlistArticleTagsRequest**](TagService.md#APIlistArticleTagsRequest)|-|-|

### Return type

ApiResponse<[**ListArticleTagsResp**](ListArticleTagsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistArticleTagsRequest"></a>
## APIlistArticleTagsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listArticleTagsReq** | [**ListArticleTagsReq**](ListArticleTagsReq.md) |  | |



## unbindArticle

> Object unbindArticle(unbindArticleRequest)



解绑文章标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        UnbindArticleTagsReq unbindArticleTagsReq = new UnbindArticleTagsReq(); // UnbindArticleTagsReq | 
        try {
            APIunbindArticleRequest request = APIunbindArticleRequest.newBuilder()
                .unbindArticleTagsReq(unbindArticleTagsReq)
                .build();
            Object result = apiInstance.unbindArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#unbindArticle");
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
| unbindArticleRequest | [**APIunbindArticleRequest**](TagService.md#APIunbindArticleRequest)|-|-|

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

## unbindArticleWithHttpInfo

> ApiResponse<Object> unbindArticleWithHttpInfo(unbindArticleRequest)



解绑文章标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        UnbindArticleTagsReq unbindArticleTagsReq = new UnbindArticleTagsReq(); // UnbindArticleTagsReq | 
        try {
            APIunbindArticleRequest request = APIunbindArticleRequest.newBuilder()
                .unbindArticleTagsReq(unbindArticleTagsReq)
                .build();
            ApiResponse<Object> response = apiInstance.unbindArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#unbindArticle");
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
| unbindArticleRequest | [**APIunbindArticleRequest**](TagService.md#APIunbindArticleRequest)|-|-|

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


<a id="APIunbindArticleRequest"></a>
## APIunbindArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unbindArticleTagsReq** | [**UnbindArticleTagsReq**](UnbindArticleTagsReq.md) |  | |



## update

> UpdateTagResp update(updateRequest)



更新标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        UpdateTagReq updateTagReq = new UpdateTagReq(); // UpdateTagReq | 
        try {
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateTagReq(updateTagReq)
                .build();
            UpdateTagResp result = apiInstance.update(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#update");
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
| updateRequest | [**APIupdateRequest**](TagService.md#APIupdateRequest)|-|-|

### Return type

[**UpdateTagResp**](UpdateTagResp.md)


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

> ApiResponse<UpdateTagResp> updateWithHttpInfo(updateRequest)



更新标签。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        UpdateTagReq updateTagReq = new UpdateTagReq(); // UpdateTagReq | 
        try {
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateTagReq(updateTagReq)
                .build();
            ApiResponse<UpdateTagResp> response = apiInstance.updateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#update");
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
| updateRequest | [**APIupdateRequest**](TagService.md#APIupdateRequest)|-|-|

### Return type

ApiResponse<[**UpdateTagResp**](UpdateTagResp.md)>


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
| **updateTagReq** | [**UpdateTagReq**](UpdateTagReq.md) |  | |


