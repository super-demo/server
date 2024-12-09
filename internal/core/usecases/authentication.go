package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthenticationUsecase interface {
	GoogleSignIn(token string) (*models.TokenResponse, error)
	RefreshToken(refreshToken string) (*models.AccessTokenResponse, error)
}

type authenticationUsecase struct {
	userRepo   repositories.UserRepository
	googleRepo repositories.AuthenticationRepository
}

func NewAuthenticationUsecase(
	userRepo repositories.UserRepository,
	authenticationRepo repositories.AuthenticationRepository,
) AuthenticationUsecase {
	return &authenticationUsecase{userRepo, authenticationRepo}
}

func (u *authenticationUsecase) GoogleSignIn(token string) (*models.TokenResponse, error) {
	userInfo, err := u.googleRepo.GetUserInfoByAccessToken(token)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(userInfo.Email, true)
	if err != nil {
		return nil, app.ErrUserNotFound
	}

	if user.GoogleToken == "" {
		user.GoogleToken = userInfo.Id
	}

	user.AvatarUrl = userInfo.Picture

	if _, err := u.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	payload := models.JwtPayload{
		UserId: user.UserId,
		// FYI: This is a dummy value
		// TODO: UserLevelId should be fetched from the database organization_user table
		UserLevelId: 1,
		Email:       user.Email,
		Name:        user.Name,
	}

	result, err := utils.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *authenticationUsecase) RefreshToken(refreshToken string) (*models.AccessTokenResponse, error) {
	claims := &models.RefreshTokenClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, app.ErrInvalidToken
		}
		return []byte(app.Config.Jwt.JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, app.ErrInvalidToken
	}

	expirationTime := time.Unix(claims.ExpiresAt, 0)
	if expirationTime.Before(time.Now()) {
		return nil, app.ErrTokenExpired
	}

	user, err := u.userRepo.GetUserById(claims.UserId)
	if err != nil {
		return nil, err
	}

	accessTokenExpire := utils.CalculateExpiration(app.Config.Jwt.JwtTokenExpire)

	payload := models.JwtPayload{
		UserId:      user.UserId,
		UserLevelId: 1,
		Email:       user.Email,
		Name:        user.Name,
	}

	signedAccessToken, err := utils.GenerateAccessToken(payload, accessTokenExpire)
	if err != nil {
		return nil, err
	}

	response := &models.AccessTokenResponse{
		AccessToken: signedAccessToken,
		ExpiresAt:   accessTokenExpire,
	}

	return response, nil
}
