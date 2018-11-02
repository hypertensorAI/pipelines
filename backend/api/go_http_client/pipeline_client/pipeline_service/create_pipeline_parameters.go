// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

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

	pipeline_model "github.com/kubeflow/pipelines/backend/api/go_http_client/pipeline_model"
)

// NewCreatePipelineParams creates a new CreatePipelineParams object
// with the default values initialized.
func NewCreatePipelineParams() *CreatePipelineParams {
	var ()
	return &CreatePipelineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePipelineParamsWithTimeout creates a new CreatePipelineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePipelineParamsWithTimeout(timeout time.Duration) *CreatePipelineParams {
	var ()
	return &CreatePipelineParams{

		timeout: timeout,
	}
}

// NewCreatePipelineParamsWithContext creates a new CreatePipelineParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePipelineParamsWithContext(ctx context.Context) *CreatePipelineParams {
	var ()
	return &CreatePipelineParams{

		Context: ctx,
	}
}

// NewCreatePipelineParamsWithHTTPClient creates a new CreatePipelineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePipelineParamsWithHTTPClient(client *http.Client) *CreatePipelineParams {
	var ()
	return &CreatePipelineParams{
		HTTPClient: client,
	}
}

/*CreatePipelineParams contains all the parameters to send to the API endpoint
for the create pipeline operation typically these are written to a http.Request
*/
type CreatePipelineParams struct {

	/*Body*/
	Body *pipeline_model.APIURL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) WithTimeout(timeout time.Duration) *CreatePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pipeline params
func (o *CreatePipelineParams) WithContext(ctx context.Context) *CreatePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pipeline params
func (o *CreatePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) WithHTTPClient(client *http.Client) *CreatePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pipeline params
func (o *CreatePipelineParams) WithBody(body *pipeline_model.APIURL) *CreatePipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pipeline params
func (o *CreatePipelineParams) SetBody(body *pipeline_model.APIURL) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
