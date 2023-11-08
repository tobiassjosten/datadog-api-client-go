// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthenticationValidationResponse Represent validation endpoint responses.
type AuthenticationValidationResponse struct {
	// Return `true` if the authentication response is valid.
	Valid *bool `json:"valid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAuthenticationValidationResponse instantiates a new AuthenticationValidationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthenticationValidationResponse() *AuthenticationValidationResponse {
	this := AuthenticationValidationResponse{}
	return &this
}

// NewAuthenticationValidationResponseWithDefaults instantiates a new AuthenticationValidationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthenticationValidationResponseWithDefaults() *AuthenticationValidationResponse {
	this := AuthenticationValidationResponse{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *AuthenticationValidationResponse) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationValidationResponse) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *AuthenticationValidationResponse) HasValid() bool {
	return o != nil && o.Valid != nil
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *AuthenticationValidationResponse) SetValid(v bool) {
	o.Valid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthenticationValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthenticationValidationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Valid *bool `json:"valid,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"valid"})
	} else {
		return err
	}
	o.Valid = all.Valid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
