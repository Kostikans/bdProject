package forumDelivery

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"
	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/forum"
)

type ForumHandler struct {
	ForumUseCase forum.UseCase
	log          *logger.CustomLogger
}

func NewForumHandler(r *mux.Router, fs forum.UseCase, lg *logger.CustomLogger) {
	handler := ForumHandler{
		ForumUseCase: fs,
		log:          lg,
	}
	r.HandleFunc("/api/forum/create", handler.CreateForum).Methods("POST")
	r.HandleFunc("/api/forum/{slug}/info}", handler.GetForumInfo).Methods("GET")
	r.HandleFunc("/api/forum/{slug}/create", handler.CreateThread).Methods("POST")
}

func (h *ForumHandler) CreateForum(w http.ResponseWriter, r *http.Request) {

	forum := models.Forum{}
	json.NewDecoder(r.Body).Decode(&forum)

	forum, err := h.ForumUseCase.CreateForum(forum)
	if err != nil {
		if customerror.ParseCode(err) == 409 {
			h.log.LogError(r.Context(), err)
			w.WriteHeader(customerror.ParseCode(err))
			json.NewEncoder(w).Encode(&forum)
			return
		} else {
			h.log.LogError(r.Context(), err)
			w.WriteHeader(customerror.ParseCode(err))
			err := models.Error{Message: "fdsfsd"}
			json.NewEncoder(w).Encode(&err)
			return
		}
	}
	json.NewEncoder(w).Encode(&forum)
	w.WriteHeader(http.StatusOK)
}

func (h *ForumHandler) GetForumInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	forum, err := h.ForumUseCase.GetForumInfo(slug)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return
	}
	json.NewEncoder(w).Encode(&forum)
	w.WriteHeader(http.StatusOK)
}

func (h *ForumHandler) CreateThread(w http.ResponseWriter, r *http.Request) {

}
