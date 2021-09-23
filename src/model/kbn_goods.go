package model

type kbn_goods struct {
	id           string `gorm:"primaryKey;comment:证件号;size:20"`
	floor        string `gorm:"not null;uniqueIndex;comment:学号;size:20"`
	show         bool
	author_name  string
	contact_info string
	content      string
	created_time string
}
