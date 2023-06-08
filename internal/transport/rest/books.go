package handler

import (
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createBook(c *gin.Context) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book domain.Book
	if err = json.Unmarshal(reqBytes, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Create(context.TODO(), book)
	if err != nil {
		log.Println("createBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Books.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
	return
}

func (h *Handler) deleteBook(c *gin.Context) {
	id,
	if err != nil {
		log.Println("deleteBook() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("deleteBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.booksService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) updateBook(c *gin.Context) {
	id, err := getIdFomRequest(r)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqByts, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.UpdateBookInput
	if err = json.Unmarshal(reqByts, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Update(context.TODO(), id, inp)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
