package repo

import (
	"micro/model"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type IProductRepo interface {
	List(page common.Pagination) (model.Products, error)
	Get(id int) (*model.Product, error)
	Add(form *model.ProductForm) (*model.Product, error)
	Update(form *model.ProductForm, id int) (*model.Product, error)
	Patch(form *model.ProductPatchForm, id int) (*model.Product, error)
	Delete(id int) (*model.Product, error)
}

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		DB: db,
	}
}

func (r ProductRepo) List(page common.Pagination) (model.Products, error) {

	products := make([]*model.Product, 0)
	if err := r.DB.Scopes(common.Paginate(page)).Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (r ProductRepo) Get(id int) (*model.Product, error) {
	product := new(model.Product)
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepo) Delete(id int) (*model.Product, error) {
	product := new(model.Product)
	if err := r.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
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

func (r ProductRepo) Update(form *model.ProductForm, id int) (*model.Product, error) {
	product, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepo) Patch(form *model.ProductPatchForm, id int) (*model.Product, error) {
	product, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
