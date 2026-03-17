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
	// Загружаем .env: сначала рядом с бинарём, потом из корня проекта
	if err := godotenv.Load(".env"); err != nil {
		godotenv.Load("../.env")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer database.Close()

	h := &handlers.Handler{DB: database}
	r := gin.Default()

	// CORS для фронтенда
	allowedOrigins := []string{
		"http://localhost:5173",
		os.Getenv("FRONTEND_URL"),
	}
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, allowed := range allowedOrigins {
			if allowed != "" && origin == allowed {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
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
