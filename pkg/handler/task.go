package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllTasks(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTaskById(c *gin.Context) {

}

func (h *Handler) getTasksByTag(c *gin.Context) {

}

func (h *Handler) getTasksByDate(c *gin.Context) {

}

func (h *Handler) createTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}
