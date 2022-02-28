package service

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
)

type PackageService interface {
	FindAll() []web.PackageResponse
	FindById(packageId int64) (entity.Package, error)
	Create(request web.PackageCreateRequest) entity.Package
	Update(request web.PackageUpdateRequest) entity.Package
	Delete(packageId int64) (entity.Package, error)
}
