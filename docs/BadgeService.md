# BadgeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteBadgeDraft**](BadgeService.md#deleteBadgeDraft) | **POST** /v1/badge/draft/delete |  |
| [**deleteBadgeDraftWithHttpInfo**](BadgeService.md#deleteBadgeDraftWithHttpInfo) | **POST** /v1/badge/draft/delete |  |
| [**listBadgeDrafts**](BadgeService.md#listBadgeDrafts) | **POST** /v1/badge/draft/list |  |
| [**listBadgeDraftsWithHttpInfo**](BadgeService.md#listBadgeDraftsWithHttpInfo) | **POST** /v1/badge/draft/list |  |
| [**renderMakerBadge**](BadgeService.md#renderMakerBadge) | **GET** /v1/badge/maker/render |  |
| [**renderMakerBadgeWithHttpInfo**](BadgeService.md#renderMakerBadgeWithHttpInfo) | **GET** /v1/badge/maker/render |  |
| [**renderPublishedBadge**](BadgeService.md#renderPublishedBadge) | **GET** /v1/badge/render |  |
| [**renderPublishedBadgeWithHttpInfo**](BadgeService.md#renderPublishedBadgeWithHttpInfo) | **GET** /v1/badge/render |  |
| [**upsertBadgeDraft**](BadgeService.md#upsertBadgeDraft) | **POST** /v1/badge/draft/upsert |  |
| [**upsertBadgeDraftWithHttpInfo**](BadgeService.md#upsertBadgeDraftWithHttpInfo) | **POST** /v1/badge/draft/upsert |  |



## deleteBadgeDraft

> Object deleteBadgeDraft(deleteBadgeDraftRequest)



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
            APIdeleteBadgeDraftRequest request = APIdeleteBadgeDraftRequest.newBuilder()
                .deleteBadgeDraftReq(deleteBadgeDraftReq)
                .build();
            Object result = apiInstance.deleteBadgeDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#deleteBadgeDraft");
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
| deleteBadgeDraftRequest | [**APIdeleteBadgeDraftRequest**](BadgeService.md#APIdeleteBadgeDraftRequest)|-|-|

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

## deleteBadgeDraftWithHttpInfo

> ApiResponse<Object> deleteBadgeDraftWithHttpInfo(deleteBadgeDraftRequest)



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
            APIdeleteBadgeDraftRequest request = APIdeleteBadgeDraftRequest.newBuilder()
                .deleteBadgeDraftReq(deleteBadgeDraftReq)
                .build();
            ApiResponse<Object> response = apiInstance.deleteBadgeDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#deleteBadgeDraft");
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
| deleteBadgeDraftRequest | [**APIdeleteBadgeDraftRequest**](BadgeService.md#APIdeleteBadgeDraftRequest)|-|-|

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


<a id="APIdeleteBadgeDraftRequest"></a>
## APIdeleteBadgeDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteBadgeDraftReq** | [**DeleteBadgeDraftReq**](DeleteBadgeDraftReq.md) |  | |



## listBadgeDrafts

> ListBadgeDraftsResp listBadgeDrafts(listBadgeDraftsRequest)



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
            APIlistBadgeDraftsRequest request = APIlistBadgeDraftsRequest.newBuilder()
                .body(body)
                .build();
            ListBadgeDraftsResp result = apiInstance.listBadgeDrafts(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#listBadgeDrafts");
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
| listBadgeDraftsRequest | [**APIlistBadgeDraftsRequest**](BadgeService.md#APIlistBadgeDraftsRequest)|-|-|

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

## listBadgeDraftsWithHttpInfo

> ApiResponse<ListBadgeDraftsResp> listBadgeDraftsWithHttpInfo(listBadgeDraftsRequest)



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
            APIlistBadgeDraftsRequest request = APIlistBadgeDraftsRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ListBadgeDraftsResp> response = apiInstance.listBadgeDraftsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#listBadgeDrafts");
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
| listBadgeDraftsRequest | [**APIlistBadgeDraftsRequest**](BadgeService.md#APIlistBadgeDraftsRequest)|-|-|

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


<a id="APIlistBadgeDraftsRequest"></a>
## APIlistBadgeDraftsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## renderMakerBadge

> RenderMakerBadgeResp renderMakerBadge(renderMakerBadgeRequest)



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
        String renderSpecImageAssetId = "renderSpecImageAssetId_example"; // String | Platform 已确认可用的徽章图标资源 ID。
        String renderSpecText = "renderSpecText_example"; // String | 文字条内容。
        List<String> renderSpecBackgroundColors = Arrays.asList(); // List<String> | 背景颜色列表。
        List<String> renderSpecTextColors = Arrays.asList(); // List<String> | 文字颜色列表。
        String renderSpecBackgroundDirection = "renderSpecBackgroundDirection_example"; // String | 背景渐变方向或角度。
        String renderSpecTextDirection = "renderSpecTextDirection_example"; // String | 文字渐变方向或角度。
        Integer renderSpecSize = 56; // Integer | 图标边长。
        Integer renderSpecBorder = 56; // Integer | 图标边距。
        Integer renderSpecTextBarLength = 56; // Integer | 文字条长度。
        Integer renderSpecTextSize = 56; // Integer | 文字字号。
        Integer renderSpecTextBarRadius = 56; // Integer | 文字条圆角。
        Double renderSpecScale = 3.4D; // Double | 整体缩放比例。
        Double renderSpecShadow = 3.4D; // Double | 阴影强度。
        Double renderSpecAnimationSeconds = 3.4D; // Double | 动画时长。
        String renderSpecFontAssetId = "renderSpecFontAssetId_example"; // String | Platform 已确认可用的自定义字体资源 ID。
        try {
            APIrenderMakerBadgeRequest request = APIrenderMakerBadgeRequest.newBuilder()
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
            RenderMakerBadgeResp result = apiInstance.renderMakerBadge(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderMakerBadge");
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
| renderMakerBadgeRequest | [**APIrenderMakerBadgeRequest**](BadgeService.md#APIrenderMakerBadgeRequest)|-|-|

### Return type

[**RenderMakerBadgeResp**](RenderMakerBadgeResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renderMakerBadgeWithHttpInfo

> ApiResponse<RenderMakerBadgeResp> renderMakerBadgeWithHttpInfo(renderMakerBadgeRequest)



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
        String renderSpecImageAssetId = "renderSpecImageAssetId_example"; // String | Platform 已确认可用的徽章图标资源 ID。
        String renderSpecText = "renderSpecText_example"; // String | 文字条内容。
        List<String> renderSpecBackgroundColors = Arrays.asList(); // List<String> | 背景颜色列表。
        List<String> renderSpecTextColors = Arrays.asList(); // List<String> | 文字颜色列表。
        String renderSpecBackgroundDirection = "renderSpecBackgroundDirection_example"; // String | 背景渐变方向或角度。
        String renderSpecTextDirection = "renderSpecTextDirection_example"; // String | 文字渐变方向或角度。
        Integer renderSpecSize = 56; // Integer | 图标边长。
        Integer renderSpecBorder = 56; // Integer | 图标边距。
        Integer renderSpecTextBarLength = 56; // Integer | 文字条长度。
        Integer renderSpecTextSize = 56; // Integer | 文字字号。
        Integer renderSpecTextBarRadius = 56; // Integer | 文字条圆角。
        Double renderSpecScale = 3.4D; // Double | 整体缩放比例。
        Double renderSpecShadow = 3.4D; // Double | 阴影强度。
        Double renderSpecAnimationSeconds = 3.4D; // Double | 动画时长。
        String renderSpecFontAssetId = "renderSpecFontAssetId_example"; // String | Platform 已确认可用的自定义字体资源 ID。
        try {
            APIrenderMakerBadgeRequest request = APIrenderMakerBadgeRequest.newBuilder()
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
            ApiResponse<RenderMakerBadgeResp> response = apiInstance.renderMakerBadgeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderMakerBadge");
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
| renderMakerBadgeRequest | [**APIrenderMakerBadgeRequest**](BadgeService.md#APIrenderMakerBadgeRequest)|-|-|

### Return type

ApiResponse<[**RenderMakerBadgeResp**](RenderMakerBadgeResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenderMakerBadgeRequest"></a>
## APIrenderMakerBadgeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **renderSpecImageAssetId** | **String** | Platform 已确认可用的徽章图标资源 ID。 | [optional] |
| **renderSpecText** | **String** | 文字条内容。 | [optional] |
| **renderSpecBackgroundColors** | [**List&lt;String&gt;**](String.md) | 背景颜色列表。 | [optional] |
| **renderSpecTextColors** | [**List&lt;String&gt;**](String.md) | 文字颜色列表。 | [optional] |
| **renderSpecBackgroundDirection** | **String** | 背景渐变方向或角度。 | [optional] |
| **renderSpecTextDirection** | **String** | 文字渐变方向或角度。 | [optional] |
| **renderSpecSize** | **Integer** | 图标边长。 | [optional] |
| **renderSpecBorder** | **Integer** | 图标边距。 | [optional] |
| **renderSpecTextBarLength** | **Integer** | 文字条长度。 | [optional] |
| **renderSpecTextSize** | **Integer** | 文字字号。 | [optional] |
| **renderSpecTextBarRadius** | **Integer** | 文字条圆角。 | [optional] |
| **renderSpecScale** | **Double** | 整体缩放比例。 | [optional] |
| **renderSpecShadow** | **Double** | 阴影强度。 | [optional] |
| **renderSpecAnimationSeconds** | **Double** | 动画时长。 | [optional] |
| **renderSpecFontAssetId** | **String** | Platform 已确认可用的自定义字体资源 ID。 | [optional] |



## renderPublishedBadge

> RenderPublishedBadgeResp renderPublishedBadge(renderPublishedBadgeRequest)



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
        String code = "code_example"; // String | 已发布徽章稳定编码。
        try {
            APIrenderPublishedBadgeRequest request = APIrenderPublishedBadgeRequest.newBuilder()
                .code(code)
                .build();
            RenderPublishedBadgeResp result = apiInstance.renderPublishedBadge(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderPublishedBadge");
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
| renderPublishedBadgeRequest | [**APIrenderPublishedBadgeRequest**](BadgeService.md#APIrenderPublishedBadgeRequest)|-|-|

### Return type

[**RenderPublishedBadgeResp**](RenderPublishedBadgeResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## renderPublishedBadgeWithHttpInfo

> ApiResponse<RenderPublishedBadgeResp> renderPublishedBadgeWithHttpInfo(renderPublishedBadgeRequest)



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
        String code = "code_example"; // String | 已发布徽章稳定编码。
        try {
            APIrenderPublishedBadgeRequest request = APIrenderPublishedBadgeRequest.newBuilder()
                .code(code)
                .build();
            ApiResponse<RenderPublishedBadgeResp> response = apiInstance.renderPublishedBadgeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#renderPublishedBadge");
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
| renderPublishedBadgeRequest | [**APIrenderPublishedBadgeRequest**](BadgeService.md#APIrenderPublishedBadgeRequest)|-|-|

### Return type

ApiResponse<[**RenderPublishedBadgeResp**](RenderPublishedBadgeResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrenderPublishedBadgeRequest"></a>
## APIrenderPublishedBadgeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **code** | **String** | 已发布徽章稳定编码。 | [optional] |



## upsertBadgeDraft

> UpsertBadgeDraftResp upsertBadgeDraft(upsertBadgeDraftRequest)



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
            APIupsertBadgeDraftRequest request = APIupsertBadgeDraftRequest.newBuilder()
                .upsertBadgeDraftReq(upsertBadgeDraftReq)
                .build();
            UpsertBadgeDraftResp result = apiInstance.upsertBadgeDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#upsertBadgeDraft");
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
| upsertBadgeDraftRequest | [**APIupsertBadgeDraftRequest**](BadgeService.md#APIupsertBadgeDraftRequest)|-|-|

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

## upsertBadgeDraftWithHttpInfo

> ApiResponse<UpsertBadgeDraftResp> upsertBadgeDraftWithHttpInfo(upsertBadgeDraftRequest)



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
            APIupsertBadgeDraftRequest request = APIupsertBadgeDraftRequest.newBuilder()
                .upsertBadgeDraftReq(upsertBadgeDraftReq)
                .build();
            ApiResponse<UpsertBadgeDraftResp> response = apiInstance.upsertBadgeDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling BadgeService#upsertBadgeDraft");
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
| upsertBadgeDraftRequest | [**APIupsertBadgeDraftRequest**](BadgeService.md#APIupsertBadgeDraftRequest)|-|-|

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


<a id="APIupsertBadgeDraftRequest"></a>
## APIupsertBadgeDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **upsertBadgeDraftReq** | [**UpsertBadgeDraftReq**](UpsertBadgeDraftReq.md) |  | |


