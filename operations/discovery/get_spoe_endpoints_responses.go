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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetSpoeEndpointsOKCode is the HTTP code returned for type GetSpoeEndpointsOK
const GetSpoeEndpointsOKCode int = 200

/*GetSpoeEndpointsOK Success

swagger:response getSpoeEndpointsOK
*/
type GetSpoeEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetSpoeEndpointsOK creates GetSpoeEndpointsOK with default headers values
func NewGetSpoeEndpointsOK() *GetSpoeEndpointsOK {

	return &GetSpoeEndpointsOK{}
}

// WithPayload adds the payload to the get spoe endpoints o k response
func (o *GetSpoeEndpointsOK) WithPayload(payload models.Endpoints) *GetSpoeEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe endpoints o k response
func (o *GetSpoeEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Endpoints{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetSpoeEndpointsDefault General Error

swagger:response getSpoeEndpointsDefault
*/
type GetSpoeEndpointsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeEndpointsDefault creates GetSpoeEndpointsDefault with default headers values
func NewGetSpoeEndpointsDefault(code int) *GetSpoeEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSpoeEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) WithStatusCode(code int) *GetSpoeEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) WithConfigurationVersion(configurationVersion string) *GetSpoeEndpointsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) WithPayload(payload *models.Error) *GetSpoeEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
