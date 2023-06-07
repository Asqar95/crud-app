package service

import (
	"github.com/Asqar95/crud-app/internal/domain"
	repository "github.com/Asqar95/crud-app/internal/repository/psql"
)

type Service struct {
	Books
}

type Books interface {
	Create(book domain.Book) error
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int64) error
	Update(id int64, inp domain.UpdateBookInput) error
}

func NewService(repos repository.Repository) *Service {
	return &Service{repo: repo}
}
