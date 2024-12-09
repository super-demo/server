package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"server/infrastructure/app"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "can't read request body"})
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		modelType := reflect.TypeOf(model)
		newModel := reflect.New(modelType.Elem()).Interface()

		if err := json.Unmarshal(bodyBytes, &newModel); err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, app.ErrBadRequest.ToJSONResponse())
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, app.ErrInternalServerError.WithMessage(err.Error()).ToJSONResponse())
			}
			return
		}

		validate := validator.New()
		if err := validate.Struct(newModel); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				appError := utils.ParseValidationErrors(errs, newModel)
				c.JSON(http.StatusUnprocessableEntity, appError.ToJSONResponse())
				c.Abort()
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, app.ErrInternalServerError.WithMessage(err.Error()).ToJSONResponse())
			}
			return
		}

		c.Next()
	}
}
