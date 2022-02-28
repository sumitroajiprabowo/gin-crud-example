package helper

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
)

func ToPackageResponse(pkg *entity.Package) *web.PackageResponse {
	return &web.PackageResponse{
		Id:   pkg.Id,
		Name: pkg.Name,
	}
}

func ToPackagesResponseList(pkgs []entity.Package) []web.PackageResponse {
	var pkgResponses []web.PackageResponse
	for _, pkg := range pkgs {
		pkgResponses = append(pkgResponses, *ToPackageResponse(&pkg))
	}
	return pkgResponses
}
