package handler

import (
	"github.com/gin-gonic/gin"
	"testTask/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/person", h.GetPerson)
	router.DELETE("/person", h.DeletePerson)
	router.PUT("/person", h.UpdatePerson)
	router.POST("/person", h.AddPerson)

	return router
}
