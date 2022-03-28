package in_memory

import (
	"HTTP31/pkg/model"
	"fmt"
	"sync"
)

type Storage struct {
	data map[int]*model.User
	mtx  sync.RWMutex
}

func NewStorage(data map[int]*model.User) *Storage {
	return &Storage{
		data: data,
		mtx:  sync.RWMutex{},
	}
}

func (d *Storage) CreateUser(u model.User) (int, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	newId := 1
	for _, user := range d.data {
		if u.Name == user.Name {
			return 0, nil
		} else {
			newId += 1
		}
	}
	u.Id = newId
	d.data[u.Id] = &u
	return u.Id, nil
}

func (d *Storage) MakeFriends(u model.User) (string, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	newNameFriends := ""

	if _, ok := d.data[u.SourceId]; !ok {
		return "", fmt.Errorf("task with id=%d not found", u.SourceId)
	}
	if userTargetId, ok := d.data[u.TargetId]; !ok {
		return "", fmt.Errorf("task with id=%d not found", u.TargetId)
	} else {
		newNameFriends = userTargetId.Name
	}
	userSource := d.data[u.SourceId]

	for _, v := range userSource.Friends {
		if v == newNameFriends {
			return "the user has already been added to friends", nil
		}
	}
	appendFriends := append(userSource.Friends, newNameFriends)
	userSource.Friends = appendFriends
	d.data[u.SourceId] = userSource

	return fmt.Sprintf("id = %d and id = %d, friends", u.SourceId, u.TargetId), nil
}

func (d *Storage) GetAll() ([]*model.User, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	allUser := make([]*model.User, 0, len(d.data))
	for _, userS := range d.data {
		allUser = append(allUser, userS)
	}
	return allUser, nil
}

func (d *Storage) GetById(id int) ([]string, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var userFriends []string
	if userS, ok := d.data[id]; !ok {
		return nil, fmt.Errorf("task with id=%d not found", id)
	} else {
		userFriends = userS.Friends
	}
	return userFriends, nil
}

func (d *Storage) UpdateUser(u model.User) (string, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	if _, exist := d.data[u.Id]; !exist {
		return "", fmt.Errorf("task with id=%d not found", u.Id)
	}
	user := d.data[u.Id]
	user.Age = u.NewAge
	d.data[u.Id] = user
	name := user.Name
	return name, nil
}

func (d *Storage) DeleteUser(u model.User) (string, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	name := ""
	if userDelete, exist := d.data[u.TargetId]; !exist {
		return "", fmt.Errorf("task with id=%d not found", u.Id)
	} else {
		name = userDelete.Name
		delete(d.data, u.TargetId)
	}
	return name, nil
}
