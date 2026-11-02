# RenderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageAssetId** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**BackgroundColors** | Pointer to **[]string** |  | [optional] 
**TextColors** | Pointer to **[]string** |  | [optional] 
**BackgroundDirection** | Pointer to **string** |  | [optional] 
**TextDirection** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Border** | Pointer to **int32** |  | [optional] 
**TextBarLength** | Pointer to **int32** |  | [optional] 
**TextSize** | Pointer to **int32** |  | [optional] 
**TextBarRadius** | Pointer to **int32** |  | [optional] 
**Scale** | Pointer to **float64** |  | [optional] 
**Shadow** | Pointer to **float64** |  | [optional] 
**AnimationSeconds** | Pointer to **float64** |  | [optional] 
**FontAssetId** | Pointer to **string** |  | [optional] 

## Methods

### NewRenderSpec

`func NewRenderSpec() *RenderSpec`

NewRenderSpec instantiates a new RenderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderSpecWithDefaults

`func NewRenderSpecWithDefaults() *RenderSpec`

NewRenderSpecWithDefaults instantiates a new RenderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageAssetId

`func (o *RenderSpec) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *RenderSpec) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *RenderSpec) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *RenderSpec) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### GetText

`func (o *RenderSpec) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RenderSpec) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RenderSpec) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RenderSpec) HasText() bool`

HasText returns a boolean if a field has been set.

### GetBackgroundColors

`func (o *RenderSpec) GetBackgroundColors() []string`

GetBackgroundColors returns the BackgroundColors field if non-nil, zero value otherwise.

### GetBackgroundColorsOk

`func (o *RenderSpec) GetBackgroundColorsOk() (*[]string, bool)`

GetBackgroundColorsOk returns a tuple with the BackgroundColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColors

`func (o *RenderSpec) SetBackgroundColors(v []string)`

SetBackgroundColors sets BackgroundColors field to given value.

### HasBackgroundColors

`func (o *RenderSpec) HasBackgroundColors() bool`

HasBackgroundColors returns a boolean if a field has been set.

### GetTextColors

`func (o *RenderSpec) GetTextColors() []string`

GetTextColors returns the TextColors field if non-nil, zero value otherwise.

### GetTextColorsOk

`func (o *RenderSpec) GetTextColorsOk() (*[]string, bool)`

GetTextColorsOk returns a tuple with the TextColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColors

`func (o *RenderSpec) SetTextColors(v []string)`

SetTextColors sets TextColors field to given value.

### HasTextColors

`func (o *RenderSpec) HasTextColors() bool`

HasTextColors returns a boolean if a field has been set.

### GetBackgroundDirection

`func (o *RenderSpec) GetBackgroundDirection() string`

GetBackgroundDirection returns the BackgroundDirection field if non-nil, zero value otherwise.

### GetBackgroundDirectionOk

`func (o *RenderSpec) GetBackgroundDirectionOk() (*string, bool)`

GetBackgroundDirectionOk returns a tuple with the BackgroundDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundDirection

`func (o *RenderSpec) SetBackgroundDirection(v string)`

SetBackgroundDirection sets BackgroundDirection field to given value.

### HasBackgroundDirection

`func (o *RenderSpec) HasBackgroundDirection() bool`

HasBackgroundDirection returns a boolean if a field has been set.

### GetTextDirection

`func (o *RenderSpec) GetTextDirection() string`

GetTextDirection returns the TextDirection field if non-nil, zero value otherwise.

### GetTextDirectionOk

`func (o *RenderSpec) GetTextDirectionOk() (*string, bool)`

GetTextDirectionOk returns a tuple with the TextDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDirection

`func (o *RenderSpec) SetTextDirection(v string)`

SetTextDirection sets TextDirection field to given value.

### HasTextDirection

`func (o *RenderSpec) HasTextDirection() bool`

HasTextDirection returns a boolean if a field has been set.

### GetSize

`func (o *RenderSpec) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RenderSpec) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RenderSpec) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RenderSpec) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetBorder

`func (o *RenderSpec) GetBorder() int32`

GetBorder returns the Border field if non-nil, zero value otherwise.

### GetBorderOk

`func (o *RenderSpec) GetBorderOk() (*int32, bool)`

