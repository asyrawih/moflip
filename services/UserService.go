package services

import "github.com/hananloser/moflip/entity"

type UserService interface {

	GetAllUser() []entity.User

	FindUser(id int) entity.User
}
