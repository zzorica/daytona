/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the GithubMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubMatch{}

// GithubMatch struct for GithubMatch
type GithubMatch struct {
	Indices []int32 `json:"indices,omitempty"`
	Text    *string `json:"text,omitempty"`
}

// NewGithubMatch instantiates a new GithubMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubMatch() *GithubMatch {
	this := GithubMatch{}
	return &this
}

// NewGithubMatchWithDefaults instantiates a new GithubMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubMatchWithDefaults() *GithubMatch {
	this := GithubMatch{}
	return &this
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *GithubMatch) GetIndices() []int32 {
	if o == nil || IsNil(o.Indices) {
		var ret []int32
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubMatch) GetIndicesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Indices) {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *GithubMatch) HasIndices() bool {
	if o != nil && !IsNil(o.Indices) {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []int32 and assigns it to the Indices field.
func (o *GithubMatch) SetIndices(v []int32) {
	o.Indices = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *GithubMatch) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubMatch) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *GithubMatch) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *GithubMatch) SetText(v string) {
	o.Text = &v
}

func (o GithubMatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Indices) {
		toSerialize["indices"] = o.Indices
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableGithubMatch struct {
	value *GithubMatch
	isSet bool
}

func (v NullableGithubMatch) Get() *GithubMatch {
	return v.value
}

func (v *NullableGithubMatch) Set(val *GithubMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubMatch(val *GithubMatch) *NullableGithubMatch {
	return &NullableGithubMatch{value: val, isSet: true}
}

func (v NullableGithubMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}