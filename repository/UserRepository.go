package repository

import "github.com/hananloser/moflip/entity"

type UserRepository interface {
	GetAll() []entity.User

  FindUsernameAndPassword(email string , password string) entity.User
}
