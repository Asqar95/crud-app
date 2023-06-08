package repository

import (
	"database/sql"
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

func (b *BooksPostgres) Create(book domain.Book) {
	_, _ = b.db.Exec("INSERT INTO books (title, author, publish_date, rating) values ($1, $2,$3, $4)",
		book.Title, book.Author, book.PublishDate, book.Rating)
	return
}

func (b *BooksPostgres) GetByID(id int) (domain.Book, error) {
	var book domain.Book
	err := b.db.QueryRow("SELECT id, title, publish_date, rating FROM books WHERE id=$1", id).
		Scan(&book.ID, book.Title, &book.Author, &book.PublishDate, &book.Rating)
	if err == sql.ErrNoRows {
		return book, domain.ErrBookNotFound
	}
	return book, err
}

func (b *BooksPostgres) GetAll() ([]domain.Book, error) {
	rows, err := b.db.Query("SELECT id, title, author, publish_date, rating FROM books")
	if err != nil {
		return nil, err
	}

	books := make([]domain.Book, 0)
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating); err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, rows.Err()
}

func (b *BooksPostgres) Delete(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}

func (b *BooksPostgres) Update(id int, inp domain.UpdateBookInput) error {
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

	_, err := b.db.Exec(query, args...)
	return err
}
