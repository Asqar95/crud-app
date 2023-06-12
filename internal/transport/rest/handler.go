package handler

import (
	_ "github.com/Asqar95/crud-app/docs"
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
