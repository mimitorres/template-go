package pkg

import (
	"context"
	goerror "errors"
	"net/http"
	"template-go/src/api/pkg/model"
	"template-go/src/api/util"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type packageService interface {
	GetPackageByID(context.Context, string) (model.Package, error)
}

type PackageController struct {
	Service packageService
}

func (pc PackageController) GetPackageById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		util.AbortWithStatusCode(c, http.StatusBadRequest, "package id cannot be empty", nil)
		return
	}

	res, err := pc.Service.GetPackageByID(c, id)

	if err != nil {
		if goerror.Is(errors.Cause(err), model.ErrBadRequest) {
			util.AbortWithStatusCode(c, http.StatusBadRequest, "invalid parameters", err)
			return
		} else if goerror.Is(errors.Cause(err), model.ErrNotFound) {
			util.AbortWithStatusCode(c, http.StatusNotFound, "package not found", err)
			return
		}
		util.AbortWithStatusCode(c, http.StatusInternalServerError, "unknown error", err)
		return
	}
	c.JSON(http.StatusOK, &res)
}
