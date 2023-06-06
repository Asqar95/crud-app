package rest

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/gin-gonic/gin"
)

type BooksService interface {
	Create(book domain.Book) error
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
	Delete(id int64) error
	Update(id int64, inp domain.UpdateBookInput) error
}

type Handler struct {
	services *service.Books
}

func NewHandler(services *service.Books) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	api := router.Group("api")
	{
		books := api.Group("/books")
		{
			books.POST("/")
			books.GET("/")
			books.GET(":/id")
			books.DELETE("/:id")
			books.PUT("/:id")
		}
	}
	return router
}
