package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(handler func(ctx *gin.Context) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}
