package repository

import (
	"database/sql"
	"jabill-notes/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct{
	databaseConnection *sql.DB	
}

func NewUserRepository(databaseConnection *sql.DB) UserRepository {
    return UserRepository{
        databaseConnection: databaseConnection,
    }
}

func (userRepository *UserRepository) Index() ([]models.User, error) {
	usersQuery := "SELECT * FROM User"
	userResults, err := userRepository.databaseConnection.Query(usersQuery)

	if err != nil{
		return []models.User{}, err
	}

	var users []models.User
	var user models.User

	for userResults.Next(){
		err := userResults.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Profile_image)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	userResults.Close()
	return users, nil
}

func (userRepository *UserRepository) Show(id int) (models.User, error) {
	userQuery := "SELECT * FROM User WHERE id = ?"
	userResult := userRepository.databaseConnection.QueryRow(userQuery, id)

	var user models.User
	err := userResult.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Profile_image)

	user.Profile_image = "http://localhost:5000/" + user.Profile_image
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (userRepository* UserRepository) InsertUser(user models.User) (int64, error) {
	insertQuery := "INSERT INTO User (name, email, password, profile_image) VALUES (?, ?, ?, ?)"
	
	var userId int64
	userInsertResult, err := userRepository.databaseConnection.Exec(insertQuery, user.Name, user.Email, user.Password, user.Profile_image)
    if err != nil {
        return 0, err
    }
	
	userId, err = userInsertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (userRepository *UserRepository) UpdateUser(user models.User) (models.User, error) {
	updateQuery := "UPDATE User SET name = ?, email = ?, password = ?, profile_image = ? WHERE id = ?"
	_, err := userRepository.databaseConnection.Exec(updateQuery, user.Name, user.Email, user.Password, user.Profile_image, user.Id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (userRepository *UserRepository) DeleteUser(id int) (bool, error){
	deleteQuery := "DELETE FROM User WHERE id = ?"
	_, err := userRepository.databaseConnection.Exec(deleteQuery, id)

	if err != nil{
		return false, err
	}

	return true, nil
}