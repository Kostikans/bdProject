package threadRepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kostikans/bdProject/models"
)

type ThreadRepository struct {
	bd *sqlx.DB
}

func NewThreadRepository(sqlx *sqlx.DB) *ThreadRepository {
	return &ThreadRepository{sqlx}
}

func (r *ThreadRepository) CreateThread(thread models.Thread) (models.Thread, error) {
	return models.Thread{}, nil
}
func (r *ThreadRepository) PostUpdate(id int, update models.PostUpdate) (models.Post, error) {
	return models.Post{}, nil
}
