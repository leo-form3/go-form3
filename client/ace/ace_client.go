// Code generated by go-swagger; DO NOT EDIT.

package ace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRolesRoleIDAcesAceID deletes access control entry
*/
func (a *Client) DeleteRolesRoleIDAcesAceID(params *DeleteRolesRoleIDAcesAceIDParams) (*DeleteRolesRoleIDAcesAceIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRolesRoleIDAcesAceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRolesRoleIDAcesAceID",
		Method:             "DELETE",
		PathPattern:        "/roles/{role_id}/aces/{ace_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRolesRoleIDAcesAceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRolesRoleIDAcesAceIDNoContent), nil

}

/*
GetRolesRoleIDAces lists all access controls for role
*/
func (a *Client) GetRolesRoleIDAces(params *GetRolesRoleIDAcesParams) (*GetRolesRoleIDAcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesRoleIDAcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRolesRoleIDAces",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/aces",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRolesRoleIDAcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRolesRoleIDAcesOK), nil

}

/*
GetRolesRoleIDAcesAceID fetches access control entry
*/
func (a *Client) GetRolesRoleIDAcesAceID(params *GetRolesRoleIDAcesAceIDParams) (*GetRolesRoleIDAcesAceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesRoleIDAcesAceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRolesRoleIDAcesAceID",
		Method:             "GET",
		PathPattern:        "/roles/{role_id}/aces/{ace_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRolesRoleIDAcesAceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRolesRoleIDAcesAceIDOK), nil

}

/*
PostRolesRoleIDAces creates access control entry
*/
func (a *Client) PostRolesRoleIDAces(params *PostRolesRoleIDAcesParams) (*PostRolesRoleIDAcesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRolesRoleIDAcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRolesRoleIDAces",
		Method:             "POST",
		PathPattern:        "/roles/{role_id}/aces",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRolesRoleIDAcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRolesRoleIDAcesCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
