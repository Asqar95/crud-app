package handler

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary Create book
// @Tags Books
// @Description create book
// @ID createBook
// @Accept  json
// @Produce  json
// @Param input body domain.Book true "Book info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /book [post]
func (h *Handler) createBook(c *gin.Context) {
	var input domain.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Create(c, input)
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
// @Tags Books
// @Description Get book by ID
// @Accept  json
// @Produce  json
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /book [get]
func (h *Handler) getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getbookbyid",
			"problem": "unmarshalling request",
		}).Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Books.GetByID(c, id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getbookbyid",
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
// @Tags Books
// @Description create book
// @Accept  json
// @Produce  json
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /books [get]
func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.GetAll(c)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "getallbooks",
			"problem": "unmarshalling request",
		}).Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})

}

// @Summary Delete book
// @Tags Books
// @Description create book
// @Accept  json
// @Produce  json
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /book [delete]
func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "deletebook",
			"problem": "unmarshalling request",
		}).Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Delete(c, id)
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
// @Accept  json
// @Produce  json
// @Param input body domain.UpdateBookInput true "Book info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /book [put]
func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "updatebook",
			"problem": "unmarshalling request",
		}).Error(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		log.WithFields(log.Fields{
			"handler": "deletebook",
			"problem": "service error",
		}).Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(c, id, input); err != nil {
		log.WithFields(log.Fields{
			"handler": "deletebook",
			"problem": "service error",
		}).Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
