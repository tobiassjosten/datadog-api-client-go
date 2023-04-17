// Create or update service definition returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ServiceDefinitionsCreateRequest{
		ServiceDefinitionV2Dot1: &datadogV2.ServiceDefinitionV2Dot1{
			Application: datadog.PtrString("my-app"),
			Contacts: []datadogV2.ServiceDefinitionV2Dot1Contact{
				datadogV2.ServiceDefinitionV2Dot1Contact{
					ServiceDefinitionV2Dot1Email: &datadogV2.ServiceDefinitionV2Dot1Email{
						Contact: "contact@datadoghq.com",
						Name:    datadog.PtrString("Team Email"),
						Type:    datadogV2.SERVICEDEFINITIONV2DOT1EMAILTYPE_EMAIL,
					}},
			},
			DdService:   "my-service",
			Description: datadog.PtrString("My service description"),
			Extensions: map[string]interface{}{
				"myorg/extension": "extensionValue",
			},
			Integrations: &datadogV2.ServiceDefinitionV2Dot1Integrations{
				Opsgenie: &datadogV2.ServiceDefinitionV2Dot1Opsgenie{
					Region:     datadogV2.SERVICEDEFINITIONV2DOT1OPSGENIEREGION_US.Ptr(),
					ServiceUrl: "https://my-org.opsgenie.com/service/123e4567-e89b-12d3-a456-426614174000",
				},
				Pagerduty: &datadogV2.ServiceDefinitionV2Dot1Pagerduty{
					ServiceUrl: datadog.PtrString("https://my-org.pagerduty.com/service-directory/PMyService"),
				},
			},
			Lifecycle: datadog.PtrString("sandbox"),
			Links: []datadogV2.ServiceDefinitionV2Dot1Link{
				{
					Name:     "Runbook",
					Provider: datadog.PtrString("Github"),
					Type:     datadogV2.SERVICEDEFINITIONV2DOT1LINKTYPE_RUNBOOK,
					Url:      "https://my-runbook",
				},
			},
			SchemaVersion: datadogV2.SERVICEDEFINITIONV2DOT1VERSION_V2_1,
			Tags: []string{
				"my:tag",
				"service:tag",
			},
			Team: datadog.PtrString("my-team"),
			Tier: datadog.PtrString("High"),
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewServiceDefinitionApi(apiClient)
	resp, r, err := api.CreateOrUpdateServiceDefinitions(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceDefinitionApi.CreateOrUpdateServiceDefinitions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceDefinitionApi.CreateOrUpdateServiceDefinitions`:\n%s\n", responseContent)
}