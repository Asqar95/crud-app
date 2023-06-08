package service

import (
	"github.com/Asqar95/crud-app/internal/domain"
	repository "github.com/Asqar95/crud-app/internal/repository/psql"
)

type Service struct {
	Books
}

type Books interface {
	Create(book domain.Book)
	GetByID(id int) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int) error
	Update(id int, inp domain.UpdateBookInput) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books: NewBooksService(repos.Books),
	}
}
