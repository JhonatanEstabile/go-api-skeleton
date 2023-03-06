package user

import domain "api/domain/user"

type Service struct {
	repo Repository
	log  Logger
}

func NewService(
	logger Logger,
	repository Repository,
) *Service {
	return &Service{
		log:  logger,
		repo: repository,
	}
}

func (s *Service) Create(user *domain.User) error {
	return nil

	//_, err := c.repository.Create(&user)
	//
	//if err != nil {
	//	return err
	//}
}
