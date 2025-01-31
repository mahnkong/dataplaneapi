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

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetHTTPErrorRuleOKCode is the HTTP code returned for type GetHTTPErrorRuleOK
const GetHTTPErrorRuleOKCode int = 200

/*
GetHTTPErrorRuleOK Successful operation

swagger:response getHttpErrorRuleOK
*/
type GetHTTPErrorRuleOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHTTPErrorRuleOKBody `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleOK creates GetHTTPErrorRuleOK with default headers values
func NewGetHTTPErrorRuleOK() *GetHTTPErrorRuleOK {

	return &GetHTTPErrorRuleOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http error rule o k response
func (o *GetHTTPErrorRuleOK) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http error rule o k response
func (o *GetHTTPErrorRuleOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http error rule o k response
func (o *GetHTTPErrorRuleOK) WithPayload(payload *GetHTTPErrorRuleOKBody) *GetHTTPErrorRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http error rule o k response
func (o *GetHTTPErrorRuleOK) SetPayload(payload *GetHTTPErrorRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetHTTPErrorRuleNotFoundCode is the HTTP code returned for type GetHTTPErrorRuleNotFound
const GetHTTPErrorRuleNotFoundCode int = 404

/*
GetHTTPErrorRuleNotFound The specified resource was not found

swagger:response getHttpErrorRuleNotFound
*/
type GetHTTPErrorRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleNotFound creates GetHTTPErrorRuleNotFound with default headers values
func NewGetHTTPErrorRuleNotFound() *GetHTTPErrorRuleNotFound {

	return &GetHTTPErrorRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get Http error rule not found response
func (o *GetHTTPErrorRuleNotFound) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Http error rule not found response
func (o *GetHTTPErrorRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Http error rule not found response
func (o *GetHTTPErrorRuleNotFound) WithPayload(payload *models.Error) *GetHTTPErrorRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http error rule not found response
func (o *GetHTTPErrorRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetHTTPErrorRuleDefault General Error

swagger:response getHttpErrorRuleDefault
*/
type GetHTTPErrorRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPErrorRuleDefault creates GetHTTPErrorRuleDefault with default headers values
func NewGetHTTPErrorRuleDefault(code int) *GetHTTPErrorRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPErrorRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) WithStatusCode(code int) *GetHTTPErrorRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) WithConfigurationVersion(configurationVersion string) *GetHTTPErrorRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) WithPayload(payload *models.Error) *GetHTTPErrorRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPErrorRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
