// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdditionalInfo AdditionalInfo
//
// swagger:model AdditionalInfo
type AdditionalInfo struct {

	// Affected Endpoints
	AffectedEndpoints []string `json:"affected_endpoints"`

	// Affected Spec Paths
	AffectedSpecPaths []string `json:"affected_spec_paths"`

	// Dictionary entries
	Entries map[string]string `json:"entries,omitempty"`
}

// Validate validates this additional info
func (m *AdditionalInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this additional info based on context it is used
func (m *AdditionalInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdditionalInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdditionalInfo) UnmarshalBinary(b []byte) error {
	var res AdditionalInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
