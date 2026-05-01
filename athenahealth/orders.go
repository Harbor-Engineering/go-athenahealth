package athenahealth

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

type SignedOffOrder struct {
	AssignedTo           *string `json:"assignedto"`
	CreatedDate          *string `json:"createddate"`
	CreatedDateTime      *string `json:"createddatetime"`
	CreatedUser          *string `json:"createduser"`
	DeletedDateTime      *string `json:"deleteddatetime"`
	DepartmentID         *string `json:"departmentid"`
	Description          *string `json:"description"`
	DocumentClass        *string `json:"documentclass"`
	DocumentRoute        *string `json:"documentroute"`
	DocumentSource       *string `json:"documentsource"`
	DocumentSubclass     *string `json:"documentsubclass"`
	DocumentTypeID       *int    `json:"documenttypeid"`
	EncounterID          *string `json:"encounterid"`
	ExternalNote         *string `json:"externalnote"`
	FacilityID           *int    `json:"facilityid"`
	InternalNote         *string `json:"internalnote"`
	LastModifiedDate     *string `json:"lastmodifieddate"`
	LastModifiedDateTime *string `json:"lastmodifieddatetime"`
	LastModifiedUser     *string `json:"lastmodifieduser"`
	OrderID              *int    `json:"orderid"`
	OrderType            *string `json:"ordertype"`
	OrderTypeID          *int    `json:"ordertypeid"`
	PatientID            *int    `json:"patientid"`
	Priority             *string `json:"priority"`
	ProviderID           *int    `json:"providerid"`
	ProviderUsername     *string `json:"providerusername"`
	SignedDate           *string `json:"signeddate"`
	SignedDateTime       *string `json:"signeddatetime"`
	Status               *string `json:"status"`
	TieToOrderID         *int    `json:"tietoorderid"`
}

type ListChangedSignedOffOrdersOptions struct {
	DepartmentID               string
	LeaveUnprocessed           bool
	PatientIDs                 []string
	ShowProcessedEndDatetime   time.Time
	ShowProcessedStartDatetime time.Time

	Pagination *PaginationOptions
}

type listChangedSignedOffOrdersResponse struct {
	Orders []*SignedOffOrder `json:"orders"`

	*PaginationResponse
}

type ListChangedSignedOffOrdersResult struct {
	Orders []*SignedOffOrder

	Pagination *PaginationResult
}

// ListChangedSignedOffOrders - Get signed-off orders based on subscribed change events
//
// GET /v1/{practiceid}/orders/signedoff/changed
//
// https://docs.athenahealth.com/api/api-ref/document-type-order#Get-signed-off-orders
func (h *HTTPClient) ListChangedSignedOffOrders(ctx context.Context, opts *ListChangedSignedOffOrdersOptions) (*ListChangedSignedOffOrdersResult, error) {
	q := url.Values{}

	if opts != nil {
		if len(opts.DepartmentID) > 0 {
			q.Add("departmentid", opts.DepartmentID)
		}
		if opts.LeaveUnprocessed {
			q.Add("leaveunprocessed", strconv.FormatBool(opts.LeaveUnprocessed))
		}
		for _, id := range opts.PatientIDs {
			q.Add("patientids", id)
		}
		if !opts.ShowProcessedEndDatetime.IsZero() {
			q.Add("showprocessedenddatetime", opts.ShowProcessedEndDatetime.Format("01/02/2006 15:04:05"))
		}
		if !opts.ShowProcessedStartDatetime.IsZero() {
			q.Add("showprocessedstartdatetime", opts.ShowProcessedStartDatetime.Format("01/02/2006 15:04:05"))
		}
		if opts.Pagination != nil {
			if opts.Pagination.Limit > 0 {
				q.Add("limit", strconv.Itoa(opts.Pagination.Limit))
			}
			if opts.Pagination.Offset > 0 {
				q.Add("offset", strconv.Itoa(opts.Pagination.Offset))
			}
		}
	}

	out := &listChangedSignedOffOrdersResponse{}

	_, err := h.Get(ctx, "/orders/signedoff/changed", q, out)
	if err != nil {
		return nil, err
	}

	return &ListChangedSignedOffOrdersResult{
		Orders:     out.Orders,
		Pagination: makePaginationResult(out.Next, out.Previous, out.TotalCount),
	}, nil
}

