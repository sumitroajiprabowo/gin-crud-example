package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
)

type PackageService interface {
	FindAll() []web.PackageResponse
	FindById(packageId int64) (entity.Package, error)
	Create(request web.PackageCreateRequest) entity.Package
	Update(packageId int64, request web.PackageUpdateRequest) (entity.Package, error)
	Delete(packageId int64) error
}
