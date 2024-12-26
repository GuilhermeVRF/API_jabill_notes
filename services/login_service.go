package services

import (
	"jabill-notes/models"
	"jabill-notes/repository"
)

type LoginService struct{
	loginRepository repository.LoginRepository
}

func NewLoginService (loginRepository repository.LoginRepository) LoginService{
	return LoginService{
		loginRepository: loginRepository,
	}
}

func (loginService *LoginService) Verify(login models.Login) (string, error){
	return loginService.loginRepository.Verify(login)
}