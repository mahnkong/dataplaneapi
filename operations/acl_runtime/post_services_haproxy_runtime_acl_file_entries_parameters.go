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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v3/models"
)

// NewPostServicesHaproxyRuntimeACLFileEntriesParams creates a new PostServicesHaproxyRuntimeACLFileEntriesParams object
// no default values defined in spec.
func NewPostServicesHaproxyRuntimeACLFileEntriesParams() PostServicesHaproxyRuntimeACLFileEntriesParams {

	return PostServicesHaproxyRuntimeACLFileEntriesParams{}
}

// PostServicesHaproxyRuntimeACLFileEntriesParams contains all the bound params for the post services haproxy runtime ACL file entries operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostServicesHaproxyRuntimeACLFileEntries
type PostServicesHaproxyRuntimeACLFileEntriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ACL ID
	  Required: true
	  In: query
	*/
	ACLID string
	/*
	  Required: true
	  In: body
	*/
	Data *models.ACLFileEntry
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostServicesHaproxyRuntimeACLFileEntriesParams() beforehand.
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qACLID, qhkACLID, _ := qs.GetOK("acl_id")
	if err := o.bindACLID(qACLID, qhkACLID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ACLFileEntry
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindACLID binds and validates parameter ACLID from query.
func (o *PostServicesHaproxyRuntimeACLFileEntriesParams) bindACLID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("acl_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("acl_id", "query", raw); err != nil {
		return err
	}

	o.ACLID = raw

	return nil
}
