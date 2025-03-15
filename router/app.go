package router

import (
	"bg/middleware"
	"bg/service"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("/api/login", service.Login)
	r.POST("/api/getMessage", service.GetMessages)
	r.POST("/api/getPhones", service.GetPhones)
	return r
}
