package repository

import (
	"HTTP31/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sync"
)

type Storage struct {
	db *sqlx.DB
	mtx sync.RWMutex
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		mtx:  sync.RWMutex{},
	}
}

func (d *Storage) CreateUser (u *model.User) (int, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, age) values ($1, $2) RETURNING id", usersTable)

	row := d.db.QueryRow(query, u.Name, u.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Storage) MakeFriends (u *model.User) (string, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	userFriends:= make([]string, 0)
	var name string

	query := fmt.Sprintf(`SELECT friends FROM %s WHERE id=$1`, usersTable)

	if err := d.db.QueryRow(query, u.SourceId).Scan(pq.Array(&userFriends)); err != nil {
		return "" , fmt.Errorf("task with id=%d not found", u.SourceId)
	}

	queryTarget1 := fmt.Sprintf(`SELECT name FROM %s WHERE id=$1`, usersTable)
	err := d.db.QueryRow(queryTarget1, u.TargetId).Scan(&name)
	if err != nil {
		return "" , fmt.Errorf("task with id=%d not found", u.TargetId)
	}

	appendFriends := append(userFriends, name)

	query1 := fmt.Sprintf("UPDATE %s SET friends = $1 WHERE id=$2 ", usersTable)
	_, err = d.db.Exec(query1, pq.Array(appendFriends), u.SourceId)

	return fmt.Sprintf("id = %d and id = %d, friends", u.SourceId, u.TargetId), err
}

func (d *Storage) GetAll () ([]model.UserGet, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var user []model.UserGet

	query := fmt.Sprintf(`SELECT name, age, id, friends FROM %s`, usersTable)

	if err := d.db.Select(&user, query); err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Storage) GetById (id int) ([]string, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var userFriends []string

	query := fmt.Sprintf(`SELECT friends FROM %s WHERE id=$1`, usersTable)

	if err := d.db.QueryRow(query, id).Scan(pq.Array(&userFriends)); err != nil {
		return nil, err
	}

	return userFriends, nil
}

func (d *Storage) UpdateUser (id int, input model.UserUpdate) (int, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	newAge:= input.NewAge

	query := fmt.Sprintf("UPDATE %s SET age = %d WHERE id=$1 ", usersTable, newAge)
	_, err := d.db.Exec(query, id)
	return id, err
}


func (d *Storage) DeleteUser (id int) (int, error){
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := d.db.Exec(query, id)

	return id, err
}