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

package fcgi_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateFCGIAppHandlerFunc turns a function with the right signature into a create FCGI app handler
type CreateFCGIAppHandlerFunc func(CreateFCGIAppParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateFCGIAppHandlerFunc) Handle(params CreateFCGIAppParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateFCGIAppHandler interface for that can handle valid create FCGI app params
type CreateFCGIAppHandler interface {
	Handle(CreateFCGIAppParams, interface{}) middleware.Responder
}

// NewCreateFCGIApp creates a new http.Handler for the create FCGI app operation
func NewCreateFCGIApp(ctx *middleware.Context, handler CreateFCGIAppHandler) *CreateFCGIApp {
	return &CreateFCGIApp{Context: ctx, Handler: handler}
}

/*
	CreateFCGIApp swagger:route POST /services/haproxy/configuration/fcgi_apps FCGIApp createFcgiApp

# Add an FCGI app

Adds a new FCGI application to the configuration file.
*/
type CreateFCGIApp struct {
	Context *middleware.Context
	Handler CreateFCGIAppHandler
}

func (o *CreateFCGIApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateFCGIAppParams()
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
