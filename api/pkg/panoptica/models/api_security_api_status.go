// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// APISecurityAPIStatus ApiStatus
//
// Api status enumeration.
//
// swagger:model ApiSecurityApiStatus
type APISecurityAPIStatus string

func NewAPISecurityAPIStatus(value APISecurityAPIStatus) *APISecurityAPIStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated APISecurityAPIStatus.
func (m APISecurityAPIStatus) Pointer() *APISecurityAPIStatus {
	return &m
}

const (

	// APISecurityAPIStatusPROCESSING captures enum value "PROCESSING"
	APISecurityAPIStatusPROCESSING APISecurityAPIStatus = "PROCESSING"

	// APISecurityAPIStatusVALID captures enum value "VALID"
	APISecurityAPIStatusVALID APISecurityAPIStatus = "VALID"

	// APISecurityAPIStatusINVALID captures enum value "INVALID"
	APISecurityAPIStatusINVALID APISecurityAPIStatus = "INVALID"
)

// for schema
var apiSecurityApiStatusEnum []interface{}

func init() {
	var res []APISecurityAPIStatus
	if err := json.Unmarshal([]byte(`["PROCESSING","VALID","INVALID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiSecurityApiStatusEnum = append(apiSecurityApiStatusEnum, v)
	}
}

func (m APISecurityAPIStatus) validateAPISecurityAPIStatusEnum(path, location string, value APISecurityAPIStatus) error {
	if err := validate.EnumCase(path, location, value, apiSecurityApiStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Api security Api status
func (m APISecurityAPIStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPISecurityAPIStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Api security Api status based on context it is used
func (m APISecurityAPIStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
