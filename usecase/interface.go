package usecase

import "clean-architecture/entity"

//Repository interface
type Repository interface {
	Get(id int64) (*entity.User, error)
}

type UserUseCase interface {
	GetUser(id int64) (*entity.User, error)
}
