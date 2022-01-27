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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// GetACLHandlerFunc turns a function with the right signature into a get Acl handler
type GetACLHandlerFunc func(GetACLParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetACLHandlerFunc) Handle(params GetACLParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetACLHandler interface for that can handle valid get Acl params
type GetACLHandler interface {
	Handle(GetACLParams, interface{}) middleware.Responder
}

// NewGetACL creates a new http.Handler for the get Acl operation
func NewGetACL(ctx *middleware.Context, handler GetACLHandler) *GetACL {
	return &GetACL{Context: ctx, Handler: handler}
}

/*GetACL swagger:route GET /services/haproxy/configuration/acls/{index} ACL getAcl

Return one ACL line

Returns one ACL line configuration by it's index in the specified parent.

*/
type GetACL struct {
	Context *middleware.Context
	Handler GetACLHandler
}

func (o *GetACL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetACLParams()

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

// GetACLOKBody get ACL o k body
//
// swagger:model GetACLOKBody
type GetACLOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.ACL `json:"data,omitempty"`
}

// Validate validates this get ACL o k body
func (o *GetACLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetACLOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAclOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetACLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetACLOKBody) UnmarshalBinary(b []byte) error {
	var res GetACLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
