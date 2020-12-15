package service

import (
	"log"
	"micro/app/locale"
	"micro/model"
	"micro/repo"
)

type IProductService interface {
	GetAll() (model.ProductDtos, error)
	GetById(id int) (*model.ProductDto, error)
	Add(product *model.ProductForm) (*model.ProductDto, error)
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

	products, err := s.ProductRepo.ListProducts()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	productsDto := products.ToDto()

	return productsDto, err
}

func (s *ProductService) GetById(id int) (*model.ProductDto, error) {

	product, err := s.ProductRepo.GetById(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(product)

	productDto := product.ToDto()

	log.Println(productDto)

	return productDto, nil
}

func (s *ProductService) Add(form *model.ProductForm) (*model.ProductDto, error) {

	product, err := s.ProductRepo.Add(form)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(product)

	productDto := product.ToDto()

	log.Println(productDto)

	return productDto, nil
}
