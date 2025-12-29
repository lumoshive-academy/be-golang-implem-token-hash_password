package service

import (
	"errors"
	"session-22/model"
	"session-22/repository"
	"session-22/utils"
)

type AuthService interface {
	Login(email, password string) (*model.User, error)
}

type authService struct {
	Repo repository.Repository
}

func NewAuthService(repo repository.Repository) AuthService {
	return &authService{Repo: repo}
}

func (s *authService) Login(email, password string) (*model.User, error) {
	// user, err := s.Repo.UserRepo.FindByEmail(email)
	// if err != nil {
	// 	return nil, errors.New("user not found")
	// }

	// if user.Password != password {
	// 	return nil, errors.New("incorrect password")
	// }
	password_store := "$2a$10$sQYHDCwNYSKsOvxxw.hQZOm4SR6igvap2Qcf2xmyoR.x08GSu42BW"

	if !utils.CheckPassword(password, password_store) {
		return nil, errors.New("incorrect password")
	}

	// password_hash := utils.HashPassword(password)

	user := &model.User{
		Name:  "lumos",
		Email: "lumos@email.com",
		Role:  "admin",
	}

	return user, nil
}
