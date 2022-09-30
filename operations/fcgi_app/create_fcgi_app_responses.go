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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateFCGIAppCreatedCode is the HTTP code returned for type CreateFCGIAppCreated
const CreateFCGIAppCreatedCode int = 201

/*
CreateFCGIAppCreated Application created

swagger:response createFcgiAppCreated
*/
type CreateFCGIAppCreated struct {

	/*
	  In: Body
	*/
	Payload *models.FCGIApp `json:"body,omitempty"`
}

// NewCreateFCGIAppCreated creates CreateFCGIAppCreated with default headers values
func NewCreateFCGIAppCreated() *CreateFCGIAppCreated {

	return &CreateFCGIAppCreated{}
}

// WithPayload adds the payload to the create Fcgi app created response
func (o *CreateFCGIAppCreated) WithPayload(payload *models.FCGIApp) *CreateFCGIAppCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Fcgi app created response
func (o *CreateFCGIAppCreated) SetPayload(payload *models.FCGIApp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFCGIAppCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFCGIAppAcceptedCode is the HTTP code returned for type CreateFCGIAppAccepted
const CreateFCGIAppAcceptedCode int = 202

/*
CreateFCGIAppAccepted Configuration change accepted and reload requested

swagger:response createFcgiAppAccepted
*/
type CreateFCGIAppAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.FCGIApp `json:"body,omitempty"`
}

// NewCreateFCGIAppAccepted creates CreateFCGIAppAccepted with default headers values
func NewCreateFCGIAppAccepted() *CreateFCGIAppAccepted {

	return &CreateFCGIAppAccepted{}
}

// WithReloadID adds the reloadId to the create Fcgi app accepted response
func (o *CreateFCGIAppAccepted) WithReloadID(reloadID string) *CreateFCGIAppAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Fcgi app accepted response
func (o *CreateFCGIAppAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Fcgi app accepted response
func (o *CreateFCGIAppAccepted) WithPayload(payload *models.FCGIApp) *CreateFCGIAppAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Fcgi app accepted response
func (o *CreateFCGIAppAccepted) SetPayload(payload *models.FCGIApp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFCGIAppAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFCGIAppBadRequestCode is the HTTP code returned for type CreateFCGIAppBadRequest
const CreateFCGIAppBadRequestCode int = 400

/*
CreateFCGIAppBadRequest Bad request

swagger:response createFcgiAppBadRequest
*/
type CreateFCGIAppBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFCGIAppBadRequest creates CreateFCGIAppBadRequest with default headers values
func NewCreateFCGIAppBadRequest() *CreateFCGIAppBadRequest {

	return &CreateFCGIAppBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create Fcgi app bad request response
func (o *CreateFCGIAppBadRequest) WithConfigurationVersion(configurationVersion string) *CreateFCGIAppBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Fcgi app bad request response
func (o *CreateFCGIAppBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Fcgi app bad request response
func (o *CreateFCGIAppBadRequest) WithPayload(payload *models.Error) *CreateFCGIAppBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Fcgi app bad request response
func (o *CreateFCGIAppBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFCGIAppBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateFCGIAppConflictCode is the HTTP code returned for type CreateFCGIAppConflict
const CreateFCGIAppConflictCode int = 409

/*
CreateFCGIAppConflict The specified resource already exists

swagger:response createFcgiAppConflict
*/
type CreateFCGIAppConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFCGIAppConflict creates CreateFCGIAppConflict with default headers values
func NewCreateFCGIAppConflict() *CreateFCGIAppConflict {

	return &CreateFCGIAppConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create Fcgi app conflict response
func (o *CreateFCGIAppConflict) WithConfigurationVersion(configurationVersion string) *CreateFCGIAppConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Fcgi app conflict response
func (o *CreateFCGIAppConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Fcgi app conflict response
func (o *CreateFCGIAppConflict) WithPayload(payload *models.Error) *CreateFCGIAppConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Fcgi app conflict response
func (o *CreateFCGIAppConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFCGIAppConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*
CreateFCGIAppDefault General Error

swagger:response createFcgiAppDefault
*/
type CreateFCGIAppDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFCGIAppDefault creates CreateFCGIAppDefault with default headers values
func NewCreateFCGIAppDefault(code int) *CreateFCGIAppDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateFCGIAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create FCGI app default response
func (o *CreateFCGIAppDefault) WithStatusCode(code int) *CreateFCGIAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create FCGI app default response
func (o *CreateFCGIAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create FCGI app default response
func (o *CreateFCGIAppDefault) WithConfigurationVersion(configurationVersion string) *CreateFCGIAppDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create FCGI app default response
func (o *CreateFCGIAppDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create FCGI app default response
func (o *CreateFCGIAppDefault) WithPayload(payload *models.Error) *CreateFCGIAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create FCGI app default response
func (o *CreateFCGIAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFCGIAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
