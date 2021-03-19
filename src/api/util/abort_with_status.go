package util

import (
	"template-go/src/api/pkg/model"

	"github.com/gin-gonic/gin"
)

func AbortWithStatusCode(c *gin.Context, status int, msg string, err error) {
	var cause model.CauseList
	if err != nil {
		cause = model.CauseList{err.Error()}
	}
	c.AbortWithStatusJSON(status, model.ApiError{
		Status:  status,
		Message: msg,
		Cause:   cause,
	})
}
