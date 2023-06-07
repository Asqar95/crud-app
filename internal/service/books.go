package service

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"time"
)

type BooksService struct {
	repo repository.Repository
}

func NewBooksService(repo repository.Repository) *BooksService {
	return &BooksService{repo: repo}
}

func (b *BooksService) Create(book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return nil
}

func (b *BooksService) GetByID(id int64) (domain.Book, error) {
	return b.repo.GetByID(id)
}

func (b *BooksService) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}

func (b *BooksService) Delete(id int64) error {
	return b.repo.Delete(id)
}

func (b *BooksService) Update(id int64, inp domain.UpdateBookInput) error {
	return b.repo.Update(id, inp)
}
