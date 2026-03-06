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

// ensurePatientCaseSubscription ensures the subscription is active for patient case changes.
// Returns true if we created the subscription (caller should clean up), false if it was already active.
func ensurePatientCaseSubscription(t *testing.T, client *HTTPClient) bool {
	t.Helper()
	ctx := context.Background()

	// Check current subscription status
	sub, err := client.GetSubscription(ctx, "documents/patientcase")
	if err != nil {
		t.Logf("GetSubscription error (subscription may not exist yet): %v", err)
	}

	// If subscription is already ACTIVE, we don't need to do anything
	if sub != nil && sub.Status == "ACTIVE" {
		t.Log("Patient case subscription already ACTIVE")
		return false
	}

	if sub != nil {
		t.Logf("Current subscription status: %s - subscribing to patient case changes", sub.Status)
	} else {
		t.Log("No existing subscription - subscribing to patient case changes")
	}

	// Subscribe to all events (nil options subscribes to all)
	err = client.Subscribe(ctx, "documents/patientcase", nil)
	if err != nil {
		t.Fatalf("Failed to subscribe to patient case changes: %v", err)
	}

	// Verify subscription is now active
	sub, err = client.GetSubscription(ctx, "documents/patientcase")
	if err != nil {
		t.Fatalf("Failed to verify subscription after subscribing: %v", err)
	}

	if sub.Status != "ACTIVE" {
		t.Fatalf("Subscription status is %s after subscribing, expected ACTIVE", sub.Status)
	}

	t.Log("Successfully subscribed to patient case changes and verified ACTIVE status")
	return true
}

func TestIntegration_ListChangedPatientCases_RawResponse(t *testing.T) {
	client := IntegrationTestClient(t)

	// Ensure subscription is active
	needsCleanup := ensurePatientCaseSubscription(t, client)
	if needsCleanup {
		defer func() {
			ctx := context.Background()
			if err := client.Unsubscribe(ctx, "documents/patientcase", nil); err != nil {
				t.Logf("Warning: Failed to unsubscribe: %v", err)
			} else {
				t.Log("Successfully unsubscribed from patient case changes")
			}
		}()
	}

	t.Log("Testing GET /documents/patientcase/changed")

	params := url.Values{}
	params.Add("leaveunprocessed", "true")
	params.Add("limit", "10")

	TestRawAPIResponse(t, client, http.MethodGet, "/documents/patientcase/changed", params)
}

func TestIntegration_ListChangedPatientCases(t *testing.T) {
	client := IntegrationTestClient(t)

	// Ensure subscription is active
	needsCleanup := ensurePatientCaseSubscription(t, client)
	if needsCleanup {
		defer func() {
			ctx := context.Background()
			if err := client.Unsubscribe(ctx, "documents/patientcase", nil); err != nil {
				t.Logf("Warning: Failed to unsubscribe: %v", err)
			} else {
				t.Log("Successfully unsubscribed from patient case changes")
			}
		}()
	}

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

	// Ensure subscription is active
	needsCleanup := ensurePatientCaseSubscription(t, client)
	if needsCleanup {
		defer func() {
			ctx := context.Background()
			if err := client.Unsubscribe(ctx, "documents/patientcase", nil); err != nil {
				t.Logf("Warning: Failed to unsubscribe: %v", err)
			} else {
				t.Log("Successfully unsubscribed from patient case changes")
			}
		}()
	}

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
