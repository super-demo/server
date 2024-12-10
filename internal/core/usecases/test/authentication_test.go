package usecases_test

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/pkg/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCmsSignInWithGoogle(t *testing.T) {
	t.Run("should sign in with google successfully", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		authenticationRepo := mocks.NewAuthenticationRepository(t)
		organizationUserRepo := mocks.NewOrganizationUserRepository(t)

		authenticationUsecase := usecases.NewAuthenticationUsecase(userRepo, authenticationRepo, organizationUserRepo)

		app.Config = &app.EnvConfigs{
			Environment: "production",
		}

		expectedUser := &models.User{
			UserId:      1,
			GoogleToken: "GoogleToken000",
			AvatarUrl:   "https://example.com/avatar.jpg",
			Name:        "John Doe",
			Email:       "",
		}

		authenticationRepo.On("GetUserInfoByAccessToken", mock.Anything).Return(&models.UserInfoResponse{
			Id:            "GoogleToken000",
			Email:         "",
			VerifiedEmail: true,
			Picture:       "https://example.com/avatar.jpg",
			HD:            "example.com",
		}, nil)

		userRepo.On("GetUserByEmail", mock.Anything).Return(expectedUser, nil)
		userRepo.On("UpdateUser", mock.Anything).Return(nil, nil)

		_, err := authenticationUsecase.CmsSignInWithGoogle("googleAccessToken000")

		assert.NoError(t, err)
	})
}

func TestOrganizationSignInWithGoogle(t *testing.T) {
	t.Run("should sign in with google successfully", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		authenticationRepo := mocks.NewAuthenticationRepository(t)
		organizationUserRepo := mocks.NewOrganizationUserRepository(t)

		authenticationUsecase := usecases.NewAuthenticationUsecase(userRepo, authenticationRepo, organizationUserRepo)

		app.Config = &app.EnvConfigs{
			Environment: "production",
		}

		expectedUser := &models.User{
			UserId:      1,
			GoogleToken: "GoogleToken000",
			AvatarUrl:   "https://example.com/avatar.jpg",
			Name:        "John Doe",
			Email:       "john@example.com",
		}

		expectedOrganizationUser := &models.OrganizationUser{
			OrganizationUserId: 1,
			UserId:             1,
			OrganizationId:     1,
		}

		authenticationRepo.On("GetUserInfoByAccessToken", mock.Anything).Return(&models.UserInfoResponse{
			Id:            "GoogleToken000",
			Email:         "john@example.com",
			VerifiedEmail: true,
			Picture:       "https://example.com/avatar.jpg",
			HD:            "example.com",
		}, nil)

		userRepo.On("GetUserByEmail", mock.Anything).Return(expectedUser, nil)
		userRepo.On("UpdateUser", mock.Anything).Return(nil, nil)
		organizationUserRepo.On("GetOrganizationUserByEmail", mock.Anything).Return(expectedOrganizationUser, nil)

		_, err := authenticationUsecase.OrganizationSignInWithGoogle("googleAccessToken000")

		assert.NoError(t, err)
	})

	t.Run("should fail to sign in because user not found", func(t *testing.T) {
		userRepo := mocks.NewUserRepository(t)
		authenticationRepo := mocks.NewAuthenticationRepository(t)
		organizationUserRepo := mocks.NewOrganizationUserRepository(t)

		authenticationUsecase := usecases.NewAuthenticationUsecase(userRepo, authenticationRepo, organizationUserRepo)

		app.Config = &app.EnvConfigs{
			Environment: "production",
		}

		authenticationRepo.On("GetUserInfoByAccessToken", mock.Anything).Return(&models.UserInfoResponse{
			Id:            "GoogleToken000",
			Email:         "john@example.com",
			VerifiedEmail: true,
			Picture:       "https://example.com/avatar.jpg",
			HD:            "example.com",
		}, nil)

		userRepo.On("GetUserByEmail", mock.Anything).Return(nil, app.ErrUserNotFound)
		_, err := authenticationUsecase.OrganizationSignInWithGoogle("googleAccessToken000")

		assert.Error(t, err)
		assert.Equal(t, app.ErrUserNotFound, err)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Skip("skipping test")
}
