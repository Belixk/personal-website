package services

import (
	"fmt"
	"net/mail"
	"strings"
)

func ValidateContactForm(name, email, message string) error {
	if !isValidName(name) {
		return fmt.Errorf("Имя может содержать только буквы и пробел")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("Неккоректный email адрес")
	}

	if containsSpam(message) {
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
