# Example: Using Integration Tests to Validate Risk Contract Endpoints

This example demonstrates how to use the integration test harness to validate the Risk Contract Reference API endpoints.

## Step 1: Set up your environment

```bash
# Enable integration tests (required)
export ATHENA_RUN_INTEGRATION_TESTS=true

export ATHENA_PRACTICE_ID=your-practice-id
export ATHENA_API_KEY=your-api-key
export ATHENA_API_SECRET=your-api-secret

# Optional: Set test data if you have existing risk contracts
export ATHENA_TEST_RISK_CONTRACT_ID=123
export ATHENA_TEST_RISK_CONTRACT_NAME="Medicare Advantage Contract"
```

## Step 2: Run raw response tests to see actual API output

This is the most important step - it shows you exactly what the API returns:

```bash
# See the raw response from GET /populationmanagement/riskcontract
go test -v -run TestIntegration_GetRiskContractReference_RawResponse ./athenahealth

# See the raw response from PUT /populationmanagement/riskcontract
go test -v -run TestIntegration_UpdateRiskContractReference_RawResponse ./athenahealth
```

Example output:
```
=== RAW API RESPONSE ===
Method: GET
Path: /populationmanagement/riskcontract
Response:
{
  "description": "Medicare Advantage risk sharing contract",
  "name": "Medicare Advantage Contract",
  "riskcontractid": 123,
  "success": "true"
}
========================
```

## Step 3: Compare with your Go struct

Based on the raw response, verify your Go struct matches:

```go
type RiskContractReference struct {
    Description    string `json:"description"`
    ErrorMessage   string `json:"errormessage"`
    Name           string `json:"name"`
    RiskContractID int    `json:"riskcontractid"`
    Success        string `json:"success"`
}
```

## Step 4: Run the full integration tests

Once you've confirmed the response format, run the full tests:

```bash
# Test getting by ID
go test -v -run TestIntegration_GetRiskContractReference_ByID ./athenahealth

# Test getting by name
go test -v -run TestIntegration_GetRiskContractReference_ByName ./athenahealth

# Test creating a new contract
go test -v -run TestIntegration_UpdateRiskContractReference_Create ./athenahealth

# Test updating an existing contract
go test -v -run TestIntegration_UpdateRiskContractReference_Update ./athenahealth
```

## Step 5: Or use the helper script

```bash
# Run all integration tests
./scripts/run-integration-tests.sh

# Run only risk contract tests
./scripts/run-integration-tests.sh TestIntegration_RiskContract
```

## Troubleshooting

### Unmarshalling errors

If you get unmarshalling errors, run the `_RawResponse` tests first to see the actual API response format. Common issues:

1. **Array vs Object**: Check if the API returns `[{...}]` or just `{...}`
2. **Field types**: Verify `int` vs `string` for numeric fields
3. **Missing fields**: Check if the API returns fields not in your struct
4. **Extra fields**: Your struct can have fields the API doesn't return (they'll be zero values)

### Authentication errors

Make sure your credentials are valid and you have access to the endpoint in your practice.

### 404 errors

The endpoint might not be available in your practice's API version or configuration.

## Creating Your Own Integration Tests

Use this pattern to create integration tests for other endpoints:

```go
func TestIntegration_MyEndpoint_RawResponse(t *testing.T) {
    client := IntegrationTestClient(t)

    params := url.Values{}
    params.Add("param1", "value1")

    TestRawAPIResponse(t, client, http.MethodGet, "/path/to/endpoint", params)
}

func TestIntegration_MyEndpoint(t *testing.T) {
    client := IntegrationTestClient(t)

    result, err := client.MyMethod(context.Background(), "arg1", opts)
    if err != nil {
        t.Fatalf("MyMethod failed: %v", err)
    }

    LogResponse(t, "MyMethod Result", result)

    // Add assertions here
}
```
