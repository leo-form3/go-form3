// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostUsersUserIDRolesRoleIDReader is a Reader for the PostUsersUserIDRolesRoleID structure.
type PostUsersUserIDRolesRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersUserIDRolesRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostUsersUserIDRolesRoleIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersUserIDRolesRoleIDCreated creates a PostUsersUserIDRolesRoleIDCreated with default headers values
func NewPostUsersUserIDRolesRoleIDCreated() *PostUsersUserIDRolesRoleIDCreated {
	return &PostUsersUserIDRolesRoleIDCreated{}
}

/*PostUsersUserIDRolesRoleIDCreated handles this case with default header values.

Role set OK
*/
type PostUsersUserIDRolesRoleIDCreated struct {
}

func (o *PostUsersUserIDRolesRoleIDCreated) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/roles/{role_id}][%d] postUsersUserIdRolesRoleIdCreated ", 201)
}

func (o *PostUsersUserIDRolesRoleIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
