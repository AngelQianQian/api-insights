// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenAPISpecScoreElement open Api spec score element
//
// swagger:model OpenApiSpecScoreElement
type OpenAPISpecScoreElement struct {

	// findings
	Findings *SpecScoreFindings `json:"findings,omitempty"`

	// Name
	//
	// Name of the Object
	Name string `json:"name,omitempty"`

	// severity
	Severity APISecurityRiskSeverity `json:"severity,omitempty"`

	// spec path
	SpecPath string `json:"specPath,omitempty"`
}

// Validate validates this open Api spec score element
func (m *OpenAPISpecScoreElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenAPISpecScoreElement) validateFindings(formats strfmt.Registry) error {
	if swag.IsZero(m.Findings) { // not required
		return nil
	}

	if m.Findings != nil {
		if err := m.Findings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("findings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("findings")
			}
			return err
		}
	}

	return nil
}

func (m *OpenAPISpecScoreElement) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	if err := m.Severity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// ContextValidate validate this open Api spec score element based on the context it is used
func (m *OpenAPISpecScoreElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFindings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenAPISpecScoreElement) contextValidateFindings(ctx context.Context, formats strfmt.Registry) error {

	if m.Findings != nil {
		if err := m.Findings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("findings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("findings")
			}
			return err
		}
	}

	return nil
}

func (m *OpenAPISpecScoreElement) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Severity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenAPISpecScoreElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenAPISpecScoreElement) UnmarshalBinary(b []byte) error {
	var res OpenAPISpecScoreElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
