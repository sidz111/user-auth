package repository

import (
	"context"

	"github.com/sidz111/user-auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id uint) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *userRepository) GetUser(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	result := r.db.WithContext(ctx).Model(&model.User{}).Where("id=?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	result := r.db.WithContext(ctx).Find(model.User{}, &users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	result := r.db.WithContext(ctx).Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
