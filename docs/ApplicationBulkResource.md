# ApplicationBulkResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**SyncLevel** | Pointer to [**ApplicationSyncLevel**](ApplicationSyncLevel.md) |  | [optional] 

## Methods

### NewApplicationBulkResource

`func NewApplicationBulkResource() *ApplicationBulkResource`

NewApplicationBulkResource instantiates a new ApplicationBulkResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBulkResourceWithDefaults

`func NewApplicationBulkResourceWithDefaults() *ApplicationBulkResource`

NewApplicationBulkResourceWithDefaults instantiates a new ApplicationBulkResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ApplicationBulkResource) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ApplicationBulkResource) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ApplicationBulkResource) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ApplicationBulkResource) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *ApplicationBulkResource) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *ApplicationBulkResource) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetTags

`func (o *ApplicationBulkResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationBulkResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationBulkResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationBulkResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationBulkResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationBulkResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *ApplicationBulkResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *ApplicationBulkResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *ApplicationBulkResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *ApplicationBulkResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetSyncLevel

`func (o *ApplicationBulkResource) GetSyncLevel() ApplicationSyncLevel`

GetSyncLevel returns the SyncLevel field if non-nil, zero value otherwise.

### GetSyncLevelOk

`func (o *ApplicationBulkResource) GetSyncLevelOk() (*ApplicationSyncLevel, bool)`

GetSyncLevelOk returns a tuple with the SyncLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncLevel

`func (o *ApplicationBulkResource) SetSyncLevel(v ApplicationSyncLevel)`

SetSyncLevel sets SyncLevel field to given value.

### HasSyncLevel

`func (o *ApplicationBulkResource) HasSyncLevel() bool`

HasSyncLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


