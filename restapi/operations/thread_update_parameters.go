// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/kostikans/bdProject/models"
)

// NewThreadUpdateParams creates a new ThreadUpdateParams object
// no default values defined in spec.
func NewThreadUpdateParams() ThreadUpdateParams {

	return ThreadUpdateParams{}
}

// ThreadUpdateParams contains all the bound params for the thread update operation
// typically these are obtained from a http.Request
//
// swagger:parameters threadUpdate
type ThreadUpdateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор ветки обсуждения.
	  Required: true
	  In: path
	*/
	SlugOrID string
	/*Данные ветки обсуждения.
	  Required: true
	  In: body
	*/
	Thread *models.ThreadUpdate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewThreadUpdateParams() beforehand.
func (o *ThreadUpdateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSlugOrID, rhkSlugOrID, _ := route.Params.GetOK("slug_or_id")
	if err := o.bindSlugOrID(rSlugOrID, rhkSlugOrID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ThreadUpdate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("thread", "body", ""))
			} else {
				res = append(res, errors.NewParseError("thread", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Thread = &body
			}
		}
	} else {
		res = append(res, errors.Required("thread", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSlugOrID binds and validates parameter SlugOrID from path.
func (o *ThreadUpdateParams) bindSlugOrID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SlugOrID = raw

	return nil
}
