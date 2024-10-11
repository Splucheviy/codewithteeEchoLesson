package main

import (
	"log"

	"github.com/Splucheviy/codewithteeEchoLesson/common"
	"github.com/Splucheviy/codewithteeEchoLesson/internal/models"
)

func main() {
	db, err := common.NewMySQL()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}
	log.Println("Migration completed successfully")
}
