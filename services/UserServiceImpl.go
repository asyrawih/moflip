package services

import (
	"github.com/hananloser/moflip/entity"
	"github.com/hananloser/moflip/repository"
)

func NewUserService(userRepo *repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: *userRepo,
	}
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func (service *UserServiceImpl) FindUsernameAndPassword(email string, password string) entity.User {
	return service.userRepo.FindUsernameAndPassword(email, password)
}

func (service *UserServiceImpl) GetAllUser() []entity.User {
	result := service.userRepo.GetAll()
	return result
}

func (userserviceimpl *UserServiceImpl) FindUser(id int) entity.User {
	return entity.User{}
}
