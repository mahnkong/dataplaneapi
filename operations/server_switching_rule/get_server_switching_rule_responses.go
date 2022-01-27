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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// GetServerSwitchingRuleOKCode is the HTTP code returned for type GetServerSwitchingRuleOK
const GetServerSwitchingRuleOKCode int = 200

/*GetServerSwitchingRuleOK Successful operation

swagger:response getServerSwitchingRuleOK
*/
type GetServerSwitchingRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetServerSwitchingRuleOKBody `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleOK creates GetServerSwitchingRuleOK with default headers values
func NewGetServerSwitchingRuleOK() *GetServerSwitchingRuleOK {

	return &GetServerSwitchingRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) WithConfigurationVersion(configurationVersion string) *GetServerSwitchingRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) WithPayload(payload *GetServerSwitchingRuleOKBody) *GetServerSwitchingRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) SetPayload(payload *GetServerSwitchingRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetServerSwitchingRuleNotFoundCode is the HTTP code returned for type GetServerSwitchingRuleNotFound
const GetServerSwitchingRuleNotFoundCode int = 404

/*GetServerSwitchingRuleNotFound The specified resource was not found

swagger:response getServerSwitchingRuleNotFound
*/
type GetServerSwitchingRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleNotFound creates GetServerSwitchingRuleNotFound with default headers values
func NewGetServerSwitchingRuleNotFound() *GetServerSwitchingRuleNotFound {

	return &GetServerSwitchingRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) WithConfigurationVersion(configurationVersion string) *GetServerSwitchingRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) WithPayload(payload *models.Error) *GetServerSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetServerSwitchingRuleDefault General Error

swagger:response getServerSwitchingRuleDefault
*/
type GetServerSwitchingRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleDefault creates GetServerSwitchingRuleDefault with default headers values
func NewGetServerSwitchingRuleDefault(code int) *GetServerSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) WithStatusCode(code int) *GetServerSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) WithConfigurationVersion(configurationVersion string) *GetServerSwitchingRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) WithPayload(payload *models.Error) *GetServerSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
