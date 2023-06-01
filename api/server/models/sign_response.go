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

// SignResponse Signature of a sent operation.
//
// swagger:model SignResponse
type SignResponse struct {

	// correlation Id
	// Format: byte
	CorrelationID CorrelationID `json:"correlationId,omitempty"`

	// Public part of the key pair used to sign the operation.
	// Read Only: true
	PublicKey string `json:"publicKey,omitempty"`

	// Hash of the operation attributes encrypted with the private part of the key pair.
	// Read Only: true
	// Format: byte
	Signature strfmt.Base64 `json:"signature,omitempty"`
}

// Validate validates this sign response
func (m *SignResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorrelationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignResponse) validateCorrelationID(formats strfmt.Registry) error {
	if swag.IsZero(m.CorrelationID) { // not required
		return nil
	}

	if err := m.CorrelationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("correlationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("correlationId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this sign response based on the context it is used
func (m *SignResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCorrelationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignResponse) contextValidateCorrelationID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CorrelationID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("correlationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("correlationId")
		}
		return err
	}

	return nil
}

func (m *SignResponse) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "publicKey", "body", string(m.PublicKey)); err != nil {
		return err
	}

	return nil
}

func (m *SignResponse) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "signature", "body", strfmt.Base64(m.Signature)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignResponse) UnmarshalBinary(b []byte) error {
	var res SignResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}