package main

import (
	"fmt"
	"os"

	"github.com/Belixk/personal-website/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env файл не найден, используем переменные окружения системы")
	}

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" || os.Getenv("TELEGRAM_CHAT_ID") == "" {
		fmt.Println("ошибка: TELEGRAM_BOT_TOKEN или TELEGRAM_CHAT_ID не установлены")
		return
	}

	fmt.Println("Настройки успешно загружены!")

	r := gin.Default()

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.HomeHandler)
	r.GET("/api/skills", handlers.SkillsHandler)
	r.POST("/contact", handlers.ContactHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
