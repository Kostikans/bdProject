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

func (u *ThreadUseCase) CreateThread(thread models.Thread) (models.Thread, error) {
	return u.repository.CreateThread(thread)
}
func (u *ThreadUseCase) PostUpdate(id int, update models.PostUpdate) (models.Post, error) {
	return u.repository.PostUpdate(id, update)
}
