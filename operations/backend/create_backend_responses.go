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

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateBackendCreatedCode is the HTTP code returned for type CreateBackendCreated
const CreateBackendCreatedCode int = 201

/*CreateBackendCreated Backend created

swagger:response createBackendCreated
*/
type CreateBackendCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Backend `json:"body,omitempty"`
}

// NewCreateBackendCreated creates CreateBackendCreated with default headers values
func NewCreateBackendCreated() *CreateBackendCreated {

	return &CreateBackendCreated{}
}

// WithPayload adds the payload to the create backend created response
func (o *CreateBackendCreated) WithPayload(payload *models.Backend) *CreateBackendCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend created response
func (o *CreateBackendCreated) SetPayload(payload *models.Backend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBackendAcceptedCode is the HTTP code returned for type CreateBackendAccepted
const CreateBackendAcceptedCode int = 202

/*CreateBackendAccepted Configuration change accepted and reload requested

swagger:response createBackendAccepted
*/
type CreateBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Backend `json:"body,omitempty"`
}

// NewCreateBackendAccepted creates CreateBackendAccepted with default headers values
func NewCreateBackendAccepted() *CreateBackendAccepted {

	return &CreateBackendAccepted{}
}

// WithReloadID adds the reloadId to the create backend accepted response
func (o *CreateBackendAccepted) WithReloadID(reloadID string) *CreateBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create backend accepted response
func (o *CreateBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create backend accepted response
func (o *CreateBackendAccepted) WithPayload(payload *models.Backend) *CreateBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend accepted response
func (o *CreateBackendAccepted) SetPayload(payload *models.Backend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateBackendBadRequestCode is the HTTP code returned for type CreateBackendBadRequest
const CreateBackendBadRequestCode int = 400

/*CreateBackendBadRequest Bad request

swagger:response createBackendBadRequest
*/
type CreateBackendBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendBadRequest creates CreateBackendBadRequest with default headers values
func NewCreateBackendBadRequest() *CreateBackendBadRequest {

	return &CreateBackendBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create backend bad request response
func (o *CreateBackendBadRequest) WithConfigurationVersion(configurationVersion string) *CreateBackendBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend bad request response
func (o *CreateBackendBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend bad request response
func (o *CreateBackendBadRequest) WithPayload(payload *models.Error) *CreateBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend bad request response
func (o *CreateBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateBackendConflictCode is the HTTP code returned for type CreateBackendConflict
const CreateBackendConflictCode int = 409

/*CreateBackendConflict The specified resource already exists

swagger:response createBackendConflict
*/
type CreateBackendConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendConflict creates CreateBackendConflict with default headers values
func NewCreateBackendConflict() *CreateBackendConflict {

	return &CreateBackendConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create backend conflict response
func (o *CreateBackendConflict) WithConfigurationVersion(configurationVersion string) *CreateBackendConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend conflict response
func (o *CreateBackendConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend conflict response
func (o *CreateBackendConflict) WithPayload(payload *models.Error) *CreateBackendConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend conflict response
func (o *CreateBackendConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateBackendDefault General Error

swagger:response createBackendDefault
*/
type CreateBackendDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBackendDefault creates CreateBackendDefault with default headers values
func NewCreateBackendDefault(code int) *CreateBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create backend default response
func (o *CreateBackendDefault) WithStatusCode(code int) *CreateBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create backend default response
func (o *CreateBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create backend default response
func (o *CreateBackendDefault) WithConfigurationVersion(configurationVersion string) *CreateBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create backend default response
func (o *CreateBackendDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create backend default response
func (o *CreateBackendDefault) WithPayload(payload *models.Error) *CreateBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create backend default response
func (o *CreateBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
