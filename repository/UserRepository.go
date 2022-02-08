package repository

import "github.com/hananloser/moflip/entity"

type UserRepository interface {
	GetAll() []entity.User
}
