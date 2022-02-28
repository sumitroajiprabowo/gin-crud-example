// controllers/packages.go

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/gin-crud-example/exception"
	"github.com/sumitroajiprabowo/gin-crud-example/model/web"
	"github.com/sumitroajiprabowo/gin-crud-example/service"
)

type PackageControllerImpl struct {
	PackageService service.PackageService
}

func NewPackageController(PackageService service.PackageService) *PackageControllerImpl {
	return &PackageControllerImpl{PackageService}
}

func (cl *PackageControllerImpl) FindAllPackages(c *gin.Context) {
	packages := cl.PackageService.FindAll() // Get all packages from service

	// Create response data with status code and data payload (packages)
	webResponse := web.WebResponse{
		Code:   int(http.StatusOK), // Status code 200
		Status: "Success",          // Status message
		Data:   packages,           // Data payload
	}

	c.JSON(http.StatusOK, webResponse) // Return response

}

func (cl *PackageControllerImpl) CreatePackages(c *gin.Context) {
	// Validate input
	var input web.PackageCreateRequest

	// Validate input with binding schema
	if err := c.ShouldBindJSON(&input); err != nil {
		exception.ErrorBadRequest(c, err) // Return 400
		return
	}

	// Input data to model
	p := cl.PackageService.Create(input)

	// Create response data with status code and payload data (p)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   p,
	}

	c.JSON(http.StatusOK, webResponse) // Return response
}

// // GET /package/:id
// // Find a packages by id
// // GET /package/:id
// // Find a packages by id
func (cl *PackageControllerImpl) FindPackagesById(c *gin.Context) {
	// // Get model if exist
	// db := c.MustGet("db").(*gorm.DB)
	// var p entity.Package
	// if err := db.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {

	// 	exception.ErrorNotFound(c, err)
	// 	return
	// }

	// Check with param id with FindById method
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	p, err := cl.PackageService.FindById(int64(id))

	if err != nil {
		exception.ErrorNotFound(c, err) // Return 404
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code":   int(http.StatusOK),
		"Status": "Success",
		"data":   p,
	})
}

func (cl *PackageControllerImpl) UpdatePackage(c *gin.Context) {
	// Validate input id parameter (c.Param("id")) with FindById method
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	p, err := cl.PackageService.FindById(int64(id))

	if err != nil {
		exception.ErrorNotFound(c, err) // Return 404
		return
	}

	// Validate input
	var input web.PackageUpdateRequest

	// Validate input with binding schema
	if err := c.ShouldBindJSON(&input); err != nil {
		exception.ErrorBadRequest(c, err) // Return 400
		return
	}

	// Update model
	p = cl.PackageService.Update(input)

	// Create response data with status code and payload data (p)
	webResponse := web.WebResponse{
		Code:   int(http.StatusOK), // Status code 200
		Status: "Success",          // Status message
		Data:   p,                  // Data payload
	}

	c.JSON(http.StatusOK, webResponse) // Return response
}

// // Get model if exist
// db := c.MustGet("db").(*gorm.DB)
// var p entity.Package
// if err := db.Where("id = ?", c.Param("id")).First(&p).Error; err != nil {

// 	exception.ErrorNotFound(c, err)
// 	return
// }

// c.JSON(http.StatusOK, gin.H{
// 	"Code":   int(http.StatusOK),
// 	"Status": "Success",
// 	"data":   p,
// })

func (cl *PackageControllerImpl) DeletePackages(c *gin.Context) {
	// Validate input id parameter (c.Param("id")) with FindById method
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	p, err := cl.PackageService.FindById(int64(id))

	if err != nil {
		exception.ErrorNotFound(c, err) // Return 404
		return
	}

	// Delete model
	cl.PackageService.Delete(p.Id)

	// Create response data with status code and payload data (p)
	webResponse := web.WebResponse{
		Code:   int(http.StatusOK), // Status code 200
		Status: "Success",          // Status message
	}

	c.JSON(http.StatusOK, webResponse) // Return response
}
