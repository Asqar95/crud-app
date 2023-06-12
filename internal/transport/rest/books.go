package handler

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary Create book
// @Books lists
// @Description create book
// @ID create-list
// @Accept  json
// @Produce  json
// @Router /books [post]
func (h *Handler) createBook(c *gin.Context) {
	var input domain.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Create(input)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "reading request body",
		}).Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get book
// @Books lists
// @Description create book
// @ID create-list
// @Accept  json
// @Produce  json
// @Router /books [get]
func (h *Handler) getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "unmarshaling request",
		}).Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Books.GetByID(id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "service error",
		}).Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
	return
}

type getAllBooksResponse struct {
	Data []domain.Book `json:"data"`
}

// @Summary GetAll
// @Books lists
// @Description create book
// @ID create-list
// @Accept  json
// @Produce  json
// @Router /books [get]
func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})

}

// @Summary Delete book
// @Books lists
// @Description create book
// @ID create-list
// @Accept  json
// @Produce  json
// @Router /books [delete]
func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Update book
// @Tags Books
// @Description create book
// @ID create-list
// @Accept  json
// @Produce  json
// @Router /books [put]
func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
