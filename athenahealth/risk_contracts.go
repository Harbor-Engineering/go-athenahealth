package athenahealth

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// RiskContract represents a risk contract associated with a patient
type RiskContract struct {
	ContractName  string `json:"contractname"`
	EffectiveDate string `json:"effectivedate"`
	ExpirationDate string `json:"expirationdate"`
	RiskContractID int    `json:"riskcontractid"`
}

// ListRiskContractsOptions represents options for listing risk contracts
type ListRiskContractsOptions struct {
	DepartmentID string
	// If true, retrieve the record which indicates a risk contract is applied to all charts associated with the given patient
	AllCharts bool
}

// CreateRiskContractOptions represents options for creating/updating a risk contract
type CreateRiskContractOptions struct {
	RiskContractID int
	EffectiveDate  string // Format: MM/DD/YYYY
	ExpirationDate string // Format: MM/DD/YYYY (optional)
	DepartmentID   int    // Department ID (optional)
	// If true, apply this update to all charts associated with the given patient
	AllCharts bool
}

// ListRiskContracts - Get a list of risk contracts associated with the patient
//
// GET /v1/{practiceid}/chart/{patientid}/riskcontract
//
// https://docs.athenahealth.com/api/api-ref/patient-risk-contract
func (h *HTTPClient) ListRiskContracts(ctx context.Context, patientID string, opts *ListRiskContractsOptions) ([]*RiskContract, error) {
	out := []*RiskContract{}

	q := url.Values{}

	if opts != nil {
		if opts.DepartmentID != "" {
			q.Add("departmentid", opts.DepartmentID)
		}

		if opts.AllCharts {
			q.Add("allcharts", "true")
		}
	}

	_, err := h.Get(ctx, fmt.Sprintf("/chart/%s/riskcontract", patientID), q, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// CreateRiskContract - Create a new risk contract for the patient
//
// PUT /v1/{practiceid}/chart/{patientid}/riskcontract
//
// https://docs.athenahealth.com/api/api-ref/patient-risk-contract
func (h *HTTPClient) CreateRiskContract(ctx context.Context, patientID string, opts *CreateRiskContractOptions) error {
	if opts == nil {
		panic("opts is nil")
	}

	out := &MessageResponse{}

	form := url.Values{}
	form.Add("riskcontractid", strconv.Itoa(opts.RiskContractID))
	form.Add("effectivedate", opts.EffectiveDate)

	if opts.ExpirationDate != "" {
		form.Add("expirationdate", opts.ExpirationDate)
	}

	if opts.DepartmentID > 0 {
		form.Add("departmentid", strconv.Itoa(opts.DepartmentID))
	}

	if opts.AllCharts {
		form.Add("allcharts", "true")
	}

	_, err := h.PutForm(ctx, fmt.Sprintf("/chart/%s/riskcontract", patientID), form, &out)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRiskContractOptions represents options for deleting a risk contract
type DeleteRiskContractOptions struct {
	DepartmentID int  // Department ID (optional)
	// If true, apply this delete to all charts associated with the given patient
	AllCharts bool
}

// DeleteRiskContract - Delete a risk contract for the patient
//
// DELETE /v1/{practiceid}/chart/{patientid}/riskcontract/{riskcontractid}
//
// https://docs.athenahealth.com/api/api-ref/patient-risk-contract
func (h *HTTPClient) DeleteRiskContract(ctx context.Context, patientID string, riskContractID int, opts *DeleteRiskContractOptions) error {
	out := &MessageResponse{}

	path := fmt.Sprintf("/chart/%s/riskcontract/%d", patientID, riskContractID)

	if opts != nil {
		q := url.Values{}

		if opts.DepartmentID > 0 {
			q.Add("departmentid", strconv.Itoa(opts.DepartmentID))
		}

		if opts.AllCharts {
			q.Add("allcharts", "true")
		}

		if len(q) > 0 {
			path = fmt.Sprintf("%s?%s", path, q.Encode())
		}
	}

	_, err := h.Delete(ctx, path, nil, &out)
	if err != nil {
		return err
	}

	return nil
}
