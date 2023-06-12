package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_webapp"
)

func (h *Handler) signIn(c *gin.Context) {

}

func (h *Handler) signUp(c *gin.Context) {
	var input todo_webapp.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
