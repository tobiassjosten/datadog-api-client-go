// Update items of a dashboard list returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.DashboardListUpdateItemsRequest{
		Dashboards: &[]datadog.DashboardListItemRequest{
			{
				Id:   "q5j-nti-fv6",
				Type: datadog.DASHBOARDTYPE_HOST_TIMEBOARD,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardListsApi.UpdateDashboardListItems(ctx, 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardListItems`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.UpdateDashboardListItems`:\n%s\n", responseContent)
}