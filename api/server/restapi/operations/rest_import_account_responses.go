// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra-plugin-wallet/api/server/models"
)

// RestImportAccountOKCode is the HTTP code returned for type RestImportAccountOK
const RestImportAccountOKCode int = 200

/*
RestImportAccountOK Account imported.

swagger:response restImportAccountOK
*/
type RestImportAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewRestImportAccountOK creates RestImportAccountOK with default headers values
func NewRestImportAccountOK() *RestImportAccountOK {

	return &RestImportAccountOK{}
}

// WithPayload adds the payload to the rest import account o k response
func (o *RestImportAccountOK) WithPayload(payload *models.Account) *RestImportAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest import account o k response
func (o *RestImportAccountOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestImportAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestImportAccountBadRequestCode is the HTTP code returned for type RestImportAccountBadRequest
const RestImportAccountBadRequestCode int = 400

/*
RestImportAccountBadRequest Bad request.

swagger:response restImportAccountBadRequest
*/
type RestImportAccountBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestImportAccountBadRequest creates RestImportAccountBadRequest with default headers values
func NewRestImportAccountBadRequest() *RestImportAccountBadRequest {

	return &RestImportAccountBadRequest{}
}

// WithPayload adds the payload to the rest import account bad request response
func (o *RestImportAccountBadRequest) WithPayload(payload *models.Error) *RestImportAccountBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest import account bad request response
func (o *RestImportAccountBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestImportAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestImportAccountUnauthorizedCode is the HTTP code returned for type RestImportAccountUnauthorized
const RestImportAccountUnauthorizedCode int = 401

/*
RestImportAccountUnauthorized Unauthorized - The request requires user authentication. Only if no correlationId is provided.

swagger:response restImportAccountUnauthorized
*/
type RestImportAccountUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestImportAccountUnauthorized creates RestImportAccountUnauthorized with default headers values
func NewRestImportAccountUnauthorized() *RestImportAccountUnauthorized {

	return &RestImportAccountUnauthorized{}
}

// WithPayload adds the payload to the rest import account unauthorized response
func (o *RestImportAccountUnauthorized) WithPayload(payload *models.Error) *RestImportAccountUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest import account unauthorized response
func (o *RestImportAccountUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestImportAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestImportAccountUnprocessableEntityCode is the HTTP code returned for type RestImportAccountUnprocessableEntity
const RestImportAccountUnprocessableEntityCode int = 422

/*
RestImportAccountUnprocessableEntity Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.

swagger:response restImportAccountUnprocessableEntity
*/
type RestImportAccountUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestImportAccountUnprocessableEntity creates RestImportAccountUnprocessableEntity with default headers values
func NewRestImportAccountUnprocessableEntity() *RestImportAccountUnprocessableEntity {

	return &RestImportAccountUnprocessableEntity{}
}

// WithPayload adds the payload to the rest import account unprocessable entity response
func (o *RestImportAccountUnprocessableEntity) WithPayload(payload *models.Error) *RestImportAccountUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest import account unprocessable entity response
func (o *RestImportAccountUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestImportAccountUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestImportAccountInternalServerErrorCode is the HTTP code returned for type RestImportAccountInternalServerError
const RestImportAccountInternalServerErrorCode int = 500

/*
RestImportAccountInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response restImportAccountInternalServerError
*/
type RestImportAccountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestImportAccountInternalServerError creates RestImportAccountInternalServerError with default headers values
func NewRestImportAccountInternalServerError() *RestImportAccountInternalServerError {

	return &RestImportAccountInternalServerError{}
}

// WithPayload adds the payload to the rest import account internal server error response
func (o *RestImportAccountInternalServerError) WithPayload(payload *models.Error) *RestImportAccountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest import account internal server error response
func (o *RestImportAccountInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestImportAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}