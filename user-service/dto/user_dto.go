package dto

import "github.com/google/uuid"

// Структура для входных данных пользователя из POST-запроса
type RegisterDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"created_at"`
}
