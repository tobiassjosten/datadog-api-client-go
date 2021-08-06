/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// IncidentTeamCreateRequest Create request with an incident team payload.
type IncidentTeamCreateRequest struct {
	Data IncidentTeamCreateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentTeamCreateRequest instantiates a new IncidentTeamCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTeamCreateRequest(data IncidentTeamCreateData) *IncidentTeamCreateRequest {
	this := IncidentTeamCreateRequest{}
	this.Data = data
	return &this
}

// NewIncidentTeamCreateRequestWithDefaults instantiates a new IncidentTeamCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTeamCreateRequestWithDefaults() *IncidentTeamCreateRequest {
	this := IncidentTeamCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentTeamCreateRequest) GetData() IncidentTeamCreateData {
	if o == nil {
		var ret IncidentTeamCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentTeamCreateRequest) GetDataOk() (*IncidentTeamCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentTeamCreateRequest) SetData(v IncidentTeamCreateData) {
	o.Data = v
}

func (o IncidentTeamCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentTeamCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *IncidentTeamCreateData `json:"data"`
	}{}
	all := struct {
		Data IncidentTeamCreateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Data = all.Data
	return nil
}

type NullableIncidentTeamCreateRequest struct {
	value *IncidentTeamCreateRequest
	isSet bool
}

func (v NullableIncidentTeamCreateRequest) Get() *IncidentTeamCreateRequest {
	return v.value
}

func (v *NullableIncidentTeamCreateRequest) Set(val *IncidentTeamCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTeamCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTeamCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTeamCreateRequest(val *IncidentTeamCreateRequest) *NullableIncidentTeamCreateRequest {
	return &NullableIncidentTeamCreateRequest{value: val, isSet: true}
}

func (v NullableIncidentTeamCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTeamCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}