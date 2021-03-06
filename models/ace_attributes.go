// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AceAttributes ace attributes
// swagger:model aceAttributes
type AceAttributes struct {

	// action
	// Pattern: ^[A-Za-z]*$
	Action string `json:"action,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// record type
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`

	// role id
	RoleID strfmt.UUID `json:"role_id,omitempty"`
}

// Validate validates this ace attributes
func (m *AceAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AceAttributes) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.Pattern("action", "body", string(m.Action), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AceAttributes) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AceAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AceAttributes) UnmarshalBinary(b []byte) error {
	var res AceAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
