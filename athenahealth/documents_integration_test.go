package athenahealth

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

// Integration tests for Document endpoints (Patient Cases)
// These tests make real API calls to validate response formats
//
// To run these tests, set the following environment variables:
//   export ATHENA_RUN_INTEGRATION_TESTS=true
//   export ATHENA_PRACTICE_ID=your-practice-id
//   export ATHENA_API_KEY=your-api-key
//   export ATHENA_API_SECRET=your-api-secret
//
// Run with: go test -v -run TestIntegration_ListChangedPatientCases ./athenahealth

func TestIntegration_ListChangedPatientCases_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	t.Log("Testing GET /documents/patientcase/changed")

	params := url.Values{}
	params.Add("leaveunprocessed", "true")
	params.Add("limit", "10")

	TestRawAPIResponse(t, client, http.MethodGet, "/documents/patientcase/changed", params)
}

func TestIntegration_ListChangedPatientCases(t *testing.T) {
	client := IntegrationTestClient(t)

	ctx := context.Background()
	opts := &ListChangedPatientCasesOptions{
		LeaveUnprocessed: true,
		Pagination: &PaginationOptions{
			Limit: 10,
		},
	}

	result, err := client.ListChangedPatientCases(ctx, opts)
	if err != nil {
		t.Fatalf("ListChangedPatientCases failed: %v", err)
	}

	LogResponse(t, "ListChangedPatientCases Result", result)

	if result.Pagination != nil {
		t.Logf("Total count: %d", result.Pagination.TotalCount)
		t.Logf("Returned patient cases: %d", len(result.ChangedPatientCases))
	}

	// Validate structure if we got results
	if len(result.ChangedPatientCases) > 0 {
		first := result.ChangedPatientCases[0]
		t.Logf("First patient case:")
		t.Logf("  PatientCaseID: %s", first.PatientCaseID)
		t.Logf("  PatientID: %d", first.PatientID)
		t.Logf("  Status: %s", first.Status)
		t.Logf("  Subject: %s", first.Subject)
		t.Logf("  DepartmentID: %s", first.DepartmentID)
		t.Logf("  DocumentClass: %s", first.DocumentClass)
	}
}

func TestIntegration_ListChangedPatientCases_WithPatientIDs(t *testing.T) {
	client := IntegrationTestClient(t)

	// You can optionally set ATHENA_TEST_PATIENT_ID to filter by specific patient
	patientID := mustGetEnv(t, "ATHENA_TEST_PATIENT_ID", true)
	if patientID == "" {
		t.Skip("ATHENA_TEST_PATIENT_ID not set, skipping patient-specific test")
	}

	ctx := context.Background()
	opts := &ListChangedPatientCasesOptions{
		PatientIDs:       []string{patientID},
		LeaveUnprocessed: true,
		Pagination: &PaginationOptions{
			Limit: 5,
		},
	}

	result, err := client.ListChangedPatientCases(ctx, opts)
	if err != nil {
		t.Fatalf("ListChangedPatientCases failed: %v", err)
	}

	LogResponse(t, "ListChangedPatientCases Result (Filtered by PatientID)", result)

	t.Logf("Found %d patient cases for patient %s", len(result.ChangedPatientCases), patientID)

	// All results should be for the specified patient
	for i, pc := range result.ChangedPatientCases {
		expectedID, _ := strconv.Atoi(patientID)
		if pc.PatientID != expectedID {
			t.Errorf("Result[%d]: Expected PatientID=%s, got %d", i, patientID, pc.PatientID)
		}
	}
}
