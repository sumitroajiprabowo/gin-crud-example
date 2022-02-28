package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/exception"
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

func (s *PackageServiceImpl) FindAll() ([]entity.Package, error) {
	return s.PackageRepository.FindAll()
}

func (s *PackageServiceImpl) FindById(packageId int64) (entity.Package, error) {
	p, err := s.PackageRepository.FindById(int(packageId))
	return p, err
}

func (s *PackageServiceImpl) Create(request web.PackageCreateRequest) (entity.Package, error) {
	p := entity.Package{
		Name: request.Name,
	}
	newPackage, err := s.PackageRepository.Create(p)
	return newPackage, err
}

func (s *PackageServiceImpl) Update(request web.PackageUpdateRequest) (entity.Package, error) {
	p, err := s.PackageRepository.FindById(int(request.ID))

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update package
	p.Name = request.Name
	updatedPackage, err := s.PackageRepository.Update(p)
	return updatedPackage, err
}

func (s *PackageServiceImpl) Delete(packageId int64) (entity.Package, error) {
	p, err := s.PackageRepository.FindById(int(packageId))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	deletePackage, err := s.PackageRepository.Delete(int(p.ID))

	return deletePackage, err
}
