package pkg

import (
	"context"
	"fmt"
	"net/http"
	"template-go/src/api/pkg/model"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/pkg/errors"
)

type client interface {
	GET(ctx context.Context, URL string, structure interface{}) (*rest.Response, error)
}

type PackageRepository struct {
	Client  client
	BaseURL string
}

const (
	mockURL = "/e813b6a9-239b-4d61-9d26-ba5f4c37e268"
)

func (r PackageRepository) GetPackageByID(ctx context.Context, id int) (model.Package, error) {
	var response model.MockyResponse
	var err error
	var pkg model.Package

	queryUrl := fmt.Sprintf("%s%s", r.BaseURL, mockURL)

	res, err := r.Client.GET(ctx, queryUrl, &response)

	if err != nil {
		if res.Response.StatusCode == http.StatusBadRequest {
			return pkg, errors.Wrap(model.ErrBadRequest, err.Error())
		} else if res.Response.StatusCode == http.StatusNotFound {
			return pkg, errors.Wrap(model.ErrNotFound, err.Error())
		}
		return pkg, err
	}

	pkg = model.Package{
		ID:          id,
		Description: fmt.Sprintf("%s %d", response.Message, id),
	}

	return pkg, nil
}
