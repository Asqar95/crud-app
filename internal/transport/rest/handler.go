package rest

import (
	"context"
	_ "github.com/Asqar95/crud-app/docs"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Books interface {
	Create(ctx context.Context, book domain.Book) (int, error)
	GetByID(ctx context.Context, id int) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, inp domain.UpdateBookInput) error
}

type User interface {
	SignUp(ctx context.Context, inp domain.SignUpInput) error
	SignIn(ctx context.Context, inp domain.SignInInput) (string, error)
	ParseToken(ctx context.Context, token string) (int64, error)
}

type Handler struct {
	services    Books
	userService User
}

func NewHandler(services Books, users User) *Handler {
	return &Handler{
		services:    services,
		userService: users,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		//auth.POST("/sing-in", h.singIn)
	}

	book := router.Group("/book")
	{
		book.POST("/", h.createBook)
		book.GET("/", h.getAllBooks)
		book.GET("/:id", h.getBookByID)
		book.DELETE("/:id", h.deleteBook)
		book.PUT("/:id", h.updateBook)
	}
	return router
}
