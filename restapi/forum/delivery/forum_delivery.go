package forumDelivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	r.HandleFunc("/api/forum/{slug}/details", handler.GetForumInfo).Methods("GET")
	r.HandleFunc("/api/forum/{slug}/create", handler.CreateThread).Methods("POST")
	r.HandleFunc("/api/forum/{slug}/users", handler.GetForumUsers).Methods("GET")
	r.HandleFunc("/api/forum/{slug}/threads", handler.GetThreadsFromForum).Methods("GET")
}

func (h *ForumHandler) CreateForum(w http.ResponseWriter, r *http.Request) {

	forum := models.Forum{}
	json.NewDecoder(r.Body).Decode(&forum)

	forum, err := h.ForumUseCase.CreateForum(forum)
	if err != nil {
		if customerror.ParseCode(err) == 409 {
			h.log.LogError(r.Context(), err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customerror.ParseCode(err))
			json.NewEncoder(w).Encode(&forum)
			return
		} else if customerror.ParseCode(err) == 404 {
			h.log.LogError(r.Context(), err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customerror.ParseCode(err))
			err := models.Error{Message: "fdsfsd"}
			json.NewEncoder(w).Encode(&err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&forum)

}

func (h *ForumHandler) GetForumInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	forum, err := h.ForumUseCase.GetForumInfo(slug)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&forum)
	w.WriteHeader(http.StatusOK)
}

func (h *ForumHandler) CreateThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE THREAD")
	vars := mux.Vars(r)
	slug := vars["slug"]

	thread := models.Thread{}
	json.NewDecoder(r.Body).Decode(&thread)

	thread, err := h.ForumUseCase.CreateThread(slug, thread)
	if err != nil {
		if customerror.ParseCode(err) == 404 {
			h.log.LogError(r.Context(), err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customerror.ParseCode(err))
			err := models.Error{Message: "fdsfsd"}
			json.NewEncoder(w).Encode(&err)
			return
		} else {

			h.log.LogError(r.Context(), err)
			w.WriteHeader(customerror.ParseCode(err))
			json.NewEncoder(w).Encode(&thread)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&thread)
}

func (h *ForumHandler) GetForumUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	values := r.URL.Query()
	slug := vars["slug"]
	since := values.Get("since")
	limitVar := values.Get("limit")
	limit, _ := strconv.Atoi(limitVar)
	descVar := values.Get("desc")
	desc, _ := strconv.ParseBool(descVar)

	users, err := h.ForumUseCase.GetForumUsers(slug, limit, since, desc)
	if err != nil {
		h.log.LogError(r.Context(), err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customerror.ParseCode(err))
		err := models.Error{Message: "fdsfsd"}
		json.NewEncoder(w).Encode(&err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
	w.WriteHeader(http.StatusOK)
}

var count int

func (h *ForumHandler) GetThreadsFromForum(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Println(count)
	vars := mux.Vars(r)
	values := r.URL.Query()
	slug := vars["slug"]
	since := values.Get("since")
	limitVar := values.Get("limit")
	limit, _ := strconv.Atoi(limitVar)
	descVar := values.Get("desc")
	desc, _ := strconv.ParseBool(descVar)

	threads, err := h.ForumUseCase.GetThreadsFromForum(slug, limit, since, desc)
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
	json.NewEncoder(w).Encode(&threads)
}
