package main

import (
	"log"
	"os"

	"mood-generator/internal/db"
	"mood-generator/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env по абсолютному пути
	envPath := `C:\Users\Пользователь\Desktop\ralez\mood-generator\backend\.env`
	if err := godotenv.Load(envPath); err != nil {
		godotenv.Load() // fallback
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer database.Close()

	h := &handlers.Handler{DB: database}
	r := gin.Default()

	// CORS для фронтенда
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/mood", h.PostMood)
		api.GET("/history", h.GetHistory)
		api.GET("/mood/:id", h.GetMoodByID)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
