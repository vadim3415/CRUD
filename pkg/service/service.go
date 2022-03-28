package service

import (
	"HTTP31/pkg/model"
	"HTTP31/pkg/repository"
)

type UserAll interface {
	CreateUser(user *model.User) (int, error)
	MakeFriends(user *model.User) (string, error)
	GetAll() ([]model.UserGet, error)
	GetById(id int) ([]string, error)
	UpdateUser(id int, input model.UserUpdate) (int, error)
	DeleteUser(id int) (int, error)
}
type Service struct {
	UserAll
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserAll: NewUserService(repos.UserAll),
	}
}
