package forumUsecase

import (
	"github.com/kostikans/bdProject/models"
	"github.com/kostikans/bdProject/restapi/forum"
)

type ForumUseCase struct {
	repository forum.Repository
}

func NewForumUseCase(repo forum.Repository) *ForumUseCase {
	return &ForumUseCase{repo}
}

func (u *ForumUseCase) CreateThread(slug string, thread models.Thread) (models.Thread, error) {
	return u.repository.CreateThread(slug, thread)
}

func (u *ForumUseCase) CreateForum(forum models.Forum) (models.Forum, error) {
	return u.repository.CreateForum(forum)
}
func (u *ForumUseCase) GetForumInfo(slug string) (models.Forum, error) {
	return u.repository.GetForumInfo(slug)
}
func (u *ForumUseCase) GetForumUsers(slug string, limit int, since string, desc bool) ([]models.User, error) {
	return u.repository.GetForumUsers(slug, limit, since, desc)
}

func (u *ForumUseCase) GetThreadsFromForum(slug string) ([]models.Thread, error) {
	return u.repository.GetThreadsFromForum(slug)
}
