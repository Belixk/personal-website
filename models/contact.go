package models

type ContactForm struct {
	Name    string `json:"name" binding:"required, min=2,max=50"`
	Email   string `json:"email" binding:"required, email"`
	Message string `json:"message" binding:"required, min=10,max=550"`
}
