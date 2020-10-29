package forumRepository

import (
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

func (r *ForumRepository) CreateForum(forum models.Forum) (models.Forum, error) {
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
