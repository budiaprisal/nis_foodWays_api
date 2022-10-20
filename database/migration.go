package database

import (
	"fmt"
	"taskgolang/models"
	"taskgolang/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
