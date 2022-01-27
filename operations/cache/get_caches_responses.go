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

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetCachesOKCode is the HTTP code returned for type GetCachesOK
const GetCachesOKCode int = 200

/*GetCachesOK Successful operation

swagger:response getCachesOK
*/
type GetCachesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetCachesOKBody `json:"body,omitempty"`
}

// NewGetCachesOK creates GetCachesOK with default headers values
func NewGetCachesOK() *GetCachesOK {

	return &GetCachesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get caches o k response
func (o *GetCachesOK) WithConfigurationVersion(configurationVersion string) *GetCachesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get caches o k response
func (o *GetCachesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get caches o k response
func (o *GetCachesOK) WithPayload(payload *GetCachesOKBody) *GetCachesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get caches o k response
func (o *GetCachesOK) SetPayload(payload *GetCachesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCachesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetCachesDefault General Error

swagger:response getCachesDefault
*/
type GetCachesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCachesDefault creates GetCachesDefault with default headers values
func NewGetCachesDefault(code int) *GetCachesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCachesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get caches default response
func (o *GetCachesDefault) WithStatusCode(code int) *GetCachesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get caches default response
func (o *GetCachesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get caches default response
func (o *GetCachesDefault) WithConfigurationVersion(configurationVersion string) *GetCachesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get caches default response
func (o *GetCachesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get caches default response
func (o *GetCachesDefault) WithPayload(payload *models.Error) *GetCachesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get caches default response
func (o *GetCachesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCachesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
