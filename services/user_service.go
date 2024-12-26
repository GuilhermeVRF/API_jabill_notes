package services

import (
	"jabill-notes/models"
	"jabill-notes/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) Index() ([]models.User, error) {
	return userService.userRepository.Index()
}

func (userService *UserService) Show(id int) (models.User, error) {
	return userService.userRepository.Show(id)
}

func (userService *UserService) CreateUser(user models.User) (models.User, error) {
	userId, err := userService.userRepository.InsertUser(user)
	if err != nil {
		return models.User{}, err
	}

	user.Id = int(userId)
	return user, nil
}

func (userService *UserService) UpdateUser(user models.User) (models.User, error) {
	return userService.userRepository.UpdateUser(user)
}

func (userService *UserService) DeleteUser(id int) (bool, error) {
	return userService.userRepository.DeleteUser(id)
}