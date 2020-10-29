package forum

import "github.com/kostikans/bdProject/models"

type Repository interface {
	CreateForum(forum models.Forum) (models.Forum, error)
	GetForumInfo(slug string) (models.Forum, error)
	GetForumUsers(slug string, limit int, since string, desc bool) ([]models.User, error)
}
