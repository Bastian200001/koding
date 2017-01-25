package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJStackTemplateCreateReader is a Reader for the PostRemoteAPIJStackTemplateCreate structure.
type PostRemoteAPIJStackTemplateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJStackTemplateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJStackTemplateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJStackTemplateCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJStackTemplateCreateOK creates a PostRemoteAPIJStackTemplateCreateOK with default headers values
func NewPostRemoteAPIJStackTemplateCreateOK() *PostRemoteAPIJStackTemplateCreateOK {
	return &PostRemoteAPIJStackTemplateCreateOK{}
}

/*PostRemoteAPIJStackTemplateCreateOK handles this case with default header values.

created JStackTemplate instance
*/
type PostRemoteAPIJStackTemplateCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJStackTemplateCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.create][%d] postRemoteApiJStackTemplateCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJStackTemplateCreateUnauthorized creates a PostRemoteAPIJStackTemplateCreateUnauthorized with default headers values
func NewPostRemoteAPIJStackTemplateCreateUnauthorized() *PostRemoteAPIJStackTemplateCreateUnauthorized {
	return &PostRemoteAPIJStackTemplateCreateUnauthorized{}
}

/*PostRemoteAPIJStackTemplateCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJStackTemplateCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJStackTemplateCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.create][%d] postRemoteApiJStackTemplateCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJStackTemplateCreateBody post remote API j stack template create body
swagger:model PostRemoteAPIJStackTemplateCreateBody
*/
type PostRemoteAPIJStackTemplateCreateBody struct {

	// config
	// Required: true
	Config interface{} `json:"config"`

	// credentials
	// Required: true
	Credentials interface{} `json:"credentials"`

	// template
	// Required: true
	Template *string `json:"template"`

	// title
	// Required: true
	Title *string `json:"title"`
}
