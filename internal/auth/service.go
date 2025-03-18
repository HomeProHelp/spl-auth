package auth

import (
	"github/LissaiDev/spl-auth/utils"

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

func (s *UserService) GetUser(id uuid.UUID) (User, string) {
	return s.r.GetUserByID(id)
}

func (s *UserService) CreateUser(user *User) (User, string) {
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		return User{}, utils.AuthenticationCodes["internal_server_error"]
	}
	*user = User{Name: user.Name, Email: user.Email, Password: hashedPwd}
	return s.r.CreateUser(user)
}
