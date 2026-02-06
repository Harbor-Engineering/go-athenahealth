package athenahealth

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient_ListRiskContracts(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	departmentID := "456"

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodGet, r.Method)
		assert.Equal(departmentID, r.URL.Query().Get("departmentid"))

		b, _ := os.ReadFile("./resources/ListRiskContracts.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &ListRiskContractsOptions{
		DepartmentID: departmentID,
	}

	contracts, err := athenaClient.ListRiskContracts(context.Background(), patientID, opts)
	assert.NoError(err)
	assert.NotNil(contracts)
	assert.Len(contracts, 2)

	// Verify first contract
	assert.Equal("Medicare Advantage", contracts[0].ContractName)
	assert.Equal("01/01/2024", contracts[0].EffectiveDate)
	assert.Equal("12/31/2024", contracts[0].ExpirationDate)
	assert.Equal(123, contracts[0].RiskContractID)

	// Verify second contract
	assert.Equal("Commercial HMO", contracts[1].ContractName)
	assert.Equal("06/01/2024", contracts[1].EffectiveDate)
	assert.Equal("05/31/2025", contracts[1].ExpirationDate)
	assert.Equal(456, contracts[1].RiskContractID)
}

func TestHTTPClient_ListRiskContracts_WithoutDepartment(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodGet, r.Method)
		assert.Empty(r.URL.Query().Get("departmentid"))

		b, _ := os.ReadFile("./resources/ListRiskContracts.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	contracts, err := athenaClient.ListRiskContracts(context.Background(), patientID, nil)
	assert.NoError(err)
	assert.NotNil(contracts)
}

func TestHTTPClient_ListRiskContracts_WithAllCharts(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodGet, r.Method)
		assert.Equal("true", r.URL.Query().Get("allcharts"))

		b, _ := os.ReadFile("./resources/ListRiskContracts.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &ListRiskContractsOptions{
		AllCharts: true,
	}

	contracts, err := athenaClient.ListRiskContracts(context.Background(), patientID, opts)
	assert.NoError(err)
	assert.NotNil(contracts)
}

func TestHTTPClient_CreateRiskContract(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	opts := &CreateRiskContractOptions{
		RiskContractID: 789,
		EffectiveDate:  "01/15/2024",
		ExpirationDate: "01/14/2025",
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodPut, r.Method)

		assert.NoError(r.ParseForm())
		assert.Equal(strconv.Itoa(opts.RiskContractID), r.Form.Get("riskcontractid"))
		assert.Equal(opts.EffectiveDate, r.Form.Get("effectivedate"))
		assert.Equal(opts.ExpirationDate, r.Form.Get("expirationdate"))

		b, _ := os.ReadFile("./resources/CreateRiskContract.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	err := athenaClient.CreateRiskContract(context.Background(), patientID, opts)
	assert.NoError(err)
}

func TestHTTPClient_CreateRiskContract_WithoutExpirationDate(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	opts := &CreateRiskContractOptions{
		RiskContractID: 789,
		EffectiveDate:  "01/15/2024",
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodPut, r.Method)

		assert.NoError(r.ParseForm())
		assert.Equal(strconv.Itoa(opts.RiskContractID), r.Form.Get("riskcontractid"))
		assert.Equal(opts.EffectiveDate, r.Form.Get("effectivedate"))
		assert.Empty(r.Form.Get("expirationdate"))

		b, _ := os.ReadFile("./resources/CreateRiskContract.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	err := athenaClient.CreateRiskContract(context.Background(), patientID, opts)
	assert.NoError(err)
}

func TestHTTPClient_CreateRiskContract_WithDepartmentIDAndAllCharts(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	opts := &CreateRiskContractOptions{
		RiskContractID: 789,
		EffectiveDate:  "01/15/2024",
		ExpirationDate: "01/14/2025",
		DepartmentID:   456,
		AllCharts:      true,
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract", r.URL.Path)
		assert.Equal(http.MethodPut, r.Method)

		assert.NoError(r.ParseForm())
		assert.Equal(strconv.Itoa(opts.RiskContractID), r.Form.Get("riskcontractid"))
		assert.Equal(opts.EffectiveDate, r.Form.Get("effectivedate"))
		assert.Equal(opts.ExpirationDate, r.Form.Get("expirationdate"))
		assert.Equal(strconv.Itoa(opts.DepartmentID), r.Form.Get("departmentid"))
		assert.Equal("true", r.Form.Get("allcharts"))

		b, _ := os.ReadFile("./resources/CreateRiskContract.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	err := athenaClient.CreateRiskContract(context.Background(), patientID, opts)
	assert.NoError(err)
}

func TestHTTPClient_DeleteRiskContract(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	riskContractID := 789

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract/789", r.URL.Path)
		assert.Equal(http.MethodDelete, r.Method)

		b, _ := os.ReadFile("./resources/DeleteRiskContract.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	err := athenaClient.DeleteRiskContract(context.Background(), patientID, riskContractID, nil)
	assert.NoError(err)
}

func TestHTTPClient_DeleteRiskContract_WithOptions(t *testing.T) {
	assert := assert.New(t)

	patientID := "123"
	riskContractID := 789
	opts := &DeleteRiskContractOptions{
		DepartmentID: 456,
		AllCharts:    true,
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("/chart/123/riskcontract/789", r.URL.Path)
		assert.Equal(http.MethodDelete, r.Method)
		assert.Equal(strconv.Itoa(opts.DepartmentID), r.URL.Query().Get("departmentid"))
		assert.Equal("true", r.URL.Query().Get("allcharts"))

		b, _ := os.ReadFile("./resources/DeleteRiskContract.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	err := athenaClient.DeleteRiskContract(context.Background(), patientID, riskContractID, opts)
	assert.NoError(err)
}
