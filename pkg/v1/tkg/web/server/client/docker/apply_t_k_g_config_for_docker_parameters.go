// Code generated by go-swagger; DO NOT EDIT.

package docker

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

	models "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/web/server/models"
)

// NewApplyTKGConfigForDockerParams creates a new ApplyTKGConfigForDockerParams object
// with the default values initialized.
func NewApplyTKGConfigForDockerParams() *ApplyTKGConfigForDockerParams {
	var ()
	return &ApplyTKGConfigForDockerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplyTKGConfigForDockerParamsWithTimeout creates a new ApplyTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplyTKGConfigForDockerParamsWithTimeout(timeout time.Duration) *ApplyTKGConfigForDockerParams {
	var ()
	return &ApplyTKGConfigForDockerParams{

		timeout: timeout,
	}
}

// NewApplyTKGConfigForDockerParamsWithContext creates a new ApplyTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplyTKGConfigForDockerParamsWithContext(ctx context.Context) *ApplyTKGConfigForDockerParams {
	var ()
	return &ApplyTKGConfigForDockerParams{

		Context: ctx,
	}
}

// NewApplyTKGConfigForDockerParamsWithHTTPClient creates a new ApplyTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplyTKGConfigForDockerParamsWithHTTPClient(client *http.Client) *ApplyTKGConfigForDockerParams {
	var ()
	return &ApplyTKGConfigForDockerParams{
		HTTPClient: client,
	}
}

/*ApplyTKGConfigForDockerParams contains all the parameters to send to the API endpoint
for the apply t k g config for docker operation typically these are written to a http.Request
*/
type ApplyTKGConfigForDockerParams struct {

	/*Params
	  parameters to apply changes to TKG configuration file for Docker

	*/
	Params *models.DockerRegionalClusterParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) WithTimeout(timeout time.Duration) *ApplyTKGConfigForDockerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) WithContext(ctx context.Context) *ApplyTKGConfigForDockerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) WithHTTPClient(client *http.Client) *ApplyTKGConfigForDockerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) WithParams(params *models.DockerRegionalClusterParams) *ApplyTKGConfigForDockerParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the apply t k g config for docker params
func (o *ApplyTKGConfigForDockerParams) SetParams(params *models.DockerRegionalClusterParams) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *ApplyTKGConfigForDockerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}