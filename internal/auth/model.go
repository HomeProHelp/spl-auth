package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;not null;unique;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name" binding:"required,min=3"`
	Email     string         `gorm:"type:varchar(255);not null;unique" json:"email" binding:"required,email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password" binding:"required,min=8"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return err
}
