package thread

import "github.com/kostikans/bdProject/models"

type Repository interface {
	PostUpdate(id int, update models.PostUpdate) (models.Post, error)
	Postpost(slug_or_id string, posts []models.Post) ([]models.Post, error)
}
