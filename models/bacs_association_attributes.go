// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BacsAssociationAttributes bacs association attributes
// swagger:model BacsAssociationAttributes
type BacsAssociationAttributes struct {

	// service user number
	ServiceUserNumber string `json:"service_user_number,omitempty"`
}

// Validate validates this bacs association attributes
func (m *BacsAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BacsAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BacsAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res BacsAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
