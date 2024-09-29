package repo

import (
	"github.com/sachin-404/post-comment/common/database"
	"github.com/sachin-404/post-comment/internal/models"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
