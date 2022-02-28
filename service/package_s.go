package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
)

type PackageService interface {
	FindAll() ([]entity.Package, error)
	FindById(packageId int64) (entity.Package, error)
	Create(p web.PackageCreateRequest) (entity.Package, error)
	Update(p web.PackageUpdateRequest) (entity.Package, error)
	Delete(packageId int64) (entity.Package, error)
}
