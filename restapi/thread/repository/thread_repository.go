package threadRepository

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"
	"github.com/kostikans/bdProject/restapi"
)

type ThreadRepository struct {
	bd *sqlx.DB
}

func NewThreadRepository(sqlx *sqlx.DB) *ThreadRepository {
	return &ThreadRepository{sqlx}
}

func (r *ThreadRepository) PostUpdate(id int, update models.PostUpdate) (models.Post, error) {
	return models.Post{}, nil
}

func (r *ThreadRepository) Postpost(slug_or_id string, posts []models.Post) ([]models.Post, error) {
	var thread_id int32
	thread_id = -1
	var forum_id int64
	var forum string
	thrId, err := strconv.Atoi(slug_or_id)
	if err != nil {
		thrId = -1
	}
	err = r.bd.QueryRow(restapi.GetExistThreadToPostReuqest, slug_or_id, thrId).Scan(&thread_id, &forum_id, &forum)

	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	tx := r.bd.MustBegin()
	defer tx.Rollback()
	stmt, err := r.bd.Preparex(restapi.CreatePostRequest)
	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	time := time.Now()
	for _, item := range posts {
		stmt.MustExec(item.Parent, item.Author, item.Message, forum, thread_id, forum_id, time)

	}
	return posts, nil
}
