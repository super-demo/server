package handlers

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler interface {
	SignWithGoogle(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authenticationHandler struct {
	authenticationUsecase usecases.AuthenticationUsecase
}

func NewAuthenticationHandler(r *gin.Engine, authenticationUsecase usecases.AuthenticationUsecase, globalMiddlewares ...gin.HandlerFunc) AuthenticationHandler {
	handler := &authenticationHandler{authenticationUsecase}

	v1 := r.Group("/v1/authentications", globalMiddlewares...)

	SignWithGoogle := []gin.HandlerFunc{
		middlewares.ValidateAppSecret(),
		middlewares.ValidateRequestBody(&models.GoogleSignInRequest{}),
		handler.SignWithGoogle,
	}

	UserSignWithGoogle := []gin.HandlerFunc{
		middlewares.ValidateAppSecret(),
		middlewares.ValidateRequestBody(&models.GoogleSignInRequest{}),
		handler.UserSignWithGoogle,
	}

	v1.POST("/sign/google", SignWithGoogle...)
	v1.POST("/user/sign/google", UserSignWithGoogle...)
	v1.POST("/token/refresh", middlewares.BearerAuth(), handler.RefreshToken)

	return handler
}

func (h *authenticationHandler) SignWithGoogle(c *gin.Context) {
	var req models.GoogleSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	data, err := h.authenticationUsecase.SignWithGoogle(req.AccessToken)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, data, "Sign in successfully")
}

func (h *authenticationHandler) UserSignWithGoogle(c *gin.Context) {
	var req models.GoogleSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	data, err := h.authenticationUsecase.SignWithGoogle(req.AccessToken)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, data, "Sign in successfully")
}

func (h *authenticationHandler) RefreshToken(c *gin.Context) {
	refreshToken, exits := c.Get("token")
	if !exits {
		middlewares.ResponseError(c, app.ErrInvalidToken)
		return
	}

	refreshTokenString, ok := refreshToken.(string)
	if !ok {
		middlewares.ResponseError(c, app.ErrInvalidToken)
		return
	}

	accessTokenResponse, err := h.authenticationUsecase.RefreshToken(refreshTokenString)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, accessTokenResponse, "Token refreshed successfully")
}
