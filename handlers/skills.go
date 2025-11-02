package handlers

import (
	"github.com/gin-gonic/gin"
)

func SkillsHandler(c *gin.Context) {
	skills := []string{"Go", "Gin", "REST API", "Docker", "Git", "Linux"}
	c.JSON(200, skills)
}
