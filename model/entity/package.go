// models/entity/package.go

package entity

import (
	"time"
)

type Package struct {
	Id        int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
