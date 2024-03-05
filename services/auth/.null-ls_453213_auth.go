package auth

import (
	"GemiApp/domain/account"
	"GemiApp/types"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo account.Repository
}

var (
	ErrInvalidUsername = errors.New("Username invalid")
	ErrInvalidPassword = errors.New("Password invalid")
	ErrInvalidAccount  = errors.New("User ID invalid")
)

func NewAuthService(r account.Repository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) UserLogin(username, password string) (string, string, error) {
	acc, err := s.repo.GetAccountByUsername(username)
	if err != nil {
		return "", "", ErrInvalidUsername
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return "", "", ErrInvalidPassword
	}
	return acc.ID, acc.Role, nil
}

func (s *AuthService) UserStatus(user_id string) (*types.User, error) {
	acc, err := s.repo.GetAccountByID(user_id)
	if err != nil {
		return map[string]any{}, ErrInvalidAccount
	}

	res := map[string]any{
		"Username": acc.Username,
		"Balance":  acc.Balance,
	}
	return res, nil
}
