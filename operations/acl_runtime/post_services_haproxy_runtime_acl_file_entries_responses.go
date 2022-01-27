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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// PostServicesHaproxyRuntimeACLFileEntriesCreatedCode is the HTTP code returned for type PostServicesHaproxyRuntimeACLFileEntriesCreated
const PostServicesHaproxyRuntimeACLFileEntriesCreatedCode int = 201

/*PostServicesHaproxyRuntimeACLFileEntriesCreated ACL entry created

swagger:response postServicesHaproxyRuntimeAclFileEntriesCreated
*/
type PostServicesHaproxyRuntimeACLFileEntriesCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ACLFileEntry `json:"body,omitempty"`
}

// NewPostServicesHaproxyRuntimeACLFileEntriesCreated creates PostServicesHaproxyRuntimeACLFileEntriesCreated with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesCreated() *PostServicesHaproxyRuntimeACLFileEntriesCreated {

	return &PostServicesHaproxyRuntimeACLFileEntriesCreated{}
}

// WithPayload adds the payload to the post services haproxy runtime Acl file entries created response
func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) WithPayload(payload *models.ACLFileEntry) *PostServicesHaproxyRuntimeACLFileEntriesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post services haproxy runtime Acl file entries created response
func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) SetPayload(payload *models.ACLFileEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostServicesHaproxyRuntimeACLFileEntriesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostServicesHaproxyRuntimeACLFileEntriesBadRequestCode is the HTTP code returned for type PostServicesHaproxyRuntimeACLFileEntriesBadRequest
const PostServicesHaproxyRuntimeACLFileEntriesBadRequestCode int = 400

/*PostServicesHaproxyRuntimeACLFileEntriesBadRequest Bad request

swagger:response postServicesHaproxyRuntimeAclFileEntriesBadRequest
*/
type PostServicesHaproxyRuntimeACLFileEntriesBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostServicesHaproxyRuntimeACLFileEntriesBadRequest creates PostServicesHaproxyRuntimeACLFileEntriesBadRequest with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesBadRequest() *PostServicesHaproxyRuntimeACLFileEntriesBadRequest {

	return &PostServicesHaproxyRuntimeACLFileEntriesBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the post services haproxy runtime Acl file entries bad request response
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) WithConfigurationVersion(configurationVersion string) *PostServicesHaproxyRuntimeACLFileEntriesBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the post services haproxy runtime Acl file entries bad request response
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the post services haproxy runtime Acl file entries bad request response
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) WithPayload(payload *models.Error) *PostServicesHaproxyRuntimeACLFileEntriesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post services haproxy runtime Acl file entries bad request response
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostServicesHaproxyRuntimeACLFileEntriesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostServicesHaproxyRuntimeACLFileEntriesConflictCode is the HTTP code returned for type PostServicesHaproxyRuntimeACLFileEntriesConflict
const PostServicesHaproxyRuntimeACLFileEntriesConflictCode int = 409

/*PostServicesHaproxyRuntimeACLFileEntriesConflict The specified resource already exists

swagger:response postServicesHaproxyRuntimeAclFileEntriesConflict
*/
type PostServicesHaproxyRuntimeACLFileEntriesConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostServicesHaproxyRuntimeACLFileEntriesConflict creates PostServicesHaproxyRuntimeACLFileEntriesConflict with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesConflict() *PostServicesHaproxyRuntimeACLFileEntriesConflict {

	return &PostServicesHaproxyRuntimeACLFileEntriesConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the post services haproxy runtime Acl file entries conflict response
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) WithConfigurationVersion(configurationVersion string) *PostServicesHaproxyRuntimeACLFileEntriesConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the post services haproxy runtime Acl file entries conflict response
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the post services haproxy runtime Acl file entries conflict response
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) WithPayload(payload *models.Error) *PostServicesHaproxyRuntimeACLFileEntriesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post services haproxy runtime Acl file entries conflict response
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostServicesHaproxyRuntimeACLFileEntriesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostServicesHaproxyRuntimeACLFileEntriesDefault General Error

swagger:response postServicesHaproxyRuntimeAclFileEntriesDefault
*/
type PostServicesHaproxyRuntimeACLFileEntriesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostServicesHaproxyRuntimeACLFileEntriesDefault creates PostServicesHaproxyRuntimeACLFileEntriesDefault with default headers values
func NewPostServicesHaproxyRuntimeACLFileEntriesDefault(code int) *PostServicesHaproxyRuntimeACLFileEntriesDefault {
	if code <= 0 {
		code = 500
	}

	return &PostServicesHaproxyRuntimeACLFileEntriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) WithStatusCode(code int) *PostServicesHaproxyRuntimeACLFileEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) WithConfigurationVersion(configurationVersion string) *PostServicesHaproxyRuntimeACLFileEntriesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) WithPayload(payload *models.Error) *PostServicesHaproxyRuntimeACLFileEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post services haproxy runtime ACL file entries default response
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostServicesHaproxyRuntimeACLFileEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
