package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ContactForm struct {
	Name    string `form:"name" binding:"required,min=2,max=50"`
	Email   string `form:"email" binding:"required,email"`
	Message string `form:"message" binding:"required,min=10,max=550"`
}

var (
	telegramBotToken string
	telegramChatID   string
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  .env файл не найден, используем переменные окружения системы")
	}

	telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatID = os.Getenv("TELEGRAM_CHAT_ID")

	if telegramBotToken == "" || telegramChatID == "" {
		fmt.Println("❌ Ошибка: TELEGRAM_BOT_TOKEN или TELEGRAM_CHAT_ID не установлены")
		return
	}

	fmt.Println("Настройки Telegram загружены успешно!")

	r := gin.Default()

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/api/skills", func(c *gin.Context) {
		skills := []string{"Go", "Gin", "Backend Development", "REST API", "Docker", "Git", "Linux"}
		c.JSON(200, skills)
	})

	r.POST("/contact", func(c *gin.Context) {
		var form ContactForm

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "Ошибка валидации: " + err.Error(),
			})
			return
		}

		if err := validateContactForm(form); err != nil {
			c.JSON(400, gin.H{
				"status":  "error",
				"message": err.Error,
			})
			return
		}

		saveContactForm(form.Name, form.Email, form.Message)

		if err := sendTelegramNotification(form.Name, form.Email, form.Message); err != nil {
			fmt.Println("Ошибка отправки в Telegram:", err)
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Сообщение отправлено!",
		})
	})

	r.Run(":8080")
}

func sendTelegramNotification(name, email, message string) error {
	text := fmt.Sprintf(`*Новое сообщение с сайта*
	
	*Имя:* %s
	*Email:* %s
	*Message:* %s
	
	*Time:* %s`,
		name, email, message, time.Now().Format("2006-01-02 15:04:05"))

	msg := map[string]string{
		"chat_id":    telegramChatID,
		"text":       text,
		"parse_mode": "Markdown",
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("ошибка создания: %v", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса:%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка от Telegram API: %s", resp.Status)
	}

	fmt.Println("✅ Уведомление отправлено в Telegram!")
	return nil

}

func validateContactForm(form ContactForm) error {
	if !isValidName(form.Name) {
		return fmt.Errorf("Имя может содержать только буквы и пробелы")
	}

	if _, err := mail.ParseAddress(form.Email); err != nil {
		return fmt.Errorf("Некорректный email адрес")
	}

	if containsSpam(form.Message) {
		return fmt.Errorf("Сообщение содержит запрещённый контент")
	}

	return nil
}

func isValidName(name string) bool {
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= 'а' && char <= 'я') ||
			(char >= 'А' && char <= 'Я') ||
			char == ' ' || char == '-') {
			return false
		}
	}
	return true
}

func containsSpam(message string) bool {
	spamWords := []string{"http://", "https://", "www.", ".com", "куплю", "продам", "заработок", "sell", "buy"}
	lowerMessage := strings.ToLower(message)

	for _, word := range spamWords {
		if strings.Contains(lowerMessage, word) {
			return true
		}
	}
	return false
}
func saveContactForm(name, email, message string) {
	file, err := os.OpenFile("contact_message.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("[%s] Name: %s | Email: %s | Message: %s\n", timeStamp, name, email, message)

	_, err = file.WriteString(entry)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}

	fmt.Println("Новое сообщение сохранено:", name, email)
}
