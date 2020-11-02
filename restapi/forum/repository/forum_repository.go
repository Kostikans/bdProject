package forumRepository

import (
	"errors"
	"fmt"
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

	thread.Forum = forum
	thread.Forum_ID = forum_id
	err := r.bd.QueryRow(restapi.CreateThreadRequest, thread.Title, thread.Author, thread.Message, thread.Created, thread.ID, thread.Forum_ID, thread.Forum, thread.Slug).Scan(&thread.ID)
	if err != nil {
		r.bd.QueryRow(restapi.GetExistThreadReuqest, thread.Slug).Scan(&thread.ID, &thread.Title, &thread.Author,
			&thread.Message, &thread.Votes, &thread.Slug, &thread.Created, &thread.Forum)

		return thread, customerror.NewCustomError(err, http.StatusConflict, 1)
	}
	return thread, nil
}

func (r *ForumRepository) CreateForum(forum models.Forum) (models.Forum, error) {
	err := r.bd.QueryRow(restapi.CheckUserExist, forum.User).Scan(&forum.User)
	if err != nil {
		return forum, customerror.NewCustomError(errors.New(""), http.StatusNotFound, 1)
	}
	_, err = r.bd.Exec(restapi.PostForumRequest, forum.Slug, forum.Title, forum.User)
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
	query := restapi.GetForumUsersRequest
	if limit < 1 || limit >= 10000 {
		limit = 100
	}
	var err error
	_, err = r.GetForumInfo(slug)
	if err != nil {
		return users, err
	}
	fmt.Println(slug, limit, since, desc)
	if desc == false {
		if since != "" {
			preQuery := ` AND lower(u.nickname) > lower($2) COLLATE "C" ORDER BY u.nickname  COLLATE "C" ASC LIMIT $3`
			query += preQuery
			err = r.bd.Select(&users, query, slug, since, limit)
		} else {
			preQuery := `ORDER BY u.nickname  COLLATE "C" LIMIT $2`
			query += preQuery
			err = r.bd.Select(&users, query, slug, limit)
		}
	} else {
		if since != "" {
			preQuery := ` AND lower(u.nickname) < lower($2)  COLLATE "C" ORDER BY u.nickname COLLATE "C" DESC LIMIT $3`
			query += preQuery
			err = r.bd.Select(&users, query, slug, since, limit)
		} else {
			preQuery := ` ORDER BY u.nickname  COLLATE "C" DESC LIMIT $2`
			query += preQuery
			err = r.bd.Select(&users, query, slug, limit)
		}
	}

	if err != nil {
		return users, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	fmt.Println(query)
	usersNew := []models.User{}
	err = r.bd.Select(&usersNew, restapi.GetVoteUsersRequest, slug)

	if err != nil {
		return users, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	for _, usr := range usersNew {
		users = append(users, usr)
	}
	return users, nil
}

func (r *ForumRepository) GetThreadsFromForum(slug string, limit int, since string, desc bool) ([]models.Thread, error) {
	threads := []models.Thread{}
	if limit < 1 && limit > 10000 {
		limit = 100
	}
	var query string

	_, err := r.GetForumInfo(slug)
	if err != nil {

		return threads, err
	}
	fmt.Println(slug, limit, since, desc)
	if desc == false && since != "" {
		query = "SELECT t.forum,t.slug,thread_id,t.title,author,message,votes,t.slug,t.created " +
			"FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug = $1 and t.created >= $2 ORDER BY created ASC LIMIT $3"
		err = r.bd.Select(&threads, query, slug, since, limit)
		if err != nil {
			return threads, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
	}
	if desc == true && since != "" {
		query = "SELECT t.forum,t.slug,thread_id,t.title,author,message,votes,t.slug,t.created " +
			"FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug = $1 and t.created <= $2 ORDER BY created DESC LIMIT $3"
		err = r.bd.Select(&threads, query, slug, since, limit)
		if err != nil {
			return threads, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
	}
	if desc == true && since == "" {
		query = "SELECT t.forum,t.slug,thread_id,t.title,author,message,votes,t.slug,created " +
			"FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug=$1 ORDER BY created DESC LIMIT $2"

		err = r.bd.Select(&threads, query, slug, limit)
		if err != nil {
			return threads, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
	}

	if desc == false && since == "" {
		query = "SELECT t.forum,t.slug,thread_id,t.title,author,message,votes,t.slug,created " +
			"FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug=$1  ORDER BY created ASC LIMIT $2"
		err = r.bd.Select(&threads, query, slug, limit)
		if err != nil {
			return threads, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
	}

	return threads, nil
}
