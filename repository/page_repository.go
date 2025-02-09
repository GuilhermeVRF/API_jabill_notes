package repository

import (
	"database/sql"
	"errors"
	"jabill-notes/models"
	"log"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type PageRepository struct{
	databaseConnection *sql.DB
}

func NewPageRepository (databaseConnection *sql.DB) PageRepository{
	return PageRepository{
		databaseConnection: databaseConnection,
	}
}

func (pageRepository *PageRepository) Show(slug string, user_id int) (models.Page, error){
	pageQuery := "SELECT * FROM Page WHERE slug = ? AND user_id = ?"
	pageResult := pageRepository.databaseConnection.QueryRow(pageQuery, slug, user_id)

	log.Printf("Executando query: %s com slug: %s e user_id: %d\n", pageQuery, slug, user_id)

	var page models.Page
	err := pageResult.Scan(&page.Id, &page.Parent_id, &page.Title, &page.Cape, &page.Content, &page.Emoji, &page.Slug, &page.User_id, &page.Depth)
	if err != nil{
		log.Printf("Erro no Scan: %s", err)
		return models.Page{}, errors.New("Página solicitada não encontrada!")
	}

	return page, nil
}

func (pageRepository *PageRepository) Index (user_id int) ([]models.Page, error){
	pagesQuery := "SELECT id, title, emoji, parent_id, depth, slug FROM Page WHERE User_id = ? ORDER BY CASE WHEN parent_id IS NULL THEN id ELSE parent_id END, parent_id;"
	pagesResults, err := pageRepository.databaseConnection.Query(pagesQuery, user_id)

	if err != nil{
		return []models.Page{}, err
	}

	pages := make([]models.Page, 0)
	for pagesResults.Next() {
		var page models.Page
		err = pagesResults.Scan(&page.Id, &page.Title, &page.Emoji, &page.Parent_id, &page.Depth, &page.Slug)
		if err != nil {
			return []models.Page{}, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}


func (pageRepository *PageRepository) Store (page models.Page) (models.Page, error){
	slug, err := pageRepository.checkIfSlugExists(slug.Make(page.Title))
	page.Slug = slug
	if err != nil{
		return models.Page{}, err
	}

	var parent_id sql.NullInt64
	var depth int

	if page.Parent_id != nil {
		selectParentPageQuery := "SELECT id, depth FROM Page WHERE slug = ? AND user_id = ?"
		err = pageRepository.databaseConnection.QueryRow(selectParentPageQuery, page.Parent_id, page.User_id).Scan(&parent_id, &depth)
		
		if err != nil {
			return models.Page{}, err
		}
		depth++
	}
	
	pageQuery := "INSERT INTO Page (title, slug, emoji, parent_id, depth, content, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = pageRepository.databaseConnection.Exec(pageQuery, page.Title, page.Slug, "📃", parent_id, depth, "", page.User_id)

	if err != nil{
		return models.Page{}, err
	}
	
	page.Emoji = "📃"
    page.Depth = depth
	return page, nil
}

func (pageRepository *PageRepository) Update (field string, title string, slug string, user_id int) (error){
	titleQuery := "UPDATE Page SET "+ field +" = ? WHERE slug = ? AND user_id = ?"

	_, err := pageRepository.databaseConnection.Exec(titleQuery, title, slug, user_id)

	if err != nil{
		return err
	}

	return nil
}

func (pageRepository *PageRepository) Delete(slug string, user_id int) (error){
	deleteQuery := "DELETE FROM Page WHERE slug = ? AND user_id = ?"

	_, err := pageRepository.databaseConnection.Exec(deleteQuery, slug, user_id)

	if err != nil{
		return err
	}

	return nil
}

func (pageRepository *PageRepository) UpdateEmoji(emoji string, actualSlug string, user_id int) (string, error){
	emojiQuery := "UPDATE Page SET emoji = ? WHERE slug = ? AND user_id = ?"

	_, err := pageRepository.databaseConnection.Exec(emojiQuery, emoji, actualSlug, user_id)

	if err != nil{
		return "", err
	}

	return emoji, nil
}

func (pageRepository *PageRepository) UpdateTitle(title string, actualSlug string, user_id int) (string, string, error){
	newSlug, err := pageRepository.checkIfSlugExists(slug.Make(title))

	if err != nil{
		return "", "", errors.New("Não foi possível gerar o novo slug para a página!")
	}

	titleQuery := "UPDATE Page SET title = ?, slug = ? WHERE slug = ? AND user_id = ?"

	_, err = pageRepository.databaseConnection.Exec(titleQuery, title, newSlug, actualSlug, user_id)

	if err != nil{
		return "", "", err
	}

	return title, newSlug, nil
}

func (pageRepository *PageRepository) UpdateContent (content string, slug string, user_id int) (error){
	titleQuery := "UPDATE Page SET content = ? WHERE slug = ? AND user_id = ?"

	_, err := pageRepository.databaseConnection.Exec(titleQuery, content, slug, user_id)

	if err != nil{
		return err
	}

	return nil
}

func (pageRepository *PageRepository) checkIfSlugExists(slug string) (string, error){
	var slugExists int
	slugQuery := "SELECT COUNT(*) FROM Page WHERE slug = ?"
	err := pageRepository.databaseConnection.QueryRow(slugQuery, slug).Scan(&slugExists)

	if err != nil{
		return "", err
	}

	if slugExists > 0{
		return slug + uuid.New().String(), nil
	}

	return slug, nil
}