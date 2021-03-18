package pkg

type packageRepository interface {
}

type PackageService struct {
	Repository packageRepository
}
