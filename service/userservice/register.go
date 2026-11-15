package userservice

import (
	"fmt"
	"gameauth_v2/entity"
	"gameauth_v2/param"
	"gameauth_v2/pkg/security/password"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("can't hash password")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    hashedPassword,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.RegisterResponse{User: param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}
