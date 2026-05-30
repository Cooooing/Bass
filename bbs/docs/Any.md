# Any

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**GoogleProtobufAny**](GoogleProtobufAny.md) |  | [optional] 
**Yaml** | Pointer to **string** |  | [optional] 

## Methods

### NewAny

`func NewAny() *Any`

NewAny instantiates a new Any object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnyWithDefaults

`func NewAnyWithDefaults() *Any`

NewAnyWithDefaults instantiates a new Any object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Any) GetValue() GoogleProtobufAny`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Any) GetValueOk() (*GoogleProtobufAny, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Any) SetValue(v GoogleProtobufAny)`

SetValue sets Value field to given value.

### HasValue

`func (o *Any) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetYaml

`func (o *Any) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Any) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Any) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *Any) HasYaml() bool`

HasYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


