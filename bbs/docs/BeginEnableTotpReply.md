# BeginEnableTotpReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | TOTP 绑定 URL。 | [optional] 
**QrCode** | Pointer to **string** | TOTP 二维码图片数据。 | [optional] 

## Methods

### NewBeginEnableTotpReply

`func NewBeginEnableTotpReply() *BeginEnableTotpReply`

NewBeginEnableTotpReply instantiates a new BeginEnableTotpReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeginEnableTotpReplyWithDefaults

`func NewBeginEnableTotpReplyWithDefaults() *BeginEnableTotpReply`

NewBeginEnableTotpReplyWithDefaults instantiates a new BeginEnableTotpReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BeginEnableTotpReply) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BeginEnableTotpReply) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BeginEnableTotpReply) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BeginEnableTotpReply) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetQrCode

`func (o *BeginEnableTotpReply) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *BeginEnableTotpReply) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *BeginEnableTotpReply) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *BeginEnableTotpReply) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


