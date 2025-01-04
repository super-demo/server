package utils

import (
	"fmt"
	"server/infrastructure/app"
	"server/internal/core/models"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(payload models.JwtPayload) (*models.TokenResponse, error) {
	accessTokenExpire := CalculateExpiration(app.Config.Jwt.JwtTokenExpire)
	refreshTokenExpire := CalculateExpiration(app.Config.Jwt.JwtRefreshTokenExpire)

	signedAccessToken, err := GenerateAccessToken(payload, accessTokenExpire)
	if err != nil {
		return nil, err
	}

	signedRefreshToken, err := GenerateRefreshToken(payload, refreshTokenExpire)
	if err != nil {
		return nil, err
	}

	response := &models.TokenResponse{
		AccessTokenResponse: models.AccessTokenResponse{
			AccessToken: signedAccessToken,
			ExpiresAt:   accessTokenExpire,
		},
		RefreshTokenResponse: models.RefreshTokenResponse{
			RefreshToken:          signedRefreshToken,
			RefreshTokenExpiresAt: refreshTokenExpire,
		},
	}

	return response, nil
}

func GenerateAccessToken(payload models.JwtPayload, exp int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = payload.UserId
	claims["user_level_id"] = payload.UserLevelId
	claims["email"] = payload.Email
	claims["name"] = fmt.Sprintf("%s", payload.Name)
	claims["exp"] = exp

	return token.SignedString([]byte(app.Config.Jwt.JwtSecretKey))
}

func GenerateRefreshToken(payload models.JwtPayload, exp int64) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_id"] = payload.UserId
	rtClaims["exp"] = exp

	return refreshToken.SignedString([]byte(app.Config.Jwt.JwtSecretKey))
}
