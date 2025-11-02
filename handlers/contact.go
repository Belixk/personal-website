package handlers

import (
	"fmt"
	"net/http"

	"github.com/Belixk/personal-website/models"
	"github.com/Belixk/personal-website/services"
	"github.com/gin-gonic/gin"
)

func ContactHandler(c *gin.Context) {
	var form models.ContactForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Ошибка валидации: " + err.Error(),
		})
		return
	}

	if err := services.ValidateContactForm(form.Name, form.Email, form.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := services.SaveContactForm(form.Name, form.Email, form.Message); err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}

	if err := services.SendTelegramNotification(form.Name, form.Email, form.Message); err != nil {
		fmt.Println("Ошибка отправки в Telegram:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Сообщение отправлено",
	})

}
