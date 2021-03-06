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

// Post Сообщение внутри ветки обсуждения на форуме.
//
//
// swagger:model Post
type Post struct {

	// Автор, написавший данное сообщение.
	// Example: j.sparrow
	// Required: true
	Author string `json:"author" db:"author"`

	// Дата создания сообщения на форуме.
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty" db:"created"`

	// Идентификатор форума (slug) данного сообещния.
	// Read Only: true
	Forum string `json:"forum,omitempty" db:"forum"`

	// Идентификатор данного сообщения.
	// Read Only: true
	ID int64 `json:"id,omitempty" db:"post_id"`

	// Истина, если данное сообщение было изменено.
	// Read Only: true
	IsEdited bool `json:"isEdited,omitempty" db:"is_edited"`

	// Собственно сообщение форума.
	// Example: We should be afraid of the Kraken.
	// Required: true
	Message string `json:"message" db:"message"`

	// Идентификатор родительского сообщения (0 - корневое сообщение обсуждения).
	//
	Parent int64 `json:"parent,omitempty" db:"parent"`

	// Идентификатор ветви (id) обсуждения данного сообещния.
	// Read Only: true
	Thread int32 `json:"thread,omitempty" db:"thread_id"`

	Forum_ID int64 `json:"forum_id,omitempty" db:"forum_id"`

	Parents []int `db:"parents"`
}

// Validate validates this post
func (m *Post) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Post) validateAuthor(formats strfmt.Registry) error {

	if err := validate.RequiredString("author", "body", string(m.Author)); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post based on the context it is used
func (m *Post) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsEdited(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThread(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Post) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Post) contextValidateForum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "forum", "body", string(m.Forum)); err != nil {
		return err
	}

	return nil
}

func (m *Post) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Post) contextValidateIsEdited(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isEdited", "body", bool(m.IsEdited)); err != nil {
		return err
	}

	return nil
}

func (m *Post) contextValidateThread(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "thread", "body", int32(m.Thread)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Post) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Post) UnmarshalBinary(b []byte) error {
	var res Post
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
