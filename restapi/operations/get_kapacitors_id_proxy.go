package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetKapacitorsIDProxyHandlerFunc turns a function with the right signature into a get kapacitors ID proxy handler
type GetKapacitorsIDProxyHandlerFunc func(context.Context, GetKapacitorsIDProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetKapacitorsIDProxyHandlerFunc) Handle(ctx context.Context, params GetKapacitorsIDProxyParams) middleware.Responder {
	return fn(ctx, params)
}

// GetKapacitorsIDProxyHandler interface for that can handle valid get kapacitors ID proxy params
type GetKapacitorsIDProxyHandler interface {
	Handle(context.Context, GetKapacitorsIDProxyParams) middleware.Responder
}

// NewGetKapacitorsIDProxy creates a new http.Handler for the get kapacitors ID proxy operation
func NewGetKapacitorsIDProxy(ctx *middleware.Context, handler GetKapacitorsIDProxyHandler) *GetKapacitorsIDProxy {
	return &GetKapacitorsIDProxy{Context: ctx, Handler: handler}
}

/*GetKapacitorsIDProxy swagger:route GET /kapacitors/{id}/proxy getKapacitorsIdProxy

GET to `path` of kapacitor.  The response and status code from kapacitor is directly returned.

*/
type GetKapacitorsIDProxy struct {
	Context *middleware.Context
	Handler GetKapacitorsIDProxyHandler
}

func (o *GetKapacitorsIDProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetKapacitorsIDProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
