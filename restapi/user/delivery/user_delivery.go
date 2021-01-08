package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/fasthttp/router"

	"github.com/valyala/fasthttp"

	customerror "github.com/kostikans/bdProject/pkg/error"

	"github.com/kostikans/bdProject/models"

	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/user"
)

type UserHandler struct {
	UserUseCase user.UseCase
	log         *logger.CustomLogger
}

func NewUserHandler(r *router.Router, us user.UseCase, lg *logger.CustomLogger) {
	handler := UserHandler{
		UserUseCase: us,
		log:         lg,
	}
	r.Handle("POST", "/api/user/{nickname}/create", handler.CreateUser)
	r.Handle("GET", "/api/user/{nickname}/profile", handler.GetUser)
	r.Handle("POST", "/api/user/{nickname}/profile", handler.UpdateUser)

}

func (h *UserHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("nickname").(string)

	user := models.User{}
	json.Unmarshal(ctx.PostBody(), &user)
	user.Nickname = name
	users, err := h.UserUseCase.CreateNewUser(user)
	if err != nil {
		h.log.LogError(err)
		ctx.SetStatusCode(customerror.ParseCode(err))
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&users)
		return
	}
	user = users[0]
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusCreated)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&user)

}

func (h *UserHandler) GetUser(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("nickname").(string)

	user, err := h.UserUseCase.GetUserInfo(name)
	if err != nil {
		errMsg := models.Error{Message: "fdsfsd"}
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&errMsg)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&user)
	ctx.SetStatusCode(http.StatusOK)
}

func (h *UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("nickname").(string)

	user := models.User{}
	json.Unmarshal(ctx.PostBody(), &user)
	user.Nickname = name
	user, err := h.UserUseCase.UpdateUser(user)
	if err != nil {
		h.log.LogError(err)
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(&err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(&user)
	ctx.SetStatusCode(http.StatusOK)
}
