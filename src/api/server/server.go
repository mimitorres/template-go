package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var production = os.Getenv("GO_ENVIRONMENT") == "production"

func New() (*gin.Engine, *ControllersContainer) {
	router := gin.New()
	group := monitoredGroup(router)
	container := appendControllers(group)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("resource %s not found.", c.Request.URL.Path)))
	})

	return router, container
}

func monitoredGroup(router *gin.Engine) *gin.RouterGroup {
	group := router.Group("/")

	group.Use(gzip.Gzip(gzip.DefaultCompression))
	group.Use(gin.Logger())

	return group
}
