// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIServiceScore ApiServiceScore
//
// swagger:model ApiServiceScore
type APIServiceScore struct {

	// api
	// Required: true
	API *APIServiceScoreGrade `json:"api"`

	// id
	// Required: true
	// Format: uuid
	Identifier *strfmt.UUID `json:"identifier"`

	// provider
	// Required: true
	Provider *APIProviderScoreGrade `json:"provider"`

	// status
	// Required: true
	Status *ScoreExitStatusEnum `json:"status"`

	// Timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this Api service score
func (m *APIServiceScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceScore) validateAPI(formats strfmt.Registry) error {

	if err := validate.Required("api", "body", m.API); err != nil {
		return err
	}

	if m.API != nil {
		if err := m.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *APIServiceScore) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	if err := validate.FormatOf("identifier", "body", "uuid", m.Identifier.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceScore) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *APIServiceScore) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *APIServiceScore) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Api service score based on the context it is used
func (m *APIServiceScore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceScore) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if m.API != nil {
		if err := m.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *APIServiceScore) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.Provider != nil {
		if err := m.Provider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

func (m *APIServiceScore) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIServiceScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIServiceScore) UnmarshalBinary(b []byte) error {
	var res APIServiceScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
