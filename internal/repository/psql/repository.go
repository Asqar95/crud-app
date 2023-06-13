package repository

import (
	"context"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Books
}

type Books interface {
	Create(ctx context.Context, book domain.Book) (int, error)
	GetByID(ctx context.Context, id int) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, inp domain.UpdateBookInput) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBookPostgres(db),
	}
}
