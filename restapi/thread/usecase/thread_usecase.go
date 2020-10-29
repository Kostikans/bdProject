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
