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

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateSiteCreatedCode is the HTTP code returned for type CreateSiteCreated
const CreateSiteCreatedCode int = 201

/*CreateSiteCreated Site created

swagger:response createSiteCreated
*/
type CreateSiteCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Site `json:"body,omitempty"`
}

// NewCreateSiteCreated creates CreateSiteCreated with default headers values
func NewCreateSiteCreated() *CreateSiteCreated {

	return &CreateSiteCreated{}
}

// WithPayload adds the payload to the create site created response
func (o *CreateSiteCreated) WithPayload(payload *models.Site) *CreateSiteCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create site created response
func (o *CreateSiteCreated) SetPayload(payload *models.Site) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSiteCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSiteAcceptedCode is the HTTP code returned for type CreateSiteAccepted
const CreateSiteAcceptedCode int = 202

/*CreateSiteAccepted Configuration change accepted and reload requested

swagger:response createSiteAccepted
*/
type CreateSiteAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Site `json:"body,omitempty"`
}

// NewCreateSiteAccepted creates CreateSiteAccepted with default headers values
func NewCreateSiteAccepted() *CreateSiteAccepted {

	return &CreateSiteAccepted{}
}

// WithReloadID adds the reloadId to the create site accepted response
func (o *CreateSiteAccepted) WithReloadID(reloadID string) *CreateSiteAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create site accepted response
func (o *CreateSiteAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create site accepted response
func (o *CreateSiteAccepted) WithPayload(payload *models.Site) *CreateSiteAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create site accepted response
func (o *CreateSiteAccepted) SetPayload(payload *models.Site) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSiteAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateSiteBadRequestCode is the HTTP code returned for type CreateSiteBadRequest
const CreateSiteBadRequestCode int = 400

/*CreateSiteBadRequest Bad request

swagger:response createSiteBadRequest
*/
type CreateSiteBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSiteBadRequest creates CreateSiteBadRequest with default headers values
func NewCreateSiteBadRequest() *CreateSiteBadRequest {

	return &CreateSiteBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create site bad request response
func (o *CreateSiteBadRequest) WithConfigurationVersion(configurationVersion string) *CreateSiteBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create site bad request response
func (o *CreateSiteBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create site bad request response
func (o *CreateSiteBadRequest) WithPayload(payload *models.Error) *CreateSiteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create site bad request response
func (o *CreateSiteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSiteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateSiteConflictCode is the HTTP code returned for type CreateSiteConflict
const CreateSiteConflictCode int = 409

/*CreateSiteConflict The specified resource already exists

swagger:response createSiteConflict
*/
type CreateSiteConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSiteConflict creates CreateSiteConflict with default headers values
func NewCreateSiteConflict() *CreateSiteConflict {

	return &CreateSiteConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create site conflict response
func (o *CreateSiteConflict) WithConfigurationVersion(configurationVersion string) *CreateSiteConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create site conflict response
func (o *CreateSiteConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create site conflict response
func (o *CreateSiteConflict) WithPayload(payload *models.Error) *CreateSiteConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create site conflict response
func (o *CreateSiteConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSiteConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateSiteDefault General Error

swagger:response createSiteDefault
*/
type CreateSiteDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSiteDefault creates CreateSiteDefault with default headers values
func NewCreateSiteDefault(code int) *CreateSiteDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateSiteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create site default response
func (o *CreateSiteDefault) WithStatusCode(code int) *CreateSiteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create site default response
func (o *CreateSiteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create site default response
func (o *CreateSiteDefault) WithConfigurationVersion(configurationVersion string) *CreateSiteDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create site default response
func (o *CreateSiteDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create site default response
func (o *CreateSiteDefault) WithPayload(payload *models.Error) *CreateSiteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create site default response
func (o *CreateSiteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSiteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
