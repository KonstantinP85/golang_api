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

func (r *BookDB) GetBook(id int) (api_go.BookResponse, error) {
	var book api_go.BookResponse

	queryBook := fmt.Sprintf(`SELECT bt.id, bt.title, bt.years, at.id "author.id", at.name "author.name" FROM %s bt INNER JOIN %s at on bt.author_id=at.id WHERE bt.id=$1`, bookTable, authorTable)
	err := r.db.Get(&book, queryBook, id)

	return book, err
}

func (r *BookDB) GetBooks(authorId int) ([]api_go.BookResponse, error) {
	var books []api_go.BookResponse

	if authorId != 0 {
		query := fmt.Sprintf(`SELECT bt.id, bt.title, bt.years, at.id "author.id", at.name "author.name" FROM %s bt INNER JOIN %s at on bt.author_id=at.id WHERE at.id=$1`, bookTable, authorTable)
		err := r.db.Select(&books, query, authorId)
		return books, err
	} else {
		query := fmt.Sprintf(`SELECT bt.id, bt.title, bt.years, at.id "author.id", at.name "author.name" FROM %s bt INNER JOIN %s at on bt.author_id=at.id`, bookTable, authorTable)
		err := r.db.Select(&books, query)
		return books, err
	}
}

func (r *BookDB) CreateBook(book api_go.BookInput) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (author_id, title, years) values ($1, $2, $3) RETURNING id", bookTable)
	row := r.db.QueryRow(query, book.AuthorId, book.Title, book.Years)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
