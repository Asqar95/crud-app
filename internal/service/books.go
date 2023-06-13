package service

import (
	"context"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/Asqar95/crud-app/internal/repository/psql"
	"time"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) Create(ctx context.Context, book domain.Book) (int, error) {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}
	return s.repo.Create(ctx, book)
}

func (s *BooksService) GetByID(ctx context.Context, id int) (domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BooksService) GetAll(ctx context.Context) ([]domain.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *BooksService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *BooksService) Update(ctx context.Context, id int, inp domain.UpdateBookInput) error {
	return s.repo.Update(ctx, id, inp)
}
