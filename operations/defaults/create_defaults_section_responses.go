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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateDefaultsSectionCreatedCode is the HTTP code returned for type CreateDefaultsSectionCreated
const CreateDefaultsSectionCreatedCode int = 201

/*
CreateDefaultsSectionCreated Defaults created

swagger:response createDefaultsSectionCreated
*/
type CreateDefaultsSectionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewCreateDefaultsSectionCreated creates CreateDefaultsSectionCreated with default headers values
func NewCreateDefaultsSectionCreated() *CreateDefaultsSectionCreated {

	return &CreateDefaultsSectionCreated{}
}

// WithPayload adds the payload to the create defaults section created response
func (o *CreateDefaultsSectionCreated) WithPayload(payload *models.Defaults) *CreateDefaultsSectionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create defaults section created response
func (o *CreateDefaultsSectionCreated) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefaultsSectionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDefaultsSectionAcceptedCode is the HTTP code returned for type CreateDefaultsSectionAccepted
const CreateDefaultsSectionAcceptedCode int = 202

/*
CreateDefaultsSectionAccepted Configuration change accepted and reload requested

swagger:response createDefaultsSectionAccepted
*/
type CreateDefaultsSectionAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewCreateDefaultsSectionAccepted creates CreateDefaultsSectionAccepted with default headers values
func NewCreateDefaultsSectionAccepted() *CreateDefaultsSectionAccepted {

	return &CreateDefaultsSectionAccepted{}
}

// WithReloadID adds the reloadId to the create defaults section accepted response
func (o *CreateDefaultsSectionAccepted) WithReloadID(reloadID string) *CreateDefaultsSectionAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create defaults section accepted response
func (o *CreateDefaultsSectionAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create defaults section accepted response
func (o *CreateDefaultsSectionAccepted) WithPayload(payload *models.Defaults) *CreateDefaultsSectionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create defaults section accepted response
func (o *CreateDefaultsSectionAccepted) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefaultsSectionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateDefaultsSectionBadRequestCode is the HTTP code returned for type CreateDefaultsSectionBadRequest
const CreateDefaultsSectionBadRequestCode int = 400

/*
CreateDefaultsSectionBadRequest Bad request

swagger:response createDefaultsSectionBadRequest
*/
type CreateDefaultsSectionBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefaultsSectionBadRequest creates CreateDefaultsSectionBadRequest with default headers values
func NewCreateDefaultsSectionBadRequest() *CreateDefaultsSectionBadRequest {

	return &CreateDefaultsSectionBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create defaults section bad request response
func (o *CreateDefaultsSectionBadRequest) WithConfigurationVersion(configurationVersion string) *CreateDefaultsSectionBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create defaults section bad request response
func (o *CreateDefaultsSectionBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create defaults section bad request response
func (o *CreateDefaultsSectionBadRequest) WithPayload(payload *models.Error) *CreateDefaultsSectionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create defaults section bad request response
func (o *CreateDefaultsSectionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefaultsSectionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateDefaultsSectionConflictCode is the HTTP code returned for type CreateDefaultsSectionConflict
const CreateDefaultsSectionConflictCode int = 409

/*
CreateDefaultsSectionConflict The specified resource already exists

swagger:response createDefaultsSectionConflict
*/
type CreateDefaultsSectionConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefaultsSectionConflict creates CreateDefaultsSectionConflict with default headers values
func NewCreateDefaultsSectionConflict() *CreateDefaultsSectionConflict {

	return &CreateDefaultsSectionConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create defaults section conflict response
func (o *CreateDefaultsSectionConflict) WithConfigurationVersion(configurationVersion string) *CreateDefaultsSectionConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create defaults section conflict response
func (o *CreateDefaultsSectionConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create defaults section conflict response
func (o *CreateDefaultsSectionConflict) WithPayload(payload *models.Error) *CreateDefaultsSectionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create defaults section conflict response
func (o *CreateDefaultsSectionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefaultsSectionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
CreateDefaultsSectionDefault General Error

swagger:response createDefaultsSectionDefault
*/
type CreateDefaultsSectionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefaultsSectionDefault creates CreateDefaultsSectionDefault with default headers values
func NewCreateDefaultsSectionDefault(code int) *CreateDefaultsSectionDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDefaultsSectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create defaults section default response
func (o *CreateDefaultsSectionDefault) WithStatusCode(code int) *CreateDefaultsSectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create defaults section default response
func (o *CreateDefaultsSectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create defaults section default response
func (o *CreateDefaultsSectionDefault) WithConfigurationVersion(configurationVersion string) *CreateDefaultsSectionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create defaults section default response
func (o *CreateDefaultsSectionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create defaults section default response
func (o *CreateDefaultsSectionDefault) WithPayload(payload *models.Error) *CreateDefaultsSectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create defaults section default response
func (o *CreateDefaultsSectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefaultsSectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
