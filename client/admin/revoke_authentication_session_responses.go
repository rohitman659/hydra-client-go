// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra-client-go/models"
)

// RevokeAuthenticationSessionReader is a Reader for the RevokeAuthenticationSession structure.
type RevokeAuthenticationSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeAuthenticationSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRevokeAuthenticationSessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeAuthenticationSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeAuthenticationSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRevokeAuthenticationSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeAuthenticationSessionNoContent creates a RevokeAuthenticationSessionNoContent with default headers values
func NewRevokeAuthenticationSessionNoContent() *RevokeAuthenticationSessionNoContent {
	return &RevokeAuthenticationSessionNoContent{}
}

/* RevokeAuthenticationSessionNoContent describes a response with status code 204, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type RevokeAuthenticationSessionNoContent struct {
}

func (o *RevokeAuthenticationSessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/login][%d] revokeAuthenticationSessionNoContent ", 204)
}

func (o *RevokeAuthenticationSessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeAuthenticationSessionBadRequest creates a RevokeAuthenticationSessionBadRequest with default headers values
func NewRevokeAuthenticationSessionBadRequest() *RevokeAuthenticationSessionBadRequest {
	return &RevokeAuthenticationSessionBadRequest{}
}

/* RevokeAuthenticationSessionBadRequest describes a response with status code 400, with default header values.

genericError
*/
type RevokeAuthenticationSessionBadRequest struct {
	Payload *models.GenericError
}

func (o *RevokeAuthenticationSessionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/login][%d] revokeAuthenticationSessionBadRequest  %+v", 400, o.Payload)
}
func (o *RevokeAuthenticationSessionBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeAuthenticationSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthenticationSessionNotFound creates a RevokeAuthenticationSessionNotFound with default headers values
func NewRevokeAuthenticationSessionNotFound() *RevokeAuthenticationSessionNotFound {
	return &RevokeAuthenticationSessionNotFound{}
}

/* RevokeAuthenticationSessionNotFound describes a response with status code 404, with default header values.

genericError
*/
type RevokeAuthenticationSessionNotFound struct {
	Payload *models.GenericError
}

func (o *RevokeAuthenticationSessionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/login][%d] revokeAuthenticationSessionNotFound  %+v", 404, o.Payload)
}
func (o *RevokeAuthenticationSessionNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeAuthenticationSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthenticationSessionInternalServerError creates a RevokeAuthenticationSessionInternalServerError with default headers values
func NewRevokeAuthenticationSessionInternalServerError() *RevokeAuthenticationSessionInternalServerError {
	return &RevokeAuthenticationSessionInternalServerError{}
}

/* RevokeAuthenticationSessionInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type RevokeAuthenticationSessionInternalServerError struct {
	Payload *models.GenericError
}

func (o *RevokeAuthenticationSessionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/auth/sessions/login][%d] revokeAuthenticationSessionInternalServerError  %+v", 500, o.Payload)
}
func (o *RevokeAuthenticationSessionInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RevokeAuthenticationSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
