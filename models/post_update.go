// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostUpdate Сообщение для обновления сообщения внутри ветки на форуме.
// Пустые параметры остаются без изменений.
//
//
// swagger:model PostUpdate
type PostUpdate struct {

	// Собственно сообщение форума.
	// Example: We should be afraid of the Kraken.
	Message string `json:"message,omitempty"`
}

// Validate validates this post update
func (m *PostUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post update based on context it is used
func (m *PostUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUpdate) UnmarshalBinary(b []byte) error {
	var res PostUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
