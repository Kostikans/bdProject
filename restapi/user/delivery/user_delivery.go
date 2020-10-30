package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	customerror "github.com/kostikans/bdProject/pkg/error"

	"github.com/kostikans/bdProject/models"

	"github.com/gorilla/mux"
	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/user"
)

type UserHandler struct {
	UserUseCase user.UseCase
	log         *logger.CustomLogger
}

func NewUserHandler(r *mux.Router, us user.UseCase, lg *logger.CustomLogger) {
	handler := UserHandler{
		UserUseCase: us,
		log:         lg,
	}
	r.HandleFunc("/api/user/{nickname}/create", handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/{nickname}/profile", handler.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{nickname}/profile", handler.UpdateUser).Methods("POST")

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["nickname"]

	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Nickname = name
	users, err := h.UserUseCase.CreateNewUser(user)
	if err != nil {
		fmt.Println(users)
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		json.NewEncoder(w).Encode(&users)
		return
	}
	user = users[0]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["nickname"]

	user, err := h.UserUseCase.GetUserInfo(name)
	if err != nil {
		errMsg := models.Error{Message: "fdsfsd"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		json.NewEncoder(w).Encode(&errMsg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["nickname"]

	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Nickname = name
	user, err := h.UserUseCase.UpdateUser(user)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}
