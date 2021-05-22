package handler

import (
	"apiserver/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		users.POST("/", h.createUser)
		users.PUT("/:id", h.updateUser)
	}

	return router
}
