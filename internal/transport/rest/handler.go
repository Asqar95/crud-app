package handler

import (
	"github.com/Asqar95/crud-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()
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
