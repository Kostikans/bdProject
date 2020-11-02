package threadDelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	r.HandleFunc("/api/post/{id}/details", handler.PostInfo).Methods("GET")
	r.HandleFunc("/api/post/{id}/details", handler.ChangeMsg).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/create", handler.CreatePost).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/details", handler.ThreadInformation).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/details", handler.UpdateThread).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/posts", handler.ThreadMsgs).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/vote", handler.ThreadVote).Methods("POST")
}

func (h *ThreadHandler) PostInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVar := vars["id"]
	id, _ := strconv.Atoi(idVar)

	post := models.PostUpdate{}
	json.NewDecoder(r.Body).Decode(&post)

	postE, err := h.ThreadUseCase.PostInfo(id)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&postE)
}

func (h *ThreadHandler) ChangeMsg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVar := vars["id"]
	id, _ := strconv.Atoi(idVar)

	post := models.PostUpdate{}
	json.NewDecoder(r.Body).Decode(&post)

	postE, err := h.ThreadUseCase.PostUpdate(id, post)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&postE)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)

}

func (h *ThreadHandler) ThreadInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug_or_id"]

	thread, err := h.ThreadUseCase.GetThreadInformation(slug)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&thread)

}

func (h *ThreadHandler) UpdateThread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	slug := vars["slug_or_id"]

	thread := models.Thread{}
	json.NewDecoder(r.Body).Decode(&thread)
	posts, err := h.ThreadUseCase.ChangeThread(slug, thread)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&posts)
}

func (h *ThreadHandler) ThreadMsgs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	values := r.URL.Query()
	slug := vars["slug_or_id"]
	sinceVar := values.Get("since")
	since, _ := strconv.Atoi(sinceVar)
	limitVar := values.Get("limit")
	limit, _ := strconv.Atoi(limitVar)
	descVar := values.Get("desc")
	desc, _ := strconv.ParseBool(descVar)
	sort := values.Get("sort")

	posts, err := h.ThreadUseCase.GetThreadPosts(slug, limit, since, sort, desc)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&posts)

}

func (h *ThreadHandler) ThreadVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug_or_id"]

	vote := models.Vote{}
	json.NewDecoder(r.Body).Decode(&vote)

	users, err := h.ThreadUseCase.Vote(slug, vote)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}
