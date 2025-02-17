// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigListenAddressReader is a Reader for the FindConfigListenAddress structure.
type FindConfigListenAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigListenAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigListenAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigListenAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigListenAddressOK creates a FindConfigListenAddressOK with default headers values
func NewFindConfigListenAddressOK() *FindConfigListenAddressOK {
	return &FindConfigListenAddressOK{}
}

/*FindConfigListenAddressOK handles this case with default header values.

Config value
*/
type FindConfigListenAddressOK struct {
	Payload string
}

func (o *FindConfigListenAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigListenAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigListenAddressDefault creates a FindConfigListenAddressDefault with default headers values
func NewFindConfigListenAddressDefault(code int) *FindConfigListenAddressDefault {
	return &FindConfigListenAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigListenAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigListenAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config listen address default response
func (o *FindConfigListenAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigListenAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigListenAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigListenAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
