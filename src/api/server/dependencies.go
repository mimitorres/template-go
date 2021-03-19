package server

import (
	"net/http"
	"template-go/src/api/pkg"
	"template-go/src/api/util/restcli"
	"time"
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
	//REVIEW: baseURL, _ := conf.String("template-go.resources.clients.mocky.uri")
	baseURL := "https://run.mocky.io/v3"
	packageClient := newRestClient().Build()
	return &pkg.PackageRepository{
		BaseURL: baseURL,
		Client:  packageClient,
	}
}

func newRestClient() restcli.ClientBuilder {
	return restcli.ClientBuilder{
		Headers:        http.Header{"X-Caller-Scopes": {"admin,write"}},
		Timeout:        2 * time.Second,
		DisableTimeout: false,
	}
}
