package main

import (
	"miniProjectAPI/controllers"
	"miniProjectAPI/initializers"
	"miniProjectAPI/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Menggunakan middleware untuk menangani CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// User Model Route (Login & Register)
	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/validate", middleware.RequireAuth, controllers.Validate)

	// Get User Data
	r.GET("/api/users", controllers.UserIndex)

	// Quiz Model Route
	QuizAPI := r.Group("/api/quiz")
	{
		QuizAPI.GET("/", controllers.QuizIndex)
		QuizAPI.GET("/:id", controllers.QuizShow)
		QuizAPI.POST("/", controllers.QuizCreate)
		QuizAPI.PUT("/:id", controllers.QuizUpdate)
		QuizAPI.DELETE("/:id", controllers.QuizDelete)
	}

	// Pertanyaan Model Route
	pertanyaanAPI := r.Group("/api/pertanyaan")
	{
		pertanyaanAPI.GET("/", controllers.PertanyaanIndex)
		pertanyaanAPI.GET("/:id", controllers.PertanyaanShow)
		pertanyaanAPI.POST("/", controllers.PertanyaanCreate)
		pertanyaanAPI.PUT("/:id", controllers.PertanyaanUpdate)
		pertanyaanAPI.DELETE("/:id", controllers.PertanyaanDelete)
	}

	// Jawaban Model Route
	jawabanAPI := r.Group("/api/jawaban")
	{
		jawabanAPI.GET("/", controllers.JawabanIndex)
		jawabanAPI.GET("/:id", controllers.JawabanShow)
		jawabanAPI.POST("/", controllers.JawabanCreate)
		jawabanAPI.PUT("/:id", controllers.JawabanUpdate)
		jawabanAPI.DELETE("/:id", controllers.JawabanDelete)
	}

	r.Run()
}
