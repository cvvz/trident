// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ApplicationSnapshotRestoreReader is a Reader for the ApplicationSnapshotRestore structure.
type ApplicationSnapshotRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationSnapshotRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationSnapshotRestoreAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationSnapshotRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationSnapshotRestoreAccepted creates a ApplicationSnapshotRestoreAccepted with default headers values
func NewApplicationSnapshotRestoreAccepted() *ApplicationSnapshotRestoreAccepted {
	return &ApplicationSnapshotRestoreAccepted{}
}

/* ApplicationSnapshotRestoreAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationSnapshotRestoreAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *ApplicationSnapshotRestoreAccepted) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots/{uuid}/restore][%d] applicationSnapshotRestoreAccepted  %+v", 202, o.Payload)
}
func (o *ApplicationSnapshotRestoreAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationSnapshotRestoreAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotRestoreDefault creates a ApplicationSnapshotRestoreDefault with default headers values
func NewApplicationSnapshotRestoreDefault(code int) *ApplicationSnapshotRestoreDefault {
	return &ApplicationSnapshotRestoreDefault{
		_statusCode: code,
	}
}

/* ApplicationSnapshotRestoreDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationSnapshotRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application snapshot restore default response
func (o *ApplicationSnapshotRestoreDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationSnapshotRestoreDefault) Error() string {
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots/{uuid}/restore][%d] application_snapshot_restore default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationSnapshotRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationSnapshotRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}