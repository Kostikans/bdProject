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
	err := r.bd.QueryRow(restapi.GetUserRequest, nickname).Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
	if err != nil {
		return user, customerror.NewCustomError(err, http.StatusNotFound, 1)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) (models.User, error) {
	usr, err := r.GetUserInfo(user.Nickname)
	if err != nil {
		return usr, customerror.NewCustomError(err, customerror.ParseCode(err), 1)
	}

	if user.Email == "" {
		user.Email = usr.Email
	}
	if user.Fullname == "" {
		user.Fullname = usr.Fullname
	}
	if user.About == "" {
		user.About = usr.About
	}

	_, err = r.bd.Exec(restapi.UpdateUserRequest, user.Nickname, user.Fullname, user.Email, user.About)
	if err != nil {
		return usr, customerror.NewCustomError(err, http.StatusConflict, 1)
	}

	return user, nil
}
