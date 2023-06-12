package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
