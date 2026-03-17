package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sidz111/user-auth/model"
	"github.com/sidz111/user-auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id uint) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.UUID = uuid.NewString()
	user.Password = string(hashedPass)
	return s.repo.CreateUser(ctx, user)
}
func (s *userService) GetUser(ctx context.Context, id uint) (*model.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID should be positive")
	}
	return s.repo.GetUser(ctx, id)
}
func (s *userService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.GetAllUsers(ctx)
}
func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.repo.UpdateUser(ctx, user)
}
func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.DeleteUser(ctx, id)
}
