// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ImportAccountHandlerFunc turns a function with the right signature into a import account handler
type ImportAccountHandlerFunc func(ImportAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImportAccountHandlerFunc) Handle(params ImportAccountParams) middleware.Responder {
	return fn(params)
}

// ImportAccountHandler interface for that can handle valid import account params
type ImportAccountHandler interface {
	Handle(ImportAccountParams) middleware.Responder
}

// NewImportAccount creates a new http.Handler for the import account operation
func NewImportAccount(ctx *middleware.Context, handler ImportAccountHandler) *ImportAccount {
	return &ImportAccount{Context: ctx, Handler: handler}
}

/*
	ImportAccount swagger:route PUT /api/accounts importAccount

Import a new account
*/
type ImportAccount struct {
	Context *middleware.Context
	Handler ImportAccountHandler
}

func (o *ImportAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImportAccountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}