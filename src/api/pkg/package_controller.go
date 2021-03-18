package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type packageService interface {
}

type PackageController struct {
	Service packageService
}

type PackageResponse struct {
	Message string `json:"message",omitempty`
}

func (pc PackageController) GetPackageById(c *gin.Context) {
	id := c.Param("id")
	message := PackageResponse{
		Message: fmt.Sprintf("The id is %s", id),
	}
	c.JSON(http.StatusOK, message)
}

type CauseList []interface{}

type apiError struct {
	Status  int       `json:"status"`
	Message string    `json:"message",omitempty`
	Cause   CauseList `json:"cause"`
}

func abortWithStatusCode(c *gin.Context, status int, msg string, err error) {
	var cause CauseList
	if err != nil {
		cause = CauseList{err.Error()}
	}
	c.AbortWithStatusJSON(status, apiError{
		Status:  status,
		Message: msg,
		Cause:   cause,
	})
}
