// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesBatchAttributes The JSON:API attributes for a batched set of scorecard outcomes.
type OutcomesBatchAttributes struct {
	// Set of scorecard outcomes to update.
	Results []OutcomesBatchRequestItem `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewOutcomesBatchAttributes instantiates a new OutcomesBatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesBatchAttributes() *OutcomesBatchAttributes {
	this := OutcomesBatchAttributes{}
	return &this
}

// NewOutcomesBatchAttributesWithDefaults instantiates a new OutcomesBatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesBatchAttributesWithDefaults() *OutcomesBatchAttributes {
	this := OutcomesBatchAttributes{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OutcomesBatchAttributes) GetResults() []OutcomesBatchRequestItem {
	if o == nil || o.Results == nil {
		var ret []OutcomesBatchRequestItem
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchAttributes) GetResultsOk() (*[]OutcomesBatchRequestItem, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OutcomesBatchAttributes) HasResults() bool {
	return o != nil && o.Results != nil
}

// SetResults gets a reference to the given []OutcomesBatchRequestItem and assigns it to the Results field.
func (o *OutcomesBatchAttributes) SetResults(v []OutcomesBatchRequestItem) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesBatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutcomesBatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Results []OutcomesBatchRequestItem `json:"results,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"results"})
	} else {
		return err
	}
	o.Results = all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
