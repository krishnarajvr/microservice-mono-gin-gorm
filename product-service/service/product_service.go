package service

import (
	"log"
	"micro/app/locale"
	"micro/model"
	"micro/repo"

	common "github.com/krishnarajvr/micro-common"
)

type IProductService interface {
	List(page common.Pagination) (model.ProductDtos, error)
	Get(id int) (*model.ProductDto, error)
	Add(product *model.ProductForm) (*model.ProductDto, error)
	Update(form *model.ProductForm, id int) (*model.ProductDto, error)
	Patch(form *model.ProductPatchForm, id int) (*model.ProductDto, error)
	Delete(id int) (*model.ProductDto, error)
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

func (s *ProductService) List(page common.Pagination) (model.ProductDtos, error) {

	products, err := s.ProductRepo.List(page)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(s.Lang.Get("hi", "GetAll"))

	productsDto := products.ToDto()

	return productsDto, err
}

func (s *ProductService) Get(id int) (*model.ProductDto, error) {

	product, err := s.ProductRepo.Get(id)

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

func (s *ProductService) Update(form *model.ProductForm, id int) (*model.ProductDto, error) {

	product, err := s.ProductRepo.Update(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(product)

	productDto := product.ToDto()

	log.Println(productDto)

	return productDto, nil
}

func (s *ProductService) Patch(form *model.ProductPatchForm, id int) (*model.ProductDto, error) {

	product, err := s.ProductRepo.Patch(form, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(product)

	productDto := product.ToDto()

	log.Println(productDto)

	return productDto, nil
}

func (s *ProductService) Delete(id int) (*model.ProductDto, error) {

	product, err := s.ProductRepo.Delete(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(product)

	productDto := product.ToDto()

	log.Println(productDto)

	return productDto, nil
}
