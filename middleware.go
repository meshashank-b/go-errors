package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		if len(c.Errors) > 0 {
			var appErr *AppError
			for _, ginErr := range c.Errors {
				if errors.As(ginErr.Err, &appErr) {
					c.JSON(appErr.HTTPStatus, appErr)
					return
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}
}
