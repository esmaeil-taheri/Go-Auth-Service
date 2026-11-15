package userservice

import (
	"fmt"
	"gameauth_v2/param"
	"gameauth_v2/pkg/richerror"
	"gameauth_v2/pkg/security/password"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "userservice.Login"

	// TODO - it would be better to use two seprate method for existence check and GetUserByPhoneNumber
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).WithMeta(
			map[string]interface{}{"phone_number": req.PhoneNumber},
		)
	}

	if !password.Compare(user.Password, req.Password) {
		return param.LoginResponse{}, fmt.Errorf("username or password is incorrect")
	}

	accesstoken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshtoken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.LoginResponse{
		User: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: param.Tokens{
			AccessToken:  accesstoken,
			RefreshToken: refreshtoken,
		},
	}, nil
}
