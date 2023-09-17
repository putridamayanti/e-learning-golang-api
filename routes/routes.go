package routes

import (
	"elearning/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine)  {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}

func UserRoutes(router *gin.Engine)  {

}
