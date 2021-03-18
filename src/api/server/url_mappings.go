package server

import "github.com/gin-gonic/gin"

func appendControllers(router gin.IRouter) *ControllersContainer {
	packageController := newPackageController()

	router.GET("/package/number/:id",
		packageController.GetPackageById)

	router.GET("/package/error")

	return &ControllersContainer{
		PackageController: *packageController,
	}
}
