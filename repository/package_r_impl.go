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

func (r *RepositoryPackageImpl) FindAll() ([]entity.Package, error) {
	var p []entity.Package

	err := r.db.Find(&p).Error
	return p, err
}

func (r *RepositoryPackageImpl) FindById(id int) (entity.Package, error) {
	var p entity.Package

	// err := r.db.First(&p, id).Error
	err := r.db.Find(&p, id).Error
	return p, err
}

func (r *RepositoryPackageImpl) Create(p entity.Package) (entity.Package, error) {
	err := r.db.Create(&p).Error
	return p, err
}

func (r *RepositoryPackageImpl) Update(p entity.Package) (entity.Package, error) {
	err := r.db.Save(&p).Error
	return p, err
}

func (r *RepositoryPackageImpl) Delete(id int) (entity.Package, error) {
	var p entity.Package

	err := r.db.Delete(&p, id).Error
	return p, err
}
