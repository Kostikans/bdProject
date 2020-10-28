// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ForumGetUsersHandlerFunc turns a function with the right signature into a forum get users handler
type ForumGetUsersHandlerFunc func(ForumGetUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ForumGetUsersHandlerFunc) Handle(params ForumGetUsersParams) middleware.Responder {
	return fn(params)
}

// ForumGetUsersHandler interface for that can handle valid forum get users params
type ForumGetUsersHandler interface {
	Handle(ForumGetUsersParams) middleware.Responder
}

// NewForumGetUsers creates a new http.Handler for the forum get users operation
func NewForumGetUsers(ctx *middleware.Context, handler ForumGetUsersHandler) *ForumGetUsers {
	return &ForumGetUsers{Context: ctx, Handler: handler}
}

/*ForumGetUsers swagger:route GET /forum/{slug}/users forumGetUsers

Пользователи данного форума

Получение списка пользователей, у которых есть пост или ветка обсуждения в данном форуме.

Пользователи выводятся отсортированные по nickname в порядке возрастания.
Порядок сотрировки должен соответсвовать побайтовому сравнение в нижнем регистре.


*/
type ForumGetUsers struct {
	Context *middleware.Context
	Handler ForumGetUsersHandler
}

func (o *ForumGetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewForumGetUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
