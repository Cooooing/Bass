# CharacterService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createCharacter**](CharacterService.md#createcharacter) | **POST** /v1/game-idle/character/create |  |
| [**listCharacter**](CharacterService.md#listcharacter) | **POST** /v1/game-idle/character/list |  |



## createCharacter

> CreateCharacterResp createCharacter(createCharacterReq)



创建角色。

### Example

```ts
import {
  Configuration,
  CharacterService,
} from '@bass/bbs-sdk-fetch';
import type { CreateCharacterRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CharacterService();

  const body = {
    // CreateCharacterReq
    createCharacterReq: ...,
  } satisfies CreateCharacterRequest;

  try {
    const data = await api.createCharacter(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createCharacterReq** | [CreateCharacterReq](CreateCharacterReq.md) |  | |

### Return type

[**CreateCharacterResp**](CreateCharacterResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listCharacter

> ListCharacterResp listCharacter(body)



查询当前账号角色列表。

### Example

```ts
import {
  Configuration,
  CharacterService,
} from '@bass/bbs-sdk-fetch';
import type { ListCharacterRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CharacterService();

  const body = {
    // object
    body: Object,
  } satisfies ListCharacterRequest;

  try {
    const data = await api.listCharacter(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | `object` |  | |

### Return type

[**ListCharacterResp**](ListCharacterResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

