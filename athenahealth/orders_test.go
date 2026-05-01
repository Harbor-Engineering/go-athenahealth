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
