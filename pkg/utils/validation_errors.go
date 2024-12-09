package utils

import (
	"fmt"
	"reflect"
	"server/infrastructure/app"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(errs validator.ValidationErrors, modelType interface{}) *app.AppError {
	e := errs[0]

	fieldName := e.StructField()
	field, ok := reflect.TypeOf(modelType).Elem().FieldByName(fieldName)
	if ok {
		jsonTag := field.Tag.Get("json")
		jsonFieldName := strings.Split(jsonTag, ",")[0]
		if jsonFieldName != "" {
			fieldName = jsonFieldName
		}
	}

	var err *app.AppError
	switch e.Tag() {
	case "required":

		err = app.ErrMissingRequiredFields.WithMessage(fmt.Sprintf("%s is required", fieldName))
	case "max":
		err = app.ErrExceedCharacterLimit.WithMessage(fmt.Sprintf("%s must not exceed %s characters", fieldName, e.Param()))
	case "excluded_unless":
		err = app.ErrExcludedUnlessCondition.WithMessage(fmt.Sprintf("%s must be empty unless %s", fieldName, e.Param()))
	default:
		err = app.ErrValidationFailed.WithMessage(fmt.Sprintf("Validation failed on the %s field", fieldName))
	}

	return err
}
