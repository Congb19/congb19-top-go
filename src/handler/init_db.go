package handler

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"congb19-top-go/src/data"
	"congb19-top-go/src/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	var err error

	user := viper.GetString("db.user")
	pwd := viper.GetString("db.pwd")
	name := viper.GetString("db.name")
	ip := viper.GetString("db.ip")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, ip, name)
	data.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}
	data.DB.SingularTable(true)

	// 测试
	data.DB.AutoMigrate(&model.Persons{})
	//test := data.DB.First(&model.Persons{}, "Id = ?", 3)
	//fmt.Println(test)

	// congb19-top
	data.DB.AutoMigrate(&model.KbnUsers{})
	data.DB.AutoMigrate(&model.KbnItems{})

	// app
	data.DB.AutoMigrate(&model.Clubs{})
	data.DB.AutoMigrate(&model.Tools{})
	data.DB.AutoMigrate(&model.Articles{})
	//data.DB.AutoMigrate(&model.Pins{})
	data.DB.AutoMigrate(&model.Users{})

	//defer data.DB.Close()
}
