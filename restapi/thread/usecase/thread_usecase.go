package threadUsecase

import (
	"github.com/kostikans/bdProject/models"
	"github.com/kostikans/bdProject/restapi/thread"
)

type ThreadUseCase struct {
	repository thread.Repository
}

func NewThreadUseCase(repo thread.Repository) *ThreadUseCase {
	return &ThreadUseCase{repo}
}

func (u *ThreadUseCase) PostUpdate(id int, update models.PostUpdate) (models.Post, error) {
	return u.repository.PostUpdate(id, update)
}

func (u *ThreadUseCase) Postpost(slug_or_id string, posts []models.Post) ([]models.Post, error) {
	return u.repository.Postpost(slug_or_id, posts)
}

func (u *ThreadUseCase) Vote(slug_or_id string, vote models.Vote) (models.Thread, error) {
	return u.repository.Vote(slug_or_id, vote)
}
func (u *ThreadUseCase) GetThreadInformation(slug_or_id string) (models.Thread, error) {
	return u.repository.GetThreadInformation(slug_or_id)
}
func (u *ThreadUseCase) GetThreadPosts(slug_or_id string, limit int, since int, sort string, desc bool) ([]models.Post, error) {
	return u.repository.GetThreadPosts(slug_or_id, limit, since, sort, desc)
}
