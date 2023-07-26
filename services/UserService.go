package services

import (
	"diluan/entities"
	"diluan/repositories"
)

type UserService interface{
	FindById(id uint64) entities.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func(s *userService) FindById(id uint64) entities.User {
	return s.userRepository.FindById(id)
}