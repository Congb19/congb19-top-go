package model

import "github.com/jinzhu/gorm"

type Persons struct {
	gorm.Model
	//Id        int16 `gorm:"primary_key"`
	LastName  string
	FirstName string
	Address   string
	City      string
}
