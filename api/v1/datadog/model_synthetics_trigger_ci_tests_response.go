/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsTriggerCITestsResponse Object containing information about the tests triggered.
type SyntheticsTriggerCITestsResponse struct {
	// List of Synthetics locations.
	Locations *[]SyntheticsTriggerCITestLocation `json:"locations,omitempty"`
	// Information about the tests runs.
	Results *[]SyntheticsTriggerCITestRunResult `json:"results,omitempty"`
	// The public IDs of the Synthetics test triggered.
	TriggeredCheckIds *[]string `json:"triggered_check_ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsTriggerCITestsResponse instantiates a new SyntheticsTriggerCITestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTriggerCITestsResponse() *SyntheticsTriggerCITestsResponse {
	this := SyntheticsTriggerCITestsResponse{}
	return &this
}

// NewSyntheticsTriggerCITestsResponseWithDefaults instantiates a new SyntheticsTriggerCITestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTriggerCITestsResponseWithDefaults() *SyntheticsTriggerCITestsResponse {
	this := SyntheticsTriggerCITestsResponse{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsTriggerCITestsResponse) GetLocations() []SyntheticsTriggerCITestLocation {
	if o == nil || o.Locations == nil {
		var ret []SyntheticsTriggerCITestLocation
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerCITestsResponse) GetLocationsOk() (*[]SyntheticsTriggerCITestLocation, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsTriggerCITestsResponse) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []SyntheticsTriggerCITestLocation and assigns it to the Locations field.
func (o *SyntheticsTriggerCITestsResponse) SetLocations(v []SyntheticsTriggerCITestLocation) {
	o.Locations = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SyntheticsTriggerCITestsResponse) GetResults() []SyntheticsTriggerCITestRunResult {
	if o == nil || o.Results == nil {
		var ret []SyntheticsTriggerCITestRunResult
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerCITestsResponse) GetResultsOk() (*[]SyntheticsTriggerCITestRunResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SyntheticsTriggerCITestsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SyntheticsTriggerCITestRunResult and assigns it to the Results field.
func (o *SyntheticsTriggerCITestsResponse) SetResults(v []SyntheticsTriggerCITestRunResult) {
	o.Results = &v
}

// GetTriggeredCheckIds returns the TriggeredCheckIds field value if set, zero value otherwise.
func (o *SyntheticsTriggerCITestsResponse) GetTriggeredCheckIds() []string {
	if o == nil || o.TriggeredCheckIds == nil {
		var ret []string
		return ret
	}
	return *o.TriggeredCheckIds
}

// GetTriggeredCheckIdsOk returns a tuple with the TriggeredCheckIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerCITestsResponse) GetTriggeredCheckIdsOk() (*[]string, bool) {
	if o == nil || o.TriggeredCheckIds == nil {
		return nil, false
	}
	return o.TriggeredCheckIds, true
}

// HasTriggeredCheckIds returns a boolean if a field has been set.
func (o *SyntheticsTriggerCITestsResponse) HasTriggeredCheckIds() bool {
	if o != nil && o.TriggeredCheckIds != nil {
		return true
	}

	return false
}

// SetTriggeredCheckIds gets a reference to the given []string and assigns it to the TriggeredCheckIds field.
func (o *SyntheticsTriggerCITestsResponse) SetTriggeredCheckIds(v []string) {
	o.TriggeredCheckIds = &v
}

func (o SyntheticsTriggerCITestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.TriggeredCheckIds != nil {
		toSerialize["triggered_check_ids"] = o.TriggeredCheckIds
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsTriggerCITestsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Locations         *[]SyntheticsTriggerCITestLocation  `json:"locations,omitempty"`
		Results           *[]SyntheticsTriggerCITestRunResult `json:"results,omitempty"`
		TriggeredCheckIds *[]string                           `json:"triggered_check_ids,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Locations = all.Locations
	o.Results = all.Results
	o.TriggeredCheckIds = all.TriggeredCheckIds
	return nil
}

type NullableSyntheticsTriggerCITestsResponse struct {
	value *SyntheticsTriggerCITestsResponse
	isSet bool
}

func (v NullableSyntheticsTriggerCITestsResponse) Get() *SyntheticsTriggerCITestsResponse {
	return v.value
}

func (v *NullableSyntheticsTriggerCITestsResponse) Set(val *SyntheticsTriggerCITestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTriggerCITestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTriggerCITestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTriggerCITestsResponse(val *SyntheticsTriggerCITestsResponse) *NullableSyntheticsTriggerCITestsResponse {
	return &NullableSyntheticsTriggerCITestsResponse{value: val, isSet: true}
}

func (v NullableSyntheticsTriggerCITestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTriggerCITestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}