package model

import "github.com/jinzhu/gorm"

type Clubs struct {
	gorm.Model
	Title     string `gorm:"size:20;not null"`
	ImgSrc    string `gorm:"size:255;not null"` // 背景图片
	ValidFlag bool   `gorm:"not null"`
}
type Tools struct {
	gorm.Model
	Title     string `gorm:"size:20;not null"`
	ImgSrc    string `gorm:"size:255;not null"` // logo图片
	ValidFlag bool   `gorm:"not null"`
}
