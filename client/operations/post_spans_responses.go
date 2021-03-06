// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostSpansReader is a Reader for the PostSpans structure.
type PostSpansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSpansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPostSpansAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSpansAccepted creates a PostSpansAccepted with default headers values
func NewPostSpansAccepted() *PostSpansAccepted {
	return &PostSpansAccepted{}
}

/*PostSpansAccepted handles this case with default header values.

Accepted
*/
type PostSpansAccepted struct {
}

func (o *PostSpansAccepted) Error() string {
	return fmt.Sprintf("[POST /spans][%d] postSpansAccepted ", 202)
}

func (o *PostSpansAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
