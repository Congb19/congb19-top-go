package model

import "github.com/jinzhu/gorm"

type Articles struct {
	gorm.Model
	Title     string `gorm:"size:50;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorId  uint
	Tags      string `gorm:"size:255;not null"`
	IsPinned  bool   `gorm:"not null"`
	ValidFlag bool   `gorm:"not null"`
}
type Replies struct {
	gorm.Model
	Content   string `gorm:"size:255;not null"`
	ArticleId uint   `gorm:"not null"`
	AuthorId  uint   `gorm:"not null"`
	Floor     uint   `gorm:"not null"`
	ValidFlag bool   `gorm:"not null"`
}
type ReReplies struct {
	gorm.Model
	Content       string `gorm:"size:255;not null"`
	ReplyId       uint   `gorm:"not null"`
	ReplyAuthorId uint   `gorm:"not null"`
	AuthorId      uint   `gorm:"not null"`
	Floor         uint   `gorm:"not null"`
	ValidFlag     bool   `gorm:"not null"`
}
