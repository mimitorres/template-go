package main

import (
	"fmt"
	"net/http"
	"os"
	"template-go/src/api/server"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		logrus.Errorf("error running server", err)
	}
}

func run(port string) error {
	gin.SetMode(gin.ReleaseMode)

	s, _ := server.New()

	health := HealthChecker{}

	mapRoutes(s, health)

	return s.Run(fmt.Sprintf(":%s", port))
}

func mapRoutes(r *gin.Engine, health HealthChecker) {
	r.GET("/ping", health.PingHandler)
}

type HealthChecker struct{}

func (h *HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
