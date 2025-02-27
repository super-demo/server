package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateCSVOrXLSXFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if the file is present
		file, err := c.FormFile("file")
		if err != nil {
			ResponseError(c, fmt.Errorf("no file is received"))
			c.Abort()
			return
		}

		// validate the file size
		const maxFileSize = 10 << 20 // 10 MB
		if file.Size > maxFileSize {
			ResponseError(c, fmt.Errorf("the file is too large"))
			c.Abort()
			return
		}

		// check the file type
		if !(strings.HasSuffix(file.Filename, ".csv") || strings.HasSuffix(file.Filename, ".xlsx")) {
			ResponseError(c, fmt.Errorf("the file is not a CSV or XLSX file"))
			c.Abort()
			return
		}

		c.Next()
	}
}
