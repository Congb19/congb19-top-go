package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)
type kbn_goods struct {
	id           uint `gorm:"not null; primaryKey; comment:id;"`
	floor        int `gorm:"not null; uniqueIndex; comment:楼层号; AUTO_INCREMENT;"`
	show         bool `gorm:"not null; default: true; comment:是否展示;"`
	author_id  	 int `gorm:"not null; comment:作者id;"` // 外键 关联users
	content      string `gorm:"not null; comment:内容;"`
	created_time time.Time
}
