package service

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/Asqar95/crud-app/internal/repository/psql"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) Create(id int, book domain.Book) (int, error) {
	return s.repo.Create(id, book)
}

func (s *BooksService) GetByID(id int) (domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BooksService) GetAll() ([]domain.Book, error) {
	return s.repo.GetAll()
}

func (s *BooksService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BooksService) Update(id int, inp domain.UpdateBookInput) error {
	return s.repo.Update(id, inp)
}
