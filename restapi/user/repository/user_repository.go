package userRepository

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kostikans/bdProject/models"
	customerror "github.com/kostikans/bdProject/pkg/error"
	"github.com/kostikans/bdProject/restapi"
)

type UserRepository struct {
	bd *sqlx.DB
}

func NewUserRepository(sqlx *sqlx.DB) *UserRepository {
	return &UserRepository{sqlx}
}

func (r *UserRepository) CreateNewUser(user models.User) ([]models.User, error) {
	users := make([]models.User, 1, 1)
	_, err := r.bd.Exec(restapi.AddUserRequest, &user.Nickname, &user.Email, &user.Fullname, &user.About)
	users[0] = user
	if err != nil {
		users = []models.User{}
		r.bd.Select(&users, restapi.GetPreviousUsers, user.Nickname, user.Email)
		return users, customerror.NewCustomError(err, http.StatusConflict, 1)
	}
	return users, nil
}

func (r *UserRepository) GetUserInfo(nickname string) (models.User, error) {
	user := models.User{}
	err := r.bd.QueryRow(restapi.GetUserRequest, nickname).Scan(&user.Nickname, &user.Email, &user.Fullname, &user.About)
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) (models.User, error) {

	res, err := r.bd.Exec(restapi.UpdateUserRequest, user.Nickname, user.Email, user.Fullname, user.About)
	_, erro := res.RowsAffected()
	if erro != nil {
		return user, customerror.NewCustomError(erro, http.StatusNotFound, 1)
	}
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusConflict, 1)
	}
	return user, nil
}
