package repository

import (
	"fmt"
	api_go "github.com/KonstantinP85/api-go"
	"github.com/jmoiron/sqlx"
)

type AuthorDB struct {
	db *sqlx.DB
}

func NewAuthorDB(db *sqlx.DB) *AuthorDB {
	return &AuthorDB{db: db}
}

func (r *AuthorDB) CreateAuthor(author api_go.AuthorInput) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", authorTable)
	row := r.db.QueryRow(query, author.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorDB) GetAuthors() ([]api_go.Author, error) {
	var authors []api_go.Author

	query := fmt.Sprintf(`SELECT * FROM %s`, authorTable)
	err := r.db.Select(&authors, query)

	return authors, err
}

func (r *AuthorDB) GetAuthor(id int) (api_go.AuthorResponse, error) {
	var author api_go.AuthorResponse

	queryBook := fmt.Sprintf(`SELECT at.id, at.name, bt.id, bt.author_id "book.author_id", bt.title "book.title", bt.years "book.years" FROM %s at INNER JOIN %s bt on at.id=bt.author_id WHERE at.id=$1`, authorTable, bookTable)
	err := r.db.Get(&author, queryBook, id)

	return author, err
}