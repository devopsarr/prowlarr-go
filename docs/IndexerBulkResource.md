# IndexerBulkResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**Enable** | Pointer to **NullableBool** |  | [optional] 
**AppProfileId** | Pointer to **NullableInt32** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**MinimumSeeders** | Pointer to **NullableInt32** |  | [optional] 
**SeedRatio** | Pointer to **NullableFloat64** |  | [optional] 
**SeedTime** | Pointer to **NullableInt32** |  | [optional] 
**PackSeedTime** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIndexerBulkResource

`func NewIndexerBulkResource() *IndexerBulkResource`

NewIndexerBulkResource instantiates a new IndexerBulkResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerBulkResourceWithDefaults

`func NewIndexerBulkResourceWithDefaults() *IndexerBulkResource`

NewIndexerBulkResourceWithDefaults instantiates a new IndexerBulkResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *IndexerBulkResource) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IndexerBulkResource) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IndexerBulkResource) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *IndexerBulkResource) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *IndexerBulkResource) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *IndexerBulkResource) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetTags

`func (o *IndexerBulkResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndexerBulkResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndexerBulkResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IndexerBulkResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IndexerBulkResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IndexerBulkResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *IndexerBulkResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *IndexerBulkResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *IndexerBulkResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *IndexerBulkResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetEnable

`func (o *IndexerBulkResource) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IndexerBulkResource) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IndexerBulkResource) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IndexerBulkResource) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *IndexerBulkResource) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *IndexerBulkResource) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetAppProfileId

`func (o *IndexerBulkResource) GetAppProfileId() int32`

GetAppProfileId returns the AppProfileId field if non-nil, zero value otherwise.

### GetAppProfileIdOk

`func (o *IndexerBulkResource) GetAppProfileIdOk() (*int32, bool)`

GetAppProfileIdOk returns a tuple with the AppProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileId

`func (o *IndexerBulkResource) SetAppProfileId(v int32)`

SetAppProfileId sets AppProfileId field to given value.

### HasAppProfileId

`func (o *IndexerBulkResource) HasAppProfileId() bool`

HasAppProfileId returns a boolean if a field has been set.

### SetAppProfileIdNil

`func (o *IndexerBulkResource) SetAppProfileIdNil(b bool)`

 SetAppProfileIdNil sets the value for AppProfileId to be an explicit nil

### UnsetAppProfileId
`func (o *IndexerBulkResource) UnsetAppProfileId()`

UnsetAppProfileId ensures that no value is present for AppProfileId, not even an explicit nil
### GetPriority

`func (o *IndexerBulkResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IndexerBulkResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IndexerBulkResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IndexerBulkResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *IndexerBulkResource) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *IndexerBulkResource) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetMinimumSeeders

`func (o *IndexerBulkResource) GetMinimumSeeders() int32`

GetMinimumSeeders returns the MinimumSeeders field if non-nil, zero value otherwise.

### GetMinimumSeedersOk

`func (o *IndexerBulkResource) GetMinimumSeedersOk() (*int32, bool)`

GetMinimumSeedersOk returns a tuple with the MinimumSeeders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSeeders

`func (o *IndexerBulkResource) SetMinimumSeeders(v int32)`

SetMinimumSeeders sets MinimumSeeders field to given value.

### HasMinimumSeeders

`func (o *IndexerBulkResource) HasMinimumSeeders() bool`

HasMinimumSeeders returns a boolean if a field has been set.

### SetMinimumSeedersNil

`func (o *IndexerBulkResource) SetMinimumSeedersNil(b bool)`

 SetMinimumSeedersNil sets the value for MinimumSeeders to be an explicit nil

### UnsetMinimumSeeders
`func (o *IndexerBulkResource) UnsetMinimumSeeders()`

UnsetMinimumSeeders ensures that no value is present for MinimumSeeders, not even an explicit nil
### GetSeedRatio

`func (o *IndexerBulkResource) GetSeedRatio() float64`

GetSeedRatio returns the SeedRatio field if non-nil, zero value otherwise.

### GetSeedRatioOk

`func (o *IndexerBulkResource) GetSeedRatioOk() (*float64, bool)`

GetSeedRatioOk returns a tuple with the SeedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRatio

`func (o *IndexerBulkResource) SetSeedRatio(v float64)`

SetSeedRatio sets SeedRatio field to given value.

### HasSeedRatio

`func (o *IndexerBulkResource) HasSeedRatio() bool`

HasSeedRatio returns a boolean if a field has been set.

### SetSeedRatioNil

`func (o *IndexerBulkResource) SetSeedRatioNil(b bool)`

 SetSeedRatioNil sets the value for SeedRatio to be an explicit nil

### UnsetSeedRatio
`func (o *IndexerBulkResource) UnsetSeedRatio()`

UnsetSeedRatio ensures that no value is present for SeedRatio, not even an explicit nil
### GetSeedTime

`func (o *IndexerBulkResource) GetSeedTime() int32`

GetSeedTime returns the SeedTime field if non-nil, zero value otherwise.

### GetSeedTimeOk

`func (o *IndexerBulkResource) GetSeedTimeOk() (*int32, bool)`

GetSeedTimeOk returns a tuple with the SeedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedTime

`func (o *IndexerBulkResource) SetSeedTime(v int32)`

SetSeedTime sets SeedTime field to given value.

### HasSeedTime

`func (o *IndexerBulkResource) HasSeedTime() bool`

HasSeedTime returns a boolean if a field has been set.

### SetSeedTimeNil

`func (o *IndexerBulkResource) SetSeedTimeNil(b bool)`

 SetSeedTimeNil sets the value for SeedTime to be an explicit nil

### UnsetSeedTime
`func (o *IndexerBulkResource) UnsetSeedTime()`

UnsetSeedTime ensures that no value is present for SeedTime, not even an explicit nil
### GetPackSeedTime

`func (o *IndexerBulkResource) GetPackSeedTime() int32`

GetPackSeedTime returns the PackSeedTime field if non-nil, zero value otherwise.

### GetPackSeedTimeOk

`func (o *IndexerBulkResource) GetPackSeedTimeOk() (*int32, bool)`

GetPackSeedTimeOk returns a tuple with the PackSeedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackSeedTime

`func (o *IndexerBulkResource) SetPackSeedTime(v int32)`

SetPackSeedTime sets PackSeedTime field to given value.

### HasPackSeedTime

`func (o *IndexerBulkResource) HasPackSeedTime() bool`

HasPackSeedTime returns a boolean if a field has been set.

### SetPackSeedTimeNil

`func (o *IndexerBulkResource) SetPackSeedTimeNil(b bool)`

 SetPackSeedTimeNil sets the value for PackSeedTime to be an explicit nil

### UnsetPackSeedTime
`func (o *IndexerBulkResource) UnsetPackSeedTime()`

UnsetPackSeedTime ensures that no value is present for PackSeedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


