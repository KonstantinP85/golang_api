package api_go

type Book struct {
	Id        int `json:"id"`
	AuthorId  int `json:"author_id"`
	Title     string `json:"title"`
	Years     int `json:"years"`
}

type Author struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
}
