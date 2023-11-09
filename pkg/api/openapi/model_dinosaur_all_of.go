/*
maestro API

maestro API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the DinosaurAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DinosaurAllOf{}

// DinosaurAllOf struct for DinosaurAllOf
type DinosaurAllOf struct {
	Species   *string    `json:"species,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDinosaurAllOf instantiates a new DinosaurAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDinosaurAllOf() *DinosaurAllOf {
	this := DinosaurAllOf{}
	return &this
}

// NewDinosaurAllOfWithDefaults instantiates a new DinosaurAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDinosaurAllOfWithDefaults() *DinosaurAllOf {
	this := DinosaurAllOf{}
	return &this
}

// GetSpecies returns the Species field value if set, zero value otherwise.
func (o *DinosaurAllOf) GetSpecies() string {
	if o == nil || IsNil(o.Species) {
		var ret string
		return ret
	}
	return *o.Species
}

// GetSpeciesOk returns a tuple with the Species field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinosaurAllOf) GetSpeciesOk() (*string, bool) {
	if o == nil || IsNil(o.Species) {
		return nil, false
	}
	return o.Species, true
}

// HasSpecies returns a boolean if a field has been set.
func (o *DinosaurAllOf) HasSpecies() bool {
	if o != nil && !IsNil(o.Species) {
		return true
	}

	return false
}

// SetSpecies gets a reference to the given string and assigns it to the Species field.
func (o *DinosaurAllOf) SetSpecies(v string) {
	o.Species = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DinosaurAllOf) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinosaurAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DinosaurAllOf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DinosaurAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DinosaurAllOf) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DinosaurAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DinosaurAllOf) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DinosaurAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DinosaurAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DinosaurAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Species) {
		toSerialize["species"] = o.Species
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDinosaurAllOf struct {
	value *DinosaurAllOf
	isSet bool
}

func (v NullableDinosaurAllOf) Get() *DinosaurAllOf {
	return v.value
}

func (v *NullableDinosaurAllOf) Set(val *DinosaurAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDinosaurAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDinosaurAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDinosaurAllOf(val *DinosaurAllOf) *NullableDinosaurAllOf {
	return &NullableDinosaurAllOf{value: val, isSet: true}
}

func (v NullableDinosaurAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDinosaurAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
