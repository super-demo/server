package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	application "server/infrastructure/app"
	"server/internal/core/handlers"
	"server/internal/core/models"
	"server/pkg/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationSignInWithGoogle(t *testing.T) {

	t.Run("should login google successfully", func(t *testing.T) {
		authUsecase := mocks.NewAuthenticationUsecase(t)
		authorizeToken := "0000"
		accessToken := "1111"
		refreshToken := "1111"
		tokenExpires := time.Now().Add(time.Hour).Unix()
		application.Config = &application.EnvConfigs{
			AppSecret: "app-secret",
		}

		mockToken := &models.TokenResponse{
			AccessTokenResponse: models.AccessTokenResponse{
				AccessToken: accessToken,
				ExpiresAt:   tokenExpires,
			},
			RefreshTokenResponse: models.RefreshTokenResponse{
				RefreshToken:          refreshToken,
				RefreshTokenExpiresAt: tokenExpires,
			},
		}

		authUsecase.On("OrganizationSignInWithGoogle", authorizeToken).Return(mockToken, nil)

		gin.SetMode(gin.TestMode)
		app := gin.New()
		handlers.NewAuthenticationHandler(app, authUsecase)

		mockReqBody := models.GoogleSignInRequest{
			AccessToken: authorizeToken,
		}

		body, _ := json.Marshal(mockReqBody)

		req := httptest.NewRequest("POST", "/v1/authentication/organization/sign/google", bytes.NewReader(body))
		req.Header.Set("App-Secret", application.Config.AppSecret)

		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)

		var response map[string]interface{}

		json.Unmarshal(res.Body.Bytes(), &response)
		data, _ := response["data"].(map[string]interface{})

		assert.Equal(t, data["access_token"], accessToken)
		assert.Equal(t, int(data["expires_at"].(float64)), int(tokenExpires))

	})
	t.Run("should fail to login because invalid app secret", func(t *testing.T) {
		authUsecase := mocks.NewAuthenticationUsecase(t)
		authorizeToken := "0000"

		application.Config = &application.EnvConfigs{
			AppSecret: "app-secret",
		}

		gin.SetMode(gin.TestMode)
		app := gin.New()
		handlers.NewAuthenticationHandler(app, authUsecase)

		mockReqBody := models.GoogleSignInRequest{
			AccessToken: authorizeToken,
		}

		body, _ := json.Marshal(mockReqBody)

		req := httptest.NewRequest("POST", "/v1/authentication/organization/sign/google", bytes.NewReader(body))
		req.Header.Set("App-Secret", "invalid-app-secret")

		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)

		var response map[string]interface{}

		json.Unmarshal(res.Body.Bytes(), &response)
		status, _ := response["status"].(map[string]interface{})
		assert.Equal(t, int(status["code"].(float64)), application.ErrInvalidAppSecret.ErrCode)
		assert.Equal(t, status["message"], application.ErrInvalidAppSecret.Err.Error())

	})
}
