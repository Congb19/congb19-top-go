package model

import "github.com/jinzhu/gorm"

type Users struct {
	// 登录使用ID和password
	gorm.Model
	Username  string `gorm:"size:50;notnull"` //昵称
	Password  string `gorm:"size:255;notnull"`
	AvatarSrc string `gorm:"size:50;notnull"`
	ValidFlag bool   `gorm:"not null"`
}
type Follows struct {
	// 登录使用ID和password
	gorm.Model
	FollowerId uint `gorm:"notnull"` //昵称
	ValidFlag  bool `gorm:"not null"`
}
