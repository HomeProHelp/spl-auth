package auth

import (
	"fmt"
	"github/LissaiDev/spl-auth/pkg/hermes"

	"github.com/google/uuid"
)

type UserService struct {
	r *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) GetUser(id uuid.UUID) (User, error) {
	user, err := s.r.GetUserByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user *User) (User, error) {
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		hermes.Log(3, fmt.Sprintf("User password hashing failed: %s", user.Password), false)
		return User{}, err
	}
	*user = User{Name: user.Name, Email: user.Email, Password: hashedPwd}
	u, error := s.r.CreateUser(user)
	if error != nil {
		hermes.Log(1, fmt.Sprintf("[SERVICE]: Error occured %s", error), false)
		return User{}, error
	}
	return u, nil
}
