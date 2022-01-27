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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteSpoeScopeNoContentCode is the HTTP code returned for type DeleteSpoeScopeNoContent
const DeleteSpoeScopeNoContentCode int = 204

/*DeleteSpoeScopeNoContent Spoe scope deleted

swagger:response deleteSpoeScopeNoContent
*/
type DeleteSpoeScopeNoContent struct {
}

// NewDeleteSpoeScopeNoContent creates DeleteSpoeScopeNoContent with default headers values
func NewDeleteSpoeScopeNoContent() *DeleteSpoeScopeNoContent {

	return &DeleteSpoeScopeNoContent{}
}

// WriteResponse to the client
func (o *DeleteSpoeScopeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteSpoeScopeNotFoundCode is the HTTP code returned for type DeleteSpoeScopeNotFound
const DeleteSpoeScopeNotFoundCode int = 404

/*DeleteSpoeScopeNotFound The specified resource was not found

swagger:response deleteSpoeScopeNotFound
*/
type DeleteSpoeScopeNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSpoeScopeNotFound creates DeleteSpoeScopeNotFound with default headers values
func NewDeleteSpoeScopeNotFound() *DeleteSpoeScopeNotFound {

	return &DeleteSpoeScopeNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete spoe scope not found response
func (o *DeleteSpoeScopeNotFound) WithConfigurationVersion(configurationVersion string) *DeleteSpoeScopeNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete spoe scope not found response
func (o *DeleteSpoeScopeNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete spoe scope not found response
func (o *DeleteSpoeScopeNotFound) WithPayload(payload *models.Error) *DeleteSpoeScopeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete spoe scope not found response
func (o *DeleteSpoeScopeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSpoeScopeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteSpoeScopeDefault General Error

swagger:response deleteSpoeScopeDefault
*/
type DeleteSpoeScopeDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSpoeScopeDefault creates DeleteSpoeScopeDefault with default headers values
func NewDeleteSpoeScopeDefault(code int) *DeleteSpoeScopeDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSpoeScopeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) WithStatusCode(code int) *DeleteSpoeScopeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) WithConfigurationVersion(configurationVersion string) *DeleteSpoeScopeDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) WithPayload(payload *models.Error) *DeleteSpoeScopeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSpoeScopeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
