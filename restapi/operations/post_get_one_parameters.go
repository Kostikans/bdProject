// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPostGetOneParams creates a new PostGetOneParams object
// no default values defined in spec.
func NewPostGetOneParams() PostGetOneParams {

	return PostGetOneParams{}
}

// PostGetOneParams contains all the bound params for the post get one operation
// typically these are obtained from a http.Request
//
// swagger:parameters postGetOne
type PostGetOneParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор сообщения.
	  Required: true
	  In: path
	*/
	ID int64
	/*Включение полной информации о соответвующем объекте сообщения.

	Если тип объекта не указан, то полная информация об этих объектах не
	передаётся.

	  In: query
	*/
	Related []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostGetOneParams() beforehand.
func (o *PostGetOneParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qRelated, qhkRelated, _ := qs.GetOK("related")
	if err := o.bindRelated(qRelated, qhkRelated, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PostGetOneParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

// bindRelated binds and validates array parameter Related from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *PostGetOneParams) bindRelated(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvRelated string
	if len(rawData) > 0 {
		qvRelated = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	relatedIC := swag.SplitByFormat(qvRelated, "")
	if len(relatedIC) == 0 {
		return nil
	}

	var relatedIR []string
	for i, relatedIV := range relatedIC {
		relatedI := relatedIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "related", i), "query", relatedI, []interface{}{"user", "forum", "thread"}, true); err != nil {
			return err
		}

		relatedIR = append(relatedIR, relatedI)
	}

	o.Related = relatedIR

	return nil
}
