package repository

import (
	"HTTP31/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserAll interface {
	CreateUser(user *model.User) (int, error)
	MakeFriends(user *model.User) (string, error)
	GetAll() ([]model.UserGet, error)
	GetById(id int) ([]string, error)
	UpdateUser(id int, input model.UserUpdate) (int, error)
	DeleteUser(id int) (int, error)
}
type Repository struct {
	UserAll
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserAll: NewStorage(db),
	}
}
