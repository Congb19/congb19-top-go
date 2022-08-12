package handler

import (
	"fmt"
	"os"

	"congb19-top-go/src/data"
	"congb19-top-go/src/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	var err error
	//data.DB, err = gorm.Open("mysql", "congb19-top:QWER1234@tcp(159.75.85.197)/congb19-top?charset=utf8&parseTime=True&loc=Local")
	data.DB, err = gorm.Open("mysql", "vmrs:QWER1234@tcp(159.75.85.197)/vmrs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}
	data.DB.SingularTable(true)
	data.DB.AutoMigrate(&model.Persons{})

	//data.DB.Create(&model.Persons{Id: 3, LastName: "test", FirstName: "test", Address: "test", City: "test"})

	//test := data.DB.First(&model.Persons{}, "Id = ?", 3)
	//fmt.Println(test)

	//defer data.DB.Close()
}
