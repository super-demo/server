package app

import (
	"errors"
)

type AppError struct {
	ErrHTTPCode int
	ErrCode     int
	Err         error
}

type ErrorResponse struct {
	Data   *interface{} `json:"data"`
	Status *ErrorStatus `json:"status"`
}

type ErrorStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		ErrHTTPCode: e.ErrHTTPCode,
		ErrCode:     e.ErrCode,
		Err:         errors.New(message),
	}
}

// Use for return nil error
var ErrNil AppError

// HTTP-Based Application errors
var (
	// 400 Bad Request
	ErrBadRequest     = AppError{400, 400001, errors.New("bad request")}
	ErrInvalidUserId  = AppError{400, 400002, errors.New("invalid user id")}
	ErrInvalidRequest = AppError{400, 400003, errors.New("invalid request")}

	// 401 Unauthorized
	ErrAccessTokenEmpty  = AppError{401, 401001, errors.New("access token is empty")}
	ErrUnauthorized      = AppError{401, 401002, errors.New("unauthorized")}
	ErrTokenExpired      = AppError{401, 401003, errors.New("token is expired")}
	ErrRefreshTokenEmpty = AppError{401, 401001, errors.New("refresh token is empty")}

	// 403 Forbidden
	ErrInvalidToken             = AppError{403, 403001, errors.New("token is invalid")}
	ErrInvalidAppSecret         = AppError{403, 403002, errors.New("app secret is invalid")}
	ErrInvalidCompanyDomainName = AppError{403, 403003, errors.New("company domain name is invalid")}
	ErrInsufficientPermissions  = AppError{403, 403004, errors.New("insufficient permissions to access this resource")}
	ErrAccessAnotherUser        = AppError{403, 403005, errors.New("you don't have permission to access another user's data")}
	ErrInvalidUserLevelId       = AppError{403, 403006, errors.New("invalid user level id")}

	// 404 Not Found
	ErrNotFound = AppError{404, 404001, errors.New("not found")}

	// 409 Conflict
	ErrConflict   = AppError{409, 409001, errors.New("conflict")}
	ErrEmailExist = AppError{409, 409002, errors.New("email already exist")}
	ErrNameExist  = AppError{409, 409003, errors.New("name already exist")}

	// 422 Unprocessable Entity
	ErrUnprocessableEntity     = AppError{422, 422001, errors.New("unprocessable entity")}
	ErrMissingRequiredFields   = AppError{422, 422002, errors.New("missing required fields")}
	ErrExceedCharacterLimit    = AppError{422, 422003, errors.New("exceed character limit")}
	ErrValidationFailed        = AppError{422, 422004, errors.New("validation failed")}
	ErrExcludedUnlessCondition = AppError{422, 422005, errors.New("excluded field unless condition")}

	// 500 Internal Server Error
	ErrInternalServerError = AppError{500, 500001, errors.New("internal server error")}
)

// ToJSONResponse
//
// / helper method to use when need to return error message as json format in the delivery layer
func (e *AppError) ToJSONResponse() ErrorResponse {
	return ErrorResponse{
		Data: nil,
		Status: &ErrorStatus{
			Code:    e.ErrCode,
			Message: e.Err.Error(),
		},
	}
}
