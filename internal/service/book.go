package service

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"time"
)

type BooksRepository interface {
	Create(book domain.Book) error
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int64) error
	Update(id int64, inp domain.UpdateBookInput) error
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return b.repo.Create(book)
}

func (b *Books) GetByID(id int64) (domain.Book, error) {
	return b.repo.GetByID(id)
}

func (b *Books) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}

func (b *Books) Delete(id int64) error {
	return b.repo.Delete(id)
}

func (b *Books) Update(id int64, inp domain.UpdateBookInput) error {
	return b.repo.Update(id, inp)
}
