package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type KbnUsers struct {
	gorm.Model
	AuthorName  string `gorm:"not null; size:20"`
	ContactInfo string `gorm:"not null; size:50"`
}
