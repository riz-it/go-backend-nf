package domain

import (
	"context"

	"riz.it/nurul-faizah/internal/dto"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error)
	SignOut(ctx context.Context, userID uint) error
	Refresh(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.SignInResponse, error)
}
