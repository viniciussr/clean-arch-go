package usecase

import "clean-architecture/entity"

type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUser(id int64) (*entity.User, error) {
	return s.repo.Get(id)
}
