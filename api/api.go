package api

import (
	"log"
	"swag/api/router"

	"github.com/gin-gonic/gin"

	_ "swag/api/docs" // حتما باید import بشه

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyApp API
// @version 1.0
// @description A professional API with Swagger documentation.
// @termsOfService http://yourdomain.com/terms
// @contact.name Support
// @contact.email support@yourdomain.com
// @license.name MIT
// @host localhost:5002
// @BasePath /api/v1
func InitService() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// ثبت routeهای اصلی
	v1 := r.Group("/api/v1")
	router.GetRouter(v1)

	// ثبت Swagger route روی Engine
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running on 5002")
	r.Run(":5002")
}
