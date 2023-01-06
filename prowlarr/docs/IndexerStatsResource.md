# IndexerStatsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Indexers** | Pointer to [**[]IndexerStatistics**](IndexerStatistics.md) |  | [optional] 
**UserAgents** | Pointer to [**[]UserAgentStatistics**](UserAgentStatistics.md) |  | [optional] 
**Hosts** | Pointer to [**[]HostStatistics**](HostStatistics.md) |  | [optional] 

## Methods

### NewIndexerStatsResource

`func NewIndexerStatsResource() *IndexerStatsResource`

NewIndexerStatsResource instantiates a new IndexerStatsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerStatsResourceWithDefaults

`func NewIndexerStatsResourceWithDefaults() *IndexerStatsResource`

NewIndexerStatsResourceWithDefaults instantiates a new IndexerStatsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerStatsResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerStatsResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerStatsResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerStatsResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexers

`func (o *IndexerStatsResource) GetIndexers() []IndexerStatistics`

GetIndexers returns the Indexers field if non-nil, zero value otherwise.

### GetIndexersOk

`func (o *IndexerStatsResource) GetIndexersOk() (*[]IndexerStatistics, bool)`

GetIndexersOk returns a tuple with the Indexers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexers

`func (o *IndexerStatsResource) SetIndexers(v []IndexerStatistics)`

SetIndexers sets Indexers field to given value.

### HasIndexers

`func (o *IndexerStatsResource) HasIndexers() bool`

HasIndexers returns a boolean if a field has been set.

### SetIndexersNil

`func (o *IndexerStatsResource) SetIndexersNil(b bool)`

 SetIndexersNil sets the value for Indexers to be an explicit nil

### UnsetIndexers
`func (o *IndexerStatsResource) UnsetIndexers()`

UnsetIndexers ensures that no value is present for Indexers, not even an explicit nil
### GetUserAgents

`func (o *IndexerStatsResource) GetUserAgents() []UserAgentStatistics`

GetUserAgents returns the UserAgents field if non-nil, zero value otherwise.

### GetUserAgentsOk

`func (o *IndexerStatsResource) GetUserAgentsOk() (*[]UserAgentStatistics, bool)`

GetUserAgentsOk returns a tuple with the UserAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgents

`func (o *IndexerStatsResource) SetUserAgents(v []UserAgentStatistics)`

SetUserAgents sets UserAgents field to given value.

### HasUserAgents

`func (o *IndexerStatsResource) HasUserAgents() bool`

HasUserAgents returns a boolean if a field has been set.

### SetUserAgentsNil

`func (o *IndexerStatsResource) SetUserAgentsNil(b bool)`

 SetUserAgentsNil sets the value for UserAgents to be an explicit nil

### UnsetUserAgents
`func (o *IndexerStatsResource) UnsetUserAgents()`

UnsetUserAgents ensures that no value is present for UserAgents, not even an explicit nil
### GetHosts

`func (o *IndexerStatsResource) GetHosts() []HostStatistics`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *IndexerStatsResource) GetHostsOk() (*[]HostStatistics, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *IndexerStatsResource) SetHosts(v []HostStatistics)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *IndexerStatsResource) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *IndexerStatsResource) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *IndexerStatsResource) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


