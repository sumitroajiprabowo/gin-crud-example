// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-crud-example/controller"
	"github.com/sumitroajiprabowo/gin-crud-example/repository"
	"github.com/sumitroajiprabowo/gin-crud-example/service"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	packageRepository := repository.NewRepositoryPackage(db)
	packageService := service.NewPackageService(packageRepository)
	packageController := controller.NewPackageController(packageService)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/package", packageController.FindAllPackages)
	r.POST("/package", packageController.CreatePackages)
	r.GET("/package/:id", packageController.FindPackagesById)
	r.PATCH("/package/:id", packageController.UpdatePackages)
	r.DELETE("/package/:id", packageController.DeletePackages)
	return r
}
