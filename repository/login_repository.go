package repository

import (
	"database/sql"
	"jabill-notes/models"
)

type LoginRepository struct{
	databaseConnection *sql.DB
}

func NewLoginRepository (databaseConnection *sql.DB) LoginRepository{
	return LoginRepository{
		databaseConnection: databaseConnection,
	}
}

func (loginRepository *LoginRepository) Verify(login models.Login) (string, error){
	loginQuery := "SELECT id, name, email, password FROM User WHERE email = ? AND password = ?"

	var user models.User
	err := loginRepository.databaseConnection.QueryRow(loginQuery, login.Email, login.Password).Scan(
		&user.Id,
        &user.Name,
        &user.Email,
        &user.Password,
	)

	if err != nil{
		return "", err
	}

	userToken, err := user.GenerateToken()
	
	if err != nil{
		return "", err
	}

	return userToken, nil
}