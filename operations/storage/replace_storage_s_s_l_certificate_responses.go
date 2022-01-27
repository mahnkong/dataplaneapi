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

// ReplaceStorageSSLCertificateOKCode is the HTTP code returned for type ReplaceStorageSSLCertificateOK
const ReplaceStorageSSLCertificateOKCode int = 200

/*ReplaceStorageSSLCertificateOK SSL certificate replaced

swagger:response replaceStorageSSLCertificateOK
*/
type ReplaceStorageSSLCertificateOK struct {

	/*
	  In: Body
	*/
	Payload *models.SslCertificate `json:"body,omitempty"`
}

// NewReplaceStorageSSLCertificateOK creates ReplaceStorageSSLCertificateOK with default headers values
func NewReplaceStorageSSLCertificateOK() *ReplaceStorageSSLCertificateOK {

	return &ReplaceStorageSSLCertificateOK{}
}

// WithPayload adds the payload to the replace storage s s l certificate o k response
func (o *ReplaceStorageSSLCertificateOK) WithPayload(payload *models.SslCertificate) *ReplaceStorageSSLCertificateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage s s l certificate o k response
func (o *ReplaceStorageSSLCertificateOK) SetPayload(payload *models.SslCertificate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageSSLCertificateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageSSLCertificateAcceptedCode is the HTTP code returned for type ReplaceStorageSSLCertificateAccepted
const ReplaceStorageSSLCertificateAcceptedCode int = 202

/*ReplaceStorageSSLCertificateAccepted SSL certificate replaced and reload requested

swagger:response replaceStorageSSLCertificateAccepted
*/
type ReplaceStorageSSLCertificateAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.SslCertificate `json:"body,omitempty"`
}

// NewReplaceStorageSSLCertificateAccepted creates ReplaceStorageSSLCertificateAccepted with default headers values
func NewReplaceStorageSSLCertificateAccepted() *ReplaceStorageSSLCertificateAccepted {

	return &ReplaceStorageSSLCertificateAccepted{}
}

// WithReloadID adds the reloadId to the replace storage s s l certificate accepted response
func (o *ReplaceStorageSSLCertificateAccepted) WithReloadID(reloadID string) *ReplaceStorageSSLCertificateAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace storage s s l certificate accepted response
func (o *ReplaceStorageSSLCertificateAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace storage s s l certificate accepted response
func (o *ReplaceStorageSSLCertificateAccepted) WithPayload(payload *models.SslCertificate) *ReplaceStorageSSLCertificateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage s s l certificate accepted response
func (o *ReplaceStorageSSLCertificateAccepted) SetPayload(payload *models.SslCertificate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageSSLCertificateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceStorageSSLCertificateBadRequestCode is the HTTP code returned for type ReplaceStorageSSLCertificateBadRequest
const ReplaceStorageSSLCertificateBadRequestCode int = 400

/*ReplaceStorageSSLCertificateBadRequest Bad request

swagger:response replaceStorageSSLCertificateBadRequest
*/
type ReplaceStorageSSLCertificateBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStorageSSLCertificateBadRequest creates ReplaceStorageSSLCertificateBadRequest with default headers values
func NewReplaceStorageSSLCertificateBadRequest() *ReplaceStorageSSLCertificateBadRequest {

	return &ReplaceStorageSSLCertificateBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace storage s s l certificate bad request response
func (o *ReplaceStorageSSLCertificateBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceStorageSSLCertificateBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace storage s s l certificate bad request response
func (o *ReplaceStorageSSLCertificateBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace storage s s l certificate bad request response
func (o *ReplaceStorageSSLCertificateBadRequest) WithPayload(payload *models.Error) *ReplaceStorageSSLCertificateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage s s l certificate bad request response
func (o *ReplaceStorageSSLCertificateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageSSLCertificateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceStorageSSLCertificateNotFoundCode is the HTTP code returned for type ReplaceStorageSSLCertificateNotFound
const ReplaceStorageSSLCertificateNotFoundCode int = 404

/*ReplaceStorageSSLCertificateNotFound The specified resource was not found

swagger:response replaceStorageSSLCertificateNotFound
*/
type ReplaceStorageSSLCertificateNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStorageSSLCertificateNotFound creates ReplaceStorageSSLCertificateNotFound with default headers values
func NewReplaceStorageSSLCertificateNotFound() *ReplaceStorageSSLCertificateNotFound {

	return &ReplaceStorageSSLCertificateNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace storage s s l certificate not found response
func (o *ReplaceStorageSSLCertificateNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceStorageSSLCertificateNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace storage s s l certificate not found response
func (o *ReplaceStorageSSLCertificateNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace storage s s l certificate not found response
func (o *ReplaceStorageSSLCertificateNotFound) WithPayload(payload *models.Error) *ReplaceStorageSSLCertificateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage s s l certificate not found response
func (o *ReplaceStorageSSLCertificateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageSSLCertificateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceStorageSSLCertificateDefault General Error

swagger:response replaceStorageSSLCertificateDefault
*/
type ReplaceStorageSSLCertificateDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceStorageSSLCertificateDefault creates ReplaceStorageSSLCertificateDefault with default headers values
func NewReplaceStorageSSLCertificateDefault(code int) *ReplaceStorageSSLCertificateDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceStorageSSLCertificateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) WithStatusCode(code int) *ReplaceStorageSSLCertificateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) WithConfigurationVersion(configurationVersion string) *ReplaceStorageSSLCertificateDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) WithPayload(payload *models.Error) *ReplaceStorageSSLCertificateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage s s l certificate default response
func (o *ReplaceStorageSSLCertificateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageSSLCertificateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
