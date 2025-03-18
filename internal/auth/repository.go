package auth

import (
	"errors"
	"github/LissaiDev/spl-auth/db"
	"github/LissaiDev/spl-auth/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *User) (User, string) {
	result := db.Database.Create(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return User{}, utils.AuthenticationCodes["email_already_exists"]
		}
		if errors.Is(result.Error, gorm.ErrInvalidData) {
			return User{}, utils.AuthenticationCodes["invalid_data"]
		}
		return User{}, utils.AuthenticationCodes["internal_server_error"]
	}

	return *user, utils.AuthenticationCodes["success"]
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (User, string) {
	var user User
	result := db.Database.First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, utils.AuthenticationCodes["user_not_found"]
		}
		return User{}, utils.AuthenticationCodes["internal_server_error"]
	}
	return user, utils.AuthenticationCodes["success"]
}

func init() {
	db.Connect()
	db.Database.AutoMigrate(&User{})
}
