package repository

import (
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

func (repo *UserRepositoryImpl) GetAll() []entity.User {

	var users []entity.User

	repo.db.Find(&users)

	return users
}
