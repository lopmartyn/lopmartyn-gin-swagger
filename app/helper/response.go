package helper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// HTTPError error response result
type HTTPError struct {
	Code    int    `json:"code" example:"999"`
	Message string `json:"message" example:"example response"`
}

// ResponseCustomError response customize error message
func ResponseCustomError(ctx *gin.Context, status int, errMsg string) {
	er := HTTPError{
		Code:    status,
		Message: errMsg,
	}
	ctx.JSON(status, er)
}

// ResponseError response error
func ResponseError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// RecordNotExistError record doesn't exist
func RecordNotExistError(ctx *gin.Context, status int, id int64, model string) {
	er := HTTPError{
		Code:    status,
		Message: model + " ID doesn't not exist: " + strconv.FormatInt(id, 10),
	}
	ctx.JSON(status, er)
}
