package handler

import (
	"fmt"
	"os"

	"github.com/congb19/congb19-top-go/src/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB()  {
	var err error
	data.DB, err = gorm.Open("mysql", "congb19-top:QWER1234@159.75.85.197/congb19-top?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		fmt.Println(err)
		os.Exit(111)
	}

	defer data.DB.Close()
}