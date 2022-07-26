package repository

import (
	"clean-architecture/entity"
)

type UserInMemory struct {
	db []*entity.User
}

func NewUserInMemory() *UserInMemory {
	var users []*entity.User
	user := entity.NewUser(0, "teste@teste.com", "123", "Teste", "teste")
	users = append(users, user)

	return &UserInMemory{
		db: users,
	}
}

func (r *UserInMemory) Get(id int64) (*entity.User, error) {
	return getUser(id, r.db)
}

func getUser(id int64, db []*entity.User) (*entity.User, error) {
	return db[id], nil
}
