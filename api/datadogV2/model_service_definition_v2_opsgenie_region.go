// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ServiceDefinitionV2OpsgenieRegion Opsgenie instance region.
type ServiceDefinitionV2OpsgenieRegion string

// List of ServiceDefinitionV2OpsgenieRegion.
const (
	SERVICEDEFINITIONV2OPSGENIEREGION_US ServiceDefinitionV2OpsgenieRegion = "US"
	SERVICEDEFINITIONV2OPSGENIEREGION_EU ServiceDefinitionV2OpsgenieRegion = "EU"
)

var allowedServiceDefinitionV2OpsgenieRegionEnumValues = []ServiceDefinitionV2OpsgenieRegion{
	SERVICEDEFINITIONV2OPSGENIEREGION_US,
	SERVICEDEFINITIONV2OPSGENIEREGION_EU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2OpsgenieRegion) GetAllowedValues() []ServiceDefinitionV2OpsgenieRegion {
	return allowedServiceDefinitionV2OpsgenieRegionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2OpsgenieRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2OpsgenieRegion(value)
	return nil
}

// NewServiceDefinitionV2OpsgenieRegionFromValue returns a pointer to a valid ServiceDefinitionV2OpsgenieRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2OpsgenieRegionFromValue(v string) (*ServiceDefinitionV2OpsgenieRegion, error) {
	ev := ServiceDefinitionV2OpsgenieRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2OpsgenieRegion: valid values are %v", v, allowedServiceDefinitionV2OpsgenieRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2OpsgenieRegion) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2OpsgenieRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2OpsgenieRegion value.
func (v ServiceDefinitionV2OpsgenieRegion) Ptr() *ServiceDefinitionV2OpsgenieRegion {
	return &v
}
