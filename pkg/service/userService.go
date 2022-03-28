package service

import (
	"HTTP31/pkg/model"
	"HTTP31/pkg/repository"
)

type UserService struct {
	repo repository.UserAll
}

func NewUserService(repo repository.UserAll) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAll() ([]model.UserGet, error) {
	return s.repo.GetAll()
}
func (s *UserService) GetById(id int) ([]string, error) {
	return s.repo.GetById(id)
}
func (s *UserService) MakeFriends(user *model.User) (string, error) {
	return s.repo.MakeFriends(user)
}

func (s *UserService) UpdateUser(id int, input model.UserUpdate) (int, error) {
	return s.repo.UpdateUser(id, input)
}

func (s *UserService) DeleteUser(id int) (int, error) {
	return s.repo.DeleteUser(id)
}

