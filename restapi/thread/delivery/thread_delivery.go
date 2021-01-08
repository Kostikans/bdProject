package threadDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fasthttp/router"

	"github.com/valyala/fasthttp"

	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"

	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/thread"
)

type ThreadHandler struct {
	ThreadUseCase thread.UseCase
	log           *logger.CustomLogger
}

func NewThreadHandler(r *router.Router, fs thread.UseCase, lg *logger.CustomLogger) {
	handler := ThreadHandler{
		ThreadUseCase: fs,
		log:           lg,
	}

	r.Handle("GET", "/api/post/{id}/details", handler.PostInfo)
	r.Handle("POST", "/api/post/{id}/details", handler.ChangeMsg)
	r.Handle("POST", "/api/thread/{slug_or_id}/create", handler.CreatePost)
	r.Handle("GET", "/api/thread/{slug_or_id}/details", handler.ThreadInformation)
	r.Handle("POST", "/api/thread/{slug_or_id}/details", handler.UpdateThread)
	r.Handle("GET", "/api/thread/{slug_or_id}/posts", handler.ThreadMsgs)
	r.Handle("POST", "/api/thread/{slug_or_id}/vote", handler.ThreadVote)
}

func (h *ThreadHandler) PostInfo(ctx *fasthttp.RequestCtx) {
	idVar := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(idVar)
	related := string(ctx.URI().QueryArgs().Peek("related"))

	postE, err := h.ThreadUseCase.PostInfo(id, related)
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
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&postE)
}

func (h *ThreadHandler) ChangeMsg(ctx *fasthttp.RequestCtx) {
	idVar := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(idVar)
	post := models.PostUpdate{}
	json.Unmarshal(ctx.PostBody(), &post)

	postE, err := h.ThreadUseCase.PostUpdate(id, post)
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
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&postE)
}

func (h *ThreadHandler) CreatePost(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)

	posts := []models.Post{}
	json.Unmarshal(ctx.PostBody(), &posts)

	users, err := h.ThreadUseCase.Postpost(slug, posts)
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return

	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusCreated)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&users)

}

func (h *ThreadHandler) ThreadInformation(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)

	thread, err := h.ThreadUseCase.GetThreadInformation(slug)
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
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&thread)

}

func (h *ThreadHandler) UpdateThread(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)

	thread := models.Thread{}
	json.Unmarshal(ctx.PostBody(), &thread)
	posts, err := h.ThreadUseCase.ChangeThread(slug, thread)
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
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&posts)
}

func (h *ThreadHandler) ThreadMsgs(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)
	sinceVar := string(ctx.URI().QueryArgs().Peek("since"))
	since, _ := strconv.Atoi(sinceVar)
	limitVar := string(ctx.URI().QueryArgs().Peek("limit"))
	limit, _ := strconv.Atoi(limitVar)
	descVar := string(ctx.URI().QueryArgs().Peek("desc"))
	desc, _ := strconv.ParseBool(descVar)
	sort := string(ctx.URI().QueryArgs().Peek("sort"))

	posts, err := h.ThreadUseCase.GetThreadPosts(slug, limit, since, sort, desc)
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
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&posts)

}

func (h *ThreadHandler) ThreadVote(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)

	vote := models.Vote{}
	json.Unmarshal(ctx.PostBody(), &vote)
	users, err := h.ThreadUseCase.Vote(slug, vote)
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
