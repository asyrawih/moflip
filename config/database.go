package config

import (
	"fmt"

	"github.com/hananloser/moflip/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := "host=localhost user=hanan password='' dbname=moflip port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

  db.AutoMigrate(&entity.User{})

	return db

}
