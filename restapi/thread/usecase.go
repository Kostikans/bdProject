package thread

import "github.com/kostikans/bdProject/models"

type UseCase interface {
	CreateThread(thread models.Thread) (models.Thread, error)
	PostUpdate(id int, update models.PostUpdate) (models.Post, error)
}
