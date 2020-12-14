package service

import (
	"micro/app/locale"
	"micro/repo"
)

type ServiceConfig struct {
	ProductRepo repo.IProductRepo
	Lang        *locale.Locale
}
