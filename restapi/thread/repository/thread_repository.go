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

func (r *ThreadRepository) PostInfo(id int) (models.PostFull, error) {
	full := models.PostFull{}
	post := models.Post{}
	err := r.bd.QueryRow(restapi.CheckPostExist, id).Scan(&post.ID, &post.Author,
		&post.Message, &post.Created, &post.Forum, &post.IsEdited, &post.Parent, &post.Thread)
	if err != nil {
		return full, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	thread := models.Thread{}
	err = r.bd.QueryRow(restapi.GetExistThreadByIdToVote, post.Thread).Scan(
		&thread.ID, &thread.Title, &thread.Author, &thread.Message, &thread.Votes, &thread.Forum, &thread.Slug, &thread.Created, &thread.Forum_ID)
	if err != nil {
		return full, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	forum := models.Forum{}
	err = r.bd.QueryRow(restapi.GetForumInfoRequest, thread.Forum).Scan(
		&forum.Title, &forum.User, &forum.Slug, &forum.Posts, &forum.Threads)
	if err != nil {
		return full, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	author := models.User{}
	err = r.bd.QueryRow(restapi.GetUserRequest, thread.Author).Scan(
		&author.Nickname, &author.Fullname, &author.Email, &author.About)
	if err != nil {
		return full, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	full.Thread = &thread
	full.Forum = &forum
	full.Author = &author
	full.Post = &post
	return full, nil
}

func (r *ThreadRepository) PostUpdate(id int, update models.PostUpdate) (models.Post, error) {
	post := models.Post{}
	err := r.bd.QueryRow(restapi.CheckPostExist, id).Scan(&post.ID, &post.Author,
		&post.Message, &post.Created, &post.Forum, &post.IsEdited, &post.Parent, &post.Thread)
	if err != nil {
		return post, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	_, err = r.bd.Exec(restapi.PostUpdate, update.Message, id)
	if err != nil {
		return post, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return post, nil
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
			forum, thread_id, forum_id, create).Scan(&posts[index].ID)
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
	thread, err := u.GetThreadInformation(slug_or_id)
	if err != nil {
		return posts, err
	}
	if limit < 1 && limit > 10000 {
		limit = 100
	}
	query := ""
	if sort == "flat" || sort == "" {
		query = u.GenerateQueryFlatSort(slug_or_id, limit, since, sort, desc)
		if since == 0 {
			err = u.bd.Select(&posts, query, thread.ID, limit)
		} else {
			err = u.bd.Select(&posts, query, thread.ID, since, limit)
		}
	}
	if sort == "tree" {
		query = u.GenerateQueryTreeSort(slug_or_id, limit, since, sort, desc)
		if since == 0 {
			err = u.bd.Select(&posts, query, thread.ID, limit)
		} else {
			err = u.bd.Select(&posts, query, thread.ID, limit)
		}
	}
	if sort == "parent_tree" {
		query = u.GenerateQueryParentTreeSort(slug_or_id, limit, since, sort, desc)
		if since == 0 {
			err = u.bd.Select(&posts, query, thread.ID, limit)
		} else {
			err = u.bd.Select(&posts, query, thread.ID, limit)
		}
	}
	fmt.Println(query)
	if err != nil {
		return posts, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return posts, nil
}

func (u *ThreadRepository) GenerateQueryFlatSort(slug_or_id string, limit int, since int, sort string, desc bool) string {
	query := ""
	if desc == true {
		if since != 0 {
			query = fmt.Sprintf("SELECT author,created,forum,post_id,message,parent,thread_id from posts" +
				" WHERE thread_id=$1 and post_id < $2 ORDER BY post_id DESC LIMIT $3")
		} else {
			query = fmt.Sprintf("SELECT author,created,forum,post_id,message,parent,thread_id from posts" +
				" WHERE thread_id=$1 ORDER BY post_id DESC LIMIT $2")
		}
	} else {
		if since != 0 {
			query = fmt.Sprintf("SELECT author,created,forum,post_id,message,parent,thread_id from posts" +
				" WHERE thread_id=$1 and post_id > $2 ORDER BY post_id ASC LIMIT $3")
		} else {
			query = fmt.Sprintf("SELECT author,created,forum,post_id,message,parent,thread_id from posts" +
				" WHERE thread_id=$1 ORDER BY post_id ASC LIMIT $2")
		}
	}
	return query
}

func (u *ThreadRepository) GenerateQueryTreeSort(slug_or_id string, limit int, since int, sort string, desc bool) string {
	query := ""
	preQuery := ""
	if since != 0 {
		if desc == true {
			preQuery = "AND parents < "
		} else {
			preQuery = "AND parents > "
		}
		preQuery += fmt.Sprintf("(SELECT parents FROM posts WHERE post_id = %d)", since)
	}
	if desc == true {
		query = fmt.Sprintf(
			"SELECT author,created,forum,post_id,message,parent,thread_id FROM posts "+
				"WHERE thread_id=$1 %s ORDER BY parents DESC, post_id DESC LIMIT NULLIF($2, 0)", preQuery)
	} else {
		query = fmt.Sprintf(
			"SELECT author,created,forum,post_id,message,parent,thread_id FROM posts "+
				"WHERE thread_id=$1 %s ORDER BY parents, post_id LIMIT NULLIF($2, 0)", preQuery)
	}
	return query
}

func (u *ThreadRepository) GenerateQueryParentTreeSort(slug_or_id string, limit int, since int, sort string, desc bool) string {
	var preQuery = ""
	var query string

	if since != 0 {
		if desc == true {
			preQuery = fmt.Sprintf("AND parents[1] < ")
		} else {
			preQuery = fmt.Sprintf("AND parents[1] > ")
		}
		preQuery += fmt.Sprintf("(SELECT parents[1] FROM posts WHERE post_id = %d)", since)
	}

	preQuery2 := fmt.Sprintf(
		"SELECT post_id FROM posts WHERE thread_id = $1 AND parent = 0 %s", preQuery)

	if desc == true {
		preQuery2 += " ORDER BY post_id DESC LIMIT $2"
		query = fmt.Sprintf(
			"SELECT author,created,forum,post_id,message,parent,thread_id "+
				"FROM posts WHERE parents[1] IN (%s) ORDER BY parents[1] DESC, parents, post_id", preQuery2)
	} else {
		preQuery2 += " ORDER BY post_id ASC LIMIT $2"
		query = fmt.Sprintf(
			"SELECT author,created,forum,post_id,message,parent,thread_id "+
				"FROM posts WHERE parents[1] IN (%s) ORDER BY parents,post_id", preQuery2)
	}
	return query
}

func (u *ThreadRepository) ChangeThread(slug_or_id string, thread models.Thread) (models.Thread, error) {
	newThread, err := u.GetThreadInformation(slug_or_id)
	if err != nil {
		return newThread, err
	}
	newThread.Title = thread.Title
	newThread.Message = thread.Message
	_, err = u.bd.Exec(restapi.UpdateThread, newThread.Title, newThread.Message, newThread.ID)
	if err != nil {
		return newThread, customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}
	return newThread, nil
}
