package repo

import (
	"micro/module/product/model"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type IProductRepo interface {
	List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductsList, *common.PageResult, error)
	Add(form *model.Product) (*model.Product, error)
	Get(tenantId int, id int) (*model.ProductDetail, error)
	Patch(form *model.ProductPatchForm, id int) (*model.Product, error)
	Delete(tenantId int, id int) (*model.Product, error)
}

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		DB: db,
	}
}

func (r ProductRepo) List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductsList, *common.PageResult, error) {
	products := make([]*model.ProductList, 0)
	var totalCount int64
	var db = r.DB
	db = db.Scopes(common.Paginate(page)).
		Table("products").
		Select(
			`products.id, 
			 products.name, 
			 products.code, 
			 products.is_active,
			 products.created_at, products.updated_at`).
		Where("products.tenant_id = ? ", tenantId)

	if len(filters.Code) > 0 {
		db = db.Where(`products.code LIKE ? `, filters.Code+"%")
	}

	if len(filters.Name) > 0 {
		db = db.Where(`products.name LIKE ? `, filters.Name+"%")
	}

	err := db.Find(&products).Count(&totalCount).Error

	if err != nil {
		return nil, nil, err
	}

	pageResult := common.PageInfo(page, totalCount)

	if len(products) == 0 {
		return nil, nil, nil
	}

	return products, &pageResult, nil
}

func (r ProductRepo) Add(product *model.Product) (*model.Product, error) {
	if err := r.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepo) Get(tenantId int, id int) (*model.ProductDetail, error) {
	product := new(model.ProductDetail)

	dbError := r.DB.Table("products").
		Select(
			`products.id, 
		     products.name, 
		     products.code, 
			 products.description,
			 products.meta,
		     products.is_active,
		     products.created_at, products.updated_at`).
		Where(`products.id = ? AND products.tenant_id = ?`, id, tenantId).
		First(&product).Error

	if dbError != nil {
		return nil, dbError
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

func (r ProductRepo) Delete(tenantId int, id int) (*model.Product, error) {
	product := new(model.Product)
	// Unscoped() hard delete operation
	if err := r.DB.Unscoped().Where("id = ? and tenant_id = ?", id, tenantId).Delete(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
