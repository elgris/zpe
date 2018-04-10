// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetServicesReader is a Reader for the GetServices structure.
type GetServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServicesOK creates a GetServicesOK with default headers values
func NewGetServicesOK() *GetServicesOK {
	return &GetServicesOK{}
}

/*GetServicesOK handles this case with default header values.

Succes
*/
type GetServicesOK struct {
	Payload []string
}

func (o *GetServicesOK) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesOK  %+v", 200, o.Payload)
}

func (o *GetServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesBadRequest creates a GetServicesBadRequest with default headers values
func NewGetServicesBadRequest() *GetServicesBadRequest {
	return &GetServicesBadRequest{}
}

/*GetServicesBadRequest handles this case with default header values.

Bad Request Error
*/
type GetServicesBadRequest struct {
}

func (o *GetServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesBadRequest ", 400)
}

func (o *GetServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
