package repository

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
)

type PackageRepository interface {
	FindAll() ([]entity.Package, error)
	FindById(id int) (entity.Package, error)
	Create(p entity.Package) (entity.Package, error)
	Update(p entity.Package) (entity.Package, error)
	Delete(id int) (entity.Package, error)
}
