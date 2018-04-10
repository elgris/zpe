// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elgris/zpe/models"
)

// NewPostSpansParams creates a new PostSpansParams object
// with the default values initialized.
func NewPostSpansParams() *PostSpansParams {
	var ()
	return &PostSpansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpansParamsWithTimeout creates a new PostSpansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpansParamsWithTimeout(timeout time.Duration) *PostSpansParams {
	var ()
	return &PostSpansParams{

		timeout: timeout,
	}
}

// NewPostSpansParamsWithContext creates a new PostSpansParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpansParamsWithContext(ctx context.Context) *PostSpansParams {
	var ()
	return &PostSpansParams{

		Context: ctx,
	}
}

// NewPostSpansParamsWithHTTPClient creates a new PostSpansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpansParamsWithHTTPClient(client *http.Client) *PostSpansParams {
	var ()
	return &PostSpansParams{
		HTTPClient: client,
	}
}

/*PostSpansParams contains all the parameters to send to the API endpoint
for the post spans operation typically these are written to a http.Request
*/
type PostSpansParams struct {

	/*Spans
	  A list of spans that belong to any trace.

	*/
	Spans models.ListOfSpans

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post spans params
func (o *PostSpansParams) WithTimeout(timeout time.Duration) *PostSpansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post spans params
func (o *PostSpansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post spans params
func (o *PostSpansParams) WithContext(ctx context.Context) *PostSpansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post spans params
func (o *PostSpansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post spans params
func (o *PostSpansParams) WithHTTPClient(client *http.Client) *PostSpansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post spans params
func (o *PostSpansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSpans adds the spans to the post spans params
func (o *PostSpansParams) WithSpans(spans models.ListOfSpans) *PostSpansParams {
	o.SetSpans(spans)
	return o
}

// SetSpans adds the spans to the post spans params
func (o *PostSpansParams) SetSpans(spans models.ListOfSpans) {
	o.Spans = spans
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Spans != nil {
		if err := r.SetBodyParam(o.Spans); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
