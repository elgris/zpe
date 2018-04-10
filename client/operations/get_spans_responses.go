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

// GetSpansReader is a Reader for the GetSpans structure.
type GetSpansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSpansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSpansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpansOK creates a GetSpansOK with default headers values
func NewGetSpansOK() *GetSpansOK {
	return &GetSpansOK{}
}

/*GetSpansOK handles this case with default header values.

OK
*/
type GetSpansOK struct {
	Payload []string
}

func (o *GetSpansOK) Error() string {
	return fmt.Sprintf("[GET /spans][%d] getSpansOK  %+v", 200, o.Payload)
}

func (o *GetSpansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpansBadRequest creates a GetSpansBadRequest with default headers values
func NewGetSpansBadRequest() *GetSpansBadRequest {
	return &GetSpansBadRequest{}
}

/*GetSpansBadRequest handles this case with default header values.

Bad Request Error
*/
type GetSpansBadRequest struct {
}

func (o *GetSpansBadRequest) Error() string {
	return fmt.Sprintf("[GET /spans][%d] getSpansBadRequest ", 400)
}

func (o *GetSpansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
