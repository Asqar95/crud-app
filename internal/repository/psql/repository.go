package repository

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Books
}

type Books interface {
	Create(book domain.Book) (int, error)
	GetByID(id int) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int) error
	Update(id int, inp domain.UpdateBookInput) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBookPostgres(db),
	}
}
