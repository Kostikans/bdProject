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

// Thread Ветка обсуждения на форуме.
//
//
// swagger:model Thread
type Thread struct {

	// Пользователь, создавший данную тему.
	// Example: j.sparrow
	// Required: true
	Author string `json:"author" db:"author"`

	// Дата создания ветки на форуме.
	// Example: 2017-01-01T00:00:00.000Z
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"  db:"created"`

	// Форум, в котором расположена данная ветка обсуждения.
	// Example: pirate-stories
	// Read Only: true
	Forum string `json:"forum,omitempty"  db:"forum"`

	// Идентификатор ветки обсуждения.
	// Example: 42
	// Read Only: true
	ID int32 `json:"id,omitempty" db:"id"`

	// Описание ветки обсуждения.
	// Example: An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?
	// Required: true
	Message string `json:"message"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL).
	// В данной структуре slug опционален и не может быть числом.
	//
	// Example: jones-cache
	// Read Only: true
	// Pattern: ^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$
	Slug string `json:"slug,omitempty"`

	// Заголовок ветки обсуждения.
	// Example: Davy Jones cache
	// Required: true
	Title string `json:"title"`

	// Кол-во голосов непосредственно за данное сообщение форума.
	// Read Only: true
	Votes int32 `json:"votes,omitempty"`

	Forum_ID int64 `json:"forum_id,omitempty"`
}

// Validate validates this thread
func (m *Thread) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thread) validateAuthor(formats strfmt.Registry) error {

	if err := validate.RequiredString("author", "body", string(m.Author)); err != nil {
		return err
	}

	return nil
}

func (m *Thread) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Thread) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *Thread) validateSlug(formats strfmt.Registry) error {

	if swag.IsZero(m.Slug) { // not required
		return nil
	}

	if err := validate.Pattern("slug", "body", string(m.Slug), `^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$`); err != nil {
		return err
	}

	return nil
}

func (m *Thread) validateTitle(formats strfmt.Registry) error {

	if err := validate.RequiredString("title", "body", string(m.Title)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this thread based on the context it is used
func (m *Thread) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateForum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlug(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thread) contextValidateForum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "forum", "body", string(m.Forum)); err != nil {
		return err
	}

	return nil
}

func (m *Thread) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Thread) contextValidateSlug(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "slug", "body", string(m.Slug)); err != nil {
		return err
	}

	return nil
}

func (m *Thread) contextValidateVotes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "votes", "body", int32(m.Votes)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thread) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thread) UnmarshalBinary(b []byte) error {
	var res Thread
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
