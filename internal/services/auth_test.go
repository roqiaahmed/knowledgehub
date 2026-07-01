package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/roqiaahmed/knowledgehub/internal/dto"
	"github.com/roqiaahmed/knowledgehub/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedUserRepo struct {
	mock.Mock
}

func (m *MockedUserRepo) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockedUserRepo) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func Test_Register(t *testing.T) {
	userRepo := new(MockedUserRepo)
	authService := NewAuthService(userRepo)
	authDTO := dto.RegisterAuthRequest{
		FullName: "Roqia Ahmed",
		Email:    "roqia@gmail.com",
		Password: "Password",
	}
	user := &models.User{
		ID:        uuid.UUID{2},
		FullName:  authDTO.FullName,
		Email:     authDTO.Email,
		CreatedAt: time.Time{},
	}
	t.Run("Register succeeds when the email doesn't already exist", func(t *testing.T) {
		userRepo.On("FindByEmail", authDTO.Email).Return((*models.User)(nil), nil).Once()
		userRepo.On("Create", mock.Anything).Return(nil).Once()
		userRepo.On("FindByEmail", authDTO.Email).Return(user, nil).Once()

		currUser, err := authService.Register(&authDTO)

		assert.NoError(t, err)
		assert.Equal(t, "roqia@gmail.com", currUser.Email)

		userRepo.AssertExpectations(t)
	})
	t.Run("Register fails when the email already exist", func(t *testing.T) {
		userRepo.On("FindByEmail", authDTO.Email).Return(user, nil).Once()
		_, err := authService.Register(&authDTO)
		assert.EqualError(t, err, "user already exists")
	})
	t.Run("Register fails when the email is Invalid", func(t *testing.T) {
		invalidUser := dto.RegisterAuthRequest{
			FullName: "Roqia Ahmed",
			Email:    "user name@domain.com",
			Password: "Password",
		}
		_, err := authService.Register(&invalidUser)
		assert.EqualError(t, err, "invalid email format")
	})
	t.Run("Register fails for short password", func(t *testing.T) {
		invalidUser := dto.RegisterAuthRequest{
			FullName: "Roqia Ahmed",
			Email:    "name@domain.com",
			Password: "Pass",
		}
		_, err := authService.Register(&invalidUser)
		assert.ErrorContains(t, err, "Password is short")
	})
	t.Run("Register fails for short fullname", func(t *testing.T) {
		invalidUser := dto.RegisterAuthRequest{
			FullName: "Roq",
			Email:    "name@domain.com",
			Password: "Password",
		}
		_, err := authService.Register(&invalidUser)
		assert.ErrorContains(t, err, "Name is short")
	})
}
