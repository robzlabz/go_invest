package database

import (
	"fmt"
	"go_invest/app/model"
)

func RunDatabaseMigrations() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated")
}
