# Integration Test Harness

This directory contains integration tests that make real API calls to the athenahealth API. These tests are useful for validating the actual response formats from the API endpoints.

## Setup

Integration tests are **skipped by default** and require explicit opt-in.

Set the following environment variables before running integration tests:

```bash
export ATHENA_RUN_INTEGRATION_TESTS=true  # Required to enable integration tests
export ATHENA_PRACTICE_ID=your-practice-id
export ATHENA_API_KEY=your-api-key
export ATHENA_API_SECRET=your-api-secret
```

## Running Integration Tests

### Run all integration tests:
```bash
go test -v -run TestIntegration ./athenahealth
```

### Run specific integration tests for Risk Contracts:
```bash
# Set optional test data
export ATHENA_TEST_RISK_CONTRACT_ID=123
export ATHENA_TEST_RISK_CONTRACT_NAME="Test Contract Name"

# Run the tests
go test -v -run TestIntegration_RiskContract ./athenahealth
```

### Run raw response tests to see actual API output:
```bash
# These tests log the raw JSON response from the API
go test -v -run TestIntegration_.*_RawResponse ./athenahealth
```

## Available Integration Tests

### Risk Contract Reference Tests

- `TestIntegration_GetRiskContractReference_RawResponse` - Shows raw API response for GET endpoint
- `TestIntegration_GetRiskContractReference_ByID` - Tests getting contract by ID
- `TestIntegration_GetRiskContractReference_ByName` - Tests getting contract by name
- `TestIntegration_UpdateRiskContractReference_RawResponse` - Shows raw API response for PUT endpoint
- `TestIntegration_UpdateRiskContractReference_Create` - Tests creating a new contract
- `TestIntegration_UpdateRiskContractReference_Update` - Tests updating an existing contract

## Creating New Integration Tests

Use the `IntegrationTestClient` helper to create a properly configured client:

```go
func TestIntegration_MyNewTest(t *testing.T) {
    client := IntegrationTestClient(t)

    // Your test code here
    result, err := client.SomeMethod(ctx, opts)
    if err != nil {
        t.Fatalf("API call failed: %v", err)
    }

    // Log the result for inspection
    LogResponse(t, "Result Description", result)
}
```

Use `TestRawAPIResponse` to inspect the raw JSON response format:

```go
func TestIntegration_MyEndpoint_RawResponse(t *testing.T) {
    client := IntegrationTestClient(t)

    params := url.Values{}
    params.Add("param1", "value1")

    // This will print the raw JSON response
    TestRawAPIResponse(t, client, http.MethodGet, "/path/to/endpoint", params)
}
```

## Best Practices

1. Always use the `_RawResponse` tests first to inspect the actual API response format
2. Don't commit sensitive data or credentials
3. Use descriptive test names that indicate what endpoint is being tested
4. Log responses using `LogResponse` for easier debugging
5. Run integration tests in isolation, not as part of regular unit test runs

## Safety

- Integration tests make **real API calls** and require valid credentials
- Use a test/sandbox practice ID whenever possible
- Never commit credentials to version control
- Be aware of API rate limits when running these tests
