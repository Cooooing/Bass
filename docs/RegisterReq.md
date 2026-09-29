

# RegisterReq


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**type** | [**TypeEnum**](#TypeEnum) |  |  |
|**name** | **String** |  |  |
|**password** | **String** |  |  |
|**nickname** | **String** |  |  [optional] |
|**emailCredential** | [**ReqEmailCredential**](ReqEmailCredential.md) |  |  [optional] |
|**phoneCredential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) |  |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| REGISTER_TYPE_UNSPECIFIED | &quot;REGISTER_TYPE_UNSPECIFIED&quot; |
| REGISTER_TYPE_EMAIL | &quot;REGISTER_TYPE_EMAIL&quot; |
| REGISTER_TYPE_PHONE | &quot;REGISTER_TYPE_PHONE&quot; |



