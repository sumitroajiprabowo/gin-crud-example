package repository

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
)

type PackageRepository interface {
	FindAll() []entity.Package
	FindById(packageId int64) (entity.Package, error)
	Create(p entity.Package) entity.Package
	Update(p entity.Package) entity.Package
	Delete(packageId int) (entity.Package, error)
}
