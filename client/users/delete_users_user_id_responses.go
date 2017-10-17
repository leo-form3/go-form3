// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUsersUserIDReader is a Reader for the DeleteUsersUserID structure.
type DeleteUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsersUserIDOK creates a DeleteUsersUserIDOK with default headers values
func NewDeleteUsersUserIDOK() *DeleteUsersUserIDOK {
	return &DeleteUsersUserIDOK{}
}

/*DeleteUsersUserIDOK handles this case with default header values.

User deleted
*/
type DeleteUsersUserIDOK struct {
}

func (o *DeleteUsersUserIDOK) Error() string {
	return fmt.Sprintf("[DELETE /security/users/{user_id}][%d] deleteUsersUserIdOK ", 200)
}

func (o *DeleteUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
