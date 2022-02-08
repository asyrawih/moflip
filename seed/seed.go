package seed

import (
	"github.com/hananloser/moflip/entity"
	"gorm.io/gorm"
)


func CreateUser(db *gorm.DB, user *entity.User) error {
	return db.Create(&user).Error
}
