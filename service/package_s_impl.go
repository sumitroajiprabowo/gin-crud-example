package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/exception"
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
	p, err := s.PackageRepository.FindById(packageId)
	return p, err
}

func (s *PackageServiceImpl) Create(request web.PackageCreateRequest) entity.Package {

	p := entity.Package{
		Name: request.Name,
	}

	createdPackage := s.PackageRepository.Create(p)

	return createdPackage

}

func (s *PackageServiceImpl) Update(request web.PackageUpdateRequest) entity.Package {

	// check if package exist or not found
	_, err := s.PackageRepository.FindById(int64(request.Id))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	p := entity.Package{
		Id:   int64(request.Id),
		Name: request.Name,
	}

	updatedPackage := s.PackageRepository.Update(p)

	return updatedPackage

}

func (s *PackageServiceImpl) Delete(packageId int64) (entity.Package, error) {
	_, err := s.PackageRepository.FindById(int64(packageId))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	deletePackage, err := s.PackageRepository.Delete(int(packageId))

	return deletePackage, err
}
