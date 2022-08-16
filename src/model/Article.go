package model

import "github.com/jinzhu/gorm"

// 帖子
type Articles struct {
	gorm.Model
	Title     string `gorm:"size:50;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorId  uint   `gorm:"not null"`
	ClubId    uint   `gorm:"not null"`
	Tags      string `gorm:"size:255;not null"`
	CanReply  bool   `gorm:"not null"`
	IsPinned  bool   `gorm:"not null"`
	ValidFlag bool   `gorm:"not null"`
}

// 回复
type Replies struct {
	gorm.Model
	Content   string `gorm:"size:255;not null"`
	ArticleId uint   `gorm:"not null"`
	AuthorId  uint   `gorm:"not null"`
	Floor     uint   `gorm:"not null"`
	ValidFlag bool   `gorm:"not null"`
}

// 楼中楼
type ReReplies struct {
	gorm.Model
	Content       string `gorm:"size:255;not null"`
	ReplyId       uint   `gorm:"not null"`
	ReplyAuthorId uint   `gorm:"not null"`
	AuthorId      uint   `gorm:"not null"`
	Floor         uint   `gorm:"not null"`
	ValidFlag     bool   `gorm:"not null"`
}

// 帖子置顶关系表
//type Pins struct {
//	gorm.Model
//	ArticleId uint `gorm:"not null"`
//	ClubId    uint `gorm:"not null"`
//	ValidFlag bool `gorm:"not null"`
//}
