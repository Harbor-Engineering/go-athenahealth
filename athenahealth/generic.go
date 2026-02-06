package athenahealth

type MessageResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ErrorMessageResponse struct {
	Message string `json:"errormessage"`
	Success bool   `json:"success"`
}

// IntegerSuccessResponse is used for endpoints that return success as an integer
type IntegerSuccessResponse struct {
	Errors  []map[string]interface{} `json:"errors,omitempty"`
	Success int                      `json:"success"`
}
