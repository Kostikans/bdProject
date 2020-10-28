// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kostikans/bdProject/models"
)

// ForumGetUsersOKCode is the HTTP code returned for type ForumGetUsersOK
const ForumGetUsersOKCode int = 200

/*ForumGetUsersOK Информация о пользователях форума.


swagger:response forumGetUsersOK
*/
type ForumGetUsersOK struct {

	/*
	  In: Body
	*/
	Payload models.Users `json:"body,omitempty"`
}

// NewForumGetUsersOK creates ForumGetUsersOK with default headers values
func NewForumGetUsersOK() *ForumGetUsersOK {

	return &ForumGetUsersOK{}
}

// WithPayload adds the payload to the forum get users o k response
func (o *ForumGetUsersOK) WithPayload(payload models.Users) *ForumGetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forum get users o k response
func (o *ForumGetUsersOK) SetPayload(payload models.Users) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForumGetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Users{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ForumGetUsersNotFoundCode is the HTTP code returned for type ForumGetUsersNotFound
const ForumGetUsersNotFoundCode int = 404

/*ForumGetUsersNotFound Форум отсутсвует в системе.


swagger:response forumGetUsersNotFound
*/
type ForumGetUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewForumGetUsersNotFound creates ForumGetUsersNotFound with default headers values
func NewForumGetUsersNotFound() *ForumGetUsersNotFound {

	return &ForumGetUsersNotFound{}
}

// WithPayload adds the payload to the forum get users not found response
func (o *ForumGetUsersNotFound) WithPayload(payload *models.Error) *ForumGetUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forum get users not found response
func (o *ForumGetUsersNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForumGetUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
