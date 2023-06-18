package rest

import (
	"errors"
	"github.com/Asqar95/crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	var inp domain.SignUpInput

	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.userService.SignUp(c, inp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) singIn(c *gin.Context) {
	var inp domain.SignInInput

	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.userService.SignIn(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			handleNotFoundError(w, err)
			return
		}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
