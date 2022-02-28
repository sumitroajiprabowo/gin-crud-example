package repository

import (
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"gorm.io/gorm"
)

type RepositoryPackageImpl struct {
	db *gorm.DB
}

func NewRepositoryPackage(db *gorm.DB) *RepositoryPackageImpl {
	return &RepositoryPackageImpl{db}
}

func (r *RepositoryPackageImpl) FindAll() []entity.Package {
	var ps []entity.Package

	r.db.Find(&ps)
	return ps
}

func (r *RepositoryPackageImpl) FindById(packageId int64) (entity.Package, error) {
	var p entity.Package
	err := r.db.Where("id = ?", packageId).First(&p).Error
	return p, err
	// var p entity.Package
	// err := r.db.Find(&p, packageId).Error
	// return p, err

}

func (r *RepositoryPackageImpl) Create(p entity.Package) entity.Package {
	r.db.Create(&p)
	return p
}

func (r *RepositoryPackageImpl) Update(p entity.Package) entity.Package {
	r.db.Save(&p)
	return p
}

func (r *RepositoryPackageImpl) Delete(p entity.Package) {
	r.db.Delete(&p)
}
