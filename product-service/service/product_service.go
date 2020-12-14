package service

import (
	"log"
	"micro/app/locale"
	"micro/model"
	"micro/repo"
)

type IProductService interface {
	GetAll() (model.ProductDtos, error)
}

type ProductService struct {
	ProductRepo repo.IProductRepo
	Lang        *locale.Locale
}

func NewProductService(c *ServiceConfig) IProductService {
	return &ProductService{
		ProductRepo: c.ProductRepo,
		Lang:        c.Lang,
	}
}

func (s *ProductService) GetAll() (model.ProductDtos, error) {

	u, err := s.ProductRepo.ListProducts()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	udto := u.ToDto()

	return udto, err
}
