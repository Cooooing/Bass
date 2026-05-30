# VerifyEmailRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 验证码。 | 
**CodeToken** | **string** | 验证码令牌。 | 

## Methods

### NewVerifyEmailRegistrationRequest

`func NewVerifyEmailRegistrationRequest(code string, codeToken string, ) *VerifyEmailRegistrationRequest`

NewVerifyEmailRegistrationRequest instantiates a new VerifyEmailRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEmailRegistrationRequestWithDefaults

`func NewVerifyEmailRegistrationRequestWithDefaults() *VerifyEmailRegistrationRequest`

NewVerifyEmailRegistrationRequestWithDefaults instantiates a new VerifyEmailRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VerifyEmailRegistrationRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyEmailRegistrationRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyEmailRegistrationRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeToken

`func (o *VerifyEmailRegistrationRequest) GetCodeToken() string`

GetCodeToken returns the CodeToken field if non-nil, zero value otherwise.

### GetCodeTokenOk

`func (o *VerifyEmailRegistrationRequest) GetCodeTokenOk() (*string, bool)`

GetCodeTokenOk returns a tuple with the CodeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeToken

`func (o *VerifyEmailRegistrationRequest) SetCodeToken(v string)`

SetCodeToken sets CodeToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


