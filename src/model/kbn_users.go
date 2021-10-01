package model

import (
	_ "github.com/go-sql-driver/mysql"
)
type kbn_users struct {
	author_id  	 int `gorm:"not null; comment:作者id;"`
	author_name  string `gorm:"not null; comment:作者称呼; size:20"`
	contact_info string `gorm:"not null; comment:作者联系方式; size:50"`
}
