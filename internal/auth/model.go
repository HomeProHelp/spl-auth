package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identifier uuid.UUID `gorm:"type:uuid;not null;unique"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
}
