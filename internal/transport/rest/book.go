package rest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Asqar95/crud-app/internal/domain"
	"io/ioutil"
	"net/http"
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
func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("createBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book domain.Book
	if err = json.Unmarshal(reqBytes, &book); err != nil {
		logError("createBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Create(r.Context(), book)
	if err != nil {
		logError("createBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("getBookByID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logError("getBookByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		logError("getBookByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
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
func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.booksService.GetAll(r.Context())
	if err != nil {
		logError("getAllBooks", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		logError("getAllBooks", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
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
func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("deleteBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Delete(r.Context(), id)
	if err != nil {
		logError("deleteBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.UpdateBookInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Update(r.Context(), id, inp)
	if err != nil {
		logError("updateBook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
