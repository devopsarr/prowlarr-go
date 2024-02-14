/*
Prowlarr

Prowlarr API docs

API version: v1.13.3.4273
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// Field struct for Field
type Field struct {
	Order *int32 `json:"order,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Label NullableString `json:"label,omitempty"`
	Unit NullableString `json:"unit,omitempty"`
	HelpText NullableString `json:"helpText,omitempty"`
	HelpTextWarning NullableString `json:"helpTextWarning,omitempty"`
	HelpLink NullableString `json:"helpLink,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Advanced *bool `json:"advanced,omitempty"`
	SelectOptions []*SelectOption `json:"selectOptions,omitempty"`
	SelectOptionsProviderAction NullableString `json:"selectOptionsProviderAction,omitempty"`
	Section NullableString `json:"section,omitempty"`
	Hidden NullableString `json:"hidden,omitempty"`
	Privacy *PrivacyLevel `json:"privacy,omitempty"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	IsFloat *bool `json:"isFloat,omitempty"`
}

// NewField instantiates a new Field object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewField() *Field {
	this := Field{}
	return &this
}

// NewFieldWithDefaults instantiates a new Field object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldWithDefaults() *Field {
	this := Field{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Field) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Field) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Field) SetOrder(v int32) {
	o.Order = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Field) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Field) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Field) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Field) UnsetName() {
	o.Name.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *Field) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *Field) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *Field) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *Field) UnsetLabel() {
	o.Label.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetUnit() string {
	if o == nil || isNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetUnitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *Field) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *Field) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *Field) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *Field) UnsetUnit() {
	o.Unit.Unset()
}

// GetHelpText returns the HelpText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetHelpText() string {
	if o == nil || isNil(o.HelpText.Get()) {
		var ret string
		return ret
	}
	return *o.HelpText.Get()
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetHelpTextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HelpText.Get(), o.HelpText.IsSet()
}

// HasHelpText returns a boolean if a field has been set.
func (o *Field) HasHelpText() bool {
	if o != nil && o.HelpText.IsSet() {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given NullableString and assigns it to the HelpText field.
func (o *Field) SetHelpText(v string) {
	o.HelpText.Set(&v)
}
// SetHelpTextNil sets the value for HelpText to be an explicit nil
func (o *Field) SetHelpTextNil() {
	o.HelpText.Set(nil)
}

// UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
func (o *Field) UnsetHelpText() {
	o.HelpText.Unset()
}

// GetHelpTextWarning returns the HelpTextWarning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetHelpTextWarning() string {
	if o == nil || isNil(o.HelpTextWarning.Get()) {
		var ret string
		return ret
	}
	return *o.HelpTextWarning.Get()
}

// GetHelpTextWarningOk returns a tuple with the HelpTextWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetHelpTextWarningOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HelpTextWarning.Get(), o.HelpTextWarning.IsSet()
}

// HasHelpTextWarning returns a boolean if a field has been set.
func (o *Field) HasHelpTextWarning() bool {
	if o != nil && o.HelpTextWarning.IsSet() {
		return true
	}

	return false
}

// SetHelpTextWarning gets a reference to the given NullableString and assigns it to the HelpTextWarning field.
func (o *Field) SetHelpTextWarning(v string) {
	o.HelpTextWarning.Set(&v)
}
// SetHelpTextWarningNil sets the value for HelpTextWarning to be an explicit nil
func (o *Field) SetHelpTextWarningNil() {
	o.HelpTextWarning.Set(nil)
}

// UnsetHelpTextWarning ensures that no value is present for HelpTextWarning, not even an explicit nil
func (o *Field) UnsetHelpTextWarning() {
	o.HelpTextWarning.Unset()
}

// GetHelpLink returns the HelpLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetHelpLink() string {
	if o == nil || isNil(o.HelpLink.Get()) {
		var ret string
		return ret
	}
	return *o.HelpLink.Get()
}

// GetHelpLinkOk returns a tuple with the HelpLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetHelpLinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.HelpLink.Get(), o.HelpLink.IsSet()
}

// HasHelpLink returns a boolean if a field has been set.
func (o *Field) HasHelpLink() bool {
	if o != nil && o.HelpLink.IsSet() {
		return true
	}

	return false
}

// SetHelpLink gets a reference to the given NullableString and assigns it to the HelpLink field.
func (o *Field) SetHelpLink(v string) {
	o.HelpLink.Set(&v)
}
// SetHelpLinkNil sets the value for HelpLink to be an explicit nil
func (o *Field) SetHelpLinkNil() {
	o.HelpLink.Set(nil)
}

// UnsetHelpLink ensures that no value is present for HelpLink, not even an explicit nil
func (o *Field) UnsetHelpLink() {
	o.HelpLink.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetValueOk() (*interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Field) HasValue() bool {
	if o != nil && isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *Field) SetValue(v interface{}) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetType() string {
	if o == nil || isNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Field) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Field) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Field) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Field) UnsetType() {
	o.Type.Unset()
}

// GetAdvanced returns the Advanced field value if set, zero value otherwise.
func (o *Field) GetAdvanced() bool {
	if o == nil || isNil(o.Advanced) {
		var ret bool
		return ret
	}
	return *o.Advanced
}

// GetAdvancedOk returns a tuple with the Advanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetAdvancedOk() (*bool, bool) {
	if o == nil || isNil(o.Advanced) {
    return nil, false
	}
	return o.Advanced, true
}

// HasAdvanced returns a boolean if a field has been set.
func (o *Field) HasAdvanced() bool {
	if o != nil && !isNil(o.Advanced) {
		return true
	}

	return false
}

// SetAdvanced gets a reference to the given bool and assigns it to the Advanced field.
func (o *Field) SetAdvanced(v bool) {
	o.Advanced = &v
}

// GetSelectOptions returns the SelectOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetSelectOptions() []*SelectOption {
	if o == nil {
		var ret []*SelectOption
		return ret
	}
	return o.SelectOptions
}

// GetSelectOptionsOk returns a tuple with the SelectOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetSelectOptionsOk() ([]*SelectOption, bool) {
	if o == nil || isNil(o.SelectOptions) {
    return nil, false
	}
	return o.SelectOptions, true
}

// HasSelectOptions returns a boolean if a field has been set.
func (o *Field) HasSelectOptions() bool {
	if o != nil && isNil(o.SelectOptions) {
		return true
	}

	return false
}

// SetSelectOptions gets a reference to the given []SelectOption and assigns it to the SelectOptions field.
func (o *Field) SetSelectOptions(v []*SelectOption) {
	o.SelectOptions = v
}

// GetSelectOptionsProviderAction returns the SelectOptionsProviderAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetSelectOptionsProviderAction() string {
	if o == nil || isNil(o.SelectOptionsProviderAction.Get()) {
		var ret string
		return ret
	}
	return *o.SelectOptionsProviderAction.Get()
}

// GetSelectOptionsProviderActionOk returns a tuple with the SelectOptionsProviderAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetSelectOptionsProviderActionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SelectOptionsProviderAction.Get(), o.SelectOptionsProviderAction.IsSet()
}

// HasSelectOptionsProviderAction returns a boolean if a field has been set.
func (o *Field) HasSelectOptionsProviderAction() bool {
	if o != nil && o.SelectOptionsProviderAction.IsSet() {
		return true
	}

	return false
}

// SetSelectOptionsProviderAction gets a reference to the given NullableString and assigns it to the SelectOptionsProviderAction field.
func (o *Field) SetSelectOptionsProviderAction(v string) {
	o.SelectOptionsProviderAction.Set(&v)
}
// SetSelectOptionsProviderActionNil sets the value for SelectOptionsProviderAction to be an explicit nil
func (o *Field) SetSelectOptionsProviderActionNil() {
	o.SelectOptionsProviderAction.Set(nil)
}

// UnsetSelectOptionsProviderAction ensures that no value is present for SelectOptionsProviderAction, not even an explicit nil
func (o *Field) UnsetSelectOptionsProviderAction() {
	o.SelectOptionsProviderAction.Unset()
}

// GetSection returns the Section field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetSection() string {
	if o == nil || isNil(o.Section.Get()) {
		var ret string
		return ret
	}
	return *o.Section.Get()
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetSectionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Section.Get(), o.Section.IsSet()
}

// HasSection returns a boolean if a field has been set.
func (o *Field) HasSection() bool {
	if o != nil && o.Section.IsSet() {
		return true
	}

	return false
}

// SetSection gets a reference to the given NullableString and assigns it to the Section field.
func (o *Field) SetSection(v string) {
	o.Section.Set(&v)
}
// SetSectionNil sets the value for Section to be an explicit nil
func (o *Field) SetSectionNil() {
	o.Section.Set(nil)
}

// UnsetSection ensures that no value is present for Section, not even an explicit nil
func (o *Field) UnsetSection() {
	o.Section.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetHidden() string {
	if o == nil || isNil(o.Hidden.Get()) {
		var ret string
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetHiddenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *Field) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableString and assigns it to the Hidden field.
func (o *Field) SetHidden(v string) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *Field) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *Field) UnsetHidden() {
	o.Hidden.Unset()
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *Field) GetPrivacy() PrivacyLevel {
	if o == nil || isNil(o.Privacy) {
		var ret PrivacyLevel
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetPrivacyOk() (*PrivacyLevel, bool) {
	if o == nil || isNil(o.Privacy) {
    return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *Field) HasPrivacy() bool {
	if o != nil && !isNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given PrivacyLevel and assigns it to the Privacy field.
func (o *Field) SetPrivacy(v PrivacyLevel) {
	o.Privacy = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetPlaceholder() string {
	if o == nil || isNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetPlaceholderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *Field) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *Field) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *Field) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *Field) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetIsFloat returns the IsFloat field value if set, zero value otherwise.
func (o *Field) GetIsFloat() bool {
	if o == nil || isNil(o.IsFloat) {
		var ret bool
		return ret
	}
	return *o.IsFloat
}

// GetIsFloatOk returns a tuple with the IsFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetIsFloatOk() (*bool, bool) {
	if o == nil || isNil(o.IsFloat) {
    return nil, false
	}
	return o.IsFloat, true
}

// HasIsFloat returns a boolean if a field has been set.
func (o *Field) HasIsFloat() bool {
	if o != nil && !isNil(o.IsFloat) {
		return true
	}

	return false
}

// SetIsFloat gets a reference to the given bool and assigns it to the IsFloat field.
func (o *Field) SetIsFloat(v bool) {
	o.IsFloat = &v
}

func (o Field) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.HelpText.IsSet() {
		toSerialize["helpText"] = o.HelpText.Get()
	}
	if o.HelpTextWarning.IsSet() {
		toSerialize["helpTextWarning"] = o.HelpTextWarning.Get()
	}
	if o.HelpLink.IsSet() {
		toSerialize["helpLink"] = o.HelpLink.Get()
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !isNil(o.Advanced) {
		toSerialize["advanced"] = o.Advanced
	}
	if o.SelectOptions != nil {
		toSerialize["selectOptions"] = o.SelectOptions
	}
	if o.SelectOptionsProviderAction.IsSet() {
		toSerialize["selectOptionsProviderAction"] = o.SelectOptionsProviderAction.Get()
	}
	if o.Section.IsSet() {
		toSerialize["section"] = o.Section.Get()
	}
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	if !isNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if !isNil(o.IsFloat) {
		toSerialize["isFloat"] = o.IsFloat
	}
	return json.Marshal(toSerialize)
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


