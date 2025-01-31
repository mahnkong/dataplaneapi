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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetAWSRegionOKCode is the HTTP code returned for type GetAWSRegionOK
const GetAWSRegionOKCode int = 200

/*
GetAWSRegionOK Successful operation

swagger:response getAWSRegionOK
*/
type GetAWSRegionOK struct {

	/*
	  In: Body
	*/
	Payload *GetAWSRegionOKBody `json:"body,omitempty"`
}

// NewGetAWSRegionOK creates GetAWSRegionOK with default headers values
func NewGetAWSRegionOK() *GetAWSRegionOK {

	return &GetAWSRegionOK{}
}

// WithPayload adds the payload to the get a w s region o k response
func (o *GetAWSRegionOK) WithPayload(payload *GetAWSRegionOKBody) *GetAWSRegionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s region o k response
func (o *GetAWSRegionOK) SetPayload(payload *GetAWSRegionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSRegionNotFoundCode is the HTTP code returned for type GetAWSRegionNotFound
const GetAWSRegionNotFoundCode int = 404

/*
GetAWSRegionNotFound The specified resource was not found

swagger:response getAWSRegionNotFound
*/
type GetAWSRegionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionNotFound creates GetAWSRegionNotFound with default headers values
func NewGetAWSRegionNotFound() *GetAWSRegionNotFound {

	return &GetAWSRegionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the get a w s region not found response
func (o *GetAWSRegionNotFound) WithConfigurationVersion(configurationVersion string) *GetAWSRegionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get a w s region not found response
func (o *GetAWSRegionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get a w s region not found response
func (o *GetAWSRegionNotFound) WithPayload(payload *models.Error) *GetAWSRegionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s region not found response
func (o *GetAWSRegionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
GetAWSRegionDefault General Error

swagger:response getAWSRegionDefault
*/
type GetAWSRegionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionDefault creates GetAWSRegionDefault with default headers values
func NewGetAWSRegionDefault(code int) *GetAWSRegionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAWSRegionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get a w s region default response
func (o *GetAWSRegionDefault) WithStatusCode(code int) *GetAWSRegionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get a w s region default response
func (o *GetAWSRegionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get a w s region default response
func (o *GetAWSRegionDefault) WithConfigurationVersion(configurationVersion string) *GetAWSRegionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get a w s region default response
func (o *GetAWSRegionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get a w s region default response
func (o *GetAWSRegionDefault) WithPayload(payload *models.Error) *GetAWSRegionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s region default response
func (o *GetAWSRegionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
