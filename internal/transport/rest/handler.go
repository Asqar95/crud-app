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

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/")
			books.GET(":/id")
			books.DELETE("/:id")
			books.PUT("/:id")
		}
	}
	return router
}
