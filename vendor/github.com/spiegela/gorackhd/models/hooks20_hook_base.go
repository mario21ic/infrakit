package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Hooks20HookBase Hook base schema
// swagger:model hooks.2.0_HookBase
type Hooks20HookBase struct {

	// filters
	Filters []interface{} `json:"filters"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this hooks 2 0 hook base
func (m *Hooks20HookBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hooks20HookBase) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {

	}

	return nil
}
