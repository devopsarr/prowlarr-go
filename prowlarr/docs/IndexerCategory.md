# IndexerCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SubCategories** | Pointer to [**[]IndexerCategory**](IndexerCategory.md) |  | [optional] [readonly] 

## Methods

### NewIndexerCategory

`func NewIndexerCategory() *IndexerCategory`

NewIndexerCategory instantiates a new IndexerCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerCategoryWithDefaults

`func NewIndexerCategoryWithDefaults() *IndexerCategory`

NewIndexerCategoryWithDefaults instantiates a new IndexerCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IndexerCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexerCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexerCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexerCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IndexerCategory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IndexerCategory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *IndexerCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndexerCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndexerCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndexerCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IndexerCategory) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IndexerCategory) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubCategories

`func (o *IndexerCategory) GetSubCategories() []IndexerCategory`

GetSubCategories returns the SubCategories field if non-nil, zero value otherwise.

### GetSubCategoriesOk

`func (o *IndexerCategory) GetSubCategoriesOk() (*[]IndexerCategory, bool)`

GetSubCategoriesOk returns a tuple with the SubCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategories

`func (o *IndexerCategory) SetSubCategories(v []IndexerCategory)`

SetSubCategories sets SubCategories field to given value.

### HasSubCategories

`func (o *IndexerCategory) HasSubCategories() bool`

HasSubCategories returns a boolean if a field has been set.

### SetSubCategoriesNil

`func (o *IndexerCategory) SetSubCategoriesNil(b bool)`

 SetSubCategoriesNil sets the value for SubCategories to be an explicit nil

### UnsetSubCategories
`func (o *IndexerCategory) UnsetSubCategories()`

UnsetSubCategories ensures that no value is present for SubCategories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


