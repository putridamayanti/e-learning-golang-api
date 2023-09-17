package main

import (
	"elearning/config"
	"elearning/controllers"
	"elearning/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env file not found")
		return
	}

	if !config.Init() {
		log.Printf("Connected to MongoDB URI: Failure")
		return
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.CorsMiddleware())

	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protectedApi := api.Group("", middlewares.AuthMiddleware())
		{
			protectedApi.GET("/user", controllers.GetUserByQuery)
			protectedApi.POST("/user", controllers.CreateUser)
			protectedApi.GET("/user/:id", controllers.GetUserById)
			protectedApi.PATCH("/user/:id", controllers.UpdateUser)
			protectedApi.DELETE("/user/:id", controllers.DeleteUser)

			protectedApi.GET("/course", controllers.GetCourseByQuery)
			protectedApi.POST("/course", controllers.CreateCourse)
			protectedApi.GET("/course/:id", controllers.GetCourseById)
			protectedApi.PATCH("/course/:id", controllers.UpdateCourse)
			protectedApi.DELETE("/course/:id", controllers.DeleteCourse)

			protectedApi.POST("/course-content/:id", controllers.AddCourseContent)

			protectedApi.POST("/enroll", controllers.EnrollCourse)
		}
	}

	port := "8000"
	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
