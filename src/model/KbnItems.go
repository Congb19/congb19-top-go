package model

import "github.com/jinzhu/gorm"

import (
	_ "github.com/go-sql-driver/mysql"
)

type KbnItems struct {
	gorm.Model
	AuthorName  string `gorm:"not null; size:20"`
	Type        int    `gorm:"not null; "`
	Show        bool   `gorm:"not null; "`
	ContactInfo string `gorm:"not null; "`
	Content     string `gorm:"not null; "`
}
