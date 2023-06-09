package repository

import (
	"fmt"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
	"strings"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) Create(book domain.Book) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var bookId int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title, author, publish_date, rating) values ($1,$2,$3,$4) RETURNING id", books)
	row := tx.QueryRow(createBookQuery, book.Title, book.Author, book.PublishDate, book.Rating)
	err = row.Scan(&bookId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return bookId, tx.Commit()
}

func (r *BooksPostgres) GetByID(id int) (domain.Book, error) {
	var book domain.Book
	query := fmt.Sprintf("SELECT id, title, publish_date, rating FROM %s WHERE id=$1", books)
	if err := r.db.Get(&book, query, id); err != nil {
		return book, err
	}
	return book, nil
}

func (r *BooksPostgres) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	query := fmt.Sprintf("SELECT id, title, author, publish_date, rating FROM %s", books)
	if err := r.db.Select(&books, query); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BooksPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", books)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *BooksPostgres) Update(id int, inp domain.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
	}

	if inp.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
	}
	if inp.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *inp.PublishDate)
		argId++
	}

	if inp.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE books SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
