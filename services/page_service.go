package services

import (
	"jabill-notes/models"
	"jabill-notes/repository"
)

type PageService struct{
	pageRepository repository.PageRepository
}

func NewPageService(pageRepository repository.PageRepository) PageService{
	return PageService{
		pageRepository: pageRepository,
	}
}

func (pageService *PageService) Show (slug string, user_id int) (models.Page, error){
	return pageService.pageRepository.Show(slug, user_id)
}

func (pageService *PageService) Index (user_id int) ([]models.Page, error){
	return pageService.pageRepository.Index(user_id)
}

func (pageService *PageService) Store (page models.Page) (models.Page, error){
	return pageService.pageRepository.Store(page)
}

func (pageService *PageService) Update (field string, value string, slug string, user_id int) (error){
	return pageService.pageRepository.Update(field, value, slug, user_id)
}

func (pageService *PageService) UpdateTitle (title string, slug string, user_id int) (error){
	return pageService.pageRepository.UpdateTitle(title, slug, user_id)
}

func (pageService *PageService) UpdateContent (content string, slug string, user_id int) (error){
	return pageService.pageRepository.UpdateContent(content, slug, user_id)
}