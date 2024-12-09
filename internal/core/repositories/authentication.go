package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/infrastructure/app"
	"server/internal/core/models"
)

type AuthenticationRepository interface {
	GetUserInfoByAccessToken(token string) (*models.UserInfoResponse, error)
}

type authenticationRepository struct{}

func NewAuthenticationRepository() AuthenticationRepository {
	return &authenticationRepository{}
}

func (r *authenticationRepository) GetUserInfoByAccessToken(token string) (*models.UserInfoResponse, error) {
	const GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s"

	userInfoUrl := fmt.Sprintf(GoogleUserInfoURL, token)
	resp, err := http.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var userInfo models.UserInfoResponse

	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	if userInfo.Email == "" {
		return nil, app.ErrInvalidToken
	}

	return &userInfo, nil
}
