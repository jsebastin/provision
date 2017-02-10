package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PostDhcpLeaseReader is a Reader for the PostDhcpLease structure.
type PostDhcpLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDhcpLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDhcpLeaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostDhcpLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostDhcpLeaseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostDhcpLeaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDhcpLeaseCreated creates a PostDhcpLeaseCreated with default headers values
func NewPostDhcpLeaseCreated() *PostDhcpLeaseCreated {
	return &PostDhcpLeaseCreated{}
}

/*PostDhcpLeaseCreated handles this case with default header values.

PostDhcpLeaseCreated post dhcp lease created
*/
type PostDhcpLeaseCreated struct {
	Payload *models.DhcpLeaseInput
}

func (o *PostDhcpLeaseCreated) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postDhcpLeaseCreated  %+v", 201, o.Payload)
}

func (o *PostDhcpLeaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DhcpLeaseInput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDhcpLeaseUnauthorized creates a PostDhcpLeaseUnauthorized with default headers values
func NewPostDhcpLeaseUnauthorized() *PostDhcpLeaseUnauthorized {
	return &PostDhcpLeaseUnauthorized{}
}

/*PostDhcpLeaseUnauthorized handles this case with default header values.

PostDhcpLeaseUnauthorized post dhcp lease unauthorized
*/
type PostDhcpLeaseUnauthorized struct {
	Payload *models.Error
}

func (o *PostDhcpLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postDhcpLeaseUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDhcpLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDhcpLeaseConflict creates a PostDhcpLeaseConflict with default headers values
func NewPostDhcpLeaseConflict() *PostDhcpLeaseConflict {
	return &PostDhcpLeaseConflict{}
}

/*PostDhcpLeaseConflict handles this case with default header values.

PostDhcpLeaseConflict post dhcp lease conflict
*/
type PostDhcpLeaseConflict struct {
	Payload *models.Error
}

func (o *PostDhcpLeaseConflict) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postDhcpLeaseConflict  %+v", 409, o.Payload)
}

func (o *PostDhcpLeaseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDhcpLeaseInternalServerError creates a PostDhcpLeaseInternalServerError with default headers values
func NewPostDhcpLeaseInternalServerError() *PostDhcpLeaseInternalServerError {
	return &PostDhcpLeaseInternalServerError{}
}

/*PostDhcpLeaseInternalServerError handles this case with default header values.

PostDhcpLeaseInternalServerError post dhcp lease internal server error
*/
type PostDhcpLeaseInternalServerError struct {
	Payload *models.Error
}

func (o *PostDhcpLeaseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /leases][%d] postDhcpLeaseInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDhcpLeaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
