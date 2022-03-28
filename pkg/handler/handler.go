package handler

import (
	"HTTP31/pkg/service"
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

	AllGroup := router.Group("/")
	{
		AllGroup.POST("/create", h.createNewUser)
		AllGroup.POST("/make_friends", h.makeFriends)
		AllGroup.GET("/", h.getAllUser)
		AllGroup.GET("/friends/:id", h.getUserByID)
		AllGroup.PUT("/age_updated/:id", h.updateUser)
		AllGroup.DELETE("/delete/:id", h.deleteUser)
	}
	return router
}
