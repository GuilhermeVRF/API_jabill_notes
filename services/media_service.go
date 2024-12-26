package services

import "jabill-notes/repository"

type MediaService struct {
	mediaRepository repository.MediaRepository
}

func NewMediaService(mediaRepository repository.MediaRepository) MediaService{
	return MediaService{
		mediaRepository : mediaRepository,
	}
}

func (mediaService *MediaService) GetUserProfile(user_id int) (string, error){
	return mediaService.mediaRepository.GetUserProfile(user_id)
}