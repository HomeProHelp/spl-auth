package auth

import (
	"fmt"
	"github/LissaiDev/spl-auth/db"
	"github/LissaiDev/spl-auth/pkg/hash"
	"github/LissaiDev/spl-auth/pkg/hermes"
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
	var user User = User{Name: name, Email: email, Password: hashedPwd}
	result := db.Database.Create(&user)

	if result.Error != nil {
		hermes.Log(3, fmt.Sprintf("User creation failed: {%s, %s, %s}\nError: %s", name, email, password, result.Error), false)
		return User{}, result.Error
	}

	hermes.Log(1, fmt.Sprintf("User created successfully: %+v", user), false)
	return user, nil
}

func init() {
	db.Connect()
	db.Database.AutoMigrate(&User{})
}
