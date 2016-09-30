package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*PatchKapacitorsIDProxyNotFound Kapacitor ID does not exist.

swagger:response patchKapacitorsIdProxyNotFound
*/
type PatchKapacitorsIDProxyNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchKapacitorsIDProxyNotFound creates PatchKapacitorsIDProxyNotFound with default headers values
func NewPatchKapacitorsIDProxyNotFound() *PatchKapacitorsIDProxyNotFound {
	return &PatchKapacitorsIDProxyNotFound{}
}

// WithPayload adds the payload to the patch kapacitors Id proxy not found response
func (o *PatchKapacitorsIDProxyNotFound) WithPayload(payload *models.Error) *PatchKapacitorsIDProxyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch kapacitors Id proxy not found response
func (o *PatchKapacitorsIDProxyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchKapacitorsIDProxyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchKapacitorsIDProxyDefault Response directly from kapacitor

swagger:response patchKapacitorsIdProxyDefault
*/
type PatchKapacitorsIDProxyDefault struct {
	_statusCode int

	// In: body
	Payload models.KapacitorProxyResponse `json:"body,omitempty"`
}

// NewPatchKapacitorsIDProxyDefault creates PatchKapacitorsIDProxyDefault with default headers values
func NewPatchKapacitorsIDProxyDefault(code int) *PatchKapacitorsIDProxyDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchKapacitorsIDProxyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch kapacitors ID proxy default response
func (o *PatchKapacitorsIDProxyDefault) WithStatusCode(code int) *PatchKapacitorsIDProxyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch kapacitors ID proxy default response
func (o *PatchKapacitorsIDProxyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch kapacitors ID proxy default response
func (o *PatchKapacitorsIDProxyDefault) WithPayload(payload models.KapacitorProxyResponse) *PatchKapacitorsIDProxyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch kapacitors ID proxy default response
func (o *PatchKapacitorsIDProxyDefault) SetPayload(payload models.KapacitorProxyResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchKapacitorsIDProxyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