GetBorderOk returns a tuple with the Border field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorder

`func (o *RenderSpec) SetBorder(v int32)`

SetBorder sets Border field to given value.

### HasBorder

`func (o *RenderSpec) HasBorder() bool`

HasBorder returns a boolean if a field has been set.

### GetTextBarLength

`func (o *RenderSpec) GetTextBarLength() int32`

GetTextBarLength returns the TextBarLength field if non-nil, zero value otherwise.

### GetTextBarLengthOk

`func (o *RenderSpec) GetTextBarLengthOk() (*int32, bool)`

GetTextBarLengthOk returns a tuple with the TextBarLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBarLength

`func (o *RenderSpec) SetTextBarLength(v int32)`

SetTextBarLength sets TextBarLength field to given value.

### HasTextBarLength

`func (o *RenderSpec) HasTextBarLength() bool`

HasTextBarLength returns a boolean if a field has been set.

### GetTextSize

`func (o *RenderSpec) GetTextSize() int32`

GetTextSize returns the TextSize field if non-nil, zero value otherwise.

### GetTextSizeOk

`func (o *RenderSpec) GetTextSizeOk() (*int32, bool)`

GetTextSizeOk returns a tuple with the TextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSize

`func (o *RenderSpec) SetTextSize(v int32)`

SetTextSize sets TextSize field to given value.

### HasTextSize

`func (o *RenderSpec) HasTextSize() bool`

HasTextSize returns a boolean if a field has been set.

### GetTextBarRadius

`func (o *RenderSpec) GetTextBarRadius() int32`

GetTextBarRadius returns the TextBarRadius field if non-nil, zero value otherwise.

### GetTextBarRadiusOk

`func (o *RenderSpec) GetTextBarRadiusOk() (*int32, bool)`

GetTextBarRadiusOk returns a tuple with the TextBarRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBarRadius

`func (o *RenderSpec) SetTextBarRadius(v int32)`

SetTextBarRadius sets TextBarRadius field to given value.

### HasTextBarRadius

`func (o *RenderSpec) HasTextBarRadius() bool`

HasTextBarRadius returns a boolean if a field has been set.

### GetScale

`func (o *RenderSpec) GetScale() float64`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *RenderSpec) GetScaleOk() (*float64, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *RenderSpec) SetScale(v float64)`

SetScale sets Scale field to given value.

### HasScale

`func (o *RenderSpec) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetShadow

`func (o *RenderSpec) GetShadow() float64`

GetShadow returns the Shadow field if non-nil, zero value otherwise.

### GetShadowOk

`func (o *RenderSpec) GetShadowOk() (*float64, bool)`

GetShadowOk returns a tuple with the Shadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadow

`func (o *RenderSpec) SetShadow(v float64)`

SetShadow sets Shadow field to given value.

### HasShadow

`func (o *RenderSpec) HasShadow() bool`

HasShadow returns a boolean if a field has been set.

### GetAnimationSeconds

`func (o *RenderSpec) GetAnimationSeconds() float64`

GetAnimationSeconds returns the AnimationSeconds field if non-nil, zero value otherwise.

### GetAnimationSecondsOk

`func (o *RenderSpec) GetAnimationSecondsOk() (*float64, bool)`

GetAnimationSecondsOk returns a tuple with the AnimationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationSeconds

`func (o *RenderSpec) SetAnimationSeconds(v float64)`

SetAnimationSeconds sets AnimationSeconds field to given value.

### HasAnimationSeconds

`func (o *RenderSpec) HasAnimationSeconds() bool`

HasAnimationSeconds returns a boolean if a field has been set.

### GetFontAssetId

`func (o *RenderSpec) GetFontAssetId() string`

GetFontAssetId returns the FontAssetId field if non-nil, zero value otherwise.

### GetFontAssetIdOk

`func (o *RenderSpec) GetFontAssetIdOk() (*string, bool)`

GetFontAssetIdOk returns a tuple with the FontAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontAssetId

`func (o *RenderSpec) SetFontAssetId(v string)`

SetFontAssetId sets FontAssetId field to given value.

### HasFontAssetId

`func (o *RenderSpec) HasFontAssetId() bool`

HasFontAssetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


