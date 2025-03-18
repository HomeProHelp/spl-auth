package auth

import (
	"github/LissaiDev/spl-auth/pkg/token"
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

func (s *UserService) GetUser(id uuid.UUID) (*User, string) {
	return s.r.GetUserByID(id)
}

func (s *UserService) CreateUser(user *User) (*User, string) {
	hashedPwd, err := HashPassword(user.Password)
	if err != nil {
		return nil, utils.AuthenticationCodes["internal_server_error"]
	}
	user.Password = hashedPwd
	return s.r.CreateUser(user)
}

func (s *UserService) AuthenticateUser(email string, password string) (*string, string) {
	user, code := s.r.GetUserByEmail(email)
	if code != utils.AuthenticationCodes["success"] {
		return nil, code
	}

	if err := VerifyPassword(user.Password, password); err != nil {
		return nil, utils.AuthenticationCodes["invalid_credentials"]
	}

	token, code := token.GenerateToken(user.ID.String())
	if code != utils.AuthenticationCodes["success"] {
		return nil, code
	}

	return token, utils.AuthenticationCodes["success"]
}
