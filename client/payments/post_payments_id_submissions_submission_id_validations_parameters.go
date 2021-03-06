// Code generated by go-swagger; DO NOT EDIT.

package payments

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

	"github.com/ewilde/go-form3/models"
)

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsParams creates a new PostPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized.
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsParams() *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithTimeout creates a new PostPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithContext creates a new PostPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithContext(ctx context.Context) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithHTTPClient creates a new PostPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDSubmissionsSubmissionIDValidationsParams contains all the parameters to send to the API endpoint
for the post payments ID submissions submission ID validations operation typically these are written to a http.Request
*/
type PostPaymentsIDSubmissionsSubmissionIDValidationsParams struct {

	/*PaymentSubmissionValidationRequest*/
	PaymentSubmissionValidationRequest *models.PaymentSubmissionValidationCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithContext(ctx context.Context) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentSubmissionValidationRequest adds the paymentSubmissionValidationRequest to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithPaymentSubmissionValidationRequest(paymentSubmissionValidationRequest *models.PaymentSubmissionValidationCreation) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetPaymentSubmissionValidationRequest(paymentSubmissionValidationRequest)
	return o
}

// SetPaymentSubmissionValidationRequest adds the paymentSubmissionValidationRequest to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetPaymentSubmissionValidationRequest(paymentSubmissionValidationRequest *models.PaymentSubmissionValidationCreation) {
	o.PaymentSubmissionValidationRequest = paymentSubmissionValidationRequest
}

// WithID adds the id to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithID(id strfmt.UUID) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSubmissionID adds the submissionID to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WithSubmissionID(submissionID strfmt.UUID) *PostPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the post payments ID submissions submission ID validations params
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentSubmissionValidationRequest != nil {
		if err := r.SetBodyParam(o.PaymentSubmissionValidationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
