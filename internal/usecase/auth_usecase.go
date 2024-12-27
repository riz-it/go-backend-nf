package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"riz.it/nurul-faizah/internal/domain"
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/entity"
	"riz.it/nurul-faizah/internal/utils"
)

type AuthUseCase struct {
	DB                    *gorm.DB
	Validate              *validator.Validate
	UserAccountRepository domain.UserAccountRepository
	Log                   *logrus.Logger
	JwtHelper             domain.JWTHelper
}

func NewAuthUseCase(db *gorm.DB, log *logrus.Logger, v *validator.Validate, userAccountRepository domain.UserAccountRepository, jwtHelper domain.JWTHelper) domain.AuthUseCase {
	return &AuthUseCase{
		DB:                    db,
		Validate:              v,
		Log:                   log,
		UserAccountRepository: userAccountRepository,
		JwtHelper:             jwtHelper,
	}
}

// Create implements domain.AuthUseCase.
func (uA *AuthUseCase) SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	tx := uA.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, err := uA.UserAccountRepository.CountByEmail(tx, req.Email)
	if err != nil {
		uA.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		uA.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "email already exists")
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		uA.Log.Warnf("failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.UserAccount{
		AccountName: req.AccountName,
		Email:       req.Email,
		Password:    string(password),
	}

	if err := uA.UserAccountRepository.Create(tx, user); err != nil {
		uA.Log.Warnf("failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		uA.Log.Warnf("failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.SignUpToResponse(user), nil
}

// SignIn implements domain.AuthUseCase.
func (uA *AuthUseCase) SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error) {
	tx := uA.DB.WithContext(ctx)
	user := new(entity.UserAccount)
	err := uA.UserAccountRepository.FindByEmail(tx, user, req.Email)

	if err != nil {
		uA.Log.Warnf("Invalid email or password : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	if exist := utils.VerifyPassword(user.Password, req.Password); !exist {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	accessToken, refreshToken, err := uA.JwtHelper.GenerateTokens(uint(user.ID))
	if err != nil {
		uA.Log.WithError(err).Error("Failed to generate tokens")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	user.HashedRt = refreshToken
	if err := uA.UserAccountRepository.Update(tx, user); err != nil {
		uA.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.SignInToResponse(user, accessToken, refreshToken), nil

}

// Refresh implements domain.AuthUseCase.
func (uA *AuthUseCase) Refresh(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.SignInResponse, error) {

	userID, err := uA.JwtHelper.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid refresh token")
	}

	tx := uA.DB.WithContext(ctx)
	user := new(entity.UserAccount)
	if err := uA.UserAccountRepository.FindByID(tx, user, userID); err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if user.HashedRt != req.RefreshToken {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid refresh token")
	}

	accessToken, refreshToken, err := uA.JwtHelper.GenerateTokens(uint(user.ID))
	if err != nil {
		uA.Log.WithError(err).Error("Failed to generate tokens")
		return nil, fiber.ErrInternalServerError
	}

	user.HashedRt = refreshToken
	if err := uA.UserAccountRepository.Update(tx, user); err != nil {
		uA.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.SignInToResponse(user, accessToken, refreshToken), nil

}
