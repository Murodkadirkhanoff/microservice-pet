package entity

import (
	"errors"
	"time"

	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/dto"
	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`          // UUID для ID
	Email      string    `gorm:"size:255;uniqueIndex;not null"` // Уникальный email, не null
	Password   string    `gorm:"not null"`                      // Пароль, не null
	FirstName  string    `gorm:"size:255"`                      // Имя, максимальный размер 255
	LastName   string    `gorm:"size:255"`                      // Фамилия, максимальный размер 255
	IsActive   bool      `gorm:"default:true"`                  // Активность, по умолчанию true
	IsVerified bool      `gorm:"default:false"`                 // Подтвержден ли пользователь, по умолчанию false
	Avatar     string    `gorm:"size:255"`                      // Путь к аватару пользователя
	CreatedAt  time.Time `gorm:"autoCreateTime"`                // Автоматически заполняется текущим временем при создании записи
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`                // Автоматически обновляется текущим временем при обновлении записи
	DeletedAt  time.Time `gorm:"index"`                         // Мягкое удаление, используется индекс для быстрого поиска
}

func ValidateCredentials(db *gorm.DB, dto dto.LoginDTO) (uuid.UUID, error) {
	var user User
	db.First(&user, "email = ?", dto.Email)
	retrievedPassword := user.Password
	passwordIsValid := utils.CheckPasswordHash(dto.Password, retrievedPassword)
	if !passwordIsValid {
		return uuid.Nil, errors.New("Credentials invalid")
	}

	return user.ID, nil
}
