package service

import (
	"log"

	common "github.com/krishnarajvr/micro-common"

	"micro/app/locale"
	"micro/module/user/model"
	"micro/module/user/repo"
)

type IUserService interface {
	List(page common.Pagination) (model.UserDtos, *common.PageResult, error)
	Get(id int) (*model.UserDto, error)
	Add(user *model.UserForm) (*model.UserDto, error)
	Update(form *model.UserForm, id int) (*model.UserDto, error)
	Patch(form *model.UserPatchForm, id int) (*model.UserDto, error)
	Delete(id int) (*model.UserDto, error)
}

type ServiceConfig struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
}

type Service struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
}

func NewService(c ServiceConfig) IUserService {
	return &Service{
		UserRepo: c.UserRepo,
		Lang:     c.Lang,
	}
}

func (s *Service) List(page common.Pagination) (model.UserDtos, *common.PageResult, error) {

	users, pageResult, err := s.UserRepo.List(page)

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	//Multi language Support Demo
	log.Println(s.Lang.Get("hi", "GetAll"))

	m := map[string]interface{}{
		"Name": "John Doe",
		"Age":  66,
	}

	s.Lang.SetLang("en-US")
	log.Println(s.Lang.Get("intro", m))

	//common.Log(s.Contex, s.Lang.Get("intro", m))

	usersDto := users.ToDto()

	return usersDto, pageResult, err
}

func (s *Service) Get(id int) (*model.UserDto, error) {

	user, err := s.UserRepo.Get(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(user)

	userDto := user.ToDto()

	log.Println(userDto)

	return userDto, nil
}

func (s *Service) Add(form *model.UserForm) (*model.UserDto, error) {

	user, err := s.UserRepo.Add(form)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(user)

	userDto := user.ToDto()

	log.Println(userDto)

	return userDto, nil
}

func (s *Service) Update(form *model.UserForm, id int) (*model.UserDto, error) {

	user, err := s.UserRepo.Update(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(user)

	userDto := user.ToDto()

	log.Println(userDto)

	return userDto, nil
}

func (s *Service) Patch(form *model.UserPatchForm, id int) (*model.UserDto, error) {

	user, err := s.UserRepo.Patch(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(user)

	userDto := user.ToDto()

	log.Println(userDto)

	return userDto, nil
}

func (s *Service) Delete(id int) (*model.UserDto, error) {

	user, err := s.UserRepo.Delete(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(user)

	userDto := user.ToDto()

	log.Println(userDto)

	return userDto, nil
}
