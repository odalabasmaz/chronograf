package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetSourcesIDRolesOK An array of roles

swagger:response getSourcesIdRolesOK
*/
type GetSourcesIDRolesOK struct {

	// In: body
	Payload *models.Roles `json:"body,omitempty"`
}

// NewGetSourcesIDRolesOK creates GetSourcesIDRolesOK with default headers values
func NewGetSourcesIDRolesOK() *GetSourcesIDRolesOK {
	return &GetSourcesIDRolesOK{}
}

// WithPayload adds the payload to the get sources Id roles o k response
func (o *GetSourcesIDRolesOK) WithPayload(payload *models.Roles) *GetSourcesIDRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id roles o k response
func (o *GetSourcesIDRolesOK) SetPayload(payload *models.Roles) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDRolesDefault Unexpected internal service error

swagger:response getSourcesIdRolesDefault
*/
type GetSourcesIDRolesDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDRolesDefault creates GetSourcesIDRolesDefault with default headers values
func NewGetSourcesIDRolesDefault(code int) *GetSourcesIDRolesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesIDRolesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources ID roles default response
func (o *GetSourcesIDRolesDefault) WithStatusCode(code int) *GetSourcesIDRolesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources ID roles default response
func (o *GetSourcesIDRolesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources ID roles default response
func (o *GetSourcesIDRolesDefault) WithPayload(payload *models.Error) *GetSourcesIDRolesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources ID roles default response
func (o *GetSourcesIDRolesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDRolesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}