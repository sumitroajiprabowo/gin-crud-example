package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/helper"
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
	"github.com/sumitroajiprabowo/gin-crud-example/repository"
)

type PackageServiceImpl struct {
	PackageRepository repository.PackageRepository
}

func NewPackageService(packageRepository repository.PackageRepository) *PackageServiceImpl {
	return &PackageServiceImpl{packageRepository}
}

func (s *PackageServiceImpl) FindAll() []web.PackageResponse {
	packages := s.PackageRepository.FindAll()
	return helper.ToPackagesResponseList(packages)
}

func (s *PackageServiceImpl) FindById(packageId int64) (entity.Package, error) {
	p, err := s.PackageRepository.FindById(int64(packageId))
	return p, err
}

func (s *PackageServiceImpl) Create(request web.PackageCreateRequest) entity.Package {

	p := entity.Package{
		Name: request.Name,
	}

	createdPackage := s.PackageRepository.Create(p)

	return createdPackage
}

// Create a new package with the given request data and return the update package
func (s *PackageServiceImpl) Update(packageId int64, request web.PackageUpdateRequest) (entity.Package, error) {

	p, err := s.PackageRepository.FindById(int64(packageId))
	if err != nil {
		return p, err
	}

	p.Name = request.Name

	updatedPackage := s.PackageRepository.Update(p)

	return updatedPackage, err
}

func (s *PackageServiceImpl) Delete(packageId int64) error {

	p, err := s.PackageRepository.FindById(int64(packageId))
	if err != nil {
		return err
	}

	s.PackageRepository.Delete(p)

	return nil
}
