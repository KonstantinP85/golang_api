package api_go

type Book struct {
	Id        int `json:"id"`
	AuthorId  int `json:"author_id"`
	Title     string `json:"title"`
	Years     int `json:"years"`
}

type BookInput struct {
	AuthorId  int `json:"author_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Years     int `json:"years" binding:"required"`
}

type BookResponse struct {
	Id        int `json:"id"`
	Author    *Author
	Title     string `json:"title"`
	Years     int `json:"years"`
}

type BookInAuthor struct {
	Id        int `json:"id"`
	Title     string `json:"title"`
	Years     int `json:"years"`
}

type Author struct 	{
	Id     int `json:"id"`
	Name   string `json:"name"`
}

type AuthorInput struct {
	Name 	   string `json:"name" binding:"required"`
}

type AuthorResponse struct 	{
	Id     int `json:"id"`
	Name   string `json:"name"`
	Book   []Book
}