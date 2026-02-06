package athenahealth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
)

// IntegrationTestClient creates a real HTTP client for integration testing.
// Integration tests are SKIPPED by default and require explicit opt-in.
//
// To run integration tests, set:
//   - ATHENA_RUN_INTEGRATION_TESTS=true (required to enable)
//   - ATHENA_PRACTICE_ID
//   - ATHENA_CLIENT_ID (or ATHENA_API_KEY)
//   - ATHENA_CLIENT_SECRET (or ATHENA_API_SECRET)
func IntegrationTestClient(t *testing.T) *HTTPClient {
	t.Helper()

	if os.Getenv("ATHENA_RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test: Set ATHENA_RUN_INTEGRATION_TESTS=true to run")
	}

	practiceID := os.Getenv("ATHENA_PRACTICE_ID")

	// Support both naming conventions
	clientID := os.Getenv("ATHENA_CLIENT_ID")
	if clientID == "" {
		clientID = os.Getenv("ATHENA_API_KEY")
	}

	clientSecret := os.Getenv("ATHENA_CLIENT_SECRET")
	if clientSecret == "" {
		clientSecret = os.Getenv("ATHENA_API_SECRET")
	}

	if practiceID == "" || clientID == "" || clientSecret == "" {
		t.Fatal("Required environment variables not set: ATHENA_PRACTICE_ID, ATHENA_CLIENT_ID/ATHENA_API_KEY, ATHENA_CLIENT_SECRET/ATHENA_API_SECRET")
	}

	client := NewHTTPClient(http.DefaultClient, practiceID, clientID, clientSecret)

	return client
}

// LogResponse is a helper to log the full response for debugging
func LogResponse(t *testing.T, description string, v interface{}) {
	t.Helper()

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Logf("Error marshalling %s: %v", description, err)
		return
	}

	t.Logf("%s:\n%s", description, string(b))
}

// TestRawAPIResponse makes a raw API call and logs the response body for inspection
func TestRawAPIResponse(t *testing.T, client *HTTPClient, method, path string, params url.Values) {
	t.Helper()

	ctx := context.Background()

	var resp *http.Response
	var err error

	switch method {
	case http.MethodGet:
		resp, err = client.Get(ctx, path, params, nil)
	case http.MethodPut:
		resp, err = client.PutForm(ctx, path, params, nil)
	case http.MethodPost:
		resp, err = client.PostForm(ctx, path, params, nil)
	case http.MethodDelete:
		resp, err = client.Delete(ctx, path, nil, nil)
	default:
		t.Fatalf("Unsupported method: %s", method)
	}

	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		var responseData interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		b, _ := json.MarshalIndent(responseData, "", "  ")
		fmt.Printf("\n=== RAW API RESPONSE ===\n")
		fmt.Printf("Method: %s\n", method)
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Response:\n%s\n", string(b))
		fmt.Printf("========================\n\n")
	}
}
