package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"riz.it/nurul-faizah/internal/config"
	"riz.it/nurul-faizah/internal/domain"
)

type JWTHelperImpl struct {
	Config *config.Bootstrap
}

func NewJWTHelperImpl(config *config.Bootstrap) domain.JWTHelper {
	return &JWTHelperImpl{
		Config: config,
	}
}

// GenerateTokens implements domain.JWTHelper.
func (j *JWTHelperImpl) GenerateTokens(userID uint) (string, string, error) {
	iss := j.Config.Server.Host
	aud := j.Config.Server.Host
	exp := time.Now().Add(time.Hour * 2).Unix()
	expRefresh := time.Now().Add(time.Hour * 24).Unix()
	sub := userID

	accessKey := j.Config.Jwt.AccessTokenKey
	refreshKey := j.Config.Jwt.RefreshTokenKey

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": iss,
			"aud": aud,
			"exp": exp,
			"sub": sub,
		})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": iss,
			"aud": aud,
			"exp": expRefresh,
			"sub": sub,
		})

	signedAccessToken, err := accessToken.SignedString([]byte(accessKey))
	if err != nil {
		return "", "", err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(refreshKey))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil

}

// ValidateAccessToken implements domain.JWTHelper.
func (j *JWTHelperImpl) ValidateAccessToken(tokenString string) (uint, error) {
	panic("unimplemented")
}

// ValidateRefreshToken implements domain.JWTHelper.
func (j *JWTHelperImpl) ValidateRefreshToken(tokenString string) (uint, error) {
	panic("unimplemented")
}
