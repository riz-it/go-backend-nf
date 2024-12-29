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

func ClassToResponse(class *entity.Class) *dto.ClassResponse {
	return &dto.ClassResponse{
		ID:        class.ID,
		Name:      class.Name,
		IsActive:  class.IsActive,
		Leader:    class.Leader,
		CreatedAt: class.CreatedAt.Time.String(),
		UpdatedAt: class.CreatedAt.Time.String(),
	}
}
