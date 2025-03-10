package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;not null;unique;primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	createdAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
