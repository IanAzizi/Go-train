package router

import (
	"swag/api/handler"

	"github.com/gin-gonic/gin"
)

func GetRouter(rg *gin.RouterGroup) {
	userHandler := handler.NewUserHandler()

	rg.GET("/users", userHandler.GetUser)
}
