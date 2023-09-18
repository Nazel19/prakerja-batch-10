package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase (){
	DB, err : gorm.Open(mysql.Open("root:karya99@tcp(www.db4free.net:3306)/prakerjakar9"))
	if err != nil {
		panic(err)	
	}

	database.AutoMigarete(&Product{})

	DB = database
}