package thread

import "github.com/kostikans/bdProject/models"

type UseCase interface {
	PostUpdate(id int, update models.PostUpdate) (models.Post, error)
	Postpost(slug_or_id string, posts []models.Post) ([]models.Post, error)
	Vote(slug_or_id string, vote models.Vote) (models.Thread, error)
	GetThreadInformation(slug_or_id string) (models.Thread, error)
	GetThreadPosts(slug_or_id string, limit int, since int, sort string, desc bool) ([]models.Post, error)
}
