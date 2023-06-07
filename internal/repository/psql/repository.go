package repository

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Books
}

type Books interface {
	Create(book domain.Book) error
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int64) error
	Update(id int64, inp domain.UpdateBookInput) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBookPostgres(db),
	}
}
