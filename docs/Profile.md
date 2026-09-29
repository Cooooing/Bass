# Profile


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**account** | [**AccountProfile**](AccountProfile.md) |  | [optional] [default to undefined]
**background_url** | **string** |  | [optional] [default to undefined]
**location** | [**ProfileLocation**](ProfileLocation.md) |  | [optional] [default to undefined]
**last_success_login_at** | **string** |  | [optional] [default to undefined]
**visibility** | [**ProfileVisibility**](ProfileVisibility.md) |  | [optional] [default to undefined]
**viewer_relation** | [**ProfileRelation**](ProfileRelation.md) |  | [optional] [default to undefined]

## Example

```typescript
import { Profile } from '@bass/bbs-sdk-axios';

const instance: Profile = {
    account,
    background_url,
    location,
    last_success_login_at,
    visibility,
    viewer_relation,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
