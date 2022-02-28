// controllers/task.go

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-crud-example/exception"
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
	"github.com/sumitroajiprabowo/gin-crud-example/service"
	"gorm.io/gorm"
)

type PackageControllerImpl struct {
	PackageService service.PackageService
}

func NewPackageController(PackageService service.PackageService) *PackageControllerImpl {
	return &PackageControllerImpl{PackageService}
}

func (cl *PackageControllerImpl) FindAllPackages(c *gin.Context) {
	packages, err := cl.PackageService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := convertToPackageResponseList(packages)

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   result,
	})
}

func (cl *PackageControllerImpl) CreatePackages(c *gin.Context) {
	// Validate input
	var input web.PackageCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		exception.ErrorBadRequest(c, err)
		return
	}

	// Create package
	p := entity.Package{Name: input.Name}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&p)

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   p,
	})
}

// // GET /package/:id
// // Find a packages by id
func (cl *PackageControllerImpl) FindPackagesById(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var p entity.Package
	if err := db.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {

		exception.ErrorNotFound(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   p,
	})
}

func (cl *PackageControllerImpl) UpdatePackages(c *gin.Context) {
	// Validate input
	var input web.PackageUpdateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		exception.ErrorBadRequest(c, err)
		return
	}

	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var p entity.Package
	if err := db.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		exception.ErrorNotFound(c, err)
		return
	}

	// Update model
	p.Name = input.Name

	db.Save(&p)

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   p,
	})
}

func (cl *PackageControllerImpl) DeletePackages(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var p entity.Package
	if err := db.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {
		exception.ErrorNotFound(c, err)
		return
	}

	// Delete model
	db.Delete(&p)

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   p,
	})
}

// helper
func convertToPackageResponse(p entity.Package) web.WebResponse {
	return web.WebResponse{
		ID:   p.ID,
		Name: p.Name,
	}
}

func convertToPackageResponseList(ps []entity.Package) []web.PackageResponse {
	var result []web.PackageResponse

	for _, p := range ps {
		packageResponse := convertToPackageResponse(p)
		result = append(result, web.PackageResponse{
			ID:   packageResponse.ID,
			Name: packageResponse.Name,
		})
	}

	return result
}
