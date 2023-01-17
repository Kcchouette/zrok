// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessRequest access request
//
// swagger:model accessRequest
type AccessRequest struct {

	// env z Id
	EnvZID string `json:"envZId,omitempty"`

	// shr token
	ShrToken string `json:"shrToken,omitempty"`
}

// Validate validates this access request
func (m *AccessRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access request based on context it is used
func (m *AccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessRequest) UnmarshalBinary(b []byte) error {
	var res AccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}