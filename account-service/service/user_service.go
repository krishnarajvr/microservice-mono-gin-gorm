package service

import (
	"log"
	"micro/app/locale"
	"micro/model"
	"micro/repo"
)

type IUserService interface {
	GetAll() (model.UserDtos, error)
}

type UserService struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
}

func NewUserService(c *ServiceConfig) IUserService {
	return &UserService{
		UserRepo: c.UserRepo,
		Lang:     c.Lang,
	}
}

func (s *UserService) GetAll() (model.UserDtos, error) {

	u, err := s.UserRepo.ListUsers()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	udto := u.ToDto()

	return udto, err
}
