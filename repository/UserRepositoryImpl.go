package repository

import (
	"fmt"

	"github.com/hananloser/moflip/entity"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: *db,
	}
}

type UserRepositoryImpl struct {
	db gorm.DB
}

// FindUserByName
func (repo *UserRepositoryImpl) FindUsernameAndPassword(email string, password string) entity.User {
	var user entity.User
	err := repo.db.Where("email=? AND password=?", email, password).Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func (repo *UserRepositoryImpl) GetAll() []entity.User {

	var users []entity.User

	repo.db.Find(&users)

	return users
}
