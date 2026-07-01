package services

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/roqiaahmed/knowledgehub/internal/dto"
	"github.com/roqiaahmed/knowledgehub/internal/models"
	"github.com/roqiaahmed/knowledgehub/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(authDTO *dto.RegisterAuthRequest) (dto.RegisterAuthResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func NormalizeAndValidateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid email format")
	}
	email = strings.ToLower(email)
	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]
	domain = strings.ReplaceAll(domain, "gmial.com", "gmail.com")
	domain = strings.ReplaceAll(domain, "gmai.com", "gmail.com")
	domain = strings.ReplaceAll(domain, "yahooo.com", "yahoo.com")
	domain = strings.ReplaceAll(domain, "hotmial.com", "hotmail.com")

	if domain == "gmail.com" || domain == "googlemail.com" {
		domain = "gmail.com"
		if strings.Contains(local, "+") {
			local = strings.Split(local, "+")[0]
		}
	}

	return local + "@" + domain, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *authService) IsEmailExist(email string) bool {
	existingUser, err := s.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return true
	}
	return false
}

func (s *authService) Register(authDTO *dto.RegisterAuthRequest) (dto.RegisterAuthResponse, error) {
	var userResponse dto.RegisterAuthResponse

	if len(authDTO.FullName) <= 5 {
		return userResponse, errors.New("Name is short it has to be more than 5 letters")
	}
	if len(authDTO.Password) <= 5 {
		return userResponse, errors.New("Password is short it has to be more than 5 numbers")
	}
	email, err := NormalizeAndValidateEmail(authDTO.Email)
	if err != nil {
		return userResponse, err
	}

	if s.IsEmailExist(email) == true {
		return userResponse, errors.New("user already exists")
	}

	hPassword, err := HashPassword(authDTO.Password)
	if err != nil {
		return userResponse, err
	}
	user := &models.User{
		FullName:     strings.Trim(authDTO.FullName, " "),
		PasswordHash: hPassword,
		Email:        email,
	}
	err = s.userRepo.Create(user)
	if err != nil {
		return userResponse, err
	}

	userModel, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return userResponse, errors.New("user not found")
	}
	userResponse = dto.RegisterAuthResponse{
		ID:        userModel.ID,
		FullName:  userModel.FullName,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt,
	}
	return userResponse, nil
}
