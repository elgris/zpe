// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDependencies Returns service links derived from spans.

*/
func (a *Client) GetDependencies(params *GetDependenciesParams) (*GetDependenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDependenciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDependencies",
		Method:             "GET",
		PathPattern:        "/dependencies",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDependenciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDependenciesOK), nil

}

/*
GetServices Returns a list of all service names associated with span endpoints.

*/
func (a *Client) GetServices(params *GetServicesParams) (*GetServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServices",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServicesOK), nil

}

/*
GetSpans Get all the span names recorded by a particular service
*/
func (a *Client) GetSpans(params *GetSpansParams) (*GetSpansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSpansParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSpans",
		Method:             "GET",
		PathPattern:        "/spans",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSpansReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSpansOK), nil

}

/*
GetTraceTraceID get trace trace ID API
*/
func (a *Client) GetTraceTraceID(params *GetTraceTraceIDParams) (*GetTraceTraceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTraceTraceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTraceTraceID",
		Method:             "GET",
		PathPattern:        "/trace/{traceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTraceTraceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTraceTraceIDOK), nil

}

/*
GetTraces Invoking this request retrieves traces matching the below filters.

Results should be filtered against endTs, subject to limit and
lookback. For example, if endTs is 10:20 today, limit is 10, and
lookback is 7 days, traces returned should be those nearest to 10:20
today, not 10:20 a week ago.

Time units of endTs and lookback are milliseconds as opposed to
microseconds, the grain of Span.timestamp. Milliseconds is a more
familiar and supported granularity for query, index and windowing
functions

*/
func (a *Client) GetTraces(params *GetTracesParams) (*GetTracesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTracesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTraces",
		Method:             "GET",
		PathPattern:        "/traces",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTracesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTracesOK), nil

}

/*
PostSpans Uploads a list of spans encoded per content-type, for example json.

*/
func (a *Client) PostSpans(params *PostSpansParams) (*PostSpansAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSpansParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSpans",
		Method:             "POST",
		PathPattern:        "/spans",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostSpansReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSpansAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
