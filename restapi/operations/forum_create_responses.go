// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kostikans/bdProject/models"
)

// ForumCreateCreatedCode is the HTTP code returned for type ForumCreateCreated
const ForumCreateCreatedCode int = 201

/*ForumCreateCreated Форум успешно создан.
Возвращает данные созданного форума.


swagger:response forumCreateCreated
*/
type ForumCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Forum `json:"body,omitempty"`
}

// NewForumCreateCreated creates ForumCreateCreated with default headers values
func NewForumCreateCreated() *ForumCreateCreated {

	return &ForumCreateCreated{}
}

// WithPayload adds the payload to the forum create created response
func (o *ForumCreateCreated) WithPayload(payload *models.Forum) *ForumCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forum create created response
func (o *ForumCreateCreated) SetPayload(payload *models.Forum) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForumCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ForumCreateNotFoundCode is the HTTP code returned for type ForumCreateNotFound
const ForumCreateNotFoundCode int = 404

/*ForumCreateNotFound Владелец форума не найден.


swagger:response forumCreateNotFound
*/
type ForumCreateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewForumCreateNotFound creates ForumCreateNotFound with default headers values
func NewForumCreateNotFound() *ForumCreateNotFound {

	return &ForumCreateNotFound{}
}

// WithPayload adds the payload to the forum create not found response
func (o *ForumCreateNotFound) WithPayload(payload *models.Error) *ForumCreateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forum create not found response
func (o *ForumCreateNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForumCreateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ForumCreateConflictCode is the HTTP code returned for type ForumCreateConflict
const ForumCreateConflictCode int = 409

/*ForumCreateConflict Форум уже присутсвует в базе данных.
Возвращает данные ранее созданного форума.


swagger:response forumCreateConflict
*/
type ForumCreateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Forum `json:"body,omitempty"`
}

// NewForumCreateConflict creates ForumCreateConflict with default headers values
func NewForumCreateConflict() *ForumCreateConflict {

	return &ForumCreateConflict{}
}

// WithPayload adds the payload to the forum create conflict response
func (o *ForumCreateConflict) WithPayload(payload *models.Forum) *ForumCreateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forum create conflict response
func (o *ForumCreateConflict) SetPayload(payload *models.Forum) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForumCreateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
