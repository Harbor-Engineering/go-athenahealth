package athenahealth

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
)

// Integration tests for Risk Contract Reference endpoints
// These tests make real API calls to validate response formats
//
// To run these tests, set the following environment variables:
//   export ATHENA_INTEGRATION_TEST=true
//   export ATHENA_PRACTICE_ID=your-practice-id
//   export ATHENA_API_KEY=your-api-key
//   export ATHENA_API_SECRET=your-api-secret
//   export ATHENA_TEST_RISK_CONTRACT_ID=123  # An existing risk contract ID (optional)
//   export ATHENA_TEST_RISK_CONTRACT_NAME="Test Contract"  # An existing contract name (optional)
//
// Run with: go test -v -run TestIntegration_RiskContract ./athenahealth

func TestIntegration_GetRiskContractReference_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	riskContractID := mustGetEnvInt(t, "ATHENA_TEST_RISK_CONTRACT_ID", true)
	if riskContractID == 0 {
		t.Skip("ATHENA_TEST_RISK_CONTRACT_ID not set, skipping")
	}

	t.Logf("Testing GET /populationmanagement/riskcontract with riskcontractid=%d", riskContractID)

	params := url.Values{}
	params.Add("riskcontractid", strconv.Itoa(riskContractID))

	TestRawAPIResponse(t, client, http.MethodGet, "/populationmanagement/riskcontract", params)
}

func TestIntegration_GetRiskContractReference_ByID(t *testing.T) {
	client := IntegrationTestClient(t)

	riskContractID := mustGetEnvInt(t, "ATHENA_TEST_RISK_CONTRACT_ID", true)
	if riskContractID == 0 {
		t.Skip("ATHENA_TEST_RISK_CONTRACT_ID not set, skipping")
	}

	ctx := context.Background()
	opts := &GetRiskContractReferenceOptions{
		RiskContractID: riskContractID,
	}

	result, err := client.GetRiskContractReference(ctx, opts)
	if err != nil {
		t.Fatalf("GetRiskContractReference failed: %v", err)
	}

	LogResponse(t, "GetRiskContractReference Result", result)

	if result.RiskContractID != riskContractID {
		t.Errorf("Expected RiskContractID=%d, got %d", riskContractID, result.RiskContractID)
	}
}

func TestIntegration_GetRiskContractReference_ByName(t *testing.T) {
	client := IntegrationTestClient(t)

	contractName := mustGetEnv(t, "ATHENA_TEST_RISK_CONTRACT_NAME", true)
	if contractName == "" {
		t.Skip("ATHENA_TEST_RISK_CONTRACT_NAME not set, skipping")
	}

	ctx := context.Background()
	opts := &GetRiskContractReferenceOptions{
		Name: contractName,
	}

	result, err := client.GetRiskContractReference(ctx, opts)
	if err != nil {
		t.Fatalf("GetRiskContractReference failed: %v", err)
	}

	LogResponse(t, "GetRiskContractReference Result", result)

	if result.Name != contractName {
		t.Errorf("Expected Name=%s, got %s", contractName, result.Name)
	}
}

func TestIntegration_UpdateRiskContractReference_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	// This test creates a new risk contract reference or updates an existing one
	t.Log("Testing PUT /populationmanagement/riskcontract")

	params := url.Values{}
	params.Add("name", "Integration Test Contract")
	params.Add("description", "Created by integration test")

	// If you want to update an existing contract, uncomment and set:
	// riskContractID := mustGetEnvInt(t, "ATHENA_TEST_RISK_CONTRACT_ID", true)
	// if riskContractID > 0 {
	// 	params.Add("riskcontractid", strconv.Itoa(riskContractID))
	// }

	TestRawAPIResponse(t, client, http.MethodPut, "/populationmanagement/riskcontract", params)
}

func TestIntegration_UpdateRiskContractReference_Create(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()
	opts := &UpdateRiskContractReferenceOptions{
		Name:        "Integration Test Contract",
		Description: "Created by integration test",
	}

	result, err := client.UpdateRiskContractReference(ctx, opts)
	if err != nil {
		t.Fatalf("UpdateRiskContractReference failed: %v", err)
	}

	LogResponse(t, "UpdateRiskContractReference Result (Create)", result)

	if result.RiskContractID == 0 {
		t.Error("Expected non-zero RiskContractID after creation")
	}
}

func TestIntegration_UpdateRiskContractReference_Update(t *testing.T) {
	client := IntegrationTestClient(t)

	riskContractID := mustGetEnvInt(t, "ATHENA_TEST_RISK_CONTRACT_ID", true)
	if riskContractID == 0 {
		t.Skip("ATHENA_TEST_RISK_CONTRACT_ID not set, skipping")
	}

	ctx := context.Background()
	opts := &UpdateRiskContractReferenceOptions{
		RiskContractID: riskContractID,
		Name:           "Updated Integration Test Contract",
		Description:    "Updated by integration test",
	}

	result, err := client.UpdateRiskContractReference(ctx, opts)
	if err != nil {
		t.Fatalf("UpdateRiskContractReference failed: %v", err)
	}

	LogResponse(t, "UpdateRiskContractReference Result (Update)", result)

	if result.RiskContractID != riskContractID {
		t.Errorf("Expected RiskContractID=%d, got %d", riskContractID, result.RiskContractID)
	}
}

// Helper functions

func mustGetEnv(t *testing.T, key string, optional bool) string {
	t.Helper()
	value := os.Getenv(key)
	if value == "" && !optional {
		t.Fatalf("Required environment variable %s not set", key)
	}
	return value
}

func mustGetEnvInt(t *testing.T, key string, optional bool) int {
	t.Helper()
	value := os.Getenv(key)
	if value == "" {
		if !optional {
			t.Fatalf("Required environment variable %s not set", key)
		}
		return 0
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		t.Fatalf("Invalid integer value for %s: %s", key, value)
	}

	return intVal
}
