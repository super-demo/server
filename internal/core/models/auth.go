package models

import "github.com/golang-jwt/jwt"

type GoogleSignInRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type RefreshTokenResponse struct {
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}

type RefreshTokenClaims struct {
	UserId    int   `json:"user_id"`
	ExpiresIn int64 `json:"exp"`
	jwt.StandardClaims
}

type TokenResponse struct {
	AccessTokenResponse
	RefreshTokenResponse
}

type JwtPayload struct {
	UserId      int
	UserLevelId int
	Email       string
	Name        string
}
