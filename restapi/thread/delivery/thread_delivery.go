package threadDelivery

import (
	"net/http"

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

}

func (h *ThreadHandler) ThreadInformation(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) UpdateThread(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) ThreadMsgs(w http.ResponseWriter, r *http.Request) {

}

func (h *ThreadHandler) ThreadVote(w http.ResponseWriter, r *http.Request) {

}
