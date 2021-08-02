package repository

import (
	"fmt"
	"github.com/KonstantinP85/api-go"
	"github.com/jmoiron/sqlx"
)

type BookDB struct {
	db *sqlx.DB
}

func NewBookDB(db *sqlx.DB) *BookDB {
	return &BookDB{db: db}
}

func (r *BookDB) GetBook(id int) (api_go.Book, error) {
	var book api_go.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", bookTable)
	err := r.db.Get(&book, query, id)

	return book, err
}

func (r *BookDB) GetBooks() ([]api_go.Book, error) {
	var books []api_go.Book

	query := fmt.Sprintf("SELECT * FROM %s", bookTable)
	err := r.db.Select(&books, query)

	return books, err
}
