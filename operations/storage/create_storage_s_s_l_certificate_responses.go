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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// CreateStorageSSLCertificateCreatedCode is the HTTP code returned for type CreateStorageSSLCertificateCreated
const CreateStorageSSLCertificateCreatedCode int = 201

/*CreateStorageSSLCertificateCreated SSL certificate created

swagger:response createStorageSSLCertificateCreated
*/
type CreateStorageSSLCertificateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SslCertificate `json:"body,omitempty"`
}

// NewCreateStorageSSLCertificateCreated creates CreateStorageSSLCertificateCreated with default headers values
func NewCreateStorageSSLCertificateCreated() *CreateStorageSSLCertificateCreated {

	return &CreateStorageSSLCertificateCreated{}
}

// WithPayload adds the payload to the create storage s s l certificate created response
func (o *CreateStorageSSLCertificateCreated) WithPayload(payload *models.SslCertificate) *CreateStorageSSLCertificateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage s s l certificate created response
func (o *CreateStorageSSLCertificateCreated) SetPayload(payload *models.SslCertificate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageSSLCertificateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageSSLCertificateBadRequestCode is the HTTP code returned for type CreateStorageSSLCertificateBadRequest
const CreateStorageSSLCertificateBadRequestCode int = 400

/*CreateStorageSSLCertificateBadRequest Bad request

swagger:response createStorageSSLCertificateBadRequest
*/
type CreateStorageSSLCertificateBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageSSLCertificateBadRequest creates CreateStorageSSLCertificateBadRequest with default headers values
func NewCreateStorageSSLCertificateBadRequest() *CreateStorageSSLCertificateBadRequest {

	return &CreateStorageSSLCertificateBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create storage s s l certificate bad request response
func (o *CreateStorageSSLCertificateBadRequest) WithConfigurationVersion(configurationVersion string) *CreateStorageSSLCertificateBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage s s l certificate bad request response
func (o *CreateStorageSSLCertificateBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage s s l certificate bad request response
func (o *CreateStorageSSLCertificateBadRequest) WithPayload(payload *models.Error) *CreateStorageSSLCertificateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage s s l certificate bad request response
func (o *CreateStorageSSLCertificateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageSSLCertificateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateStorageSSLCertificateConflictCode is the HTTP code returned for type CreateStorageSSLCertificateConflict
const CreateStorageSSLCertificateConflictCode int = 409

/*CreateStorageSSLCertificateConflict The specified resource already exists

swagger:response createStorageSSLCertificateConflict
*/
type CreateStorageSSLCertificateConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageSSLCertificateConflict creates CreateStorageSSLCertificateConflict with default headers values
func NewCreateStorageSSLCertificateConflict() *CreateStorageSSLCertificateConflict {

	return &CreateStorageSSLCertificateConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create storage s s l certificate conflict response
func (o *CreateStorageSSLCertificateConflict) WithConfigurationVersion(configurationVersion string) *CreateStorageSSLCertificateConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage s s l certificate conflict response
func (o *CreateStorageSSLCertificateConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage s s l certificate conflict response
func (o *CreateStorageSSLCertificateConflict) WithPayload(payload *models.Error) *CreateStorageSSLCertificateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage s s l certificate conflict response
func (o *CreateStorageSSLCertificateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageSSLCertificateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateStorageSSLCertificateDefault General Error

swagger:response createStorageSSLCertificateDefault
*/
type CreateStorageSSLCertificateDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateStorageSSLCertificateDefault creates CreateStorageSSLCertificateDefault with default headers values
func NewCreateStorageSSLCertificateDefault(code int) *CreateStorageSSLCertificateDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateStorageSSLCertificateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) WithStatusCode(code int) *CreateStorageSSLCertificateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) WithConfigurationVersion(configurationVersion string) *CreateStorageSSLCertificateDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) WithPayload(payload *models.Error) *CreateStorageSSLCertificateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageSSLCertificateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
