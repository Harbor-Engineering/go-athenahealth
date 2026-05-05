package athenahealth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

// Integration tests for order endpoints.
// These tests make real API calls to validate response formats and path assumptions.
//
// To run these tests, set the following environment variables:
//
//	export ATHENA_RUN_INTEGRATION_TESTS=true
//	export ATHENA_PRACTICE_ID=your-practice-id
//	export ATHENA_API_KEY=your-api-key
//	export ATHENA_API_SECRET=your-api-secret
//
// Run with: go test -v -run TestIntegration_Order ./athenahealth

func TestIntegration_ListChangedOrders_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	t.Log("Testing GET /orders/changed")

	params := url.Values{}
	params.Add("leaveunprocessed", "true")

	TestRawAPIResponse(t, client, http.MethodGet, "/orders/changed", params)
}

func TestIntegration_ListChangedOrders(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()
	opts := &ListChangedOrdersOptions{
		LeaveUnprocessed: true,
	}

	result, err := client.ListChangedOrders(ctx, opts)
	if err != nil {
		t.Fatalf("ListChangedOrders failed: %v", err)
	}

	LogResponse(t, "ListChangedOrders Result", result)
	t.Logf("Total changed orders: %d", result.Pagination.TotalCount)
}

func TestIntegration_GetOrderSubscription_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	t.Log("Testing GET /orders/changed/subscription")

	TestRawAPIResponse(t, client, http.MethodGet, "/orders/changed/subscription", nil)
}

func TestIntegration_GetOrderSubscription(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()

	result, err := client.GetOrderSubscription(ctx)
	if err != nil {
		t.Fatalf("GetOrderSubscription failed: %v", err)
	}

	LogResponse(t, "GetOrderSubscription Result", result)
}

func TestIntegration_ListChangedSignedOffOrders_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	t.Log("Testing GET /orders/signedoff/changed")

	params := url.Values{}
	params.Add("leaveunprocessed", "true")

	TestRawAPIResponse(t, client, http.MethodGet, "/orders/signedoff/changed", params)
}

func TestIntegration_ListChangedSignedOffOrders(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()
	opts := &ListChangedSignedOffOrdersOptions{
		LeaveUnprocessed: true,
	}

	result, err := client.ListChangedSignedOffOrders(ctx, opts)
	if err != nil {
		t.Fatalf("ListChangedSignedOffOrders failed: %v", err)
	}

	LogResponse(t, "ListChangedSignedOffOrders Result", result)
	t.Logf("Total changed signed-off orders: %d", result.Pagination.TotalCount)
}

func TestIntegration_GetSignedOffOrderSubscription_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	t.Log("Testing GET /orders/signedoff/changed/subscription")

	TestRawAPIResponse(t, client, http.MethodGet, "/orders/signedoff/changed/subscription", nil)
}

func TestIntegration_GetSignedOffOrderSubscription(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()

	result, err := client.GetSignedOffOrderSubscription(ctx)
	if err != nil {
		t.Fatalf("GetSignedOffOrderSubscription failed: %v", err)
	}

	LogResponse(t, "GetSignedOffOrderSubscription Result", result)
}

func TestIntegration_AddOrderActionNote_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	orderID := mustGetEnvInt(t, "ATHENA_TEST_ORDER_ID", true)
	if orderID == 0 {
		t.Skip("ATHENA_TEST_ORDER_ID not set, skipping")
	}

	t.Logf("Testing POST /documents/order/%d/actions", orderID)

	params := url.Values{}
	params.Add("actionnote", "Integration test action note")

	TestRawAPIResponse(t, client, http.MethodPost, fmt.Sprintf("/documents/order/%d/actions", orderID), params)
}

func TestIntegration_AddOrderActionNote(t *testing.T) {
	client := IntegrationTestClient(t)

	orderID := mustGetEnvInt(t, "ATHENA_TEST_ORDER_ID", true)
	if orderID == 0 {
		t.Skip("ATHENA_TEST_ORDER_ID not set, skipping")
	}

	ctx := context.Background()
	opts := &AddOrderActionNoteOptions{
		ActionNote: "Integration test action note",
	}

	result, err := client.AddOrderActionNote(ctx, orderID, opts)
	if err != nil {
		t.Fatalf("AddOrderActionNote failed: %v", err)
	}

	LogResponse(t, "AddOrderActionNote Result", result)
	t.Logf("Success: %v", result.Success)
}
