package threadRepository

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"

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
	var forum_id int64
	var forum string
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		id = -1
	}
	err = r.bd.QueryRow(restapi.GetExistThreadToPostReuqest, slug_or_id, id).Scan(&thread_id, &forum_id, &forum)
	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}

	tx := r.bd.MustBegin()
	defer tx.Rollback()
	stmt, err := r.bd.Preparex(restapi.CreatePostRequest)
	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	create := strfmt.DateTime(time.Now())
	for index, _ := range posts {
		posts[index].Forum = forum
		posts[index].Created = &create
		posts[index].Thread = thread_id

		err := stmt.QueryRow(posts[index].Parent, posts[index].Author, posts[index].Message,
			forum, thread_id, forum_id, create, strconv.Itoa(int(posts[index].Parent))).Scan(&posts[index].ID)
		if err != nil {
			return posts, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}

	}
	return posts, nil
}

func (u *ThreadRepository) Vote(slug_or_id string, vote models.Vote) (models.Thread, error) {

	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		id = -1
	}
	thread := models.Thread{}

	err = u.bd.QueryRow(restapi.GetExistThreadReuqestToVote, slug_or_id, id).Scan(&thread.ID, &thread.Title, &thread.Author,
		&thread.Message, &thread.Votes, &thread.Forum, &thread.Slug, &thread.Created)

	if err != nil {
		return thread, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	var prevVoice int32
	err = u.bd.QueryRow(restapi.GetPrevVote, vote.Nickname, thread.ID).Scan(&prevVoice)

	if err == nil {
		_, err = u.bd.Exec(restapi.UpdateVote, vote.Voice, vote.Nickname, thread.ID)
		if err != nil {
			return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}
		if prevVoice == -1 && vote.Voice == 1 {
			err = u.bd.QueryRow(restapi.UpdateVoteCount, 2, thread.ID).Scan(&thread.Votes)
			thread.Votes += 2
			if err != nil {
				return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
			}
		}
		if prevVoice == 1 && vote.Voice == -1 {
			err = u.bd.QueryRow(restapi.UpdateVoteCount, -2, thread.ID).Scan(&thread.Votes)
			thread.Votes += -2
			if err != nil {
				return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
			}
		}
		if prevVoice == vote.Voice {
			fmt.Println("nothing to see here")
			return thread, nil
		}
		err = u.bd.QueryRow(restapi.UpdateVoteCount, vote.Voice, thread.ID).Scan(&thread.Votes)
		if err != nil {
			return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
		}

		return thread, nil
	}
	_, err = u.bd.Exec(restapi.AddVote, vote.Voice, vote.Nickname, thread.ID)
	if err != nil {

		return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}

	err = u.bd.QueryRow(restapi.UpdateVoteCount, vote.Voice, thread.ID).Scan(&thread.Votes)

	if err != nil {
		return thread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return thread, nil
}

func (u *ThreadRepository) GetThreadInformation(slug_or_id string) (models.Thread, error) {
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		id = -1
	}
	thread := models.Thread{}

	err = u.bd.QueryRow(restapi.GetExistThreadReuqestToVote, slug_or_id, id).Scan(&thread.ID, &thread.Title, &thread.Author,
		&thread.Message, &thread.Votes, &thread.Forum, &thread.Slug, &thread.Created)

	if err != nil {
		return thread, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return thread, nil
}

func (u *ThreadRepository) GetThreadPosts(slug_or_id string, limit int, since int, sort string, desc bool) ([]models.Post, error) {
	posts := []models.Post{}
	fmt.Println(slug_or_id)
	thread, err := u.GetThreadInformation(slug_or_id)
	if err != nil {
		return posts, err
	}
	if limit < 1 && limit > 10000 {
		limit = 100
	}
	var cursor string
	if since != 0 {
		fmt.Println("fds")
		number := strconv.Itoa(since)
		if desc == false {
			cursor = fmt.Sprintf(" and comm_id < %d ", number)
		} else {
			cursor = fmt.Sprintf(" and comm_id > %d ", number)
		}
	}
	var query string
	var order string
	if desc == false {
		order = " ASC "
	} else {
		order = " DESC "
	}
	var parent string
	if sort == "" || sort == "flat" {
		if desc == true {
			parent = "post_id DESC"
		} else {
			parent = "post_id ASC"
		}
	}
	var where string
	if sort == "tree" {
		if desc == true {

			parent = "parent_path DESC"
			where = " and  parent_path <@ '0' "

		} else {
			parent = "parent_path ASC"
			where = " and  parent_path <@ '0' "
		}
	}

	if sort == "parent_tree" {
		if desc == true {
			parent = "parent_path ASC"
			where = " and  parent_path <@ '0' "

		} else {
			parent = "parent_path ASC"
			where = " and  parent_path <@ '0' "
		}
	}
	fmt.Println(thread.ID, limit)
	query = fmt.Sprintf("SELECT author,created,forum,post_id,message,parent,thread_id"+
		" from posts WHERE thread_id=$1 %s %s ORDER BY  %s,created %s LIMIT $2", where, cursor, parent, order)

	fmt.Println(query)
	err = u.bd.Select(&posts, query, thread.ID, limit)
	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return posts, nil
}
