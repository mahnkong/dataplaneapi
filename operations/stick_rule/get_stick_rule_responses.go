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

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetStickRuleOKCode is the HTTP code returned for type GetStickRuleOK
const GetStickRuleOKCode int = 200

/*
GetStickRuleOK Successful operation

swagger:response getStickRuleOK
*/
type GetStickRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetStickRuleOKBody `json:"body,omitempty"`
}

// NewGetStickRuleOK creates GetStickRuleOK with default headers values
func NewGetStickRuleOK() *GetStickRuleOK {

	return &GetStickRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get stick rule o k response
func (o *GetStickRuleOK) WithConfigurationVersion(configurationVersion string) *GetStickRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get stick rule o k response
func (o *GetStickRuleOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get stick rule o k response
func (o *GetStickRuleOK) WithPayload(payload *GetStickRuleOKBody) *GetStickRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick rule o k response
func (o *GetStickRuleOK) SetPayload(payload *GetStickRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetStickRuleNotFoundCode is the HTTP code returned for type GetStickRuleNotFound
const GetStickRuleNotFoundCode int = 404

/*
GetStickRuleNotFound The specified resource was not found

swagger:response getStickRuleNotFound
*/
type GetStickRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStickRuleNotFound creates GetStickRuleNotFound with default headers values
func NewGetStickRuleNotFound() *GetStickRuleNotFound {

	return &GetStickRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get stick rule not found response
func (o *GetStickRuleNotFound) WithConfigurationVersion(configurationVersion string) *GetStickRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get stick rule not found response
func (o *GetStickRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get stick rule not found response
func (o *GetStickRuleNotFound) WithPayload(payload *models.Error) *GetStickRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick rule not found response
func (o *GetStickRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*
GetStickRuleDefault General Error

swagger:response getStickRuleDefault
*/
type GetStickRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStickRuleDefault creates GetStickRuleDefault with default headers values
func NewGetStickRuleDefault(code int) *GetStickRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStickRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stick rule default response
func (o *GetStickRuleDefault) WithStatusCode(code int) *GetStickRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stick rule default response
func (o *GetStickRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get stick rule default response
func (o *GetStickRuleDefault) WithConfigurationVersion(configurationVersion string) *GetStickRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get stick rule default response
func (o *GetStickRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get stick rule default response
func (o *GetStickRuleDefault) WithPayload(payload *models.Error) *GetStickRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick rule default response
func (o *GetStickRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
