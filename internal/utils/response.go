package utils

import (
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/entity"
)

func SignUpToResponse(userAccount *entity.UserAccount) *dto.SignUpResponse {
	return &dto.SignUpResponse{
		AccountName: userAccount.AccountName,
		Email:       userAccount.Email,
	}
}

func SignInToResponse(userAccount *entity.UserAccount, accessToken string, refreshToken string) *dto.SignInResponse {
	return &dto.SignInResponse{
		User: dto.CredentialData{
			Name:  userAccount.AccountName,
			Email: userAccount.Email,
		},
		Token: dto.TokenData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}
