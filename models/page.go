package models

type Page struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Cape string `json:"cape"`
	Emoji string `json:"emoji"`
	Parent_id interface{} `json:"parent_id"`
	User_id int `json:"user_id"`
	Content string `json:"content"`
	Slug string `json:"slug"`
}