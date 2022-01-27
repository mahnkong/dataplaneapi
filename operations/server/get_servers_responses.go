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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetServersOKCode is the HTTP code returned for type GetServersOK
const GetServersOKCode int = 200

/*GetServersOK Successful operation

swagger:response getServersOK
*/
type GetServersOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetServersOKBody `json:"body,omitempty"`
}

// NewGetServersOK creates GetServersOK with default headers values
func NewGetServersOK() *GetServersOK {

	return &GetServersOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get servers o k response
func (o *GetServersOK) WithConfigurationVersion(configurationVersion string) *GetServersOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get servers o k response
func (o *GetServersOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get servers o k response
func (o *GetServersOK) WithPayload(payload *GetServersOKBody) *GetServersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get servers o k response
func (o *GetServersOK) SetPayload(payload *GetServersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServersDefault General Error

swagger:response getServersDefault
*/
type GetServersDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServersDefault creates GetServersDefault with default headers values
func NewGetServersDefault(code int) *GetServersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get servers default response
func (o *GetServersDefault) WithStatusCode(code int) *GetServersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get servers default response
func (o *GetServersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get servers default response
func (o *GetServersDefault) WithConfigurationVersion(configurationVersion string) *GetServersDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get servers default response
func (o *GetServersDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get servers default response
func (o *GetServersDefault) WithPayload(payload *models.Error) *GetServersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get servers default response
func (o *GetServersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
