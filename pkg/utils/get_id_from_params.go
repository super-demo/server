package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParams(c *gin.Context) int {
	id := 0
	if idParam, err := strconv.Atoi(c.Param("id")); err == nil {
		id = idParam
	}
	return id
}
