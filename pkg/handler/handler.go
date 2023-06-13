package handler

import (
	"github.com/gin-gonic/gin"
	"todo_webapp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	tasks := router.Group("/task", h.userIdentity)
	{
		tasks.GET("/", h.getAllTasks)
		tasks.POST("/", h.createTask)
		tasks.DELETE("/", h.deleteTask)
		tasks.GET("/:task_id", h.getTaskById)

		tasks.GET("/due/:yy/:mm/:dd", h.getTasksByDate)
		tasks.GET("/tag/:tag_name", h.getTasksByTag)
	}

	return router
}
