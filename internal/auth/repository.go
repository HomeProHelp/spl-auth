package auth

import (
	"errors"
	"fmt"
	"github/LissaiDev/spl-auth/db"
	"github/LissaiDev/spl-auth/pkg/hash"
	"github/LissaiDev/spl-auth/pkg/hermes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func GetUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(name, email, password string) (User, error) {
	hashedPwd, err := hash.HashPassword(password)
	if err != nil {
		hermes.Log(3, fmt.Sprintf("User password hashing failed: %s", password), false)
		return User{}, err
	}
	var user User = User{Identifier: uuid.New(), Name: name, Email: email, Password: hashedPwd}
	result := db.Database.Create(&user)

	if result.Error != nil {
		hermes.Log(3, fmt.Sprintf("User creation failed: {%s, %s, %s}\nError: %s", name, email, password, result.Error), false)
		return User{}, result.Error
	}

	hermes.Log(1, fmt.Sprintf("User created successfully: %+v", user), false)
	return user, nil
}

func (r *UserRepository) GetUserByIdentifier(id uuid.UUID) (User, error) {
	var user User
	result := db.Database.First(&user, "identifier = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			hermes.Log(1, fmt.Sprintf("User with ID: %s Not Found", id), false)
			return User{}, result.Error
		}
		hermes.Log(3, "Error retrieving user", true)
		return User{}, result.Error
	}
	hermes.Log(1, fmt.Sprintf("User with ID: %s Sucessfully retrieved", id), false)
	return user, nil
}

func init() {
	db.Connect()
	db.Database.AutoMigrate(&User{})
}
