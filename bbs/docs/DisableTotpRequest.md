# DisableTotpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | TOTP 验证码。 | 

## Methods

### NewDisableTotpRequest

`func NewDisableTotpRequest(code string, ) *DisableTotpRequest`

NewDisableTotpRequest instantiates a new DisableTotpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableTotpRequestWithDefaults

`func NewDisableTotpRequestWithDefaults() *DisableTotpRequest`

NewDisableTotpRequestWithDefaults instantiates a new DisableTotpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DisableTotpRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DisableTotpRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DisableTotpRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


