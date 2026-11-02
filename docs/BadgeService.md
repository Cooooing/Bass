# BadgeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteDraft**](BadgeService.md#deleteDraft) | **POST** /v1/badge/draft/delete |  |
| [**deleteDraftWithHttpInfo**](BadgeService.md#deleteDraftWithHttpInfo) | **POST** /v1/badge/draft/delete |  |
| [**listDrafts**](BadgeService.md#listDrafts) | **POST** /v1/badge/draft/list |  |
| [**listDraftsWithHttpInfo**](BadgeService.md#listDraftsWithHttpInfo) | **POST** /v1/badge/draft/list |  |
| [**renderMaker**](BadgeService.md#renderMaker) | **GET** /v1/badge/maker/render |  |
| [**renderMakerWithHttpInfo**](BadgeService.md#renderMakerWithHttpInfo) | **GET** /v1/badge/maker/render |  |
| [**renderPublished**](BadgeService.md#renderPublished) | **GET** /v1/badge/render |  |
| [**renderPublishedWithHttpInfo**](BadgeService.md#renderPublishedWithHttpInfo) | **GET** /v1/badge/render |  |
| [**upsertDraft**](BadgeService.md#upsertDraft) | **POST** /v1/badge/draft/upsert |  |
| [**upsertDraftWithHttpInfo**](BadgeService.md#upsertDraftWithHttpInfo) | **POST** /v1/badge/draft/upsert |  |



## deleteDraft

> Object deleteDraft(deleteDraftRequest)



DeleteDraft 删除一个未发布草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        DeleteBadgeDraftReq deleteBadgeDraftReq = new DeleteBadgeDraftReq(); // DeleteBadgeDraftReq | 
        try {
            APIdeleteDraftRequest request = APIdeleteDraftRequest.newBuilder()
                .deleteBadgeDraftReq(deleteBadgeDraftReq)
                .build();
            Object result = apiInstance.deleteDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#deleteDraft");
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
| deleteDraftRequest | [**APIdeleteDraftRequest**](BadgeService.md#APIdeleteDraftRequest)|-|-|

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

## deleteDraftWithHttpInfo

> ApiResponse<Object> deleteDraftWithHttpInfo(deleteDraftRequest)



DeleteDraft 删除一个未发布草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        DeleteBadgeDraftReq deleteBadgeDraftReq = new DeleteBadgeDraftReq(); // DeleteBadgeDraftReq | 
        try {
            APIdeleteDraftRequest request = APIdeleteDraftRequest.newBuilder()
                .deleteBadgeDraftReq(deleteBadgeDraftReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#deleteDraft");
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
| deleteDraftRequest | [**APIdeleteDraftRequest**](BadgeService.md#APIdeleteDraftRequest)|-|-|

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


<a id="APIdeleteDraftRequest"></a>
## APIdeleteDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteBadgeDraftReq** | [**DeleteBadgeDraftReq**](DeleteBadgeDraftReq.md) |  | |



## listDrafts

> ListBadgeDraftsResp listDrafts(listDraftsRequest)



ListDrafts 返回当前登录用户的徽章草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        Object body = null; // Object | 
        try {
            APIlistDraftsRequest request = APIlistDraftsRequest.newBuilder()
                .body(body)
                .build();
            ListBadgeDraftsResp result = apiInstance.listDrafts(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#listDrafts");
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
| listDraftsRequest | [**APIlistDraftsRequest**](BadgeService.md#APIlistDraftsRequest)|-|-|

### Return type

[**ListBadgeDraftsResp**](ListBadgeDraftsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listDraftsWithHttpInfo

> ApiResponse<ListBadgeDraftsResp> listDraftsWithHttpInfo(listDraftsRequest)



ListDrafts 返回当前登录用户的徽章草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        Object body = null; // Object | 
        try {
            APIlistDraftsRequest request = APIlistDraftsRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ListBadgeDraftsResp> response = apiInstance.listDraftsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#listDrafts");
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
| listDraftsRequest | [**APIlistDraftsRequest**](BadgeService.md#APIlistDraftsRequest)|-|-|

### Return type

ApiResponse<[**ListBadgeDraftsResp**](ListBadgeDraftsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistDraftsRequest"></a>
## APIlistDraftsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## renderMaker

> ImageResp renderMaker(renderMakerRequest)



RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        String renderSpecImageAssetId = "renderSpecImageAssetId_example"; // String | 
        String renderSpecText = "renderSpecText_example"; // String | 
        List<String> renderSpecBackgroundColors = Arrays.asList(); // List<String> | 
        List<String> renderSpecTextColors = Arrays.asList(); // List<String> | 
        String renderSpecBackgroundDirection = "renderSpecBackgroundDirection_example"; // String | 
        String renderSpecTextDirection = "renderSpecTextDirection_example"; // String | 
        Integer renderSpecSize = 56; // Integer | 
        Integer renderSpecBorder = 56; // Integer | 
        Integer renderSpecTextBarLength = 56; // Integer | 
        Integer renderSpecTextSize = 56; // Integer | 
        Integer renderSpecTextBarRadius = 56; // Integer | 
        Double renderSpecScale = 3.4D; // Double | 
        Double renderSpecShadow = 3.4D; // Double | 
        Double renderSpecAnimationSeconds = 3.4D; // Double | 
        String renderSpecFontAssetId = "renderSpecFontAssetId_example"; // String | 
        try {
            APIrenderMakerRequest request = APIrenderMakerRequest.newBuilder()
                .renderSpecImageAssetId(renderSpecImageAssetId)
                .renderSpecText(renderSpecText)
                .renderSpecBackgroundColors(renderSpecBackgroundColors)
                .renderSpecTextColors(renderSpecTextColors)
                .renderSpecBackgroundDirection(renderSpecBackgroundDirection)
                .renderSpecTextDirection(renderSpecTextDirection)
                .renderSpecSize(renderSpecSize)
                .renderSpecBorder(renderSpecBorder)
                .renderSpecTextBarLength(renderSpecTextBarLength)
                .renderSpecTextSize(renderSpecTextSize)
                .renderSpecTextBarRadius(renderSpecTextBarRadius)
                .renderSpecScale(renderSpecScale)
                .renderSpecShadow(renderSpecShadow)
                .renderSpecAnimationSeconds(renderSpecAnimationSeconds)
                .renderSpecFontAssetId(renderSpecFontAssetId)
                .build();
            ImageResp result = apiInstance.renderMaker(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderMaker");
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
| renderMakerRequest | [**APIrenderMakerRequest**](BadgeService.md#APIrenderMakerRequest)|-|-|

### Return type

[**ImageResp**](ImageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renderMakerWithHttpInfo

> ApiResponse<ImageResp> renderMakerWithHttpInfo(renderMakerRequest)



RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        String renderSpecImageAssetId = "renderSpecImageAssetId_example"; // String | 
        String renderSpecText = "renderSpecText_example"; // String | 
        List<String> renderSpecBackgroundColors = Arrays.asList(); // List<String> | 
        List<String> renderSpecTextColors = Arrays.asList(); // List<String> | 
        String renderSpecBackgroundDirection = "renderSpecBackgroundDirection_example"; // String | 
        String renderSpecTextDirection = "renderSpecTextDirection_example"; // String | 
        Integer renderSpecSize = 56; // Integer | 
        Integer renderSpecBorder = 56; // Integer | 
        Integer renderSpecTextBarLength = 56; // Integer | 
        Integer renderSpecTextSize = 56; // Integer | 
        Integer renderSpecTextBarRadius = 56; // Integer | 
        Double renderSpecScale = 3.4D; // Double | 
        Double renderSpecShadow = 3.4D; // Double | 
        Double renderSpecAnimationSeconds = 3.4D; // Double | 
        String renderSpecFontAssetId = "renderSpecFontAssetId_example"; // String | 
        try {
            APIrenderMakerRequest request = APIrenderMakerRequest.newBuilder()
                .renderSpecImageAssetId(renderSpecImageAssetId)
                .renderSpecText(renderSpecText)
                .renderSpecBackgroundColors(renderSpecBackgroundColors)
                .renderSpecTextColors(renderSpecTextColors)
                .renderSpecBackgroundDirection(renderSpecBackgroundDirection)
                .renderSpecTextDirection(renderSpecTextDirection)
                .renderSpecSize(renderSpecSize)
                .renderSpecBorder(renderSpecBorder)
                .renderSpecTextBarLength(renderSpecTextBarLength)
                .renderSpecTextSize(renderSpecTextSize)
                .renderSpecTextBarRadius(renderSpecTextBarRadius)
                .renderSpecScale(renderSpecScale)
                .renderSpecShadow(renderSpecShadow)
                .renderSpecAnimationSeconds(renderSpecAnimationSeconds)
                .renderSpecFontAssetId(renderSpecFontAssetId)
                .build();
            ApiResponse<ImageResp> response = apiInstance.renderMakerWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderMaker");
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
| renderMakerRequest | [**APIrenderMakerRequest**](BadgeService.md#APIrenderMakerRequest)|-|-|

### Return type

ApiResponse<[**ImageResp**](ImageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenderMakerRequest"></a>
## APIrenderMakerRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renderSpecImageAssetId** | **String** |  | [optional] |
| **renderSpecText** | **String** |  | [optional] |
| **renderSpecBackgroundColors** | [**List&lt;String&gt;**](String.md) |  | [optional] |
| **renderSpecTextColors** | [**List&lt;String&gt;**](String.md) |  | [optional] |
| **renderSpecBackgroundDirection** | **String** |  | [optional] |
| **renderSpecTextDirection** | **String** |  | [optional] |
| **renderSpecSize** | **Integer** |  | [optional] |
| **renderSpecBorder** | **Integer** |  | [optional] |
| **renderSpecTextBarLength** | **Integer** |  | [optional] |
| **renderSpecTextSize** | **Integer** |  | [optional] |
| **renderSpecTextBarRadius** | **Integer** |  | [optional] |
| **renderSpecScale** | **Double** |  | [optional] |
| **renderSpecShadow** | **Double** |  | [optional] |
| **renderSpecAnimationSeconds** | **Double** |  | [optional] |
| **renderSpecFontAssetId** | **String** |  | [optional] |



## renderPublished

> ImageResp renderPublished(renderPublishedRequest)



RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        String code = "code_example"; // String | 
        try {
            APIrenderPublishedRequest request = APIrenderPublishedRequest.newBuilder()
                .code(code)
                .build();
            ImageResp result = apiInstance.renderPublished(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderPublished");
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
| renderPublishedRequest | [**APIrenderPublishedRequest**](BadgeService.md#APIrenderPublishedRequest)|-|-|

### Return type

[**ImageResp**](ImageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renderPublishedWithHttpInfo

> ApiResponse<ImageResp> renderPublishedWithHttpInfo(renderPublishedRequest)



RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        String code = "code_example"; // String | 
        try {
            APIrenderPublishedRequest request = APIrenderPublishedRequest.newBuilder()
                .code(code)
                .build();
            ApiResponse<ImageResp> response = apiInstance.renderPublishedWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderPublished");
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
| renderPublishedRequest | [**APIrenderPublishedRequest**](BadgeService.md#APIrenderPublishedRequest)|-|-|

### Return type

ApiResponse<[**ImageResp**](ImageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenderPublishedRequest"></a>
## APIrenderPublishedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **code** | **String** |  | [optional] |



## upsertDraft

> UpsertBadgeDraftResp upsertDraft(upsertDraftRequest)



UpsertDraft 保存当前登录用户的徽章草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        UpsertBadgeDraftReq upsertBadgeDraftReq = new UpsertBadgeDraftReq(); // UpsertBadgeDraftReq | 
        try {
            APIupsertDraftRequest request = APIupsertDraftRequest.newBuilder()
                .upsertBadgeDraftReq(upsertBadgeDraftReq)
                .build();
            UpsertBadgeDraftResp result = apiInstance.upsertDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#upsertDraft");
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
| upsertDraftRequest | [**APIupsertDraftRequest**](BadgeService.md#APIupsertDraftRequest)|-|-|

### Return type

[**UpsertBadgeDraftResp**](UpsertBadgeDraftResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## upsertDraftWithHttpInfo

> ApiResponse<UpsertBadgeDraftResp> upsertDraftWithHttpInfo(upsertDraftRequest)



UpsertDraft 保存当前登录用户的徽章草稿。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.BadgeService;
import com.bass.bbs.api.BadgeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        BadgeService apiInstance = new BadgeService(defaultClient);
        UpsertBadgeDraftReq upsertBadgeDraftReq = new UpsertBadgeDraftReq(); // UpsertBadgeDraftReq | 
        try {
            APIupsertDraftRequest request = APIupsertDraftRequest.newBuilder()
                .upsertBadgeDraftReq(upsertBadgeDraftReq)
                .build();
            ApiResponse<UpsertBadgeDraftResp> response = apiInstance.upsertDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#upsertDraft");
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
| upsertDraftRequest | [**APIupsertDraftRequest**](BadgeService.md#APIupsertDraftRequest)|-|-|

### Return type

ApiResponse<[**UpsertBadgeDraftResp**](UpsertBadgeDraftResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupsertDraftRequest"></a>
## APIupsertDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **upsertBadgeDraftReq** | [**UpsertBadgeDraftReq**](UpsertBadgeDraftReq.md) |  | |


