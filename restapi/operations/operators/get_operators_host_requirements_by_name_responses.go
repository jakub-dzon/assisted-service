// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetOperatorsHostRequirementsByNameOKCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameOK
const GetOperatorsHostRequirementsByNameOKCode int = 200

/*GetOperatorsHostRequirementsByNameOK Success.

swagger:response getOperatorsHostRequirementsByNameOK
*/
type GetOperatorsHostRequirementsByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.OperatorHostRequirements `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameOK creates GetOperatorsHostRequirementsByNameOK with default headers values
func NewGetOperatorsHostRequirementsByNameOK() *GetOperatorsHostRequirementsByNameOK {

	return &GetOperatorsHostRequirementsByNameOK{}
}

// WithPayload adds the payload to the get operators host requirements by name o k response
func (o *GetOperatorsHostRequirementsByNameOK) WithPayload(payload *models.OperatorHostRequirements) *GetOperatorsHostRequirementsByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name o k response
func (o *GetOperatorsHostRequirementsByNameOK) SetPayload(payload *models.OperatorHostRequirements) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOperatorsHostRequirementsByNameUnauthorizedCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameUnauthorized
const GetOperatorsHostRequirementsByNameUnauthorizedCode int = 401

/*GetOperatorsHostRequirementsByNameUnauthorized Unauthorized.

swagger:response getOperatorsHostRequirementsByNameUnauthorized
*/
type GetOperatorsHostRequirementsByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameUnauthorized creates GetOperatorsHostRequirementsByNameUnauthorized with default headers values
func NewGetOperatorsHostRequirementsByNameUnauthorized() *GetOperatorsHostRequirementsByNameUnauthorized {

	return &GetOperatorsHostRequirementsByNameUnauthorized{}
}

// WithPayload adds the payload to the get operators host requirements by name unauthorized response
func (o *GetOperatorsHostRequirementsByNameUnauthorized) WithPayload(payload *models.InfraError) *GetOperatorsHostRequirementsByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name unauthorized response
func (o *GetOperatorsHostRequirementsByNameUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOperatorsHostRequirementsByNameForbiddenCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameForbidden
const GetOperatorsHostRequirementsByNameForbiddenCode int = 403

/*GetOperatorsHostRequirementsByNameForbidden Forbidden.

swagger:response getOperatorsHostRequirementsByNameForbidden
*/
type GetOperatorsHostRequirementsByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameForbidden creates GetOperatorsHostRequirementsByNameForbidden with default headers values
func NewGetOperatorsHostRequirementsByNameForbidden() *GetOperatorsHostRequirementsByNameForbidden {

	return &GetOperatorsHostRequirementsByNameForbidden{}
}

// WithPayload adds the payload to the get operators host requirements by name forbidden response
func (o *GetOperatorsHostRequirementsByNameForbidden) WithPayload(payload *models.InfraError) *GetOperatorsHostRequirementsByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name forbidden response
func (o *GetOperatorsHostRequirementsByNameForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOperatorsHostRequirementsByNameNotFoundCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameNotFound
const GetOperatorsHostRequirementsByNameNotFoundCode int = 404

/*GetOperatorsHostRequirementsByNameNotFound Not found.

swagger:response getOperatorsHostRequirementsByNameNotFound
*/
type GetOperatorsHostRequirementsByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameNotFound creates GetOperatorsHostRequirementsByNameNotFound with default headers values
func NewGetOperatorsHostRequirementsByNameNotFound() *GetOperatorsHostRequirementsByNameNotFound {

	return &GetOperatorsHostRequirementsByNameNotFound{}
}

// WithPayload adds the payload to the get operators host requirements by name not found response
func (o *GetOperatorsHostRequirementsByNameNotFound) WithPayload(payload *models.Error) *GetOperatorsHostRequirementsByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name not found response
func (o *GetOperatorsHostRequirementsByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOperatorsHostRequirementsByNameMethodNotAllowedCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameMethodNotAllowed
const GetOperatorsHostRequirementsByNameMethodNotAllowedCode int = 405

/*GetOperatorsHostRequirementsByNameMethodNotAllowed Method Not Allowed.

swagger:response getOperatorsHostRequirementsByNameMethodNotAllowed
*/
type GetOperatorsHostRequirementsByNameMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameMethodNotAllowed creates GetOperatorsHostRequirementsByNameMethodNotAllowed with default headers values
func NewGetOperatorsHostRequirementsByNameMethodNotAllowed() *GetOperatorsHostRequirementsByNameMethodNotAllowed {

	return &GetOperatorsHostRequirementsByNameMethodNotAllowed{}
}

// WithPayload adds the payload to the get operators host requirements by name method not allowed response
func (o *GetOperatorsHostRequirementsByNameMethodNotAllowed) WithPayload(payload *models.Error) *GetOperatorsHostRequirementsByNameMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name method not allowed response
func (o *GetOperatorsHostRequirementsByNameMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOperatorsHostRequirementsByNameInternalServerErrorCode is the HTTP code returned for type GetOperatorsHostRequirementsByNameInternalServerError
const GetOperatorsHostRequirementsByNameInternalServerErrorCode int = 500

/*GetOperatorsHostRequirementsByNameInternalServerError Error.

swagger:response getOperatorsHostRequirementsByNameInternalServerError
*/
type GetOperatorsHostRequirementsByNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOperatorsHostRequirementsByNameInternalServerError creates GetOperatorsHostRequirementsByNameInternalServerError with default headers values
func NewGetOperatorsHostRequirementsByNameInternalServerError() *GetOperatorsHostRequirementsByNameInternalServerError {

	return &GetOperatorsHostRequirementsByNameInternalServerError{}
}

// WithPayload adds the payload to the get operators host requirements by name internal server error response
func (o *GetOperatorsHostRequirementsByNameInternalServerError) WithPayload(payload *models.Error) *GetOperatorsHostRequirementsByNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get operators host requirements by name internal server error response
func (o *GetOperatorsHostRequirementsByNameInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOperatorsHostRequirementsByNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
