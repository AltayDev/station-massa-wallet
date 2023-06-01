// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRestAccountGetParams creates a new RestAccountGetParams object
// with the default values initialized.
func NewRestAccountGetParams() RestAccountGetParams {

	var (
		// initialize parameters with default values

		cipheredDefault = bool(true)
	)

	return RestAccountGetParams{
		Ciphered: &cipheredDefault,
	}
}

// RestAccountGetParams contains all the bound params for the rest account get operation
// typically these are obtained from a http.Request
//
// swagger:parameters restAccountGet
type RestAccountGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*whether to return the data ciphered or not
	  In: query
	  Default: true
	*/
	Ciphered *bool
	/*Account's short name.
	  Required: true
	  In: path
	*/
	Nickname string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRestAccountGetParams() beforehand.
func (o *RestAccountGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCiphered, qhkCiphered, _ := qs.GetOK("ciphered")
	if err := o.bindCiphered(qCiphered, qhkCiphered, route.Formats); err != nil {
		res = append(res, err)
	}

	rNickname, rhkNickname, _ := route.Params.GetOK("nickname")
	if err := o.bindNickname(rNickname, rhkNickname, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCiphered binds and validates parameter Ciphered from query.
func (o *RestAccountGetParams) bindCiphered(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewRestAccountGetParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("ciphered", "query", "bool", raw)
	}
	o.Ciphered = &value

	return nil
}

// bindNickname binds and validates parameter Nickname from path.
func (o *RestAccountGetParams) bindNickname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Nickname = raw

	return nil
}