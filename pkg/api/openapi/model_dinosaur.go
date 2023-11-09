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

// checks if the Dinosaur type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dinosaur{}

// Dinosaur struct for Dinosaur
type Dinosaur struct {
	Id        *string    `json:"id,omitempty"`
	Kind      *string    `json:"kind,omitempty"`
	Href      *string    `json:"href,omitempty"`
	Species   *string    `json:"species,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDinosaur instantiates a new Dinosaur object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDinosaur() *Dinosaur {
	this := Dinosaur{}
	return &this
}

// NewDinosaurWithDefaults instantiates a new Dinosaur object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDinosaurWithDefaults() *Dinosaur {
	this := Dinosaur{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dinosaur) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dinosaur) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dinosaur) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Dinosaur) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Dinosaur) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Dinosaur) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Dinosaur) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Dinosaur) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Dinosaur) SetHref(v string) {
	o.Href = &v
}

// GetSpecies returns the Species field value if set, zero value otherwise.
func (o *Dinosaur) GetSpecies() string {
	if o == nil || IsNil(o.Species) {
		var ret string
		return ret
	}
	return *o.Species
}

// GetSpeciesOk returns a tuple with the Species field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetSpeciesOk() (*string, bool) {
	if o == nil || IsNil(o.Species) {
		return nil, false
	}
	return o.Species, true
}

// HasSpecies returns a boolean if a field has been set.
func (o *Dinosaur) HasSpecies() bool {
	if o != nil && !IsNil(o.Species) {
		return true
	}

	return false
}

// SetSpecies gets a reference to the given string and assigns it to the Species field.
func (o *Dinosaur) SetSpecies(v string) {
	o.Species = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Dinosaur) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Dinosaur) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Dinosaur) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Dinosaur) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dinosaur) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Dinosaur) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Dinosaur) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Dinosaur) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dinosaur) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
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

type NullableDinosaur struct {
	value *Dinosaur
	isSet bool
}

func (v NullableDinosaur) Get() *Dinosaur {
	return v.value
}

func (v *NullableDinosaur) Set(val *Dinosaur) {
	v.value = val
	v.isSet = true
}

func (v NullableDinosaur) IsSet() bool {
	return v.isSet
}

func (v *NullableDinosaur) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDinosaur(val *Dinosaur) *NullableDinosaur {
	return &NullableDinosaur{value: val, isSet: true}
}

func (v NullableDinosaur) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDinosaur) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
