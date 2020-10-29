package threadDelivery

import (
	"encoding/json"
	"net/http"

	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"

	"github.com/gorilla/mux"
	"github.com/kostikans/bdProject/pkg/logger"
	"github.com/kostikans/bdProject/restapi/thread"
)

type ThreadHandler struct {
	ThreadUseCase thread.UseCase
	log           *logger.CustomLogger
}

func NewThreadHandler(r *mux.Router, fs thread.UseCase, lg *logger.CustomLogger) {
	handler := ThreadHandler{
		ThreadUseCase: fs,
		log:           lg,
	}

	r.HandleFunc("/api/post/{id}/details", handler.ChangeMsg).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/create", handler.CreatePost).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/details", handler.ThreadInformation).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/details", handler.UpdateThread).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/posts", handler.ThreadMsgs).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/vote", handler.ThreadVote).Methods("POST")
}

func (h *ThreadHandler) ChangeMsg(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug_or_id"]

	posts := []models.Post{}
	json.NewDecoder(r.Body).Decode(&posts)

	users, err := h.ThreadUseCase.Postpost(slug, posts)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	json.NewEncoder(w).Encode(&users)
	w.WriteHeader(http.StatusOK)

}

func (h *ThreadHandler) ThreadInformation(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) UpdateThread(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) ThreadMsgs(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) ThreadVote(w http.ResponseWriter, r *http.Request) {

}
