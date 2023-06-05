package rest

import (
	"context"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Books interface {
	Create(book domain.Book) error
	GetByID(id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error
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
	r := gin.New()

	api := r.Group("api")
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
	return r
	//r := mux.NewRouter()
	//r.Use(loggingMiddleware)
	//
	//books := r.PathPrefix("/books").Subrouter()
	//{
	//	books.HandleFunc("", h.createBook).Methods(http.MethodPost)
	//	books.HandleFunc("", h.getAllBooks).Methods(http.MethodGet)
	//	books.HandleFunc("/{id:[0-9]+}", h.getBookByID).Methods(http.MethodGet)
	//	books.HandleFunc("/{id:[0-9]+}", h.deleteBook).Methods(http.MethodDelete)
	//	books.HandleFunc("/{id:[0-9]+}", h.updateBook).Methods(http.MethodPut)
	//}
}
