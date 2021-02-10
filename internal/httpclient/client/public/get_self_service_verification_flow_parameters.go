// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSelfServiceVerificationFlowParams creates a new GetSelfServiceVerificationFlowParams object
// with the default values initialized.
func NewGetSelfServiceVerificationFlowParams() *GetSelfServiceVerificationFlowParams {
	var ()
	return &GetSelfServiceVerificationFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSelfServiceVerificationFlowParamsWithTimeout creates a new GetSelfServiceVerificationFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSelfServiceVerificationFlowParamsWithTimeout(timeout time.Duration) *GetSelfServiceVerificationFlowParams {
	var ()
	return &GetSelfServiceVerificationFlowParams{

		timeout: timeout,
	}
}

// NewGetSelfServiceVerificationFlowParamsWithContext creates a new GetSelfServiceVerificationFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSelfServiceVerificationFlowParamsWithContext(ctx context.Context) *GetSelfServiceVerificationFlowParams {
	var ()
	return &GetSelfServiceVerificationFlowParams{

		Context: ctx,
	}
}

// NewGetSelfServiceVerificationFlowParamsWithHTTPClient creates a new GetSelfServiceVerificationFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSelfServiceVerificationFlowParamsWithHTTPClient(client *http.Client) *GetSelfServiceVerificationFlowParams {
	var ()
	return &GetSelfServiceVerificationFlowParams{
		HTTPClient: client,
	}
}

/*GetSelfServiceVerificationFlowParams contains all the parameters to send to the API endpoint
for the get self service verification flow operation typically these are written to a http.Request
*/
type GetSelfServiceVerificationFlowParams struct {

	/*ID
	  The Flow ID

	The value for this parameter comes from `request` URL Query parameter sent to your
	application (e.g. `/verification?flow=abcde`).

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) WithTimeout(timeout time.Duration) *GetSelfServiceVerificationFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) WithContext(ctx context.Context) *GetSelfServiceVerificationFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) WithHTTPClient(client *http.Client) *GetSelfServiceVerificationFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) WithID(id string) *GetSelfServiceVerificationFlowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get self service verification flow params
func (o *GetSelfServiceVerificationFlowParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSelfServiceVerificationFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
