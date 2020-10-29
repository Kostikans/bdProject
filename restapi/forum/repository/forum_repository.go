package forumRepository

import (
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"
	"github.com/kostikans/bdProject/restapi"
)

type ForumRepository struct {
	bd *sqlx.DB
}

func NewForumRepository(sqlx *sqlx.DB) *ForumRepository {
	return &ForumRepository{sqlx}
}

func (r *ForumRepository) CreateThread(slug string, thread models.Thread) (models.Thread, error) {
	forum := ""
	var forum_id int64
	row, _ := r.bd.Exec(restapi.CheckUserExist, thread.Author)
	r.bd.QueryRow(restapi.CheckForumExist, slug).Scan(&forum, &forum_id)
	count, _ := row.RowsAffected()
	if count == 0 || forum == "" {
		return thread, customerror.NewCustomError(errors.New(""), http.StatusNotFound, 1)
	}
	_, err := r.bd.Exec(restapi.CreateThreadRequest, thread.Title, thread.Author, thread.Message, thread.Created, forum_id)
	if err != nil {
		r.bd.QueryRow(restapi.GetExistThreadReuqest, thread.Title).Scan(&thread.ID, &thread.Title, &thread.Author,
			&thread.Message, &thread.Votes, &thread.Slug, &thread.Created)
		thread.Forum = forum
		return thread, customerror.NewCustomError(err, http.StatusConflict, 1)
	}
	thread.Forum = forum
	return thread, nil
}

func (r *ForumRepository) CreateForum(forum models.Forum) (models.Forum, error) {
	row, _ := r.bd.Exec(restapi.CheckUserExist, forum.User)
	count, _ := row.RowsAffected()
	if count == 0 {
		return forum, customerror.NewCustomError(errors.New(""), http.StatusNotFound, 1)
	}
	_, err := r.bd.Exec(restapi.PostForumRequest, forum.Slug, forum.Title, forum.User)
	if err != nil {
		forum, _ = r.GetForumInfo(forum.Slug)
		return forum, customerror.NewCustomError(err, http.StatusConflict, 1)

	}
	return forum, nil
}

func (r *ForumRepository) GetForumInfo(slug string) (models.Forum, error) {
	forum := models.Forum{}
	err := r.bd.QueryRow(restapi.GetForumInfoRequest, slug).Scan(&forum.Title, &forum.User, &forum.Slug, &forum.Posts, &forum.Threads)
	if err != nil {
		return forum, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return forum, nil
}

func (r *ForumRepository) GetForumUsers(slug string, limit int, since string, desc bool) ([]models.User, error) {
	users := []models.User{}
	err := r.bd.Select(&users, restapi.GetForumUsersRequest, slug, limit, since)
	if err != nil {
		return users, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return users, nil
}

func (r *ForumRepository) GetThreadsFromForum(slug string) ([]models.Thread, error) {
	threads := []models.Thread{}
	err := r.bd.Select(&threads, restapi.GetThreadsFromForum, slug)
	if err != nil {
		return threads, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	if len(threads) == 0 {
		return threads, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return threads, nil
}
