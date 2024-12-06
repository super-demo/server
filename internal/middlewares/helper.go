package middlewares

import (
	"encoding/json"
	"net/http"
	"server/infrastructure/app"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ReturnResponse struct {
	Status StatusResponse `json:"status"`
	Data   interface{}    `json:"data"`
}

func ResponseError(c *gin.Context, err error) {
	switch v := err.(type) {
	case app.AppError:
		c.JSON(v.ErrHTTPCode, v.ToJSONResponse())
	case error:
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": gin.H{
				"code":    500001,
				"message": err.Error(),
			},
			"data": nil,
		})
	}
}

func ResponseErrorCustom(c *gin.Context, statusCode int, errCode int, errMessage string) {
	c.JSON(statusCode, ReturnResponse{
		Status: StatusResponse{
			Code:    errCode,
			Message: errMessage,
		},
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}, opts ...interface{}) {
	defaultCode := 200
	defaultMessage := "Success"

	code := defaultCode
	message := defaultMessage

	for _, opt := range opts {
		switch v := opt.(type) {
		case string:
			message = v
		case int:
			code = v
		}
	}

	status := StatusResponse{
		Code:    code,
		Message: message,
	}

	_, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": gin.H{
				"code":    500001,
				"message": err.Error(),
			},
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, ReturnResponse{
		Status: status,
		Data:   data,
	})
}
