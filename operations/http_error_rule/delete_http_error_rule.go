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

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteHTTPErrorRuleHandlerFunc turns a function with the right signature into a delete HTTP error rule handler
type DeleteHTTPErrorRuleHandlerFunc func(DeleteHTTPErrorRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHTTPErrorRuleHandlerFunc) Handle(params DeleteHTTPErrorRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteHTTPErrorRuleHandler interface for that can handle valid delete HTTP error rule params
type DeleteHTTPErrorRuleHandler interface {
	Handle(DeleteHTTPErrorRuleParams, interface{}) middleware.Responder
}

// NewDeleteHTTPErrorRule creates a new http.Handler for the delete HTTP error rule operation
func NewDeleteHTTPErrorRule(ctx *middleware.Context, handler DeleteHTTPErrorRuleHandler) *DeleteHTTPErrorRule {
	return &DeleteHTTPErrorRule{Context: ctx, Handler: handler}
}

/*
	DeleteHTTPErrorRule swagger:route DELETE /services/haproxy/configuration/http_error_rules/{index} HTTPErrorRule deleteHttpErrorRule

# Delete a HTTP Error Rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.
*/
type DeleteHTTPErrorRule struct {
	Context *middleware.Context
	Handler DeleteHTTPErrorRuleHandler
}

func (o *DeleteHTTPErrorRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteHTTPErrorRuleParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
