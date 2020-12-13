package service

import (
	"micro/app/locale"
	"micro/repo"
)

type ServiceConfig struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
}
