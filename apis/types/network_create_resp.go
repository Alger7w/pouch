// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkCreateResp contains the response for the remote API: POST /networks/create
// swagger:model NetworkCreateResp

type NetworkCreateResp struct {

	// ID is the id of the network.
	ID string `json:"ID,omitempty"`

	// Warning means the message of create network result.
	Warning string `json:"Warning,omitempty"`
}

/* polymorph NetworkCreateResp ID false */

/* polymorph NetworkCreateResp Warning false */

// Validate validates this network create resp
func (m *NetworkCreateResp) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreateResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreateResp) UnmarshalBinary(b []byte) error {
	var res NetworkCreateResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
