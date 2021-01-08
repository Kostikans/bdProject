package forumDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fasthttp/router"

	"github.com/valyala/fasthttp"

	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"
	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/forum"
)

type ForumHandler struct {
	ForumUseCase forum.UseCase
	log          *logger.CustomLogger
}

func NewForumHandler(r *router.Router, fs forum.UseCase, lg *logger.CustomLogger) {
	handler := ForumHandler{
		ForumUseCase: fs,
		log:          lg,
	}
	r.Handle("POST", "/api/forum/create", handler.CreateForum)
	r.Handle("GET", "/api/forum/{slug}/details", handler.GetForumInfo)
	r.Handle("POST", "/api/forum/{slug}/create", handler.CreateThread)
	r.Handle("GET", "/api/forum/{slug}/users", handler.GetForumUsers)
	r.Handle("GET", "/api/forum/{slug}/threads", handler.GetThreadsFromForum)
	r.Handle("GET", "/api/service/status", handler.GetStatusServer)
	r.Handle("POST", "/api/service/clear", handler.Clear)
}
func (h *ForumHandler) Clear(ctx *fasthttp.RequestCtx) {
	err := h.ForumUseCase.Clear()
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		return

	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
}
func (h *ForumHandler) CreateForum(ctx *fasthttp.RequestCtx) {

	forum := models.Forum{}
	json.Unmarshal(ctx.PostBody(), &forum)

	forum, err := h.ForumUseCase.CreateForum(forum)
	if err != nil {
		if customerror.ParseCode(err) == 409 {
			h.log.LogError(err)
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(customerror.ParseCode(err))
			json.NewEncoder(ctx.Response.BodyWriter()).Encode(&forum)
			return
		} else if customerror.ParseCode(err) == 404 {
			h.log.LogError(err)
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(customerror.ParseCode(err))
			err := models.Error{Message: "fdsfsd"}
			json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
			return
		}
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusCreated)

	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&forum)

}

func (h *ForumHandler) GetForumInfo(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug").(string)
	forum, err := h.ForumUseCase.GetForumInfo(slug)
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&forum)
	ctx.SetStatusCode(http.StatusOK)
}

func (h *ForumHandler) CreateThread(ctx *fasthttp.RequestCtx) {

	slug := ctx.UserValue("slug").(string)

	thread := models.Thread{}
	json.Unmarshal(ctx.PostBody(), &thread)

	thread, err := h.ForumUseCase.CreateThread(slug, thread)
	if err != nil {
		if customerror.ParseCode(err) == 404 {
			h.log.LogError(err)
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(customerror.ParseCode(err))
			err := models.Error{Message: "fdsfsd"}
			json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
			return
		} else {

			h.log.LogError(err)
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(customerror.ParseCode(err))
			json.NewEncoder(ctx.Response.BodyWriter()).Encode(&thread)
			return
		}
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusCreated)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&thread)
}

func (h *ForumHandler) GetForumUsers(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug").(string)
	since := string(ctx.URI().QueryArgs().Peek("since"))
	limitVar := string(ctx.URI().QueryArgs().Peek("limit"))
	limit, _ := strconv.Atoi(limitVar)
	descVar := string(ctx.URI().QueryArgs().Peek("desc"))
	desc, _ := strconv.ParseBool(descVar)

	users, err := h.ForumUseCase.GetForumUsers(slug, limit, since, desc)
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return

	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&users)
}

var count int

func (h *ForumHandler) GetThreadsFromForum(ctx *fasthttp.RequestCtx) {
	count++
	slug := ctx.UserValue("slug").(string)
	since := string(ctx.URI().QueryArgs().Peek("since"))
	limitVar := string(ctx.URI().QueryArgs().Peek("limit"))
	limit, _ := strconv.Atoi(limitVar)
	descVar := string(ctx.URI().QueryArgs().Peek("desc"))
	desc, _ := strconv.ParseBool(descVar)

	threads, err := h.ForumUseCase.GetThreadsFromForum(slug, limit, since, desc)
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return

	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&threads)
}

func (h *ForumHandler) GetStatusServer(ctx *fasthttp.RequestCtx) {

	threads, err := h.ForumUseCase.GetServerStatus()
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&threads)
}
