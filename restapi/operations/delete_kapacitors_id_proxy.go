package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteKapacitorsIDProxyHandlerFunc turns a function with the right signature into a delete kapacitors ID proxy handler
type DeleteKapacitorsIDProxyHandlerFunc func(context.Context, DeleteKapacitorsIDProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteKapacitorsIDProxyHandlerFunc) Handle(ctx context.Context, params DeleteKapacitorsIDProxyParams) middleware.Responder {
	return fn(ctx, params)
}

// DeleteKapacitorsIDProxyHandler interface for that can handle valid delete kapacitors ID proxy params
type DeleteKapacitorsIDProxyHandler interface {
	Handle(context.Context, DeleteKapacitorsIDProxyParams) middleware.Responder
}

// NewDeleteKapacitorsIDProxy creates a new http.Handler for the delete kapacitors ID proxy operation
func NewDeleteKapacitorsIDProxy(ctx *middleware.Context, handler DeleteKapacitorsIDProxyHandler) *DeleteKapacitorsIDProxy {
	return &DeleteKapacitorsIDProxy{Context: ctx, Handler: handler}
}

/*DeleteKapacitorsIDProxy swagger:route DELETE /kapacitors/{id}/proxy deleteKapacitorsIdProxy

DELETE to `path` of kapacitor.  The response and status code from kapacitor is directly returned.

*/
type DeleteKapacitorsIDProxy struct {
	Context *middleware.Context
	Handler DeleteKapacitorsIDProxyHandler
}

func (o *DeleteKapacitorsIDProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteKapacitorsIDProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
