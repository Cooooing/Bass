# HolidayService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentHoliday**](HolidayService.md#getCurrentHoliday) | **POST** /v1/platform/holiday/get-current |  |
| [**getCurrentHolidayWithHttpInfo**](HolidayService.md#getCurrentHolidayWithHttpInfo) | **POST** /v1/platform/holiday/get-current |  |



## getCurrentHoliday

> GetCurrentHolidayResp getCurrentHoliday(getCurrentHolidayRequest)



获取当前节假日状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.HolidayService;
import com.bass.bbs.api.HolidayService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        HolidayService apiInstance = new HolidayService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentHolidayRequest request = APIgetCurrentHolidayRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentHolidayResp result = apiInstance.getCurrentHoliday(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling HolidayService#getCurrentHoliday");
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
| getCurrentHolidayRequest | [**APIgetCurrentHolidayRequest**](HolidayService.md#APIgetCurrentHolidayRequest)|-|-|

### Return type

[**GetCurrentHolidayResp**](GetCurrentHolidayResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getCurrentHolidayWithHttpInfo

> ApiResponse<GetCurrentHolidayResp> getCurrentHolidayWithHttpInfo(getCurrentHolidayRequest)



获取当前节假日状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.HolidayService;
import com.bass.bbs.api.HolidayService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        HolidayService apiInstance = new HolidayService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentHolidayRequest request = APIgetCurrentHolidayRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentHolidayResp> response = apiInstance.getCurrentHolidayWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling HolidayService#getCurrentHoliday");
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
| getCurrentHolidayRequest | [**APIgetCurrentHolidayRequest**](HolidayService.md#APIgetCurrentHolidayRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentHolidayResp**](GetCurrentHolidayResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetCurrentHolidayRequest"></a>
## APIgetCurrentHolidayRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |


