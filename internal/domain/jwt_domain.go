package domain

type JWTHelper interface {
	GenerateTokens(userID uint) (string, string, error)
	ValidateAccessToken(tokenString string) (uint, error)
	ValidateRefreshToken(tokenString string) (uint, error)
}
