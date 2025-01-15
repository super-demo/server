package handlers

import (
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserById(c *gin.Context)
}

type userHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase usecases.UserUsecase, globalMiddlewares ...gin.HandlerFunc) UserHandler {
	handler := &userHandler{userUsecase}

	v1 := r.Group("/v1/users", globalMiddlewares...)

	getUserById := []gin.HandlerFunc{
		handler.GetUserById,
	}

	v1.GET("/:id/profile", getUserById...)

	return handler
}

func (h *userHandler) GetUserById(c *gin.Context) {
	userId := utils.GetIdFromParams(c)
	requesterUserId := c.MustGet("user_id").(int)

	user, err := h.userUsecase.GetUserById(userId, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, user, "User retrieved successfully")
}
