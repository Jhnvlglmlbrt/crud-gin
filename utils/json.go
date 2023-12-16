package utils

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/crud-gin/data/response"
	"github.com/gin-gonic/gin"
)

func RespondWithJson(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, response.Response{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	})
}
