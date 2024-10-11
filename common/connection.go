package common

import (
	"fmt"
	"log"

	"github.com/Splucheviy/codewithteeEchoLesson/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL() (*gorm.DB, error) {
	env := configs.NewEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Default().Println("Database connection established")
	return db, nil
}
