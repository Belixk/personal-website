package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func SendTelegramNotification(name, email, message string) error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if token == "" || chatID == "" {
		return fmt.Errorf("Telegram конфигурация не найдена")
	}

	text := fmt.Sprintf(`Новое сообщение с сайта*
	*Name:* %s
	*Email:* %s
	*Message:* %s
	
	*Time:* %s`,
		name, email, message, time.Now().Format("2006-01-02 15:04:05"))

	msg := map[string]string{
		"chat_ID":    chatID,
		"text":       text,
		"parse_mode": "Markdown",
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Ошибка сереализации данных: %v", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("Ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Ошибка от Telegram API: %s", resp.Status)
	}

	fmt.Println("Уведомление успешно отправлено в Telegram!")
	return nil
}
