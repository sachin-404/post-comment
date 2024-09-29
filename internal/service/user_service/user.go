package user_service

import (
	"errors"
	"github.com/sachin-404/post-comment/internal/models"
	"github.com/sachin-404/post-comment/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req *UserSignupRequest) error {
	existingUser, err := s.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}
	user := models.NewUser(req.Name, req.Email, req.Password)
	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
