# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalDocs** | Pointer to [**ExternalDocs**](ExternalDocs.md) |  | [optional] 
**SpecificationExtension** | Pointer to [**[]NamedAny**](NamedAny.md) |  | [optional] 

## Methods

### NewTag

`func NewTag() *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Tag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDocs

`func (o *Tag) GetExternalDocs() ExternalDocs`

GetExternalDocs returns the ExternalDocs field if non-nil, zero value otherwise.

### GetExternalDocsOk

`func (o *Tag) GetExternalDocsOk() (*ExternalDocs, bool)`

GetExternalDocsOk returns a tuple with the ExternalDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocs

`func (o *Tag) SetExternalDocs(v ExternalDocs)`

SetExternalDocs sets ExternalDocs field to given value.

### HasExternalDocs

`func (o *Tag) HasExternalDocs() bool`

HasExternalDocs returns a boolean if a field has been set.

### GetSpecificationExtension

`func (o *Tag) GetSpecificationExtension() []NamedAny`

GetSpecificationExtension returns the SpecificationExtension field if non-nil, zero value otherwise.

### GetSpecificationExtensionOk

`func (o *Tag) GetSpecificationExtensionOk() (*[]NamedAny, bool)`

GetSpecificationExtensionOk returns a tuple with the SpecificationExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationExtension

`func (o *Tag) SetSpecificationExtension(v []NamedAny)`

SetSpecificationExtension sets SpecificationExtension field to given value.

### HasSpecificationExtension

`func (o *Tag) HasSpecificationExtension() bool`

HasSpecificationExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


