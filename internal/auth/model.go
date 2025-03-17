package auth

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;not null;unique;primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;unique"`
	Password  string         `gorm:"type:varchar(255);not null"`
	createdAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type userValidation struct {
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,strong_password"`
}

func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	strongPasswordRegex := `^(?=.*[A-Z])(?=.*\d).{8,}$`
	match, _ := regexp.MatchString(strongPasswordRegex, password)

	fmt.Printf("Match: %t\nPassword: %s", match, password)

	return match
}

var errorMessages map[string]string = map[string]string{
	"Name":     "Name is required and must be at least 3 characters long",
	"Email":    "Email is required and must be a valid email address",
	"Password": "Password is required and must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character",
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	validate := validator.New()
	validate.RegisterValidation("strong_password", passwordValidator)

	err = validate.Struct(userValidation{Name: u.Name, Email: u.Email, Password: u.Password})
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(errorMessages[err.Field()])
		}
	}

	return err
}