// GetSignedOffOrderSubscription - Get list of signed-off order change subscriptions
//
// GET /v1/{practiceid}/orders/signedoff/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/document-type-order#Get-list-of-signed-off-order-change-subscription(s)
func (h *HTTPClient) GetSignedOffOrderSubscription(ctx context.Context) (*Subscription, error) {
	out := &Subscription{}

	_, err := h.Get(ctx, "/orders/signedoff/changed/subscription", nil, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type SubscribeSignedOffOrdersOptions struct {
	EventName string
}

// SubscribeSignedOffOrders - Subscribe to all/specific change events for signed-off orders
//
// POST /v1/{practiceid}/orders/signedoff/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/document-type-order#Subscribe-to-all/specific-change-events-for-signed-off-orders
func (h *HTTPClient) SubscribeSignedOffOrders(ctx context.Context, opts *SubscribeSignedOffOrdersOptions) error {
	var form url.Values

	if opts != nil && len(opts.EventName) > 0 {
		form = url.Values{}
		form.Add("eventname", opts.EventName)
	}

	_, err := h.PostForm(ctx, "/orders/signedoff/changed/subscription", form, nil)
	return err
}

// UnsubscribeSignedOffOrders - Unsubscribe from all/specific change events for signed-off orders
//
// DELETE /v1/{practiceid}/orders/signedoff/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/document-type-order#Subscribe-to-all/specific-change-events-for-signed-off-orders
func (h *HTTPClient) UnsubscribeSignedOffOrders(ctx context.Context, opts *SubscribeSignedOffOrdersOptions) error {
	var form url.Values

	if opts != nil && len(opts.EventName) > 0 {
		form = url.Values{}
		form.Add("eventname", opts.EventName)
	}

	_, err := h.DeleteForm(ctx, "/orders/signedoff/changed/subscription", form, nil)
	return err
}

type ChangedOrder struct {
	AssignedTo           *string `json:"assignedto"`
	CreatedDate          *string `json:"createddate"`
	CreatedDateTime      *string `json:"createddatetime"`
	CreatedUser          *string `json:"createduser"`
	DeletedDateTime      *string `json:"deleteddatetime"`
	DepartmentID         *string `json:"departmentid"`
	Description          *string `json:"description"`
	DocumentClass        *string `json:"documentclass"`
	DocumentRoute        *string `json:"documentroute"`
	DocumentSource       *string `json:"documentsource"`
	DocumentSubclass     *string `json:"documentsubclass"`
	DocumentTypeID       *int    `json:"documenttypeid"`
	EncounterID          *string `json:"encounterid"`
	ExternalNote         *string `json:"externalnote"`
	FacilityID           *int    `json:"facilityid"`
	InternalNote         *string `json:"internalnote"`
	LastModifiedDate     *string `json:"lastmodifieddate"`
	LastModifiedDateTime *string `json:"lastmodifieddatetime"`
	LastModifiedUser     *string `json:"lastmodifieduser"`
	OrderID              *int    `json:"orderid"`
	OrderType            *string `json:"ordertype"`
	OrderTypeID          *int    `json:"ordertypeid"`
	PatientID            *int    `json:"patientid"`
	Priority             *string `json:"priority"`
	ProviderID           *int    `json:"providerid"`
	ProviderUsername     *string `json:"providerusername"`
	Status               *string `json:"status"`
	TieToOrderID         *int    `json:"tietoorderid"`
}

type ListChangedOrdersOptions struct {
	DepartmentID               string
	LeaveUnprocessed           bool
	PatientIDs                 []string
	ShowProcessedEndDatetime   time.Time
	ShowProcessedStartDatetime time.Time

	Pagination *PaginationOptions
}

type listChangedOrdersResponse struct {
	Orders []*ChangedOrder `json:"orders"`

	*PaginationResponse
}

type ListChangedOrdersResult struct {
	Orders []*ChangedOrder

	Pagination *PaginationResult
}

// ListChangedOrders - Get changed orders based on subscribed change events
//
// GET /v1/{practiceid}/orders/changed
//
// https://docs.athenahealth.com/api/api-ref/order#Get-changed-orders
func (h *HTTPClient) ListChangedOrders(ctx context.Context, opts *ListChangedOrdersOptions) (*ListChangedOrdersResult, error) {
	q := url.Values{}

	if opts != nil {
		if len(opts.DepartmentID) > 0 {
			q.Add("departmentid", opts.DepartmentID)
		}
		if opts.LeaveUnprocessed {
			q.Add("leaveunprocessed", strconv.FormatBool(opts.LeaveUnprocessed))
		}
		for _, id := range opts.PatientIDs {
			q.Add("patientids", id)
		}
		if !opts.ShowProcessedEndDatetime.IsZero() {
			q.Add("showprocessedenddatetime", opts.ShowProcessedEndDatetime.Format("01/02/2006 15:04:05"))
		}
		if !opts.ShowProcessedStartDatetime.IsZero() {
			q.Add("showprocessedstartdatetime", opts.ShowProcessedStartDatetime.Format("01/02/2006 15:04:05"))
		}
		if opts.Pagination != nil {
			if opts.Pagination.Limit > 0 {
				q.Add("limit", strconv.Itoa(opts.Pagination.Limit))
			}
			if opts.Pagination.Offset > 0 {
				q.Add("offset", strconv.Itoa(opts.Pagination.Offset))
			}
		}
	}

	out := &listChangedOrdersResponse{}

	_, err := h.Get(ctx, "/orders/changed", q, out)
	if err != nil {
		return nil, err
	}

	return &ListChangedOrdersResult{
		Orders:     out.Orders,
		Pagination: makePaginationResult(out.Next, out.Previous, out.TotalCount),
	}, nil
}

// GetOrderSubscription - Get list of order change subscriptions
//
// GET /v1/{practiceid}/orders/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/order#Get-list-of-order-change-subscriptions
func (h *HTTPClient) GetOrderSubscription(ctx context.Context) (*Subscription, error) {
	out := &Subscription{}

	_, err := h.Get(ctx, "/orders/changed/subscription", nil, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type SubscribeOrdersOptions struct {
	EventName string
}

// SubscribeOrders - Subscribe to all/specific change events for orders
//
// POST /v1/{practiceid}/orders/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/order#Subscribe-to-all/specific-change-events-for-orders
func (h *HTTPClient) SubscribeOrders(ctx context.Context, opts *SubscribeOrdersOptions) error {
	var form url.Values

	if opts != nil && len(opts.EventName) > 0 {
		form = url.Values{}
		form.Add("eventname", opts.EventName)
	}

	_, err := h.PostForm(ctx, "/orders/changed/subscription", form, nil)
	return err
}

// UnsubscribeOrders - Unsubscribe from all/specific change events for orders
//
// DELETE /v1/{practiceid}/orders/changed/subscription
//
// https://docs.athenahealth.com/api/api-ref/order#Subscribe-to-all/specific-change-events-for-orders
func (h *HTTPClient) UnsubscribeOrders(ctx context.Context, opts *SubscribeOrdersOptions) error {
	var form url.Values

	if opts != nil && len(opts.EventName) > 0 {
		form = url.Values{}
		form.Add("eventname", opts.EventName)
	}

	_, err := h.DeleteForm(ctx, "/orders/changed/subscription", form, nil)
	return err
}
