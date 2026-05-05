package athenahealth

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient_ListChangedSignedOffOrders(t *testing.T) {
	assert := assert.New(t)

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("true", r.URL.Query().Get("leaveunprocessed"))

		b, _ := os.ReadFile("./resources/ListChangedSignedOffOrders.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &ListChangedSignedOffOrdersOptions{
		LeaveUnprocessed: true,
	}

	res, err := athenaClient.ListChangedSignedOffOrders(context.Background(), opts)

	assert.NoError(err)
	assert.Len(res.Orders, 2)
	assert.Equal(2, res.Pagination.TotalCount)
}

func TestHTTPClient_GetSignedOffOrderSubscription(t *testing.T) {
	assert := assert.New(t)

	h := func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("./resources/GetSignedOffOrderSubscription.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	subscription, err := athenaClient.GetSignedOffOrderSubscription(context.Background())

	assert.NoError(err)
	assert.NotNil(subscription)
	assert.Equal("ACTIVE", subscription.Status)
	assert.Len(subscription.Subscriptions, 1)
}

func TestHTTPClient_SubscribeSignedOffOrders(t *testing.T) {
	assert := assert.New(t)

	called := false
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		assert.Contains(string(reqBody), "eventname=SignOffOrder")

		called = true
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &SubscribeSignedOffOrdersOptions{
		EventName: "SignOffOrder",
	}
	err := athenaClient.SubscribeSignedOffOrders(context.Background(), opts)

	assert.NoError(err)
	assert.True(called)
}

func TestHTTPClient_UnsubscribeSignedOffOrders(t *testing.T) {
	assert := assert.New(t)

	called := false
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		assert.Contains(string(reqBody), "eventname=SignOffOrder")

		called = true
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &SubscribeSignedOffOrdersOptions{
		EventName: "SignOffOrder",
	}
	err := athenaClient.UnsubscribeSignedOffOrders(context.Background(), opts)

	assert.NoError(err)
	assert.True(called)
}

func TestHTTPClient_ListChangedOrders(t *testing.T) {
	assert := assert.New(t)

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal("true", r.URL.Query().Get("leaveunprocessed"))

		b, _ := os.ReadFile("./resources/ListChangedOrders.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &ListChangedOrdersOptions{
		LeaveUnprocessed: true,
	}

	res, err := athenaClient.ListChangedOrders(context.Background(), opts)

	assert.NoError(err)
	assert.Len(res.Orders, 2)
	assert.Equal(2, res.Pagination.TotalCount)
}

func TestHTTPClient_GetOrderSubscription(t *testing.T) {
	assert := assert.New(t)

	h := func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("./resources/GetOrderSubscription.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	subscription, err := athenaClient.GetOrderSubscription(context.Background())

	assert.NoError(err)
	assert.NotNil(subscription)
	assert.Equal("ACTIVE", subscription.Status)
	assert.Len(subscription.Subscriptions, 1)
}

func TestHTTPClient_SubscribeOrders(t *testing.T) {
	assert := assert.New(t)

	called := false
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		assert.Contains(string(reqBody), "eventname=UpdateOrder")

		called = true
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &SubscribeOrdersOptions{
		EventName: "UpdateOrder",
	}
	err := athenaClient.SubscribeOrders(context.Background(), opts)

	assert.NoError(err)
	assert.True(called)
}

func TestHTTPClient_UnsubscribeOrders(t *testing.T) {
	assert := assert.New(t)

	called := false
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		assert.Contains(string(reqBody), "eventname=UpdateOrder")

		called = true
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &SubscribeOrdersOptions{
		EventName: "UpdateOrder",
	}
	err := athenaClient.UnsubscribeOrders(context.Background(), opts)

	assert.NoError(err)
	assert.True(called)
}

func TestHTTPClient_AddOrderActionNote(t *testing.T) {
	assert := assert.New(t)

	orderID := 12345

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(http.MethodPost, r.Method)
		assert.Equal("/documents/order/12345/actions", r.URL.Path)

		assert.NoError(r.ParseForm())
		assert.Equal("Test action note", r.Form.Get("actionnote"))

		b, _ := os.ReadFile("./resources/AddOrderActionNote.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	opts := &AddOrderActionNoteOptions{
		ActionNote: "Test action note",
	}

	result, err := athenaClient.AddOrderActionNote(context.Background(), orderID, opts)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal("true", result.Success)
	assert.NotNil(result.NewDocumentID)
	assert.Equal("98765", *result.NewDocumentID)
	assert.NotNil(result.VersionToken)
	assert.Equal("abc123token", *result.VersionToken)
}

func TestHTTPClient_AddOrderActionNote_NilOpts(t *testing.T) {
	assert := assert.New(t)

	orderID := 12345

	h := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(http.MethodPost, r.Method)
		assert.Equal("/documents/order/12345/actions", r.URL.Path)

		assert.NoError(r.ParseForm())
		assert.Empty(r.Form.Get("actionnote"))

		b, _ := os.ReadFile("./resources/AddOrderActionNote.json")
		w.Write(b)
	}

	athenaClient, ts := testClient(h)
	defer ts.Close()

	result, err := athenaClient.AddOrderActionNote(context.Background(), orderID, nil)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal("true", result.Success)
}
