// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// NewPutMmdsConfigParams creates a new PutMmdsConfigParams object
// with the default values initialized.
func NewPutMmdsConfigParams() *PutMmdsConfigParams {
	var ()
	return &PutMmdsConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMmdsConfigParamsWithTimeout creates a new PutMmdsConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMmdsConfigParamsWithTimeout(timeout time.Duration) *PutMmdsConfigParams {
	var ()
	return &PutMmdsConfigParams{

		timeout: timeout,
	}
}

// NewPutMmdsConfigParamsWithContext creates a new PutMmdsConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMmdsConfigParamsWithContext(ctx context.Context) *PutMmdsConfigParams {
	var ()
	return &PutMmdsConfigParams{

		Context: ctx,
	}
}

// NewPutMmdsConfigParamsWithHTTPClient creates a new PutMmdsConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMmdsConfigParamsWithHTTPClient(client *http.Client) *PutMmdsConfigParams {
	var ()
	return &PutMmdsConfigParams{
		HTTPClient: client,
	}
}

/*PutMmdsConfigParams contains all the parameters to send to the API endpoint
for the put mmds config operation typically these are written to a http.Request
*/
type PutMmdsConfigParams struct {

	/*Body
	  The MMDS configuration as JSON.

	*/
	Body *models.MmdsConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put mmds config params
func (o *PutMmdsConfigParams) WithTimeout(timeout time.Duration) *PutMmdsConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put mmds config params
func (o *PutMmdsConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put mmds config params
func (o *PutMmdsConfigParams) WithContext(ctx context.Context) *PutMmdsConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put mmds config params
func (o *PutMmdsConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put mmds config params
func (o *PutMmdsConfigParams) WithHTTPClient(client *http.Client) *PutMmdsConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put mmds config params
func (o *PutMmdsConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put mmds config params
func (o *PutMmdsConfigParams) WithBody(body *models.MmdsConfig) *PutMmdsConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put mmds config params
func (o *PutMmdsConfigParams) SetBody(body *models.MmdsConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutMmdsConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
