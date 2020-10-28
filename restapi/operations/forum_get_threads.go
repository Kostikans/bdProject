// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ForumGetThreadsHandlerFunc turns a function with the right signature into a forum get threads handler
type ForumGetThreadsHandlerFunc func(ForumGetThreadsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ForumGetThreadsHandlerFunc) Handle(params ForumGetThreadsParams) middleware.Responder {
	return fn(params)
}

// ForumGetThreadsHandler interface for that can handle valid forum get threads params
type ForumGetThreadsHandler interface {
	Handle(ForumGetThreadsParams) middleware.Responder
}

// NewForumGetThreads creates a new http.Handler for the forum get threads operation
func NewForumGetThreads(ctx *middleware.Context, handler ForumGetThreadsHandler) *ForumGetThreads {
	return &ForumGetThreads{Context: ctx, Handler: handler}
}

/*ForumGetThreads swagger:route GET /forum/{slug}/threads forumGetThreads

Список ветвей обсужления форума

Получение списка ветвей обсужления данного форума.

Ветви обсуждения выводятся отсортированные по дате создания.


*/
type ForumGetThreads struct {
	Context *middleware.Context
	Handler ForumGetThreadsHandler
}

func (o *ForumGetThreads) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewForumGetThreadsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
