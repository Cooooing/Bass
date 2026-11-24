# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**bindArticleTags**](TagService.md#bindArticleTags) | **POST** /v1/content/tag/bind-article |  |
| [**bindArticleTagsWithHttpInfo**](TagService.md#bindArticleTagsWithHttpInfo) | **POST** /v1/content/tag/bind-article |  |
| [**createTag**](TagService.md#createTag) | **POST** /v1/content/tag/create |  |
| [**createTagWithHttpInfo**](TagService.md#createTagWithHttpInfo) | **POST** /v1/content/tag/create |  |
| [**listArticleTags**](TagService.md#listArticleTags) | **POST** /v1/content/tag/list-article-tags |  |
| [**listArticleTagsWithHttpInfo**](TagService.md#listArticleTagsWithHttpInfo) | **POST** /v1/content/tag/list-article-tags |  |
| [**listTags**](TagService.md#listTags) | **POST** /v1/content/tag/list |  |
| [**listTagsWithHttpInfo**](TagService.md#listTagsWithHttpInfo) | **POST** /v1/content/tag/list |  |
| [**unbindArticleTags**](TagService.md#unbindArticleTags) | **POST** /v1/content/tag/unbind-article |  |
| [**unbindArticleTagsWithHttpInfo**](TagService.md#unbindArticleTagsWithHttpInfo) | **POST** /v1/content/tag/unbind-article |  |
| [**updateTag**](TagService.md#updateTag) | **POST** /v1/content/tag/update |  |
| [**updateTagWithHttpInfo**](TagService.md#updateTagWithHttpInfo) | **POST** /v1/content/tag/update |  |



## bindArticleTags

> Object bindArticleTags(bindArticleTagsRequest)



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
            APIbindArticleTagsRequest request = APIbindArticleTagsRequest.newBuilder()
                .bindArticleTagsReq(bindArticleTagsReq)
                .build();
            Object result = apiInstance.bindArticleTags(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#bindArticleTags");
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
| bindArticleTagsRequest | [**APIbindArticleTagsRequest**](TagService.md#APIbindArticleTagsRequest)|-|-|

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

## bindArticleTagsWithHttpInfo

> ApiResponse<Object> bindArticleTagsWithHttpInfo(bindArticleTagsRequest)



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
            APIbindArticleTagsRequest request = APIbindArticleTagsRequest.newBuilder()
                .bindArticleTagsReq(bindArticleTagsReq)
                .build();
            ApiResponse<Object> response = apiInstance.bindArticleTagsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#bindArticleTags");
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
| bindArticleTagsRequest | [**APIbindArticleTagsRequest**](TagService.md#APIbindArticleTagsRequest)|-|-|

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


<a id="APIbindArticleTagsRequest"></a>
## APIbindArticleTagsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **bindArticleTagsReq** | [**BindArticleTagsReq**](BindArticleTagsReq.md) |  | |



## createTag

> CreateTagResp createTag(createTagRequest)



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
            APIcreateTagRequest request = APIcreateTagRequest.newBuilder()
                .createTagReq(createTagReq)
                .build();
            CreateTagResp result = apiInstance.createTag(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#createTag");
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
| createTagRequest | [**APIcreateTagRequest**](TagService.md#APIcreateTagRequest)|-|-|

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

## createTagWithHttpInfo

> ApiResponse<CreateTagResp> createTagWithHttpInfo(createTagRequest)



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
            APIcreateTagRequest request = APIcreateTagRequest.newBuilder()
                .createTagReq(createTagReq)
                .build();
            ApiResponse<CreateTagResp> response = apiInstance.createTagWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#createTag");
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
| createTagRequest | [**APIcreateTagRequest**](TagService.md#APIcreateTagRequest)|-|-|

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


<a id="APIcreateTagRequest"></a>
## APIcreateTagRequest
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



## listTags

> ListTagsResp listTags(listTagsRequest)



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
            APIlistTagsRequest request = APIlistTagsRequest.newBuilder()
                .listTagsReq(listTagsReq)
                .build();
            ListTagsResp result = apiInstance.listTags(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#listTags");
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
| listTagsRequest | [**APIlistTagsRequest**](TagService.md#APIlistTagsRequest)|-|-|

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

## listTagsWithHttpInfo

> ApiResponse<ListTagsResp> listTagsWithHttpInfo(listTagsRequest)



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
            APIlistTagsRequest request = APIlistTagsRequest.newBuilder()
                .listTagsReq(listTagsReq)
                .build();
            ApiResponse<ListTagsResp> response = apiInstance.listTagsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#listTags");
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
| listTagsRequest | [**APIlistTagsRequest**](TagService.md#APIlistTagsRequest)|-|-|

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


<a id="APIlistTagsRequest"></a>
## APIlistTagsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listTagsReq** | [**ListTagsReq**](ListTagsReq.md) |  | |



## unbindArticleTags

> Object unbindArticleTags(unbindArticleTagsRequest)



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
            APIunbindArticleTagsRequest request = APIunbindArticleTagsRequest.newBuilder()
                .unbindArticleTagsReq(unbindArticleTagsReq)
                .build();
            Object result = apiInstance.unbindArticleTags(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#unbindArticleTags");
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
| unbindArticleTagsRequest | [**APIunbindArticleTagsRequest**](TagService.md#APIunbindArticleTagsRequest)|-|-|

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

## unbindArticleTagsWithHttpInfo

> ApiResponse<Object> unbindArticleTagsWithHttpInfo(unbindArticleTagsRequest)



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
            APIunbindArticleTagsRequest request = APIunbindArticleTagsRequest.newBuilder()
                .unbindArticleTagsReq(unbindArticleTagsReq)
                .build();
            ApiResponse<Object> response = apiInstance.unbindArticleTagsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#unbindArticleTags");
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
| unbindArticleTagsRequest | [**APIunbindArticleTagsRequest**](TagService.md#APIunbindArticleTagsRequest)|-|-|

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


<a id="APIunbindArticleTagsRequest"></a>
## APIunbindArticleTagsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **unbindArticleTagsReq** | [**UnbindArticleTagsReq**](UnbindArticleTagsReq.md) |  | |



## updateTag

> UpdateTagResp updateTag(updateTagRequest)



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
            APIupdateTagRequest request = APIupdateTagRequest.newBuilder()
                .updateTagReq(updateTagReq)
                .build();
            UpdateTagResp result = apiInstance.updateTag(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#updateTag");
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
| updateTagRequest | [**APIupdateTagRequest**](TagService.md#APIupdateTagRequest)|-|-|

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

## updateTagWithHttpInfo

> ApiResponse<UpdateTagResp> updateTagWithHttpInfo(updateTagRequest)



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
            APIupdateTagRequest request = APIupdateTagRequest.newBuilder()
                .updateTagReq(updateTagReq)
                .build();
            ApiResponse<UpdateTagResp> response = apiInstance.updateTagWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#updateTag");
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
| updateTagRequest | [**APIupdateTagRequest**](TagService.md#APIupdateTagRequest)|-|-|

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


<a id="APIupdateTagRequest"></a>
## APIupdateTagRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateTagReq** | [**UpdateTagReq**](UpdateTagReq.md) |  | |


