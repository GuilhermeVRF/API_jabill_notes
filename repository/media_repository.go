package repository

import (
	"database/sql"
	"encoding/base64"
	"os"
)

type MediaRepository struct{
	databaseConnection *sql.DB
}

func NewMediaRepository (databaseConnection *sql.DB) MediaRepository{
	return MediaRepository{
		databaseConnection: databaseConnection,
	}
}

func (mediaRepository *MediaRepository) GetUserProfile(user_id int) (string, error){
	userProfileQuery := "SELECT profile_image FROM User WHERE id = ?"

	var imagePath string
	err := mediaRepository.databaseConnection.QueryRow(userProfileQuery, user_id).Scan(&imagePath)

	if err != nil{
		return "", err
	}

	binaryImage, err := os.ReadFile(imagePath)

	if err != nil{
		return "", err
	}

	encodedImage := base64.StdEncoding.EncodeToString(binaryImage)

	return "data:image/jpeg;base64," + encodedImage, nil;
}