package service

import (
	"encoding/json"
	"micro/module/product/model"
	"micro/module/product/repo"

	"gorm.io/datatypes"

	common "github.com/krishnarajvr/micro-common"

	"github.com/krishnarajvr/micro-common/locale"
)

type IProductService interface {
	List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductListDtos, *common.PageResult, error)
	Get(tenantId int, id int) (*model.ProductDetailDto, error)
	Add(tenantId int, content *model.ProductForm) (*model.ProductDto, error)
	Patch(form *model.ProductPatchForm, id int) (*model.ProductDto, error)
	ConvertMapToJson(metaMap *map[string]interface{}) (datatypes.JSON, error)
	CheckProductSchema(tenantId int, jsonData []byte, validateFor string) (map[string]interface{}, *map[string]interface{}, error)
	Delete(tenantId int, id int) (*model.ProductDto, error)
}

type ServiceConfig struct {
	ProductRepo repo.IProductRepo
	Lang        *locale.Locale
}

type Service struct {
	ProductRepo repo.IProductRepo
	Lang        *locale.Locale
}

func NewService(c ServiceConfig) IProductService {
	return &Service{
		ProductRepo: c.ProductRepo,
		Lang:        c.Lang,
	}
}

func (s *Service) List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductListDtos, *common.PageResult, error) {
	productsList, pageResult, err := s.ProductRepo.List(tenantId, page, filters)

	if err != nil {
		return nil, nil, err
	}

	productListDtos := productsList.ToDto()

	return productListDtos, pageResult, err
}

func (s *Service) Add(tenantId int, form *model.ProductForm) (*model.ProductDto, error) {
	productModel, err := form.ToModel()
	productModel.TenantId = tenantId

	if err != nil {
		return nil, err
	}

	product, err := s.ProductRepo.Add(productModel)

	if err != nil {
		return nil, err
	}

	productDto := product.ToDto()

	return productDto, nil
}

func (s *Service) Get(tenantId int, id int) (*model.ProductDetailDto, error) {
	productDetail, err := s.ProductRepo.Get(tenantId, id)

	if err != nil {
		return nil, err
	}

	productDetailDto := productDetail.ToProductDetailDto()

	return productDetailDto, nil
}

func (s *Service) Patch(form *model.ProductPatchForm, id int) (*model.ProductDto, error) {
	product, err := s.ProductRepo.Patch(form, id)

	if err != nil {
		return nil, err
	}

	productDto := product.ToDto()

	return productDto, nil
}

func (s *Service) CheckProductSchema(tenantId int, jsonData []byte, validateFor string) (map[string]interface{}, *map[string]interface{}, error) {
	// currently JSONSCHEMA is hardcoded constant value, feature will read from database
	const METAJSONSCHEMA = `{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"title": "product",
		"description": "product creation",
		"type": "object",
		"properties": {
			"author": {
				"description": "Product Author",
				"type": "string",
				"minLength": 1
			},
			"country": {
				"description": "Product Country",
				"type": "string",
				"minLength": 1
			}
		},
		"required": ["author"]
	}`

	var schemaDef map[string]interface{}
	var formMeta *map[string]interface{}

	if validateFor == "validateMetaData" {

		// convert jsonschema to map[string]interface{}
		err := json.Unmarshal([]byte(METAJSONSCHEMA), &schemaDef)

		if err != nil {
			return nil, nil, err
		}

	}
	err := json.Unmarshal([]byte(jsonData), &formMeta)

	return schemaDef, formMeta, err
}

func (s *Service) ConvertMapToJson(metaMap *map[string]interface{}) (datatypes.JSON, error) {
	metaJson, err := json.Marshal(&metaMap)

	if err != nil {
		return nil, err
	}
	return metaJson, nil
}

func (s *Service) Delete(tenantId int, id int) (*model.ProductDto, error) {
	product, err := s.ProductRepo.Delete(tenantId, id)

	if err != nil {
		return nil, err
	}

	productDto := product.ToDto()

	return productDto, nil
}
