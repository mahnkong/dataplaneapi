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

// GetSpoeAgentsOKCode is the HTTP code returned for type GetSpoeAgentsOK
const GetSpoeAgentsOKCode int = 200

/*GetSpoeAgentsOK Successful operation

swagger:response getSpoeAgentsOK
*/
type GetSpoeAgentsOK struct {
	/*Spoe configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetSpoeAgentsOKBody `json:"body,omitempty"`
}

// NewGetSpoeAgentsOK creates GetSpoeAgentsOK with default headers values
func NewGetSpoeAgentsOK() *GetSpoeAgentsOK {

	return &GetSpoeAgentsOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe agents o k response
func (o *GetSpoeAgentsOK) WithConfigurationVersion(configurationVersion string) *GetSpoeAgentsOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe agents o k response
func (o *GetSpoeAgentsOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe agents o k response
func (o *GetSpoeAgentsOK) WithPayload(payload *GetSpoeAgentsOKBody) *GetSpoeAgentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe agents o k response
func (o *GetSpoeAgentsOK) SetPayload(payload *GetSpoeAgentsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeAgentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetSpoeAgentsDefault General Error

swagger:response getSpoeAgentsDefault
*/
type GetSpoeAgentsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeAgentsDefault creates GetSpoeAgentsDefault with default headers values
func NewGetSpoeAgentsDefault(code int) *GetSpoeAgentsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeAgentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe agents default response
func (o *GetSpoeAgentsDefault) WithStatusCode(code int) *GetSpoeAgentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe agents default response
func (o *GetSpoeAgentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe agents default response
func (o *GetSpoeAgentsDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeAgentsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe agents default response
func (o *GetSpoeAgentsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe agents default response
func (o *GetSpoeAgentsDefault) WithPayload(payload *models.Error) *GetSpoeAgentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe agents default response
func (o *GetSpoeAgentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeAgentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
