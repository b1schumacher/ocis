/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the SharingInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharingInvitation{}

// SharingInvitation invitation-related data items
type SharingInvitation struct {
	InvitedBy *IdentitySet `json:"invitedBy,omitempty"`
}

// NewSharingInvitation instantiates a new SharingInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingInvitation() *SharingInvitation {
	this := SharingInvitation{}
	return &this
}

// NewSharingInvitationWithDefaults instantiates a new SharingInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingInvitationWithDefaults() *SharingInvitation {
	this := SharingInvitation{}
	return &this
}

// GetInvitedBy returns the InvitedBy field value if set, zero value otherwise.
func (o *SharingInvitation) GetInvitedBy() IdentitySet {
	if o == nil || IsNil(o.InvitedBy) {
		var ret IdentitySet
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingInvitation) GetInvitedByOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.InvitedBy) {
		return nil, false
	}
	return o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *SharingInvitation) HasInvitedBy() bool {
	if o != nil && !IsNil(o.InvitedBy) {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given IdentitySet and assigns it to the InvitedBy field.
func (o *SharingInvitation) SetInvitedBy(v IdentitySet) {
	o.InvitedBy = &v
}

func (o SharingInvitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharingInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvitedBy) {
		toSerialize["invitedBy"] = o.InvitedBy
	}
	return toSerialize, nil
}

type NullableSharingInvitation struct {
	value *SharingInvitation
	isSet bool
}

func (v NullableSharingInvitation) Get() *SharingInvitation {
	return v.value
}

func (v *NullableSharingInvitation) Set(val *SharingInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingInvitation(val *SharingInvitation) *NullableSharingInvitation {
	return &NullableSharingInvitation{value: val, isSet: true}
}

func (v NullableSharingInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}