// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v3/models"
)

// GetCachesHandlerFunc turns a function with the right signature into a get caches handler
type GetCachesHandlerFunc func(GetCachesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCachesHandlerFunc) Handle(params GetCachesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCachesHandler interface for that can handle valid get caches params
type GetCachesHandler interface {
	Handle(GetCachesParams, interface{}) middleware.Responder
}

// NewGetCaches creates a new http.Handler for the get caches operation
func NewGetCaches(ctx *middleware.Context, handler GetCachesHandler) *GetCaches {
	return &GetCaches{Context: ctx, Handler: handler}
}

/*GetCaches swagger:route GET /services/haproxy/configuration/caches Cache getCaches

Return an array of caches

Returns an array of all configured caches.

*/
type GetCaches struct {
	Context *middleware.Context
	Handler GetCachesHandler
}

func (o *GetCaches) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCachesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetCachesOKBody get caches o k body
//
// swagger:model GetCachesOKBody
type GetCachesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Caches `json:"data"`
}

// Validate validates this get caches o k body
func (o *GetCachesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCachesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getCachesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getCachesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCachesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCachesOKBody) UnmarshalBinary(b []byte) error {
	var res GetCachesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
