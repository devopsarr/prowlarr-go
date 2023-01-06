# IndexerProxyResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ConfigContract** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to [**ProviderMessage**](ProviderMessage.md) |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presets** | Pointer to [**[]IndexerProxyResource**](IndexerProxyResource.md) |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**OnHealthIssue** | Pointer to **bool** |  | [optional] 
**SupportsOnHealthIssue** | Pointer to **bool** |  | [optional] 
**IncludeHealthWarnings** | Pointer to **bool** |  | [optional] 
**TestCommand** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIndexerProxyResource

`func NewIndexerProxyResource() *IndexerProxyResource`

NewIndexerProxyResource instantiates a new IndexerProxyResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerProxyResourceWithDefaults

`func NewIndexerProxyResourceWithDefaults() *IndexerProxyResource`

NewIndexerProxyResourceWithDefaults instantiates a new IndexerProxyResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerProxyResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerProxyResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerProxyResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerProxyResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IndexerProxyResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexerProxyResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexerProxyResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexerProxyResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IndexerProxyResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IndexerProxyResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *IndexerProxyResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexerProxyResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexerProxyResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IndexerProxyResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *IndexerProxyResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *IndexerProxyResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *IndexerProxyResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *IndexerProxyResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *IndexerProxyResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *IndexerProxyResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *IndexerProxyResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *IndexerProxyResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *IndexerProxyResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *IndexerProxyResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *IndexerProxyResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *IndexerProxyResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *IndexerProxyResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *IndexerProxyResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *IndexerProxyResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *IndexerProxyResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *IndexerProxyResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *IndexerProxyResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *IndexerProxyResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *IndexerProxyResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *IndexerProxyResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *IndexerProxyResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *IndexerProxyResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *IndexerProxyResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *IndexerProxyResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *IndexerProxyResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *IndexerProxyResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IndexerProxyResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IndexerProxyResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IndexerProxyResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *IndexerProxyResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndexerProxyResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndexerProxyResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IndexerProxyResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IndexerProxyResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IndexerProxyResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *IndexerProxyResource) GetPresets() []IndexerProxyResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *IndexerProxyResource) GetPresetsOk() (*[]IndexerProxyResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *IndexerProxyResource) SetPresets(v []IndexerProxyResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *IndexerProxyResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *IndexerProxyResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *IndexerProxyResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetLink

`func (o *IndexerProxyResource) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IndexerProxyResource) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IndexerProxyResource) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *IndexerProxyResource) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *IndexerProxyResource) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *IndexerProxyResource) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetOnHealthIssue

`func (o *IndexerProxyResource) GetOnHealthIssue() bool`

GetOnHealthIssue returns the OnHealthIssue field if non-nil, zero value otherwise.

### GetOnHealthIssueOk

`func (o *IndexerProxyResource) GetOnHealthIssueOk() (*bool, bool)`

GetOnHealthIssueOk returns a tuple with the OnHealthIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHealthIssue

`func (o *IndexerProxyResource) SetOnHealthIssue(v bool)`

SetOnHealthIssue sets OnHealthIssue field to given value.

### HasOnHealthIssue

`func (o *IndexerProxyResource) HasOnHealthIssue() bool`

HasOnHealthIssue returns a boolean if a field has been set.

### GetSupportsOnHealthIssue

`func (o *IndexerProxyResource) GetSupportsOnHealthIssue() bool`

GetSupportsOnHealthIssue returns the SupportsOnHealthIssue field if non-nil, zero value otherwise.

### GetSupportsOnHealthIssueOk

`func (o *IndexerProxyResource) GetSupportsOnHealthIssueOk() (*bool, bool)`

GetSupportsOnHealthIssueOk returns a tuple with the SupportsOnHealthIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOnHealthIssue

`func (o *IndexerProxyResource) SetSupportsOnHealthIssue(v bool)`

SetSupportsOnHealthIssue sets SupportsOnHealthIssue field to given value.

### HasSupportsOnHealthIssue

`func (o *IndexerProxyResource) HasSupportsOnHealthIssue() bool`

HasSupportsOnHealthIssue returns a boolean if a field has been set.

### GetIncludeHealthWarnings

`func (o *IndexerProxyResource) GetIncludeHealthWarnings() bool`

GetIncludeHealthWarnings returns the IncludeHealthWarnings field if non-nil, zero value otherwise.

### GetIncludeHealthWarningsOk

`func (o *IndexerProxyResource) GetIncludeHealthWarningsOk() (*bool, bool)`

GetIncludeHealthWarningsOk returns a tuple with the IncludeHealthWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHealthWarnings

`func (o *IndexerProxyResource) SetIncludeHealthWarnings(v bool)`

SetIncludeHealthWarnings sets IncludeHealthWarnings field to given value.

### HasIncludeHealthWarnings

`func (o *IndexerProxyResource) HasIncludeHealthWarnings() bool`

HasIncludeHealthWarnings returns a boolean if a field has been set.

### GetTestCommand

`func (o *IndexerProxyResource) GetTestCommand() string`

GetTestCommand returns the TestCommand field if non-nil, zero value otherwise.

### GetTestCommandOk

`func (o *IndexerProxyResource) GetTestCommandOk() (*string, bool)`

GetTestCommandOk returns a tuple with the TestCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCommand

`func (o *IndexerProxyResource) SetTestCommand(v string)`

SetTestCommand sets TestCommand field to given value.

### HasTestCommand

`func (o *IndexerProxyResource) HasTestCommand() bool`

HasTestCommand returns a boolean if a field has been set.

### SetTestCommandNil

`func (o *IndexerProxyResource) SetTestCommandNil(b bool)`

 SetTestCommandNil sets the value for TestCommand to be an explicit nil

### UnsetTestCommand
`func (o *IndexerProxyResource) UnsetTestCommand()`

UnsetTestCommand ensures that no value is present for TestCommand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


