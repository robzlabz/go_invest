package database

import "go_invest/app/model"

func RunDatabaseMigrations() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
