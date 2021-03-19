package restcli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

type ClientBuilder struct {
	Headers            http.Header
	Timeout            time.Duration
	ConnectTimeout     time.Duration
	BaseURL            string
	ContentType        rest.ContentType
	EnableCache        bool
	DisableTimeout     bool
	FollowRedirect     bool
	CustomPool         *rest.CustomPool
	BasicAuth          *rest.BasicAuth
	UserAgent          string
	Client             *http.Client
	UncompressResponse bool
}

func (cb ClientBuilder) Build() client {
	return client{RestClient: &rest.RequestBuilder{
		Headers:        cb.Headers,
		Timeout:        cb.Timeout,
		ConnectTimeout: cb.ConnectTimeout,
		BaseURL:        cb.BaseURL,
		ContentType:    cb.ContentType,
		DisableTimeout: cb.DisableTimeout,
		FollowRedirect: cb.FollowRedirect,
		CustomPool:     cb.CustomPool,
		BasicAuth:      cb.BasicAuth,
		UserAgent:      cb.UserAgent,
		Client:         cb.Client,
	}}
}

type client struct {
	RestClient *rest.RequestBuilder
}

func (c client) GET(ctx context.Context, URL string, structure interface{}) (*rest.Response, error) {
	response := c.RestClient.Get(URL)
	return response, c.unmarshalResponse(response, URL, structure)
}

func (c client) POST(ctx context.Context, URL string, body interface{}, structure interface{}) (*rest.Response, error) {
	response := c.RestClient.Post(URL, body)
	return response, c.unmarshalResponse(response, URL, structure)
}
func (c client) PUT(ctx context.Context, URL string, body interface{}, structure interface{}) (*rest.Response, error) {
	response := c.RestClient.Put(URL, body)
	return response, c.unmarshalResponse(response, URL, structure)
}

func (c client) DELETE(ctx context.Context, URL string, structure interface{}) (*rest.Response, error) {
	response := c.RestClient.Delete(URL)
	return response, c.unmarshalResponse(response, URL, structure)
}

type externalAPIError struct {
	Message  string `json:"message"`
	ErrorStr string `json:"error"`
	Cause    string `json:"cause"`
}

func (c client) unmarshalResponse(response *rest.Response, URL string, structure interface{}) error {
	if response.Response == nil {
		response.Response = &http.Response{}
	}

	if response.Err != nil {
		return response.Err
	}

	if response.StatusCode >= 300 {

		if response.StatusCode == http.StatusNotFound {
			return fmt.Errorf("resource not found %s", URL)
		}

		upstreamErr := &externalAPIError{}
		if err := json.Unmarshal(response.Bytes(), upstreamErr); err != nil {
			return fmt.Errorf("unknown error - resource: %s - code: %d", URL, response.StatusCode)
		}

		return fmt.Errorf("%s - resource: %s - code: %d - error: %s - cause: %s", upstreamErr.Message, URL,
			response.StatusCode, upstreamErr.ErrorStr, upstreamErr.Cause)
	}

	if structure == nil {
		return nil
	}

	if unmarshalError := json.Unmarshal(response.Bytes(), structure); unmarshalError != nil {
		return unmarshalError
	}

	return nil
}
