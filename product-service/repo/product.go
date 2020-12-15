package repo

import (
	"gorm.io/gorm"

	"micro/model"
)

type IProductRepo interface {
	ListProducts() (model.Products, error)
	GetById(id int) (*model.Product, error)
	Add(form *model.ProductForm) (*model.Product, error)
}

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		DB: db,
	}
}

func (r ProductRepo) ListProducts() (model.Products, error) {
	products := make([]*model.Product, 0)
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (r ProductRepo) GetById(id int) (*model.Product, error) {
	product := new(model.Product)
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepo) Add(form *model.ProductForm) (*model.Product, error) {
	product, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
