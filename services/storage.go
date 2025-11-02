package services

import (
	"fmt"
	"os"
	"time"
)

func SaveContactForm(name, email, message string) error {
	file, err := os.OpenFile("contact_message.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("[%s] Name: %s | Email: %s | Message: %s\n", timeStamp, name, email, message)

	_, err = file.WriteString(entry)
	if err != nil {
		return fmt.Errorf("ошибка записи: %v", err)
	}

	fmt.Println("Новое сообщение сохранено:", name, email)
	return nil
}
