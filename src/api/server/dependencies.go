package server

import (
	"template-go/src/api/pkg"
)

type ControllersContainer struct {
	//Your Controllers here
	PackageController pkg.PackageController
}

func newPackageController() *pkg.PackageController {
	return &pkg.PackageController{
		Service: newPackageService(),
	}
}

func newPackageService() *pkg.PackageService {
	return &pkg.PackageService{
		Repository: newPackageRepository(),
	}
}

func newPackageRepository() *pkg.PackageRepository {
	baseURL, _ := conf.String("resources.clients.example-client")
	return &pkg.PackageRepository{
		BaseURL: baseURL,
	}
}
