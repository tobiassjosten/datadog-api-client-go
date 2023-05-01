// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansMetricResponseCompute The compute rule to compute the span-based metric.
type SpansMetricResponseCompute struct {
	// The type of aggregation to use.
	AggregationType *SpansMetricComputeAggregationType `json:"aggregation_type,omitempty"`
	// Toggle to include or exclude percentile aggregations for distribution metrics.
	// Only present when the `aggregation_type` is `distribution`.
	IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	// The path to the value the span-based metric will aggregate on (only used if the aggregation type is a "distribution").
	Path *string `json:"path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansMetricResponseCompute instantiates a new SpansMetricResponseCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansMetricResponseCompute() *SpansMetricResponseCompute {
	this := SpansMetricResponseCompute{}
	return &this
}

// NewSpansMetricResponseComputeWithDefaults instantiates a new SpansMetricResponseCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansMetricResponseComputeWithDefaults() *SpansMetricResponseCompute {
	this := SpansMetricResponseCompute{}
	return &this
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *SpansMetricResponseCompute) GetAggregationType() SpansMetricComputeAggregationType {
	if o == nil || o.AggregationType == nil {
		var ret SpansMetricComputeAggregationType
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseCompute) GetAggregationTypeOk() (*SpansMetricComputeAggregationType, bool) {
	if o == nil || o.AggregationType == nil {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *SpansMetricResponseCompute) HasAggregationType() bool {
	return o != nil && o.AggregationType != nil
}

// SetAggregationType gets a reference to the given SpansMetricComputeAggregationType and assigns it to the AggregationType field.
func (o *SpansMetricResponseCompute) SetAggregationType(v SpansMetricComputeAggregationType) {
	o.AggregationType = &v
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *SpansMetricResponseCompute) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseCompute) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *SpansMetricResponseCompute) HasIncludePercentiles() bool {
	return o != nil && o.IncludePercentiles != nil
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *SpansMetricResponseCompute) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SpansMetricResponseCompute) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricResponseCompute) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SpansMetricResponseCompute) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SpansMetricResponseCompute) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansMetricResponseCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AggregationType != nil {
		toSerialize["aggregation_type"] = o.AggregationType
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansMetricResponseCompute) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AggregationType    *SpansMetricComputeAggregationType `json:"aggregation_type,omitempty"`
		IncludePercentiles *bool                              `json:"include_percentiles,omitempty"`
		Path               *string                            `json:"path,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_type", "include_percentiles", "path"})
	} else {
		return err
	}
	if v := all.AggregationType; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.AggregationType = all.AggregationType
	o.IncludePercentiles = all.IncludePercentiles
	o.Path = all.Path
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
