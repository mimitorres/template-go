package pkg

import (
	"context"
	"errors"
	"strconv"
	"template-go/src/api/pkg/model"
)

type packageRepository interface {
	GetPackageByID(context.Context, int) (model.Package, error)
}

type PackageService struct {
	Repository packageRepository
}

func (s PackageService) GetPackageByID(ctx context.Context, id string) (model.Package, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return model.Package{}, errors.New("cannot parse id")
	}

	res, err := s.Repository.GetPackageByID(ctx, parsedID)
	if err != nil {
		return model.Package{}, err
	}

	return res, nil
}
