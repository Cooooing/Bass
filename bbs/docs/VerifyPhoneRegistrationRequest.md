# VerifyPhoneRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 验证码。 | 
**CodeToken** | **string** | 验证码令牌。 | 

## Methods

### NewVerifyPhoneRegistrationRequest

`func NewVerifyPhoneRegistrationRequest(code string, codeToken string, ) *VerifyPhoneRegistrationRequest`

NewVerifyPhoneRegistrationRequest instantiates a new VerifyPhoneRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyPhoneRegistrationRequestWithDefaults

`func NewVerifyPhoneRegistrationRequestWithDefaults() *VerifyPhoneRegistrationRequest`

NewVerifyPhoneRegistrationRequestWithDefaults instantiates a new VerifyPhoneRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VerifyPhoneRegistrationRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyPhoneRegistrationRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyPhoneRegistrationRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeToken

`func (o *VerifyPhoneRegistrationRequest) GetCodeToken() string`

GetCodeToken returns the CodeToken field if non-nil, zero value otherwise.

### GetCodeTokenOk

`func (o *VerifyPhoneRegistrationRequest) GetCodeTokenOk() (*string, bool)`

GetCodeTokenOk returns a tuple with the CodeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeToken

`func (o *VerifyPhoneRegistrationRequest) SetCodeToken(v string)`

SetCodeToken sets CodeToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


